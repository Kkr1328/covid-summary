package main

import (
    "covid-summary/handlers"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    router.GET("/covid/summary", handlers.GetCovidSummary)
    router.Run(":8080")
}
