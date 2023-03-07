/*
这是包注释go

godoc 既是一个程序，又是一个 Web 服务器，它对 Go 的源码进行处理，并提取包中的文档内容。

出现在顶级声明之前，且与该声明之间没有空行的注释，将与该声明一起被提取出来，作为该条目的说明文档。

这些注释的类型和风格决定了 godoc生成的文档质量。
*/
package main

import (
	"demotest"
	"fmt"
	ginpprof "github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
	"runtime"
	"runtime/pprof"
)

var db = make(map[string]string)

// setupRouter 用于路由设置
func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "I have these api: /ping, /user/:name, /serialize, /testmap, /testslice, /testgenerics, /testExcelize")
	})

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	r.GET("/serialize", func(c *gin.Context) {
		value, err := demotest.Serialize()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"err": err})
		} else {
			c.JSON(http.StatusOK, gin.H{"value": value})
		}
	})

	r.GET("/testmap", func(c *gin.Context) {
		value, err := demotest.TestMap()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"err": err})
		} else {
			c.JSON(http.StatusOK, gin.H{"value": value})
		}
	})

	r.GET("/testslice", func(c *gin.Context) {
		value, err := demotest.TestSlice()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"err": err})
		} else {
			c.JSON(http.StatusOK, gin.H{"value": value})
		}
	})

	r.GET("/testgenerics", func(c *gin.Context) {
		value := demotest.TestGenerics()
		c.JSON(http.StatusOK, gin.H{"value": value})
	})

	r.GET("/testExcelize", func(c *gin.Context) {
		demotest.TestExcelize()
		c.JSON(http.StatusOK, gin.H{"value": "void"})
	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	/* example curl for /admin with basicauth header
	   Zm9vOmJhcg== is base64("foo:bar")

		curl -X POST \
	  	http://localhost:8080/admin \
	  	-H 'authorization: Basic Zm9vOmJhcg==' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'
	*/
	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return r
}

// linux tmp 系统重启时，默认清理"tmp目录下"10天未用的文件

var logFile = "/tmp/mGo.log"
var cpuprofile = "/tmp/cpu.prof"
var memprofile = "/tmp/mem.prof"

// initGoToolPprofConfig 初始化pprof配置
func initGoToolPprofConfig() {
	f, err := os.Create(cpuprofile)
	if err != nil {
		logrus.Fatal("could not create CPU profile: ", err)
	}
	defer f.Close() // error handling omitted for example
	if err := pprof.StartCPUProfile(f); err != nil {
		logrus.Fatal("could nolst start CPU profile: ", err)
	}
	defer pprof.StopCPUProfile()

	f2, err2 := os.Create(memprofile)
	if err2 != nil {
		logrus.Fatal("could not create memory profile: ", err2)
	}
	defer f2.Close() // error handling omitted for example
	runtime.GC()     // get up-to-date statistics
	if err2 := pprof.WriteHeapProfile(f2); err2 != nil {
		logrus.Fatal("could not write memory profile: ", err2)
	}
}

//  接口 strong check

// Shape 形状接口
type Shape interface {
	Sides() int
	Area() int
}

type Square struct {
	len int
}

func (s *Square) Area() int {
	//TODO implement me
	panic("implement me")
}

func (s *Square) Sides() int {
	return 4
}

// init main方法执行前初始化
func init() {
	// logrus 设置日志时间output
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
}

// main main包的main方法
func main() {
	// 推荐将错误消息发送值UNIX机器上的日志服务，防止发用不必要的数据填写日志文件
	// 日志配置 0644：UNIX文件权限
	f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	iLog := log.New(f, "customLogLineNumber", log.LstdFlags)
	// 第二个参数为输出行号
	iLog.SetFlags(log.LstdFlags | log.Lshortfile)

	iLog.Println("server starting...")
	iLog.Println("server started!")

	// logrus 日志库
	// 若需输出到文件，仍采用上面的方式输出到文件
	logrus.WithFields(logrus.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")

	// 校验是否完全实现接口方法，在 Go 语言编程圈里有一个比较标准的作法如下
	var _ Shape = (*Square)(nil)
	s := Square{len: 5}
	logrus.WithFields(logrus.Fields{"sides": s.len}).Info("Square Sides")

	// go tool pprof 信息采集至文件
	// cmd:go tool pprof cpu.prof
	/*initGoToolPprofConfig()*/

	r := setupRouter()
	// ginpprof
	ginpprof.Register(r)
	// Listen and Server in 0.0.0.0:8080
	r.Run(":18080")
}
