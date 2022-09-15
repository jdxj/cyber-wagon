package web

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type RegisterRoute func(gin.IRouter)

func Start(addr string, rr RegisterRoute) {
	r := gin.Default()
	rr(r)

	server := &http.Server{
		Addr:    addr,
		Handler: r,
	}
	err := server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		logrus.Errorf("listen and serve: %s", err)
	}
}
