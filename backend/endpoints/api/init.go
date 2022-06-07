package api

import (
	"github.com/gofiber/fiber/v2"

	apiAccount "backend/endpoints/api/account"
	apiRecord "backend/endpoints/api/record"
	"backend/loaders/fiber/middlewares"
)

func Init(router fiber.Router) {
	// Account Group
	account := router.Group("account/")
	account.Post("session/new", apiAccount.NewSessionHandler)

	record := router.Group("record/")
	record.Get("ios-profile", apiRecord.GenerateIosProfileHandler)

	authenticatedRecord := record.Use(middlewares.Authen)
	authenticatedRecord.Get("query-history", apiRecord.QueryHistoryHandler)
}
