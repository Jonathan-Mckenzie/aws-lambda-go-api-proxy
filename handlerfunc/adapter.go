package handlerfunc

import (
	"net/http"

	"github.com/Jonathan-Mckenzie/aws-lambda-go-api-proxy/httpadapter"
)

type HandlerFuncAdapter = httpadapter.HandlerAdapter

func New(handlerFunc http.HandlerFunc) *HandlerFuncAdapter {
	return httpadapter.New(handlerFunc)
}
