package main

import (
	"context"
	"encoding/json"
	"github.com/tencentyun/scf-go-lib/cloudfunction"
	"github.com/tencentyun/scf-go-lib/events"
)

type Body struct {
	Code int
	Msg  string
}

func hello(ctx context.Context, req events.APIGatewayRequest) (events.APIGatewayResponse, error) {
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
