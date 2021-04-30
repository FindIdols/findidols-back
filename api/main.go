package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/FindIdols/findidols-back/api/handler"
	"github.com/FindIdols/findidols-back/api/middleware"
	"github.com/FindIdols/findidols-back/config"
	"github.com/FindIdols/findidols-back/infrastructure/repository"
	"github.com/FindIdols/findidols-back/usecase/bankaccount"
	"github.com/FindIdols/findidols-back/usecase/idol"
	"github.com/FindIdols/findidols-back/usecase/order"
	"github.com/FindIdols/findidols-back/usecase/pricecontent"
	"github.com/FindIdols/findidols-back/usecase/socialnetworks"
	"github.com/FindIdols/findidols-back/usecase/user"
	"github.com/FindIdols/findidols-back/usecase/video"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	dataSourceName := fmt.Sprintf(
		"host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
		config.DB_HOST,
		config.DB_PORT,
		config.DB_USER,
		config.DB_PASSWORD,
		config.DB_DATABASE,
	)

	db, err := sql.Open("postgres", dataSourceName)

	if err != nil {
		fmt.Println("erro ao conectar no banco")
		log.Fatal(err.Error())
	}

	defer db.Close()

	userRepo := repository.NewUserPostgres(db)
	userService := user.NewService(userRepo)

	orderRepo := repository.NewOrderPostgres(db)
	orderService := order.NewService(orderRepo)

	idolRepo := repository.NewIdolPostgres(db)
	idolService := idol.NewService(idolRepo)

	pricePerContentRepo := repository.NewPriceContentPostgres(db)
	pricePerContentService := pricecontent.NewService(pricePerContentRepo)

	videoRepo := repository.NewVideoPostgres(db)
	videoService := video.NewService(videoRepo)

	socialNetworksRepo := repository.NewSocialNetworksPostgres(db)
	socialNetworksService := socialnetworks.NewService(socialNetworksRepo)

	bankAccountRepo := repository.NewBankAccountPostgres(db)
	bankAccountService := bankaccount.NewService(bankAccountRepo)

	r := mux.NewRouter()

	n := negroni.New(
		negroni.HandlerFunc(middleware.Cors),
		negroni.NewLogger(),
	)

	handler.MakeOrderHandlers(r, *n, orderService, userService)
	handler.MakeIdolHandlers(r, *n, idolService, userService, socialNetworksService, videoService, bankAccountService, pricePerContentService)
	handler.MakeBillingHandlers(r, *n)

	http.Handle("/", r)

	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)

	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":" + strconv.Itoa(config.API_PORT),
		Handler:      context.ClearHandler(http.DefaultServeMux),
		ErrorLog:     logger,
	}

	err = srv.ListenAndServe()

	if err != nil {
		fmt.Println("erro que nao sei")
		log.Fatal(err.Error())
	}
}
