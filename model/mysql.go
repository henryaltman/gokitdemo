package model

import (
	"fmt"
	"gokitdemo/core"
	"log"
	"strings"
	"time"
)
//`id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
//`a` bigint NOT NULL DEFAULT '0' COMMENT 'userId',
//`b` int DEFAULT '0' COMMENT '主题id',
//`c` varchar(256) DEFAULT NULL COMMENT '主题列表',
//`d` int DEFAULT '0' COMMENT '头像挂饰',
//`status` int unsigned NOT NULL DEFAULT '0' COMMENT '状态(未使用=0, 已使用=1)',

type TestTable struct {
	id int
	a int
	b int
	c string
	d int
	status int
}
type Field struct {
	fieldName string
	fieldDesc string
	dataType  string
	isNull    string
	length    int
}

type TestTableName struct {
	id int
	a int
	b int
	c string
	d int
	status int
	create_time time.Time
	update_time time.Time
}



func FieldInfo(dbName,tableName string) map[string]Field{
	sqlStr := `SELECT COLUMN_NAME fName,column_comment fDesc,DATA_TYPE dataType,
						IS_NULLABLE isNull,IFNULL(CHARACTER_MAXIMUM_LENGTH,0) sLength
			FROM information_schema.columns 
			WHERE table_schema = ? AND table_name = ?`


	 result := map[string]Field{}
	db := core.Instance().Db
	rows, err := db.Query(sqlStr,dbName,tableName)
	checkErr(err)

	for rows.Next() {
		var f Field
		err = rows.Scan(&f.fieldName, &f.fieldDesc, &f.dataType, &f.isNull, &f.length)
		checkErr(err)
		result[f.fieldName] = f
	}
	return result
}

//getone func is get one row
func GetOne(dbNAME ,tableNAME string ,field string,query map[string]interface{}) (interface{},error) {

	if dbNAME == "" ||  tableNAME == "" {

	}

	queryField := strings.Split(field,",")

	fields := FieldInfo(dbNAME,tableNAME)
	needQueryFields := []string{}
	for _,field := range queryField {
		for _, f := range fields {
            if field ==  f.fieldName  {
				//fmt.Println("v.fieldName ,v.dataType ",f.fieldName ,f.dataType)
				needQueryFields = append(needQueryFields,field)
			}
		}
	}

	queryString  := []string{}
	queryStr := ""
	if len(query) > 0 {
		for k,v := range query{
			queryString = append(queryString,fmt.Sprintf("%s=%v",k,v))
		}
		fmt.Println("queryString",queryString)
		if len(queryString) > 0 {
			queryStr = strings.Join(queryString," and ")
		}
	}

	needQueryFieldsStr := strings.Join(needQueryFields,",")
	if len(needQueryFieldsStr) == 0 {
		needQueryFieldsStr = "*"
	}
	sqlStr := ""
	if queryStr != "" {
		sqlStr = fmt.Sprintf("select id,a,b,c,d,status from %s where %s" , tableNAME,queryStr)
	}else{
		sqlStr =  fmt.Sprintf("select * from %s limit 100" , tableNAME)
	}

	fmt.Println("sql",sqlStr)

    rows ,err :=  core.Instance().Db.Query(sqlStr)
    defer rows.Close()
    if err != nil {

	}

	result := []TestTableName{}
	for rows != nil && rows.Next() {
		tt := TestTableName{}
		if err := rows.Scan(&tt.id,&tt.a,&tt.b,&tt.c,&tt.d,&tt.status); err != nil {
			log.Fatal(err)
		}
		result = append(result, tt)
	}
	for _,v := range result {
		fmt.Println(fmt.Sprintf("v=%v",v))
	}
   return result,err
}

func  checkErr(err error)  {

}