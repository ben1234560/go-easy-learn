package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// main作为运行的主函数，启动时，run整个book目录
func main() {
	// 首先运行初始化函数，这样才有操作数据库的权限
	err := initDB()
	if err != nil {
		// panic是一个报错函数，会在终端输出并终端，return则只会返回值
		panic(err)
	}
	// 初始化gin，这样才能用gin
	engine := gin.Default()
	// 加载模板，绝对路径是从src里算起，我的路径是src/go/book/templates
	engine.LoadHTMLGlob("./book/templates/*")
	// GET请求访问/book/list则触发回调函数，我们把回调函数写下面
	engine.GET("/book/list", bookListHandler)
	engine.GET("/book/new",newBookHandler)
	engine.POST("/book/new",createBookHandler)
	engine.GET("/book/delete",deleteBookHandler)
	// 运行，监听8000端口（监听哪个端口都可以，不冲突就行）
	err = engine.Run(":8000")
	if err != nil{
		return
	}
}

// 查询的回调函数
func bookListHandler(c *gin.Context) {
	// 调用数据库的查询函数（在db.go文件）
	bookList, err := queryAllBook()
	// 处理异常情况
	if err != nil {
		// 与前端的交互要用json传输
		c.JSON(401, gin.H{"code": 1, "msg": err})
		return
	}
	// 如果没异常
	// StatusOK，返回html模板，并返回bookList数据给data（data是在html页面写定了是data）
	c.HTML(http.StatusOK, "book_list.html", gin.H{"code": 0, "data": bookList})
}

// 添加的回调函数（1）
func newBookHandler(c *gin.Context)  {
	// 点击添加按钮是跳转到其它页面，所以我们写的跳转, nil表示不需要传值
	c.HTML(http.StatusOK, "new_book.html",nil)
}

// 添加的回调函数（2）
func createBookHandler(c *gin.Context)  {
	// 接收表单传过来的数据，获取到书名和价格
	titleVal := c.PostForm("title")
	priceVal := c.PostForm("price")
	// 传过来的price是字符串类型，我们用的是int，需要传格式
	price, err := strconv.ParseInt(priceVal, 10, 64)
	if err != nil{
		c.JSON(401, gin.H{"code": 1, "msg": err})
		return
	}
	// 调用数据库的添加函数
	err = insertBook(titleVal, price)
	if err != nil{
		c.JSON(401, gin.H{"code": 1, "msg": err})
		return
	}
	// 添加完成后转回主页，重定向
	c.Redirect(http.StatusMovedPermanently, "/book/list")
}

// 删除的回调函数
func deleteBookHandler(c *gin.Context)  {
	// 接收传过来的id，和上面的表单不一样，它是url形式的
	idStr := c.Query("id")
	// 传传过来的id是字符串类型，我们用的是int，需要传格式
	idVal, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil{
		c.JSON(401, gin.H{"code": 1, "msg": err})
		return
	}
	// 调用数据库的删除函数
	err = deleteBook(idVal)
	if err != nil{
		c.JSON(401, gin.H{"code": 1, "msg": err})
		return
	}
	// 删除完跳转回主页
	c.Redirect(http.StatusMovedPermanently, "/book/list")
}