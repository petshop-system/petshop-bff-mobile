package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/petshop-system/petshop-bff-mobile/handler/iphone/customer"
	"go.uber.org/zap"
)

type Router struct {
	ContextPath string
	chiRouter   chi.Router
	LoggerSugar *zap.SugaredLogger
}

func GetNewRouter(loggerSugar *zap.SugaredLogger) Router {
	router := chi.NewRouter()
	return Router{
		chiRouter:   router,
		LoggerSugar: loggerSugar,
	}
}

func (router Router) GetChiRouter() chi.Router {
	return router.chiRouter
}

func (router Router) AddGroupHandlerHealthCheck(ah *Generic) func(r chi.Router) {
	return func(r chi.Router) {
		r.Route("/health-check", func(r chi.Router) {
			r.Get("/", ah.HealthCheck)
		})
	}
}

func (router Router) AddGroupHandlerIPhoneCustomer(rh *customer.IPhoneCustomerHandler) func(r chi.Router) {
	return func(r chi.Router) {
		r.Route("/iphone/customer", func(r chi.Router) {
			r.Get("/create-screen", rh.CreateScreen)
			r.Post("/create", rh.Create)
		})
	}
}
