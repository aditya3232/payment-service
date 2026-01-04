package routes

import (
	controllers "payment-service/controllers/http"

	"github.com/gin-gonic/gin"
)

type PaymentRoute struct {
	controller controllers.IControllerRegistry
	group      *gin.RouterGroup
}

type IPaymentRoute interface {
	Run()
}

func NewPaymentRoute(controller controllers.IControllerRegistry, group *gin.RouterGroup) IPaymentRoute {
	return &PaymentRoute{controller: controller, group: group}
}

func (r *PaymentRoute) Run() {
	group := r.group.Group("/payments")
	group.GET("", r.controller.GetPayment().FindAllWithoutPagination)
	group.POST("", r.controller.GetPayment().Create)
}
