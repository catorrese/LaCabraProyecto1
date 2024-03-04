package structs

import (
	uuid "github.com/google/uuid"
)

type SportmenSport struct {
	Sport string `json:"sport" validate:"oneof=basketball cycling"`
}

type Sportmen struct {
	UserId           uuid.UUID       `json:"user"`
	Name             string          `json:"name" validate:"required"`
	LastName         string          `json:"last_name" validate:"required"`
	Age              int             `json:"age" validate:"required"`
	Weight           int             `json:"weight" validate:"required"`
	Height           int             `json:"height" validate:"required"`
	CountryBirth     string          `json:"country_birth" validate:"required"`
	CityBirth        string          `json:"city_birth" validate:"required"`
	CountryResidence string          `json:"country_residence" validate:"required"`
	CityResidence    string          `json:"city_residence" validate:"required"`
	LengthResidence  int             `json:"length_residence" validate:"required"`
	Sport            []SportmenSport `validate:"dive"`
}

type Sportmen1 struct {
	UserId           uuid.UUID       `json:"-"`
	Name             string          `json:"name"`
	LastName         string          `json:"last_name"`
	Age              int             `json:"age"`
	Weight           int             `json:"weight"`
	Height           int             `json:"height"`
	CountryBirth     string          `json:"country_birth"`
	CityBirth        string          `json:"city_birth"`
	CountryResidence string          `json:"country_residence"`
	CityResidence    string          `json:"city_residence"`
	LengthResidence  int             `json:"length_residence"`
	Sport            []SportmenSport `validate:"dive"`
}

type SportmenResponse struct {
	Sportmen Sportmen1 `json:"sportmen"`
	Status   string    `json:"status"`
}
