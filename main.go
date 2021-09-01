package main

import (
	"context"
	"github.com/tencentyun/scf-go-lib/cloudfunction"
	"github.com/tencentyun/scf-go-lib/events"
)

func hello(ctx context.Context, req events.APIGatewayRequest) events.APIGatewayResponse {
	return events.APIGatewayResponse{
		IsBase64Encoded: false,
		StatusCode:      200,
		Headers:         nil,
		Body:            "{msg: success}",
	}
}

func main() {
	cloudfunction.Start(hello)
}
