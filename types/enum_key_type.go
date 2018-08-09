package types

type EnumKey string

func (k EnumKey) IsValid(enumable []string) bool {
	for _, v := range enumable {
		if string(k) == v {
			return true
		}
	}

	return false
}

func (k EnumKey) String() string {
	return string(k)
}