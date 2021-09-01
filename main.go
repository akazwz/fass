package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/tencentyun/scf-go-lib/cloudfunction"
	"github.com/tencentyun/scf-go-lib/events"
)

type Body struct {
	Code int
	Msg  string
}

func hello() (events.APIGatewayResponse, error) {
	body := Body{
		Code: 200,
		Msg:  "success",
	}
	marshal, err := json.Marshal(body)
	if err != nil {
		return events.APIGatewayResponse{}, err
	}

	if err != nil {
		return events.APIGatewayResponse{}, err
	}
	server := gin.Default()
	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "success",
		})
		return
	})
	err = server.Run(":9000")
	if err != nil {
		return events.APIGatewayResponse{}, err
	}
	return events.APIGatewayResponse{
		IsBase64Encoded: false,
		StatusCode:      200,
		Headers:         nil,
		Body:            string(marshal),
	}, nil
}

func main() {
	cloudfunction.Start(hello)
}
