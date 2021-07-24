package handler

//go get -u github.com/gin-gonic/gin
import (
	"github.com/gin-gonic/gin"
)

type Handler struct {

}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group(relativePath: "/auth")
	{
		auth.POST(relativePath: "/sign-up")
		auth.POST(relativePath: "/sign-in")
	}

	/*api := router.Group(relativePath: "/api"){

	}*/
}