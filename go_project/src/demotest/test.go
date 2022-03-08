package demotest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Student struct {
	Name   string   `json:"name"`   // 姓名
	Age    int      `json:"age"`    // 年龄
	Gender string   `json:"gender"` // 性别
	Score  float64  `json:"score"`  // 分数
	Course []string `json:"course"` // 课程
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func fileHandler(w http.ResponseWriter, r *http.Request) {
	filePath := "/Users/sunwenhao/Downloads/temp/1.txt"
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(data)
	fmt.Fprintf(w, filePath)
}

func serialize() {
	stu := Student{
		"张三",
		20,
		"男",
		78.6,
		[]string{"语文", "数学", "音乐"},
	}

	// 序列化
	data, err := json.Marshal(&stu)
	if err != nil {
		fmt.Println("序列化错误", err)
	} else {
		fmt.Println("序列化结果：" + string(data))
	}

	// 反序列化
	var stu2 Student
	err2 := json.Unmarshal(data, &stu2)
	if err2 != nil {
		fmt.Println("反序列化失败", err2)
	} else {
		fmt.Printf("反序列化结果：")
		fmt.Println(stu2)
	}
}
