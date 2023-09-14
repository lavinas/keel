package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-chi/chi/v5"

	"github.com/lavinas/keel/internal/example/infra/akafka"
	"github.com/lavinas/keel/internal/example/infra/repository"
	"github.com/lavinas/keel/internal/example/infra/web"
	"github.com/lavinas/keel/internal/example/service"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(host.docker.internal:3306)/products")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repository := repository.NewProductRepositoryMysql(db)
	createProductService := service.NewCreateProductService(repository)
	listProductService := service.NewListProductService(repository)

	productHandler := web.NewProductService(createProductService, listProductService)
	r := chi.NewRouter()
	r.Post("/products", productHandler.CreateProductHandler)
	r.Get("/products", productHandler.ListProductHandler)

	go http.ListenAndServe(":8000", r)

	msgChan := make(chan *kafka.Message)
	go akafka.Consume([]string{"products"}, "host.docker.internal:9094", msgChan)

	for msg := range msgChan {
		dto := service.CreateProductInputDto{}
		err := json.Unmarshal(msg.Value, &dto)
		if err != nil {
			panic(err)
		}
		_, err = createProductService.Execute(dto)
		if err != nil {
			panic(err)
		}

	}

}
