package models

type Tag struct {
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"` // c.JSON时会自动转换格式
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
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
