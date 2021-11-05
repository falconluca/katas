# ORM 相关概念

## 1. 级联 

级联删除、更新、创建

## 2. 一对一、一对多、多对多

TODO

## 3. 外键、引用

```go
type User struct {
  gorm.Model
  Name      string
  CompanyID int // 外键
  Company   Company `gorm:"foreignKey:CompanyID; references:ID"`
}

type Company struct {
  ID   int // 引用
  Name string
}
```