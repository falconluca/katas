package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Tag struct {
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"` // c.JSON时会自动转换格式
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

// BeforeCreate gorm创建数据的回调方法
func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	// TODO CreatedOn 或 created_on 都行
	scope.SetColumn("created_on", time.Now().Unix())
	return nil
}

// BeforeUpdate gorm更新数据的回调方法
func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("modified_on", time.Now().Unix())
	return nil
}

func GetTags(page int, size int, maps interface{}) (tags []Tag) {
	// TODO 因为在同一个models包下, 因此db *gorm.DB是可以直接使用的
	db.Where(maps).
		Offset(page).
		Limit(size).
		Find(&tags)
	return
}

func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).
		Where(maps).
		Count(&count)
	return
}

// ExistTagByName 是否已经存在这个名称的标签
func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name = ?", name).First(&tag)
	return tag.Id > 0
}

// AddTag 创建标签
func AddTag(name string, state int, createdBy string) bool {
	db.Create(&Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	})
	return true
}
