package goset

type StringSet map[string]int

func str(arg interface{}) string {
	return arg.(string)
}

func (set *StringSet) Add(arg interface{}) {
	(*set)[str(arg)] = 1
}

func (set *StringSet) Del(arg interface{}) {
	delete(*set, str(arg))
}

func (set *StringSet) IsExist(arg interface{}) bool {
	if _, ok := (*set)[str(arg)]; ok {
		return true
	}
	return false
}

func NewStringSet(args ...string) *StringSet {
	strSet := StringSet{}
	for _, v := range args {
		strSet[v] = 1
	}
	return &strSet
}
