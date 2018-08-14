/*
Copyright 2018 JD-Tiger
created by lvxin  at 18-8-13 下午12:03
*/
package main

import (
	example "github.com/lvxin1986/protobuf-golang/proto"
	"log"
	"github.com/golang/protobuf/proto"
	"fmt"
)


func main() {
	// 为 AllPerson 填充数据
	//使用protobuf的封装类型定义
	p1 := example.Person{
		Id:*proto.Int32(1),
		Name:*proto.String("lvxin"),
	}
	//使用golang的原始类型定义
	p2 := example.Person{
		Id:2,
		Name:"gopher",
	}

	all_p := example.AllPerson{
		Per:[]*example.Person{&p1, &p2},
	}

	// 对数据进行序列化
	data, err := proto.Marshal(&all_p)
	if err != nil {
		log.Fatalln("Mashal data error:", err)
	}

	// 对已经序列化的数据进行反序列化
	var target example.AllPerson
	err = proto.Unmarshal(data, &target)
	if err != nil{
		log.Fatalln("UnMashal data error:", err)
	}
	for k,v := range target.Per {
		fmt.Println("person[",k,"]:",v.Name)
	}
}
