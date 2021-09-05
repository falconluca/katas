package mysql

import (
	"log"
)

func SaveAlbum(title string, artist string, price float32) {
	stmt, err := db.Prepare("INSERT INTO album SET title=?, artist=?, price=?")
	if err != nil {
		log.Fatal("编译新增唱片SQL失败 ", err)
	}

	res, err := stmt.Exec(title, artist, price)
	if err != nil {
		log.Fatal("新增唱片失败 ", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("获取新增唱片Id失败 ", err)
	}
	log.Printf("新增的唱片Id为: %v", id)
}

func UpdateAlbum(title string, id int) {
	stmt, err := db.Prepare("update album set title=? where id=?")
	if err != nil {
		log.Fatal("编译更新唱片SQL失败 ", err)
	}

	res, err := stmt.Exec(title, id)
	if err != nil {
		log.Fatal("更新唱片失败 ", err)
	}

	affectRows, err := res.RowsAffected()
	if err != nil {
		log.Fatal("更新唱片失败 ", err)
	}
	log.Printf("更新了%v张唱片, 唱片id:%v", affectRows, id)
}

func DeleteAlbum(id int) {
	stmt, err := db.Prepare("delete from album where id=?")
	if err != nil {
		log.Fatal("编译删除唱片SQL失败 ", err)
	}

	res, err := stmt.Exec(id)
	if err != nil {
		log.Fatal("删除唱片失败 ", err)
	}

	affectRows, err := res.RowsAffected()
	if err != nil {
		log.Fatal("删除唱片失败 ", err)
	}
	log.Printf("删除了%v张唱片, 唱片id: %v", affectRows, id)
}

func ListAllAlbums() ([]Album, error) {
	var result []Album

	rows, err := db.Query("SELECT * FROM album")
	if err != nil {
		log.Fatal("查询唱片列表失败  ", err)
	}
	defer rows.Close()

	for rows.Next() {
		var album Album

		if err := rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
			log.Fatal("获取唱片失败 ", err)
		}
		result = append(result, album)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return result, nil
}
