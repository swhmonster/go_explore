package demotest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
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

func Serialize() (Student, error) {
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
	return stu, err
}

// basetest
func TestMap() (map[string]string, error) {
	m := make(map[string]string)
	m["key1"] = "value1"
	m["key2"] = "value2"
	return m, nil
}

func TestSlice() ([]string, error) {
	s := make([]string, 1, 5)
	s = append(s, "value1")
	s = append(s, "value2", "value3")
	return s, nil
}

func TestGenerics() string {
	// Initialize a map for the integer values
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	return "Generic Sums, type parameters inferred:" + strconv.Itoa(int(SumIntsOrFloats(ints))) + " and " + strconv.FormatFloat(SumIntsOrFloats(floats), 'f', 2, 64)
}

// SumIntsOrFloats sums the values of map m. It supports both floats and integers
// as map values.
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
