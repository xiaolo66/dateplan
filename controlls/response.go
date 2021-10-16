package controlls

import (
	"dateplan/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int
	Message string
	Data  interface{}
}

func Result(c *gin.Context,data interface{},err error){
	if err!=nil{
		fail(c,err)
	}else {
		Success(c,data)
	}
}

func Success(c *gin.Context,data interface{}){
	utils.Log.Infof("response:%#v", data)
	c.JSON(http.StatusOK,Response{
		Code:200,
		Message: "success",
		Data: data,
	})
}

func fail(c *gin.Context,err error){
	utils.Log.Errorf("response:%#v", err)
	c.JSON(http.StatusOK,Response{
		Code:200,
		Message: fmt.Sprintf("err:%s",err),
		Data: nil,
	})
}