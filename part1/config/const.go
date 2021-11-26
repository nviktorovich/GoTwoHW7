package config

import "reflect"

var (
	String       = reflect.TypeOf("")
	NumericInt   = reflect.TypeOf(12)
	NumericFloat = reflect.TypeOf(12.2)
	Bool         = reflect.TypeOf(true)
	ChanInt      = reflect.TypeOf(make(chan int))
	ChanString   = reflect.TypeOf(make(chan int))
	ChanBool     = reflect.TypeOf(make(chan bool))
	ChanFloat    = reflect.TypeOf(make(chan float64))
	ChanByte     = reflect.TypeOf(make(chan byte))
	ChanStruct     = reflect.TypeOf(make(chan struct{}))
)
