package gorm_ext

import (
	"github.com/jinzhu/gorm"
	"github.com/drskur/gorm-ext/types"
)

type Oper int

const (
	Eq   Oper = iota
	Like
	In
	Gt
	Gte
	Lt
	Lte
)

func (dbq DBQuery) Query() string {
	var query string
	switch dbq.Operator {
	case Eq:
		query = dbq.Field + " = ?"
	case Like:
		query = dbq.Field + " LIKE ?"
	case In:
		query = dbq.Field + " IN (?)"
	case Gt:
		query = dbq.Field + " > ?"
	case Gte:
		query = dbq.Field + " >= ?"
	case Lt:
		query = dbq.Field + " < ?"
	case Lte:
		query = dbq.Field + " <= ?"
	}
	return query
}

type DBQuery struct {
	Field    string
	Operator Oper
	Arg      interface{}
}

// WhereOpt returns gorm DB
// if args don't contain zero value (nil, '', 0) then apply where query.
func WhereOpt(db *gorm.DB, query string, args ...interface{}) *gorm.DB {
	for _, v := range args {
		if isBlank(v) {
			return db
		}
	}

	return db.Where(query, args...)
}

func WhereArrOpt(db *gorm.DB, queryArr []DBQuery) *gorm.DB {
	for _, q := range queryArr {
		if !isBlank(q.Arg) {
			var arg interface{}

			switch v := q.Arg.(type) {
			case types.BoolKey:
				if ok, err := v.Parse(); err != nil {
					continue
				} else {
					arg = ok
				}

			default:
				arg = q.Arg

			}
			db = db.Where(q.Query(), arg)
		}
	}

	return db
}
