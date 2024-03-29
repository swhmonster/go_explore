package demotest

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/xuri/excelize/v2"
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

var logrusTestInstance = logrus.New()
var logFile = "/tmp/go_project.log"

func init() {
	// logrus 设置日志时间output
	f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	logrusTestInstance.Out = f
	logrusTestInstance.SetLevel(logrus.DebugLevel)
	logrusTestInstance.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
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
	logrusTestInstance.Debug(stu)
	return stu, err
}

// basetest
func TestMap() (map[string]string, error) {
	m := make(map[string]string)
	m["key1"] = "value1"
	m["key2"] = "value2"
	logrusTestInstance.Debug(m)
	return m, nil
}

func TestSlice() ([]string, error) {
	s := make([]string, 1, 5)
	s = append(s, "value1")
	s = append(s, "value2", "value3")
	logrusTestInstance.Debug(s)
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
	logrusTestInstance.Debug("Generic Sums, type parameters inferred:" + strconv.Itoa(int(SumIntsOrFloats(ints))) + " and " + strconv.FormatFloat(SumIntsOrFloats(floats), 'f', 2, 64))
	return "Generic Sums, type parameters inferred:" + strconv.Itoa(int(SumIntsOrFloats(ints))) + " and " + strconv.FormatFloat(SumIntsOrFloats(floats), 'f', 2, 64)
}

// SumIntsOrFloats sums the values of map m. It supports both floats and integers
// as map values.
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	logrusTestInstance.Debug(s)
	return s
}

// type assert demo
// Container is a generic container, accepting anything.
type Container []interface{}

// Put adds an element to the container.
func (c *Container) Put(elem interface{}) {
	*c = append(*c, elem)
}

// Get gets an element from the container.
func (c *Container) Get() interface{} {
	elem := (*c)[0]
	*c = (*c)[1:]
	return elem
}

/*
// common use
intContainer := &Container{}
intContainer.Put(7)
intContainer.Put(42)

// type assert use
// assert that the actual type is int
elem, ok := intContainer.Get().(int)

	if !ok {
	    fmt.Println("Unable to read an int from intContainer")
	}

fmt.Printf("assertExample: %d (%T)\n", elem, elem)
*/
func TestExcelize() {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// 创建一个工作表
	index, err := f.NewSheet("Sheet2")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 设置单元格的值
	f.SetCellValue("Sheet2", "A2", "Hello world.")
	f.SetCellValue("Sheet1", "B2", 100)
	// 设置工作簿的默认工作表
	f.SetActiveSheet(index)
	// 根据指定路径保存文件
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
}
