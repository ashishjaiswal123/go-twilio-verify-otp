package api

import (
	"context"
	"fmt"
	"net/http"
	"sms-verify/data"
	"time"

	"github.com/gin-gonic/gin"
)

const appTimeOut = time.Second * 10

func (app *Config) sendSMS() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), appTimeOut)
		var payload data.OTPData
		defer cancel()

		app.validateBody(ctx, &payload)

		newData := data.OTPData{
			PhoneNumber: payload.PhoneNumber,
		}

		_, err := app.twilioSendOTP(newData.PhoneNumber)
		if err != nil {
			app.errorJSON(ctx, err)
			return
		}

		app.writeJSON(ctx, http.StatusAccepted, "otp sent successfully!")
	}
}

func (app *Config) verifySMS() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), appTimeOut)
		var payload data.VerifyData
		defer cancel()

		app.validateBody(ctx, &payload)

		newData := data.VerifyData{
			User: payload.User,
			Code: payload.Code,
		}

		err := app.twilioVerifyOTP(newData.User.PhoneNumber, newData.Code)
		fmt.Println("err: ", err)
		if err != nil {
			app.errorJSON(ctx, err)
			return
		}

		app.writeJSON(ctx, http.StatusAccepted, "OTP verified successfully!")
	}
}
