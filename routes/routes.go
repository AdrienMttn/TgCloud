package routes

import (
	"net/http"
	"telegram-cloud/controllers"
)



type Controllers struct {
	Img *controllers.ImgController
	AcceptFormat *controllers.AcceptFormatController
	Stat *controllers.StatController
}

func SetupRouter(c Controllers) http.Handler {
	router := http.NewServeMux()

	registerStat(router, *c.Stat)
	registerImg(router, *c.Img)
	registerAcceptFormat(router, *c.AcceptFormat)

	return router
}