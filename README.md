# gorm-ext
This is helper functions for gorm

## Usage
```go
import (
	ext "github.com/drskur/gorm-ext"
)

db, err := gorm.Open("provider", "connection string")
```

### Page
```go
page := 2
limit := 10
db = ext.Page(db, page, limit)

// LIMIT 10 OFFSET 10
```

### WhereOpt
if args is zero value, then ignore query.
```go
a := "John"
b := ""

db = ext.WhereOpt(db, "name = ?", a)
// WHERE name = 'John'

db = ext.WhereOpt(db, "name = ?", b)
// ignore this query
```

### WhereArrOpt
```go
db = ext.WhereArrOpt(db, []ext.DBQuery{
    {Field: "name", Operator: ext.Like, Arg: "%John%"}, // name LIKE '%John%'
    {Field: "department", Operator: ext.Eq, Arg: "sales"}, // department = 'sales'
    {Field: "email", Operator: ext.Eq, Arg: ""}, // ignore
    {Field: "age", Operator: ext.Gt, Arg: 30}, // age > 30
})
```