package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"reflectProject/config"
)

//Написать функцию, которая принимает на вход структуру in (struct или кастомную
//struct) и values map[string]interface{} (key - название поля структуры,
//которому нужно присвоить value этой мапы). Необходимо по значениям из мапы
//изменить входящую структуру in с помощью пакета reflect. Функция может
//возвращать только ошибку error. Написать к данной функции тесты (чем больше,
//тем лучше - зачтется в плюс).


type New struct {
	key interface{}
}


func typeTester(i interface{}) (ans bool) {
	valueSlice := []reflect.Type{
		config.String,
		config.NumericInt,
		config.NumericFloat,
		config.Bool,
		config.ChanInt,
		config.ChanString,
		config.ChanBool,
		config.ChanFloat ,
		config.ChanByte,
		config.ChanStruct,
	}
	for _, v := range valueSlice{
		if v == reflect.TypeOf(i){
			return true
		}

	}
	return false
}

func Remake(in *New, values map[string]interface{}) error {
	for _, v := range values{
		if typeTester(v){
			in.key = reflect.ValueOf(v)
			return nil
		}

	}
	return errors.New("неизвестный тип")
}

func main()  {
	var a New
	b := make(map[string]interface{})
	b["new"] = make(chan int)

	if err := Remake(&a, b); err != nil{
		log.Println("ошибка типа,", err)
		os.Exit(1)
	}
	fmt.Println(a.key)

}