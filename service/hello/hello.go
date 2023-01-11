package hello

import (
	"github.com/gin-gonic/gin"
	"github.com/rodert/hepburn/model/request"
	"github.com/rodert/hepburn/model/response"
)

func Hello(c *gin.Context, req *request.HelloRequest) (response.HelloResponse, error) {
	hr := response.HelloResponse{
		Message: "hello, hepburn ",
	}
	return hr, nil
}
