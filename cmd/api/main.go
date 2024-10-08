package main

import (
	// db "calapp/internal/db"
	// models "calapp/internal/models"
	cal "calapp/internal/calculator"
	"fmt"
	// "github.com/gin-gonic/gin"
)

func main() {
	// DB Init:
	// db.Connect()
	var c cal.CalObj
	res := c.Calculator(2, 2, "DIV")
	fmt.Println(res)
	// Router Init:
	// router := gin.Default()
}
