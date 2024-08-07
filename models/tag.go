package models

import "log"

type Tag struct {
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func GetTags(pageNum int, pageSize int, maps any) (tags []Tag) {
	if db == nil {
		log.Fatalf("null pointer")
	}
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)

	return
}

func GetTagTotal(maps any) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)

	return
}
