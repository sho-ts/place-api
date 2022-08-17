package util

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetLimitAndOffset(c *gin.Context) (int, int) {
	limit, err := strconv.Atoi(c.Query("limit"))

	if err != nil || limit > 30 {
		limit = 10
	}

	offset, err := strconv.Atoi(c.Query("offset"))

	if err != nil {
		offset = 0
	}

	return limit, offset
}
