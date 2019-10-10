package main

import (
	"fmt"
	"net/http"
	"reflect"

	"gopkg.in/go-playground/validator.v9"

	"github.com/labstack/echo"
)

/*
	curl -X POST \
		http://localhost:1323/data \
		-H 'Content-Type: application/json' \
		-d '{
		"name": "golf",
		"type": "campaign_stat",
		"timestamp": 1566801149000,
		"error": {
			"code": "MKT-0001",
			"timestamp": 1566801149000
		}
	}'
*/

type Data struct {
	Name      string     `json:"name" validate:"required"`
	Type      string     `json:"type" valdate:"required"`
	Timestamp int64      `json:"timestamp" validate:"requiredTimestamp"`
	Error     *ErrorData `json:"error,omitempty"`
}

type ErrorData struct {
	Code      string `json:"code,omitempty" validate:"required"`
	Timestamp int64  `json:"timestamp" validate:"requiredTimeStampOnError"`
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

// RequiredValidateTimestamp handles for basic event
func RequiredValidateTimestamp(timestamp validator.FieldLevel) bool {
	// return value from pointer https://golang.org/pkg/reflect/#Indirect
	campaignType := reflect.Indirect(timestamp.Top()).FieldByName("Type")
	if campaignType.Len() == 0 {
		return false
	}

	fmt.Println(campaignType.String())
	fmt.Println(timestamp.Field().Kind())
	fmt.Println(timestamp.Field().Int())

	if campaignType.String() == "campaign_stat" && timestamp.Field().Int() == 0 {
		return false
	}

	// fmt.Printf("%+v\n", timestamp.Parent())
	// fmt.Printf("%+v\n", timestamp.Field())
	// fmt.Printf("%+v\n", timestamp.FieldName())
	// fmt.Printf("%+v\n", timestamp.StructFieldName())

	return true
}

// RequiredValidateTimestampOnError handles for error event
func RequiredValidateTimestampOnError(timestamp validator.FieldLevel) bool {
	// return value from pointer https://golang.org/pkg/reflect/#Indirect
	campaignType := reflect.Indirect(timestamp.Top()).FieldByName("Type")
	if campaignType.Len() == 0 {
		return false
	}

	if campaignType.String() == "error" && (timestamp.Field().Kind() != reflect.Int64 || timestamp.Field().Int() == 0) {
		return false
	}

	return true
}

func main() {
	e := echo.New()

	customValidator := validator.New()
	customValidator.RegisterValidation("requiredTimestamp", RequiredValidateTimestamp)
	customValidator.RegisterValidation("requiredTimeStampOnError", RequiredValidateTimestampOnError)

	e.Validator = &CustomValidator{validator: customValidator}
	e.POST("/data", func(c echo.Context) (err error) {
		var data Data

		if err = c.Bind(&data); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		if err = c.Validate(&data); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("%+v", err))
		}

		return c.JSON(http.StatusOK, &data)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
