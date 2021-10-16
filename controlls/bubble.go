package controlls

import (
	"dateplan/dto"
	"dateplan/servers"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Bubble struct {

}

func NewBubble()*Bubble{
	return &Bubble{}
}


func (b *Bubble)GetBubbles()gin.HandlerFunc{
	return func(c *gin.Context) {
		list,err:=servers.BubbleSvr.GetAllBubble()
		if err!=nil{
			Result(c,nil,err)
			return
		}
		c.JSON(http.StatusOK,list)
	}
}

func(b *Bubble)AddBubble()gin.HandlerFunc{
	return func(c *gin.Context) {
		var task dto.Bubble
		err:=c.Bind(&task)
		fmt.Println(task)
		if err!=nil{
			Result(c,nil,err)
		}
		err=servers.BubbleSvr.AddBubble(task)
		Result(c,nil,err)
	}
}

func(b *Bubble)PutBubbles()gin.HandlerFunc{
	return func(c *gin.Context) {
		id,ok:=c.Params.Get("id")
		if !ok{
			Result(c,nil,fmt.Errorf("id not exist"))
		}
		err:=servers.BubbleSvr.UpdateBubble(id)
		Result(c,nil,err)
	}
}

func(b *Bubble)DeleteBubbles()gin.HandlerFunc{
	return func(c *gin.Context) {
		id,ok:=c.Params.Get("id")
		if !ok{
			Result(c,nil,fmt.Errorf("id not exist"))
		}
		err:=servers.BubbleSvr.DeleteBubble(id)
		Result(c,nil,err)
	}
}