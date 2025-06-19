package app

import (
	"GonIO/internal/dal"
	"GonIO/internal/handlers"
	"GonIO/internal/service"
	"log"
	"log/slog"
	"net/http"
	"os"

	httpswagger "github.com/swaggo/http-swagger"
)

func SetHandler() *http.ServeMux {
	mux := http.NewServeMux()
	SetSwagger(mux)

	objectDal := dal.NewObjectCSVRepo()
	objectServ := service.NewObjectService(*objectDal)
	objectHandler := handlers.NewObjectHandler(objectServ)

	bucketDal := dal.NewBucketXMLRepo()
	bucketServ := service.NewBucketService(*bucketDal)
	bucketHandler := handlers.NewBucketHandler(bucketServ)

	healthHandler := handlers.NewHealthHandler()

	mux.HandleFunc("GET /PING", healthHandler.Ping) // Healthcheck

	mux.HandleFunc("GET /", bucketHandler.BucketListsHandler)                 // Bucket list
	mux.HandleFunc("PUT /{BucketName}", bucketHandler.CreateBucketHandler)    // Create bucket
	mux.HandleFunc("DELETE /{BucketName}", bucketHandler.DeleteBucketHandler) // Delete bucket

	mux.HandleFunc("GET /{BucketName}", objectHandler.GetObjectList)               // Get object list in bucket
	mux.HandleFunc("GET /{BucketName}/{ObjectKey}", objectHandler.RetrieveObject)  // Retrieve an object
	mux.HandleFunc("PUT /{BucketName}/{ObjectKey}", objectHandler.UpdateObject)    // Upload an object
	mux.HandleFunc("DELETE /{BucketName}/{ObjectKey}", objectHandler.DeleteObject) // Delete an object

	return mux
}

func SetSwagger(mux *http.ServeMux) {
	swaggerBytes, err := os.ReadFile("docs/swagger.json")
	if err != nil {
		log.Fatal("Failed to read swagger file: ", err)
	}

	mux.HandleFunc("GET /api/bro/docs/swagger/", httpswagger.Handler(
		httpswagger.URL("/docs/swagger.json"),
	))

	mux.HandleFunc("GET  /docs/swagger.json", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		if _, err := writer.Write(swaggerBytes); err != nil {
			slog.Error("Failed to send swagger file: ", "error", err)
			writer.WriteHeader(http.StatusInternalServerError)
		}
	})
}
