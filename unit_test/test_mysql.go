package main

import (
	"gokitdemo/model"
)

func main() {
	//index := 0
	//for {
	//
	//	if index > 9 {
	//		break
	//	}
	//	core.Instance()
	//	index++
	//}

	//model.GetInstance()
	queryMap := map[string]interface{}{"a":1}
	 model.GetOne("gokitdemo", "testtable","a,b,c",queryMap)

	//fmt.Println("fieldInfo", fieldInfo)
	//for _, ff := range fieldInfo {
    //     fmt.Println("ff",ff)
	//}
}