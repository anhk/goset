package goset

type Set interface {
	Add(interface{})
	Del(interface{})
	IsExist(interface{}) bool
}

func __for_test() {
	var _ Set = &StringSet{}
}
