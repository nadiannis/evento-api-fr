package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func ReadIDParam(c *gin.Context) (int64, error) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, ErrInvalidID
	}

	return id, nil
}
