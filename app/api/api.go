package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"qsoft-go-test/app/middleware"
	"strconv"
	"time"
)

var sugar = zap.NewExample().Sugar()

// InitHandlers - is a function to init our (only one) handler
func InitHandlers() *gin.Engine {
	r := gin.New()
	r.Use(middleware.HeaderCheck())

	r.GET("/when/:year", whenHandler)
	r.GET("/info", infoHandler)
	r.NoRoute(notFound)

	return r
}

// whenHandler - is a handler function to which will show how many days are left or have passed before January 1st of the year specified in the route parameter to the current time
func whenHandler(c *gin.Context) {
	timeNow := time.Now()
	userYear, err := strconv.Atoi(c.Param("year"))
	if err != nil || userYear == 0 { //There are no 0 year
		c.String(http.StatusBadRequest, "Вы указали некорректный год")
		sugar.Info("The user entered incorrect data. Specified value: ", c.Param("year"), " time of:", time.Now().Format("02.01.2006.15.04.05"))
	} else {
		userTime := time.Date(userYear, 1, 1, 0, 0, 0, 0, time.UTC)
		dayValue := ""

		if timeNow.After(userTime) {
			timeNow, userTime = userTime, timeNow
		}

		days := -timeNow.YearDay()

		for year := timeNow.Year(); year < userTime.Year(); year++ {
			days += time.Date(year, time.December, 31, 0, 0, 0, 0, time.UTC).YearDay()
		}

		days += userTime.YearDay()

		checkDay := days % 10
		switch {
		case checkDay == 1:
			dayValue = "день"
		case checkDay > 1 && checkDay < 5:
			dayValue = "дня"
		default:
			dayValue = "дней"
		}
		if days%100 > 10 && days%100 < 15 {
			dayValue = "дней"
		}

		if userYear > timeNow.Year() {
			c.String(http.StatusOK, "Осталось %d %s до 01.01.%d", days, dayValue, userYear)
		} else if userYear < 0 {
			c.String(http.StatusOK, "Прошло %d %s с 01.01.%d до н.э.", days, dayValue, userYear/-1)
		} else {
			c.String(http.StatusOK, "Прошло %d %s с 01.01.%d", days, dayValue, userYear)
		}

		sugar.Info("The function for calculating the number of days worked correctly")
	}
}

// infoHandler Handler for router path /info which include information about application
func infoHandler(c *gin.Context) {
	c.String(http.StatusOK, "В данном приложение работает только путь /when/... где '...' обозначает указаный вами год."+
		" Функционал заключается в том, что приложение посчитает количество дней прошедших или оставшихся c первого января указаного вами года до наших дней."+
		" Попробуйте перейти по адресу с указанием интересующего вас года. Так же при добавление header запроса X-PING со значение ping вы получите ответ в header X-PONG со значение pong.")
	sugar.Info("Connecting to /info")
}

// notFound Handler for any incorrect path
func notFound(c *gin.Context) {
	url := c.Request.URL.Path
	c.String(http.StatusNotFound, "По такому адресу ничего нет :( \nПерейдите на страницу информации - localhost:8080/info")
	sugar.Info("Attempt to connect to a non-existent url:", url, " time of: ", time.Now().Format("02.01.2006.15.04.05"))
}
