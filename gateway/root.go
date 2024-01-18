package gateway

import "net/http"

type GatewayService interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}
