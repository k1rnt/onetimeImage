package handler

import (
	"net/http"

	"github.com/k1rnt/onetimeImage/usecase"

	"github.com/labstack/echo/v4"
)

type HealthHandler interface {
	Post() echo.HandlerFunc
}

type healthHandler struct {
	healthUsecase usecase.HealthUsecase
}

func NewHealthHandler(healthUsecase usecase.HealthUsecase) HealthHandler {
	return &healthHandler{healthUsecase: healthUsecase}
}

type responseHealth struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// Handler
func (hu *healthHandler) Post() echo.HandlerFunc {
	return func(c echo.Context) error {
		createdHealth, err := hu.healthUsecase.Create("ok", "👍")
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		json := responseHealth{
			Status:  createdHealth.Status,
			Message: createdHealth.Message,
		}

		return c.JSON(http.StatusOK, json)
	}
}
