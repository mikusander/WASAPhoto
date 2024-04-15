package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {

	//login
	rt.router.POST("/session", rt.wrap(rt.doLogin))

	//User
	rt.router.PUT("/users/:user_name/username", rt.wrap(rt.setMyUserName))

	//Photo
	rt.router.POST("/users/:username/photo", rt.wrap(rt.uploadPhoto))

	// Register routes
	rt.router.GET("/", rt.getHelloWorld)
	rt.router.GET("/context", rt.wrap(rt.getContextReply))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
