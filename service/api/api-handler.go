package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {

	// login
	rt.router.POST("/session", rt.wrap(rt.doLogin))

	// User
	rt.router.PUT("/users/:username/username", rt.wrap(rt.setMyUserName))
	rt.router.GET("/users/:username/profile", rt.wrap(rt.getUserProfile))
	rt.router.GET("/users/:username/stream", rt.wrap(rt.getMyStream))
	rt.router.GET("/users/:username/id", rt.wrap(rt.getIDUser))

	// Photo
	rt.router.POST("/users/:username/photo", rt.wrap(rt.uploadPhoto))
	rt.router.DELETE("/users/:username/photo/:photoid", rt.wrap(rt.deletePhoto))

	// follow
	rt.router.PUT("/users/:username/follow/:followid", rt.wrap(rt.followUser))
	rt.router.DELETE("/users/:username/follow/:followid", rt.wrap(rt.unfollowUser))
	rt.router.GET("/users/:username/follow/:followid", rt.wrap(rt.isFollowUser))

	// ban
	rt.router.PUT("/users/:username/ban/:banid", rt.wrap(rt.banUser))
	rt.router.DELETE("/users/:username/ban/:banid", rt.wrap(rt.unbanUser))
	rt.router.GET("/users/:username/ban/:banid", rt.wrap(rt.isBanUser))

	// comment
	rt.router.POST("/users/:username/photo/:photoid/comment", rt.wrap(rt.commentPhoto))
	rt.router.DELETE("/users/:username/photo/:photoid/comment/:commentid", rt.wrap(rt.uncommentPhoto))

	// like
	rt.router.PUT("/users/:username/photo/:photoid/like/:likeid", rt.wrap(rt.likePhoto))
	rt.router.DELETE("/users/:username/photo/:photoid/like/:likeid", rt.wrap(rt.unlikePhoto))
	rt.router.GET("/users/:username/photo/:photoid/like/:likeid", rt.wrap(rt.isLikePhoto))

	// Register routes
	rt.router.GET("/", rt.getHelloWorld)
	rt.router.GET("/context", rt.wrap(rt.getContextReply))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
