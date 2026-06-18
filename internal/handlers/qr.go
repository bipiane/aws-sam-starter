package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func ListQR(_ context.Context, _ events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	return jsonResponse(http.StatusOK, `{"qr_codes":[]}`)
}

func CreateQR(_ context.Context, _ events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	return jsonResponse(http.StatusCreated, `{"id":"","url":""}`)
}

func GetQR(_ context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	id := req.PathParameters["id"]
	return jsonResponse(http.StatusOK, fmt.Sprintf(`{"id":%q,"url":""}`, id))
}

func DeleteQR(_ context.Context, _ events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	return events.APIGatewayV2HTTPResponse{StatusCode: http.StatusNoContent}, nil
}

func jsonResponse(status int, body string) (events.APIGatewayV2HTTPResponse, error) {
	return events.APIGatewayV2HTTPResponse{
		StatusCode: status,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       body,
	}, nil
}
