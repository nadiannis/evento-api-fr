package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/nadiannis/evento-api-fr/internal/domain/response"
)

func ReadJSON(c *gin.Context, dst any) error {
	err := c.BindJSON(dst)
	if err != nil {
		var unmarshalTypeError *json.UnmarshalTypeError

		switch {
		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect JSON type for field %q", unmarshalTypeError.Field)
			}
			return fmt.Errorf("body contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)
		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")
		default:
			return err
		}
	}

	return nil
}

func WriteJSON(c *gin.Context, status int, data response.Response) {
	c.IndentedJSON(status, data)

	var message any
	switch v := data.(type) {
	case response.SuccessResponse:
		message = v.Message
	case response.ErrorResponse:
		message = v.Detail
	default:
		message = "request processed"
	}

	SetLogMessage(c, message)
}
