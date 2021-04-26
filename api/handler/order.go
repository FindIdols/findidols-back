package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/FindIdols/findidols-back/api/presenter"
	"github.com/FindIdols/findidols-back/entity"
	"github.com/FindIdols/findidols-back/usecase/order"
	"github.com/FindIdols/findidols-back/usecase/user"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func createOrder(orderService order.UseCase, userService user.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error adding order"

		var input struct {
			Email       string `json:"email"`
			Phone       string `json:"phone"`
			FirstName   string `json:"firstName"`
			LastName    string `json:"lastName"`
			Genre       string `json:"genre"`
			Category    string `json:"category"`
			Usage       string `json:"usage"`
			Subject     string `json:"subject"`
			Instruction string `json:"instruction"`
			TermsOfUse  bool   `json:"termsOfUse"`
			IdolID      string `json:"idolId"`
		}

		err := json.NewDecoder(r.Body).Decode(&input)

		if err != nil {
			fmt.Println("erro decode")
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		fmt.Println(input)

		user, err := userService.Create(
			input.FirstName,
			input.LastName,
			input.Email,
			input.Phone,
			input.Genre,
			input.Category,
		)

		if err != nil {
			fmt.Println("erro criacao user")
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		order, err := orderService.CreateOrder(
			user,
			input.Usage,
			input.Subject,
			input.Instruction,
			input.TermsOfUse,
			input.IdolID,
		)

		if err != nil {
			fmt.Println("erro criacao order")
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		toJ := &presenter.Order{
			ID:          user.ID,
			OrderNumber: order.OrderNumber,
		}

		w.WriteHeader(http.StatusCreated)

		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

func getOrder(service order.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading order again"

		vars := mux.Vars(r)

		id, err := entity.StringToID(vars["id"])

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		data, err := service.GetOrder(id)

		w.Header().Set("Content-Type", "application/json")

		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}

		toJ := &presenter.Order{
			ID:          data.ID,
			OrderNumber: data.OrderNumber,
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
	})
}

//MakeOrderHandlers make url handlers
func MakeOrderHandlers(r *mux.Router, n negroni.Negroni, OrderService order.UseCase, UserService user.UseCase) {
	r.Handle("/v1/order", n.With(
		negroni.Wrap(createOrder(OrderService, UserService)),
	)).Methods("POST", "OPTIONS").Name("createOrder")

	r.Handle("/v1/order/{id}", n.With(
		negroni.Wrap(getOrder(OrderService)),
	)).Methods("GET", "OPTIONS").Name("getOrder")

}
