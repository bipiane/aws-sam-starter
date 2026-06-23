package main

import (
	"context"
	"net/http"

	"qr_cloud/internal/handlers"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func router(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	switch req.RouteKey {
	case "GET /api/qr":
		return handlers.ListQR(ctx, req)
	case "POST /api/qr":
		return handlers.CreateQR(ctx, req)
	case "GET /api/qr/{id}":
		return handlers.GetQR(ctx, req)
	case "DELETE /api/qr/{id}":
		return handlers.DeleteQR(ctx, req)
	default:
		return events.APIGatewayV2HTTPResponse{StatusCode: http.StatusNotFound}, nil
	}
}

func main() {
	lambda.Start(router)
}
