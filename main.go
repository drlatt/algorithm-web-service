package main

import (
	"net/http"
	"strconv"

	"github.com/Depado/ginprom"
	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func simpleFib(n int) int64 {
	if n >= 0 && n <= 1 {
		return int64(n)
	}
	if n < 0 {
		return simpleFib(n+2) - simpleFib(n+1)
	}
	return simpleFib(n-1) + simpleFib(n-2)
}

func optimizedFib(n int) int64 {
	fibSlice := []int64{0, 1}

	if n == 0 || n == 1 {
		return int64(n)
	}
	if n < 0 {
		a, b := 0, 1
		for i := n; i < 0; i++ {
			a, b = b, a-b
		}
		return int64(a)
	}
	for i := 2; i < n+1; i++ {
		result := fibSlice[i-1] + fibSlice[i-2]
		fibSlice = append(fibSlice, result)
	}
	return fibSlice[n]
}

func simpleAckermann(m, n int) int64 {
	var result int64
	if m == 0 {
		result = int64(n + 1)
	}
	if m > 0 && n == 0 {
		result = simpleAckermann(m-1, 1)
	}
	if m > 0 && n > 0 {
		result = simpleAckermann(m-1, int(simpleAckermann(m, n-1)))
	}
	return result
}

func simpleFactorial(n int) uint64 {
	var result uint64
	if n == 0 || n == 1 {
		result = 1
		return result
	}
	result = uint64(n) * simpleFactorial(n-1)
	return result
}

func optimizedFactorial(n int) uint64 {
	var result uint64
	if n == 0 || n == 1 {
		result = 1
		return result
	}
	result = uint64(1)
	for i := n; i > 1; i-- {
		result = result * uint64(i)
	}
	return result
}

func main() {

	// initialize Gin router
	router := gin.Default()

	// Prometheus instrumentation
	prom := ginprom.New(
		ginprom.Engine(router),
		ginprom.Subsystem("gin"),
		ginprom.Path("/metrics"),
	)
	router.Use(prom.Instrument())

	// configure logger
	log.SetFormatter(&log.JSONFormatter{})
	log.SetReportCaller(true)

	// BodyF struct for factorial and fibonacci endpoints
	type BodyF struct {
		N string `json:"n" binding:"required"`
	}

	// BodyA struct for ackermann endpoint
	type BodyA struct {
		M string `json:"m" binding:"required"`
		N string `json:"n" binding:"required"`
	}

	v1 := router.Group("/api/v1")
	{
		// simple fibonacci
		v1.POST("/fibonacci", func(c *gin.Context) {
			var body BodyF
			if err := c.ShouldBindJSON(&body); err != nil {
				log.Errorf("error parsing request body: %s", err)
				c.JSON(http.StatusBadRequest, gin.H{"error": "missing parameters"})
				return
			}

			n, err := strconv.Atoi(body.N)
			if err != nil {
				log.Errorf("error converting to integer: %s", err)
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "enter a valid integer",
				})
				return
			}

			result := simpleFib(n)
			c.JSON(http.StatusOK, gin.H{
				"result": result,
			})
		})

		// simple Ackermann
		v1.POST("/ackermann", func(c *gin.Context) {
			var body BodyA
			if err := c.ShouldBindJSON(&body); err != nil {
				log.Errorf("error parsing request body: %s", err)
				c.JSON(http.StatusBadRequest, gin.H{"error": "missing parameters"})
				return
			}

			// check that m is non-negative
			if body.M < "0" {
				log.Errorf("wrong parameter passed: negative value %s provided", body.M)
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "negative value for m provided",
				})
				return
			}

			m, err := strconv.Atoi(body.M)
			if err != nil {
				log.Errorf("error converting to integer: %s", err)
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "invalid value for m provided",
				})
				return
			}

			// check that n is non-negative
			if body.N < "0" {
				log.Errorf("wrong parameter passed: negative value %s provided", body.N)
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "negative value for n provided",
				})
				return
			}

			n, err := strconv.Atoi(body.N)
			if err != nil {
				log.Errorf("error converting to integer: %s", err)
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "invalid value for n provided",
				})
				return
			}

			result := simpleAckermann(m, n)
			if result == 0 {
				log.Errorf("could not handle request")
				// c.JSON(http.status)
			}
			c.JSON(http.StatusOK, gin.H{
				"result": result,
			})
		})

		// simple factorial
		v1.POST("/factorial", func(c *gin.Context) {
			var body BodyF
			if err := c.ShouldBindJSON(&body); err != nil {
				log.Errorf("error parsing request body: %s", err)
				c.JSON(http.StatusBadRequest, gin.H{"error": "missing parameter"})
				return
			}

			// check that n is non-negative
			if body.N < "0" {
				log.Errorf("wrong parameter passed: negative value %s provided", body.N)
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "negative value for n provided",
				})
				return
			}

			n, err := strconv.Atoi(body.N)
			if err != nil {
				log.Errorf("error converting to integer: %s", err)
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "invalid value for n provided",
				})
				return
			}
			result := simpleFactorial(n)
			c.JSON(http.StatusOK, gin.H{
				"result": result,
			})
		})
	}

	v2 := router.Group("/api/v2")
	{

		// optimized fib
		v2.POST("/fibonacci", func(c *gin.Context) {
			var body BodyF
			if err := c.ShouldBindJSON(&body); err != nil {
				log.Errorf("error parsing request body: %s", err)
				c.JSON(http.StatusBadRequest, gin.H{"error": "missing parameters"})
				return
			}

			n, err := strconv.Atoi(body.N)
			if err != nil {
				log.Errorf("error converting to integer: %s", err)
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "enter a valid integer",
				})
				return
			}
			result := optimizedFib(n)
			c.JSON(http.StatusOK, gin.H{
				"result": result,
			})
		})

		// optimized factorial
		v2.POST("/factorial", func(c *gin.Context) {
			var body BodyF
			if err := c.ShouldBindJSON(&body); err != nil {
				log.Errorf("error parsing request body: %s", err)
				c.JSON(http.StatusBadRequest, gin.H{"error": "missing parameter"})
				return
			}

			// check that n is non-negative
			if body.N < "0" {
				log.Errorf("wrong parameter passed: negative value %s provided", body.N)
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "negative value for n provided",
				})
				return
			}

			n, err := strconv.Atoi(body.N)
			if err != nil {
				log.Errorf("error converting to integer: %s", err)
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "invalid value for n provided",
				})
				return
			}
			result := optimizedFactorial(n)
			c.JSON(http.StatusOK, gin.H{
				"result": result,
			})
		})
	}

	router.Run(":3000")
}
