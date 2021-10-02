package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Update all users
func UpdateUserController(c echo.Context) error {
	user := User{}
	c.Bind(&user)

	// id := c.Param("id")
	id, _ := strconv.Atoi(c.Param("id"))

	// if err := DB.First(&users, id).Error; err != nil {
	// 	return c.JSON(http.StatusBadRequest, map[string]interface{}{
	// 		"message": "User not found",
	// 	})
	// }

	// DB.Save(&user)
	// if err := DB.Model(&id).Updates(&user).Error; err != nil {
	// 	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	// }
	if err := DB.Where("id = ?", id).Updates(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user",
		"user":    user,
	})
}
