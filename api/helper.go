package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type jsonResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

var validate = validator.New()

func (app *Config) validateBody(ctx *gin.Context, data any) error {
	if err := ctx.BindJSON(&data); err != nil {
		return err
	}

	if err := validate.Struct(&data); err != nil {
		return err
	}

	return nil
}

func (app *Config) writeJSON(ctx *gin.Context, status int, data any) {
	ctx.JSON(status, jsonResponse{
		Status:  status,
		Message: "success",
		Data:    data,
	})
}

func (app *Config) errorJSON(ctx *gin.Context, err error, status ...int) {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}
	ctx.JSON(statusCode, jsonResponse{
		Status:  statusCode,
		Message: err.Error(),
	})
}
