package models

type Tag struct {
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"` // c.JSON时会自动转换格式
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

/*
// BeforeCreate gorm创建数据的回调方法
func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	// tag.CreatedOn, CreatedOn 或 created_on 都行, 参看注释
	scope.SetColumn(tag.CreatedOn, time.Now().Unix())
	return nil
}

// BeforeUpdate gorm更新数据的回调方法
func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn(tag.ModifiedOn, time.Now().Unix())
	return nil
}
*/

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

func ExistTagById(id int) bool {
	var tag Tag
	db.Select("id").Where("id = ?", id).First(&tag)
	return tag.Id > 0
}

func DeleteTag(id int) bool {
	db.Where("id = ?", id).Delete(&Tag{})
	return true
}

func EditTag(id int, data interface{}) bool {
	db.Model(&Tag{}).Where("id = ?", id).Updates(data)
	return true
}

func CleanDeletedTags() bool {
	db.Unscoped().
		Where("deleted_on != ?", 0).
		Delete(&Tag{})
	return true
}
