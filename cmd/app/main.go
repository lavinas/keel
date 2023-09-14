package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"encoding/json"
	"net/http"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-chi/chi/v5"

	"github.com/lavinas/keel/internal/example/infra/akafka"
	"github.com/lavinas/keel/internal/example/infra/repository"
	"github.com/lavinas/keel/internal/example/infra/web"
	"github.com/lavinas/keel/internal/example/usecase"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(host.docker.internal:3306)/products")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repository := repository.NewProductRepositoryMysql(db)
	createProductUseCase := usecase.NewCreateProductUseCase(repository)
	listProductUseCase := usecase.NewListProductUseCase(repository)

	productHandler := web.NewProductUseCase(createProductUseCase, listProductUseCase)
	r := chi.NewRouter()
	r.Post("/products", productHandler.CreateProductHandler)
	r.Get("/products", productHandler.ListProductHandler)

	go http.ListenAndServe(":8000", r)

	msgChan := make(chan *kafka.Message)
	go akafka.Consume([]string{"products"}, "host.docker.internal:9094", msgChan)

	for msg := range msgChan {
		dto := usecase.CreateProductInputDto{}
		err := json.Unmarshal(msg.Value, &dto)
		if err != nil {
			panic(err)
		}
		_, err = createProductUseCase.Execute(dto)
		if err != nil {
			panic(err)
		}

	}

}
