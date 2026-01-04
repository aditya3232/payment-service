package routes

import (
	controllers "payment-service/controllers/http"
	routes "payment-service/routes/payment"

	"github.com/gin-gonic/gin"
)

type Registry struct {
	controller controllers.IControllerRegistry
	group      *gin.RouterGroup
}

type IRouteRegister interface {
	Serve()
}

func NewRouteRegistry(controller controllers.IControllerRegistry, group *gin.RouterGroup) IRouteRegister {
	return &Registry{controller: controller, group: group}
}

func (r *Registry) paymentRoute() routes.IPaymentRoute {
	return routes.NewPaymentRoute(r.controller, r.group)
}

func (r *Registry) Serve() {
	r.paymentRoute().Run()
}
