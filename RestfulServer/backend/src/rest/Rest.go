package rest

import (
	"github.com/gin-gonic/gin"
)

func RunAPI(address string) error {
	r := gin.Default()
	h, _ := NewHandler()
	r.GET("/products", h.GetProducts)
	r.GET("/promos", h.GetPromos)

	userGroup := r.Group("/user")
	{
		userGroup.POST("/:id/signout", h.SignOut)
		userGroup.GET("/:id/orders", h.GetOrders)
	}

	usersGroup := r.Group("/users")
	{
		usersGroup.POST("/signin", h.SignIn)
		usersGroup.POST("", h.AddUser)
		usersGroup.POST("/charge", h.Charge)
	}

	return r.Run(address)
}
