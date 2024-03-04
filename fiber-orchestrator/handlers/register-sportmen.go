package handlers

import (
	"fiber-orchestrator/structs"
	"fiber-orchestrator/utils"

	"github.com/gofiber/fiber/v2"

	uuid "github.com/google/uuid"

	"encoding/json"
)

func RegisterSportmen(c *fiber.Ctx) error {
	UserUuid := uuid.New()

	response_sport, response_user, response_sub := utils.GoRequests(c, UserUuid)

	if response_sport.Err != nil || response_user.Err != nil || response_sub.Err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "failure"})
	}

	if response_sport.StatusCode != 201 || response_user.StatusCode != 201 || response_sub.StatusCode != 200 {
		return c.Status(400).JSON(fiber.Map{"status": "failure"})
	}

	var SportmenJson structs.SportmenResponse
	err := json.Unmarshal(response_sport.Data, &SportmenJson)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "failure"})
	}

	var UserJson structs.UserRespose
	err_1 := json.Unmarshal(response_user.Data, &UserJson)
	if err_1 != nil {
		return c.Status(400).JSON(fiber.Map{"status": "failure"})
	}

	var SubJson []structs.Sub
	err_2 := json.Unmarshal(response_sub.Data, &SubJson)
	if err_2 != nil {
		println("holi")
		return c.Status(400).JSON(fiber.Map{"status": "failure"})
	}

	return c.Status(201).JSON(fiber.Map{"status": "success", "sportmen": SportmenJson.Sportmen, "user": UserJson, "sub": SubJson})
}
