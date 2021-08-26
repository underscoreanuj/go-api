package main

import (
	"os"

	"github.com/underscoreanuj/mux_api/cache"
	"github.com/underscoreanuj/mux_api/controller"
	"github.com/underscoreanuj/mux_api/http"
	"github.com/underscoreanuj/mux_api/repository"
	"github.com/underscoreanuj/mux_api/service"
)

var (
	postRepository repository.PostRepository = repository.NewSQLiteRepository()
	postService    service.PostService       = service.NewPostService(postRepository)
	postCache      cache.PostCache           = cache.NewRedisCache("localhost:6379", 1, 10)
	postController controller.PostController = controller.NewPostController(postService, postCache)
	httpRouter     router.Router             = router.NewChiRouter()
)

func main() {
	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)
	httpRouter.GET("/posts/{id}", postController.GetPostById)

	httpRouter.SERVE(os.Getenv("GO_API_PORT"))
}
