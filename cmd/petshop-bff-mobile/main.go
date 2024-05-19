package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/kelseyhightower/envconfig"
	"github.com/petshop-system/petshop-bff-mobile/configuration/environment"
	"github.com/petshop-system/petshop-bff-mobile/handler"
	"github.com/petshop-system/petshop-bff-mobile/handler/iphone/customer"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
	"os"
)

var loggerSugar *zap.SugaredLogger

func init() {

	err := envconfig.Process("setting", &environment.Setting)
	if err != nil {
		panic(err.Error())
	}

	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	jsonEncoder := zapcore.NewJSONEncoder(config)
	core := zapcore.NewTee(
		zapcore.NewCore(jsonEncoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel),
	)
	logger := zap.New(core, zap.AddCaller())
	defer logger.Sync() // flushes buffer, if any
	loggerSugar = logger.Sugar()

}

func main() {

	genericHandler := &handler.Generic{
		LoggerSugar: loggerSugar,
	}

	iPhoneRegisterHandler := customer.NewRegisterHandler(loggerSugar)

	contextPath := environment.Setting.Server.Context
	newRouter := handler.GetNewRouter(loggerSugar)
	newRouter.GetChiRouter().Route(fmt.Sprintf("/%s", contextPath), func(r chi.Router) {
		r.NotFound(genericHandler.NotFound)
		r.Group(newRouter.AddGroupHandlerHealthCheck(genericHandler))
		r.Group(newRouter.AddGroupHandlerIPhoneCustomer(&iPhoneRegisterHandler))
	})

	serverHttp := &http.Server{
		Addr:           fmt.Sprintf(":%s", environment.Setting.Server.Port),
		Handler:        newRouter.GetChiRouter(),
		ReadTimeout:    environment.Setting.Server.ReadTimeout,
		WriteTimeout:   environment.Setting.Server.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	loggerSugar.Infow("server started", "port", serverHttp.Addr,
		"contextPath", contextPath)

	if err := serverHttp.ListenAndServe(); err != nil {
		loggerSugar.Errorw("error to listen and starts server", "port", serverHttp.Addr,
			"contextPath", contextPath, "err", err.Error())
		panic(err.Error())
	}

}
