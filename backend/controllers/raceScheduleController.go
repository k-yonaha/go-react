package controllers

import (
	"backend/database"
	"backend/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

// 次のレースを取得するAPI
func GetNextRaceByCourse(c echo.Context) error {
	courseName := c.QueryParam("course_name")
	if courseName == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "course_nameが必要です"})
	}

	// レーススケジュールを取得
	raceSchedule, err := services.GetRaceScheduleByCourseName(database.DB, courseName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, raceSchedule)
}

func GetAllRaceSchedules(c echo.Context) error {

	raceSchedulesByCourse, err := services.GetRaceSchedulesByDate(database.DB)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, raceSchedulesByCourse)
}
