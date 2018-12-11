package models

import "fmt"

type Tag struct {
	Model

	Name string `json:"name"`
	CreatedBy string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State int `json:"state"`
}

func GetTags(pageNum int, pageSize int, maps interface{}, params ...interface{})(tags[]Tag){
	Db.Where(maps, params).Offset(pageNum - 1).Limit(pageSize).Find(&tags)
	fmt.Println("jason09990:", &tags)
	return
}

