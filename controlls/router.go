package controlls

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterRouter(r *gin.Engine){

	r.Static("/static","static")
	//告诉gin去哪里找文件 *表示template文件夹下的所有文件
	r.LoadHTMLGlob("template/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.html",nil)
	})

	v1Group:=r.Group("v1")
	var BubbleCtr=NewBubble()
	{
		v1Group.GET("/todo",BubbleCtr.GetBubbles())
		v1Group.POST("/todo",BubbleCtr.AddBubble())
		v1Group.PUT("/todo/:id",BubbleCtr.PutBubbles())
		v1Group.DELETE("/todo/:id",BubbleCtr.DeleteBubbles())
	}

}
