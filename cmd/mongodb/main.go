package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/TobiasGleiter/go-databases/internal/data/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type application struct {
	logger *slog.Logger
	users  models.UserModelInterface
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "mongodb://localhost:27017", "MongoDB data source name")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	app := &application{
		logger: logger,
		users:  &models.MongoUserModel{DB: db},
	}

	srv := &http.Server{
		Addr:    *addr,
		Handler: app.routes(),
	}

	logger.Info("starting server", "addr", srv.Addr)

	err = srv.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(uri string) (*mongo.Client, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	serverAPI.SetStrict(false)
	serverAPI.SetDeprecationErrors(false)

	// credential := options.Credential{
	//	AuthMechanism: "SCRAM-SHA-256",
	// 	AuthSource: "<authenticationDb>",
	// 	Username:   "<username>",
	// 	Password:   "<password>",
	// }

	opts := options.Client().ApplyURI(uri) //.SetAuth(credential)
	opts.SetServerAPIOptions(serverAPI)
	//opts.SetTLSConfig(&tls.Config{})
	opts.SetConnectTimeout(10 * time.Second)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	var result bson.M
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	return client, nil
}
