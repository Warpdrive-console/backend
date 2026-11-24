package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// login
	// most of the actual auth can handle client side, server only has to handle validators
	// server stores games and other data like that so we need ways to upload, probably through cf signed urls and shit
	// games likely to be binary blobs, versioning required?
	// will probably make the wd app able to keep a secret so its not a public app

	r.Run()
}
