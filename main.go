package main

import (
	"context"
	"github.com/tencentyun/scf-go-lib/cloudfunction"
	"github.com/tencentyun/scf-go-lib/events"
)

func hello(ctx context.Context, req events.APIGatewayRequest) string {
	return req.Method
}

func main() {
	cloudfunction.Start(hello)
}
