package main

import (
    "net/http"
    "sync"
		"strconv"
    "github.com/gin-gonic/gin"
)

// メモリ上に保存するデータ構造
type Memo struct {
    ID   int    `json:"id"`
    Text string `json:"text"`
}

// メモリ上に保存するデータ
var (
    memos []Memo
    mu    sync.Mutex
    idSeq int
)

func getMemosHandler(c *gin.Context) {
	c.JSON(http.StatusOK, memos)
}

func createMemoHandler(c *gin.Context) {
	var memo Memo
	if err := c.ShouldBindJSON(&memo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mu.Lock()
	idSeq++
	memo.ID = idSeq
	memos = append(memos, memo)
	mu.Unlock()

	c.JSON(http.StatusCreated, memo)
}

func updateMemoHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
			return
	}

	var newMemo Memo
	if err := c.ShouldBindJSON(&newMemo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}

	mu.Lock()
	defer mu.Unlock()
	for i, memo := range memos {
			if memo.ID == id {
					memos[i].Text = newMemo.Text // Update the memo's text
					c.JSON(http.StatusOK, memos[i])
					return
			}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Memo not found"})
}

func deleteMemoHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	mu.Lock()
	defer mu.Unlock()
	for i, memo := range memos {
		if memo.ID == id {
			memos = append(memos[:i], memos[i+1:]...)
			c.Status(http.StatusNoContent)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Memo not found"})
}

func main() {
	router := gin.Default()

	router.GET("/memos", getMemosHandler)
	router.POST("/memos", createMemoHandler)
	router.DELETE("/memos/:id", deleteMemoHandler)
	router.PUT("/memos/:id", updateMemoHandler)

	router.Run(":8080")
}
