package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)
// album represents data about a record album
// album 的数据结构
type album struct {
	ID		string	`json:"id"`
	Title	string	`json:"title"`
	Artist	string	`json:"artist"`
	Price	float64	`json:"price"`
}

// album slice to seed record album data.
// 示例数据 albums
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// main
func main () {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumById)
	
	router.Run("localhost:8082")
}

// getAlbums responds with the list of all albums as JSON
// getAlbums 方法以json格式返回所有示例数据（列表）
func getAlbums (c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums) // 响应码 200
}

// postAlbums adds an album from JSON received in the request body.
// 通过客户端的post请求体中json格式的album进行添加album
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to newAlbum.
	// 将客户端发来的json格式的album数据绑定到newAlbum对象
	if err := c.BindJSON(&newAlbum); err != nil {
		return
		// return err  该写法将无法正常启动程序
	}

	// add the new album to the slice.
	// 将newAlbum对象添加到之前的albums中
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum) // 响应码 201
	// c.IndentedJSON(http.StatusOK, "good job!") // 测试代码
}

// getAlbumByID locates the album whose ID value matches the id parameter sent by the client, the returns that album as a response,
// 依据客户端提供的album的id 返回指定的album
func getAlbumById(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for an album whose ID value matches the parameter.
	// 遍历已有列表中的album，对比albums的id，返回id相同的album
	for _, a := range albums{
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message":"the album that you provide id of not found"})
}