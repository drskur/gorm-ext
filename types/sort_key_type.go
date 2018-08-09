package types

import "strings"

type SortKey string

func (k SortKey) IsValid(sortable []string) bool {
	for _, v := range sortable {
		if string(k) == v {
			return true
		}
	}

	return false
}

func (k SortKey) Sql() string {
	str := string(k)
	if strings.HasPrefix(str,"-") {
		return str[1:] + " DESC"
	}

	return string(k)
}

func (k SortKey) SqlPrefix(prefix string) string {
	strKey := string(k)
	if strings.HasPrefix(strKey, "-") {
		return prefix + k.Sql()
	}

	return prefix + strKey
}