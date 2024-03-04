package utils

import (
	"fiber-orchestrator/structs"

	"github.com/gofiber/fiber/v2"
	uuid "github.com/google/uuid"
)

func GoRequests(c *fiber.Ctx, UserUuid uuid.UUID) (structs.Response, structs.Response, structs.Response) {
	response_sport := make(chan structs.Response)
	response_user := make(chan structs.Response)
	response_sub := make(chan structs.Response)
	go func() {
		data, statusCode, err := CreateSportmen(c, UserUuid)
		response_sport <- structs.Response{Data: data, StatusCode: statusCode, Err: err}
	}()
	go func() {
		data, statusCode, err := CreateUser(c, UserUuid)
		response_user <- structs.Response{Data: data, StatusCode: statusCode, Err: err}
	}()
	go func() {
		data, statusCode, err := GetSub()
		response_sub <- structs.Response{Data: data, StatusCode: statusCode, Err: err}
	}()
	return <-response_sport, <-response_user, <-response_sub
}
