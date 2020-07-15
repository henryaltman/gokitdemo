package core

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func CallReflect(any interface{}, name string, args ...interface{}) []reflect.Value {
	inputs := make([]reflect.Value, len(args))
	for i, _ := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}
	fmt.Println("inputs", inputs)

	if v := reflect.ValueOf(any).MethodByName(name); v.String() == "<invalid Value>" {
		fmt.Println("<invalid Value>")
		return nil
	} else {
		ret := v.Call(inputs)
		retByte, _ := json.Marshal(ret[0].Interface())
		fmt.Println("ret", string(retByte))
		return ret
	}
}
