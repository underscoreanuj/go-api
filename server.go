package main

import (
	"os"

	"github.com/underscoreanuj/mux_api/controller"
	"github.com/underscoreanuj/mux_api/http"
	"github.com/underscoreanuj/mux_api/repository"
	"github.com/underscoreanuj/mux_api/service"
)

var (
	postRepository repository.PostRepository = repository.NewSQLiteRepository()
	postService    service.PostService       = service.NewPostService(postRepository)
	postController controller.PostController = controller.NewPostController(postService)
	httpRouter     router.Router             = router.NewChiRouter()
)

func main() {
	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)

	httpRouter.SERVE(os.Getenv("GO_API_PORT"))
}
