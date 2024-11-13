package dlog

import (
	"encoding/json"
	"fmt"
	"runtime"

	"github.com/sirupsen/logrus"
)

func LogJson(x ...interface{}) {
	var fields []interface{}

	_, fn, line, _ := runtime.Caller(1)
	fields = append(fields, fmt.Sprintf("log on: %s:%d |", fn, line))
	for _, data := range x {
		d, _ := json.Marshal(data)
		fields = append(fields, fmt.Sprintf(" %s |", string(d)))
	}
	logrus.Info(fields...)
}
