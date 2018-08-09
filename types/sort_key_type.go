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
	return string(k)
}

func (k SortKey) SqlPrefix(prefix string) string {
	strKey := string(k)
	if strings.HasPrefix(strKey, "-") {
		return "-" + prefix + strKey[1:]
	}

	return prefix + strKey
}