package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// update user by id
func DeleteUserController(c echo.Context) error {

	users := User{}
	id, _ := strconv.Atoi(c.Param("id"))

	if err := DB.First(&users, id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "User not found",
		})
	}

	if err := DB.Delete(&users, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusBadRequest, map[string]interface{}{
		"message": "Success delete user",
	})
}
