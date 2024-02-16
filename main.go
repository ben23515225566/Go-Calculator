package main

import (
	myMath "Go_Calculator/calculate" // 同python的import Go_Calculator/calculate as myMath
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)
type Data struct {
	Ans float64
}
var temp float64 = 0
var operator string = ""
var ans float64 = 0

func index(c *gin.Context) {
	value := temp
	data := Data{value}
	c.HTML(http.StatusOK, "Go_calculator.html", data)
}

func operation(c *gin.Context){
	operator_temp, _ := c.GetPostForm("submit")
	strValue, _ := c.GetPostForm("value")
	floatValue, _ := strconv.ParseFloat(strValue, 64)
	if(operator == ""){
		operator = operator_temp	
		temp= floatValue
	}else{
		switch operator {
		case "+":
			ans = myMath.Add(temp, floatValue)
		case "-":
			ans = myMath.Sub(temp, floatValue)
        case "*":
			ans = myMath.Mul(temp, floatValue)
		case "/":
			ans = myMath.Div(temp, floatValue)
		case "%":
			ans = myMath.Mod(temp, floatValue)
		case "^":
			ans = myMath.Pow(temp, floatValue)
		}
		operator = ""
	}
	data := Data{ans}
	ans = 0
	c.HTML(http.StatusOK, "Go_calculator.html", data)
}

func main(){
	router := gin.Default()
	router.LoadHTMLGlob("./Template/html/*")
	router.Static("/assets", "./Template/assets")
	router.GET("/", index)
	router.POST("/", operation)
	router.Run(":8888")
}