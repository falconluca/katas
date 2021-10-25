package models

type Article struct {
	Model // 类似继承

	TagId      int    `json:"tag_id" gorm:"index"` // 声明这是个索引字段. 使用自动迁移功能则会有所影响，不使用则无影响
	Tag        Tag    `json:"tag"`                 // TODO 嵌套struct 关联查询
	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

/*
func (a *Article) BeforeCreate(s *gorm.Scope) error {
	s.SetColumn("created_on", time.Now().Unix()) // 当前时间戳
	return nil
}

func (a *Article) BeforeUpdate(s *gorm.Scope) error {
	s.SetColumn("modified_on", time.Now().Unix())
	return nil
}
*/

func ExistArticleById(id int) bool {
	var article Article
	db.Select("id").Where("id = ?", id).First(&article)
	return article.Id > 0
}

func GetArticleTotal(maps interface{}) (count int) {
	db.Model(&Article{}).Where(maps).Count(&count)
	return
}

func GetArticles(page int, size int, maps interface{}) (articles []Article) {
	db.Preloads("Tag"). // 关联查询
				Where(maps).
				Offset(page).
				Limit(size).
				Find(&articles)
	return
}

func GetArticle(id int) (article Article) {
	db.Where("id = ?", id).First(&article)
	db.Model(&article).Related(&article.Tag) // TODO 关联查询
	return
}

func EditArticle(id int, data map[string]interface{}) bool {
	db.Model(&Article{}).Where("id = ?", id).Updates(data)
	return true
}

func AddArticle(data map[string]interface{}) bool {
	db.Create(&Article{
		// TODO 使用interface做类型转换 类型断言
		// 是否可以把interface{}理解成Java里的Object对象
		TagId:     data["tag_id"].(int),
		Title:     data["title"].(string),
		Desc:      data["desc"].(string),
		Content:   data["content"].(string),
		CreatedBy: data["created_by"].(string),
		State:     data["state"].(int),
	})
	return true
}

func DeleteArticle(id int) bool {
	db.Where("id = ?", id).Delete(Article{}) // TODO 值or指针
	return true
}

func CleanDeletedArticles() bool {
	db.Unscoped().
		Where("deleted_on != ?", 0).
		Delete(&Article{})
	return true
}
