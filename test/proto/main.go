package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
	pb "proto/lucapb"
)

func main() {
	p := pb.Person{Name: "Luca", Id: 12, Email: "luca@gmail.com", Phones: []*pb.Person_PhoneNumber{
		{
			Number: "123",
			Type:   pb.Person_MOBILE,
		},
		{
			Number: "123444",
			Type:   pb.Person_WORK,
		},
	}}
	fmt.Println(p)
	fmt.Println(&p)

	// 写数据
	b := &pb.AddressBook{People: []*pb.Person{
		&p,
	}}
	out, err := proto.Marshal(b)
	if err != nil {
		log.Fatalln("序列化AddressBook失败, err: ", err)
	}
	if err = ioutil.WriteFile("./lucapb/address.json", out, 0644); err != nil { // TODO ioutil.WriteFile ioutil.ReadFile
		log.Fatalln("将内存的数据写入文件失败, err: ", err)
	}

	// 读数据
	in, err := ioutil.ReadFile("./lucapb/address.json")
	if err != nil {
		log.Fatalln("将磁盘文件的数据读入内存失败, err: ", err)
	}
	book := &pb.AddressBook{}
	if err = proto.Unmarshal(in, book); err != nil { // TODO ioutil.WriteFile ioutil.ReadFile
		log.Fatalln("将数据写入文件失败, err: ", err)
	}
	fmt.Printf("读取文件里的数据: ")
	fmt.Println(book)
}
