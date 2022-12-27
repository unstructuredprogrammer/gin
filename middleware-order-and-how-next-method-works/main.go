package main

import (
	"fmt"
	"github.com/alirezaeftekhari/gin"
	"log"
	"strings"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(strings.Repeat("*", 20))
		fmt.Println("main.Logger")
		fmt.Println(strings.Repeat("*", 20))

		t := time.Now()

		// Set example variable
		c.Set("example", "12345")
		//fmt.Println(c.HandlerName())

		// before request
		/*
			I should uncomment this c.Next() method to see the order of execution
			When we call c.Next() method the next middleware will go to executing simultaneously
		*/
		c.Next()

		// after request
		latency := time.Since(t)
		log.Println("latency", latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println("status", status)
	}
}

func Logger2() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(strings.Repeat("*", 20))
		fmt.Println("main.Logger2")
		fmt.Println(strings.Repeat("*", 20))

		t := time.Now()

		// Set example variable
		c.Set("example2", "54321")
		//fmt.Println(c.HandlerName())

		// before request
		/*
			I should uncomment this c.Next() method to see the order of execution
			When we call c.Next() method the next middleware will go to executing simultaneously
		*/
		c.Next()

		// after request
		latency := time.Since(t)
		log.Println("latency2", latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println("status2", status)
	}
}

func main() {
	r := gin.Default()

	/*
		router is using the Logger() func that defined above
	*/
	r.Use(Logger())
	r.Use(Logger2())

	r.GET("/test", func(c *gin.Context) {
		example := c.MustGet("example").(string)
		example2 := c.MustGet("example2").(string)

		// it would print: "12345"
		log.Println("example", example)
		log.Println("example2", example2)
		c.JSON(200, gin.H{
			"ex":  example,
			"ex2": example2,
		})
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
