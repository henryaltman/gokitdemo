package core

import (
	"fmt"
	"reflect"
	"time"
)

func CallReflect(any interface{}, name string, args ...interface{}) []reflect.Value {
	startTs := time.Now().UnixNano()
	//defer func(sTS int64) {
	//	endTs := time.Now().UnixNano()
	//	fmt.Println("endTs - startTs",endTs - startTs)
	//}(startTs)
	inputs := make([]reflect.Value, len(args))
	for i, _ := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}
	endTs1 := time.Now().UnixNano()
	fmt.Sprintf("name=%s,endTs - startTs", name, endTs1-startTs)
	fmt.Println(fmt.Sprintf("name=%s,endTs - startTs=%d", name, endTs1-startTs))
	if v := reflect.ValueOf(any).MethodByName(name); v.String() == "<invalid Value>" {
		return nil
	} else {
		ret := v.Call(inputs)
		return ret
	}
}
