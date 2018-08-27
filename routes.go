package main

func initializeRoutes() {
	//router
	router.GET("/", showIndexPage)
	router.GET("/article/view/:article_id", getArticle)
}
