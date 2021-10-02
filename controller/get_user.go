package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// get user by id
func GetUserController(c echo.Context) error {
	user := User{}
	c.Bind(&user)

	id, _ := strconv.Atoi(c.Param("id"))

	if err := DB.First(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user",
		"user":    user,
	})
}
