package main

import (
	"net/http"
	"task1/sumutil"

	"github.com/gin-gonic/gin"
)

func main() {
	
	r := gin.Default()
	// Load template từ file HTML
	r.LoadHTMLFiles("templates/index.html")
	// Route để hiển thị form
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"result": ""})
	})
	// Route xử lý form submit
	r.POST("/calculate", func(c *gin.Context) {
		result := make([]string, 0)
		// Lấy giá trị từ form
		n1 := c.PostForm("num1")
		n2 := c.PostForm("num2")
		// Tính toán kết quả
		sumutil.SumlogHaiso(n1, n2, len(n1)-1, len(n2)-1, func(n1, n2 string) int {
			if len(n1) > len(n2) {
				return len(n1) - 1
			}
			return len(n2) - 1
		}(n1, n2), 0, &result)
		sumutil.ReverseSlice(&result)
		num, err := sumutil.Display(&result)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		// Hiển thị lại form với kết quả

		c.HTML(http.StatusOK, "index.html", gin.H{"result": num})
		// Chạy server trên cổng 8080

	})
	///
	r.Run(":3000")
}
