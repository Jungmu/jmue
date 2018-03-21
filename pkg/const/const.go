package jmueConst

//Mode : logging Mode
type Mode int

const (
	//DEBUG = 1
	DEBUG Mode = 1 + iota
	//SERVICE = 2
	SERVICE
	//TEST = 3
	TEST
)
