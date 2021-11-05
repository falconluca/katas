package main

import (
	"fmt"
	"gorm.io/gorm/logger"
	"log"
	"strconv"
	"time"

	"github.com/fatih/color"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type House struct {
	// 包括哪些公共字段?
	gorm.Model
	Section string `gorm:"column:section"`
	HouseId uint   `gorm:"column:house_id"`
}

var (
	db *gorm.DB
)

func main() {
	db = initDb()

	// 更新表结构
	//db.AutoMigrate(&House{})

	//logGreenStr("查询房源列表")
	//houses := getHousesBySection("万达")
	//for i, house := range *houses {
	//	log.Printf("第%v套房源: %+v\n", i+1, house)
	//}
	//
	//logGreenStr("创建房源")
	//newHouse := &House{HouseId: 7, Section: "万科"}
	//addNewHouse(newHouse)
	//log.Printf("新创建的房源: %+v\n", newHouse)
	//
	//logGreenStr("删除房源(软删除: 因为有deletedAt字段)")
	//deleteHouseByHouseId(7)
	//
	//logGreenStr("获取所有房源列表")
	//allHouses := findAllHouses()
	//for i, house := range *allHouses {
	//	log.Printf("第%v套房源: %+v\n", i+1, house)
	//}

	//logGreenStr("查询房源")
	//house := getHouseByHouseId(12)
	//log.Printf("查询房源: %+v\n", house)

	//logGreenStr("更新房源")
	//house.Section = "万科"
	//updateHouse(house)

	logGreenStr("分页获取房源")
	housePage := getHousePage(0, 2)
	for i, h := range *housePage {
		log.Printf("第%v套房源: %+v\n", i+1, h)
	}

	// 原生SQL
	//var h House
	//db.
	//	Debug(). // 打印SQL
	//	Raw("SELECT h1.section, h2.house_id, h2.updated_at FROM houses h1 "+
	//		"left join houses h2 on h1.house_id = h2.house_id  WHERE h1.house_id = ?", 13).Scan(&h)
	//log.Printf("查询房源: %+v\n", h)

	// TODO gorm钩子
}

func getHousePage(page int, size int) *[]House {
	houses := make([]House, 0)
	// var houses []*House = make([]*House, 0)
	if err := db.
		Where("house_id is not null").
		Offset(page).
		Limit(size).
		Order("house_id desc").
		Find(&houses).
		Error; err != nil {
		log.Fatalf("获取房源分页数据失败! err: %v", err)
	}
	return &houses
}

func updateHouse(house *House) {
	if err := db.
		Save(house).
		Error; err != nil {
		log.Fatalf("更新房源失败! err: %v", err)
	}
}

func deleteHouseByHouseId(houseId uint) {
	houseIdStr := strconv.Itoa(int(houseId))
	// TODO 这个SQL好残疾 :) 只能这样写么?
	if err := db.
		Where("house_id = ?", houseIdStr).
		//Unscoped(). // 永久删除
		Delete(&House{}).
		Error; err != nil {
		log.Fatalf("删除房源失败! err: %v", err)
	}
}

func addNewHouse(house *House) {
	if err := db.Create(house).Error; err != nil {
		log.Fatalf("创建房源失败! err: %v", err)
	}
}

func logGreenStr(str string) {
	log.Println(color.GreenString(str))
}

func initDb() *gorm.DB {
	dsn := fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s`,
		"root", "root", "127.0.0.1:3306", "test", true, "Local")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 默认是Warn, 即只打印慢查询
	})
	if err != nil {
		panic("failed to connect db")
	}

	// 配置连接池
	// FIXME 可视化连接池的各种概念 mysql redis
	datasource, err := db.DB()
	if err != nil {
		panic("create connections pool failed")
	}
	datasource.SetMaxIdleConns(100) // 最大空闲连接数
	datasource.SetMaxOpenConns(100) // 最大连接数
	datasource.SetConnMaxLifetime(time.Second * 10)
	return db
}

func getHouseByHouseId(houseId uint) *House {
	house := House{}
	log.Printf("&House{}: %v", &house) // &{{0 0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC {0001-01-01 00:00:00 +0000 UTC false}}  0}

	p := &house
	log.Printf("&&House{}: %v", &p) // 0xc0000a2018
	// 这是啥的地址呀??? 没整明白 TODO

	houseIdStr := strconv.Itoa(int(houseId))
	if err := db.
		Where("house_id = ?", houseIdStr).
		First(&house).
		Error; err != nil {
		log.Fatalf("get house error: %v", err)
	}
	return &house
}

func getHousesBySection(section string) *[]House { // 返回数组的地址...
	var houses []House
	if err := db.
		//Select("house_id"). // 默认情况下, GORM 在查询时会选择所有的字段, 可以使用 Select 来指定您想要的字段
		// 或者使用智能选择字段有点像Mybatis的resultType
		Where("section = ?", section).
		Find(&houses).
		Error; // 查看finisher_api.go
	err != nil {
		log.Fatalf("get house error: %v", err)
	}
	return &houses
}

func findAllHouses() *[]House {
	var houses []House
	if err := db.
		Find(&houses).
		Error; err != nil {
		log.Fatalf("find all houses error: %v", err)
	}
	return &houses
}
