package tests

import (
    "testing"
    "covid-summary/handlers"
    "github.com/gin-gonic/gin"
    "net/http"
    "net/http/httptest"
    "github.com/stretchr/testify/assert"
)

func TestGetCovidSummary(t *testing.T) {
    router := gin.Default()
    router.GET("/covid/summary", handlers.GetCovidSummary)

    response := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/covid/summary", nil)
    router.ServeHTTP(response, req)

    assert.Equal(t, http.StatusOK, response.Code)
	assert.Contains(t, response.Body.String(), "province")
	assert.Contains(t, response.Body.String(), "ageGroup")
}
