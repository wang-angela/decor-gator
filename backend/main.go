package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/decor-gator/backend/pkg/configs"
	"github.com/decor-gator/backend/pkg/controllers"
	"github.com/decor-gator/backend/pkg/routes"
	"github.com/decor-gator/backend/pkg/utils"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	folder := "test-pictures"

	controllers.Init()
	controllers.CreateBucket()

	for _, bucket := range controllers.ListBuckets().Buckets {
		fmt.Println(*bucket.Name)
	}

	files, _ := ioutil.ReadDir(folder)
	for _, file := range files {
		if file.IsDir() {
			continue
		} else {
			controllers.UploadObject(folder + "/" + file.Name())
		}
	}

	fmt.Println(controllers.ListObjects())

	for _, object := range controllers.ListObjects().Contents {
		controllers.GetObject(*object.Key)
		controllers.DeleteObject(*object.Key)
	}

	fmt.Println(controllers.ListObjects())

	err := godotenv.Load()
	if err != nil {
		log.Print("Error loading .env")
	}

	r := mux.NewRouter()

	// Connect database
	configs.ConnectDB()
	utils.InitDB("data")

	// Routes
	routes.UserRoutes(r)
	routes.PostRoutes(r)
	routes.ImageRoutes(r)
	routes.EmailRoutes(r)
	routes.JwtRoutes(r)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
