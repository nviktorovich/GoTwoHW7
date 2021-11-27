package main

//go:generate stringer -type=Page

type Page int

//const (
//	Ya       Page = "https://market.yandex.ru"
//	RBT      Page = "https://www.rbt.ru"
//	Eldorado Page = "https://www.eldorado.ru"
//	MVideo   Page = "https://www.mvideo.ru"
//)

const (
	Ya       Page = iota
	RBT
	Eldorado
	MVideo
)