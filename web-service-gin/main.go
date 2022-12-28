package main

import "example/web-service-gin/api"

func main() {
	router := api.Configure()

	_ = router.Run("localhost:8080")
}
