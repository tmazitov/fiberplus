package main

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"

	fiberplus "github.com/tmazitov/fiberplus"
	"github.com/tmazitov/fiberplus/examples/handlers"
	"github.com/tmazitov/fiberplus/examples/services"
)

func main() {
	var (
		innerServices = &services.Services{
			Database: services.Database{},
		}
		a = fiberplus.NewApp(&fiberplus.AppConfig[services.Services]{
			Services: innerServices,
		})
	)
	a.Core().Use(logger.New())

	a.Group("/test").
		Add(&fiberplus.Endpoint[services.Services]{Method: "POST", Route: "/in", Handler: &handlers.ReadHandlerExample{}}).
		Add(&fiberplus.Endpoint[services.Services]{Method: "GET", Route: "/out", Handler: &handlers.WriteHandlerExample{}}).
		Add(&fiberplus.Endpoint[services.Services]{Method: "POST", Route: "/io", Handler: &handlers.CoreHandlerExample{}})

	log.Errorf("app error : ", a.Run(":8070"))
}
