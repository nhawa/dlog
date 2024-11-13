package dlog

import (
	"encoding/json"
	"fmt"
	"reflect"
	"runtime"

	"github.com/sirupsen/logrus"
)

func PrintStructFields(x ...interface{}) {
	for _, s := range x {
		v := reflect.ValueOf(s)
		t := reflect.TypeOf(s)

		fmt.Println("=======================================================================================================")
		fmt.Printf("Fields of struct %s:\n", t.Name())
		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ *** ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			fieldName := t.Field(i).Name
			fieldValue := field.Interface()

			fmt.Printf("%s\t: %+v\n", fieldName, fieldValue)
		}
		fmt.Println("=======================================================================================================")

	}
}

func LogJson(x ...interface{}) {
	_, fn, line, _ := runtime.Caller(1)
	for _, data := range x {
		d, _ := json.Marshal(data)
		logrus.Info(fmt.Sprintf("[error] %s:%d", fn, line), string(d))
	}
}
