package counterd

import (
	"fmt"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

// Service is a struct that defines handlers to handle all routes
type Service struct {
	num int
}

// InitializeRoutes defines all the routes and function mappings
func InitializeRoutes() *gin.Engine {
	r := gin.Default()
	fmt.Println("Starting up tinderd")
	service, err := InitService()
	if err != nil {
		fmt.Printf("Fatal Error: %s", err.Error())

		return nil
	}

	r.GET("/count", service.Count)
	r.NoRoute(func(c *gin.Context) {
		c.JSON(
			http.StatusNotFound,
			gin.H{"code": "PAGE_NOT_FOUND", "message": "Route does not exist"},
		)
	})
	fmt.Println("Quoted started. Ready to take requests")
	return r
}

func (s *Service) Count(c *gin.Context) {
	increment := 1
	incrementStr := c.Query("i")
	if incrementStr != "" {
		var err error
		increment, err = strconv.Atoi(incrementStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}
	s.num += increment
	fmt.Printf("New num %d\n", s.num)
	c.JSON(http.StatusOK, gin.H{"count": s.num})
	return
}

// InitService returns an instantiation of Service object
func InitService() (*Service, error) {
	return &Service{}, nil
}
