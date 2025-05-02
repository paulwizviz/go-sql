package structtag

import (
	"fmt"
	"reflect"
	"strings"
)

type LivingThing interface {
	Human | Animal
}

type StructTag struct {
	FieldName string
	Tag       string
}

func DBTags[T LivingThing](tagName string, typ *T) []StructTag {
	ev := reflect.Indirect(reflect.ValueOf(typ))
	tags := []StructTag{}
	for i := 0; i < ev.Type().NumField(); i++ {
		tag := StructTag{}
		tag.FieldName = ev.Type().Field(i).Name
		t := ev.Type().Field(i).Tag
		tElems := strings.Split(string(t), " ")
		for _, tElem := range tElems {
			if strings.Contains(tElem, tagName) {
				sElems := strings.Split(tElem, ":")
				tag.Tag = sElems[1][1 : len(sElems[1])-1]
			}
		}
		tags = append(tags, tag)
	}
	return tags
}

type Human struct {
	ID        int    `json:"id" sqlite:"id,INTEGER,PRIMARY_KEY"`
	FirstName string `json:"first_name" sqlite:"first_name,TEXT"`
	Surname   string `json:"surname" sqlite:"surname,TEXT"`
}

type Animal struct {
	ID      int    `json:"id" sqlite:"id,INTEGER,PRIMARY_KEY"`
	Species string `json:"species" sqlite:"species,TEXT"`
	Name    string `json:"name" sqlite:"name,TEXT"`
}

func SQLiteCreateTblStmtStr[T LivingThing](v *T) string {
	t := reflect.TypeOf(v)
	tblName := strings.Split(fmt.Sprintf("%v", t), ".")[1][0:]
	tags := DBTags("sqlite", v)

	stmt := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (", strings.ToLower(tblName))
	for _, tag := range tags {
		tagElemes := strings.Split(tag.Tag, ",")
		switch len(tagElemes) {
		case 2:
			stmt = fmt.Sprintf("%s %s %s,", stmt, tagElemes[0], tagElemes[1])
		case 3:
			pkstmt := tagElemes[2]
			pkElems := strings.Split(pkstmt, "_")
			stmt = fmt.Sprintf("%s %s %s %s %s,", stmt, tagElemes[0], tagElemes[1], pkElems[0], pkElems[1])
		}
	}
	stmt = stmt[0 : len(stmt)-1] // remove last comma
	stmt = fmt.Sprintf(`%s )`, stmt)
	return stmt
}

func SQLiteInsertStmtStr[T LivingThing](v *T) string {
	t := reflect.TypeOf(v)
	tblName := strings.Split(fmt.Sprintf("%v", t), ".")[1][0:]
	tags := DBTags("sqlite", v)

	stmt := fmt.Sprintf("INSERT INTO %s (", strings.ToLower(tblName))
	for _, tag := range tags {
		t1 := strings.Split(tag.Tag, ",")
		stmt = fmt.Sprintf("%s %s,", stmt, t1[0])
	}

	stmt = stmt[0 : len(stmt)-1]
	stmt = fmt.Sprintf("%s) VALUES (", stmt)
	for range tags {
		stmt = fmt.Sprintf("%s ?,", stmt)
	}
	stmt = stmt[0 : len(stmt)-1]
	stmt = fmt.Sprintf("%s )", stmt)
	return stmt
}
