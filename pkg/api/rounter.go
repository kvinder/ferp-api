package api

import "github.com/ant0ine/go-json-rest/rest"
import "log"

func NewRouter() rest.App {
	router, err := rest.MakeRouter(
		rest.Post("/api/register", register),
		rest.Post("/api/upload", upload),
	)
	if err != nil {
		log.Fatal(err)
	}
	return router
}

func NewAPI(router rest.App) (api *rest.Api) {
	api = rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	allowedMethods := []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	allowedHeaders := []string{
		"Accept",
		"Authorization",
		"X-Real-IP",
		"Content-Type",
		"X-Custom-Header",
		"Language",
		"Origin",
	}
	api.Use(&rest.CorsMiddleware{
		RejectNonCorsRequests: false,
		OriginValidator: func(origin string, request *rest.Request) bool {
			return true
		},
		AllowedMethods:                allowedMethods,
		AllowedHeaders:                allowedHeaders,
		AccessControlAllowCredentials: true,
		AccessControlMaxAge:           3600,
	})

	loginMiddle := &LoginMiddleware{}
	api.Use(loginMiddle)

	api.SetApp(router)
	return
}

type LoginMiddleware struct {
}

var pathNeedLogin = map[string]bool{
	"/api/user": true,
}

func (login *LoginMiddleware) MiddlewareFunc(handler rest.HandlerFunc) rest.HandlerFunc {
	return func(w rest.ResponseWriter, r *rest.Request) {
		if pathNeedLogin[r.URL.Path] {
			// err := TokenValidator(strings.Replace(r.Header.Get("Authorization"), "Bearer ", "", -1))
			// if err != nil {
			// 	w.WriteHeader(http.StatusUnauthorized)
			// 	w.WriteJson(map[string]string{"error": err.Error()})
			// 	return
			// }
		}
		handler(w, r)
	}
}
