package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/FindIdols/findidols-back/api/presenter"
	"github.com/FindIdols/findidols-back/entity"
	"github.com/FindIdols/findidols-back/usecase/bankaccount"
	"github.com/FindIdols/findidols-back/usecase/idol"
	"github.com/FindIdols/findidols-back/usecase/socialnetworks"
	"github.com/FindIdols/findidols-back/usecase/user"
	"github.com/FindIdols/findidols-back/usecase/video"
	"github.com/codegangsta/negroni"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func createIdol(
	idolService idol.UseCase,
	userService user.UseCase,
	socialNetworksService socialnetworks.UseCase,
	bankAccountService bankaccount.UseCase,
) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error adding idol"

		var input struct {
			Email        string  `json:"email"`
			Phone        string  `json:"phone"`
			FirstName    string  `json:"first_name"`
			LastName     string  `json:"last_name"`
			Genre        string  `json:"genre"`
			Category     string  `json:"category"`
			ArtisticName string  `json:"artistic_name"`
			Profession   string  `json:"profession"`
			Description  string  `json:"description"`
			Value        float64 `json:"value"`
			Deadline     int16   `json:"deadline"`
			BankName     string  `json:"bankName"`
			TypeAccount  string  `json:"accountType"`
			Agency       string  `json:"agency"`
			Operation    string  `json:"operation"`
			Account      string  `json:"account"`
			Digit        string  `json:"digit"`
			Youtube      string  `json:"youtube"`
			Instagram    string  `json:"instagram"`
			Twitter      string  `json:"twitter"`
			TikTok       string  `json:"tiktok"`
			Secret       string  `json:"secret"`
		}

		err := json.NewDecoder(r.Body).Decode(&input)

		if err != nil {
			fmt.Println("erro decode idol")
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		if input.Secret != "extroida" {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("invalid_secret"))
			return
		}

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

		socialNetworksID, err := socialNetworksService.CreateSocialNetworks(
			input.Youtube,
			input.Instagram,
			input.Twitter,
			input.TikTok,
		)

		if err != nil {
			fmt.Println("erro criacao idol")
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		bankAccountID, err := bankAccountService.CreateBankAccount(
			input.BankName,
			input.TypeAccount,
			input.Agency,
			input.Operation,
			input.Account,
			input.Digit,
		)

		if err != nil {
			fmt.Println("erro criacao idol")
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		_, err = idolService.CreateIdol(
			input.ArtisticName,
			input.Profession,
			input.Description,
			input.Value,
			input.Deadline,
			user.ID.String(),
			socialNetworksID.String(),
			bankAccountID.String(),
		)

		if err != nil {
			fmt.Println("erro criacao idol")
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		toJ := &presenter.Idol{
			ID:           user.ID,
			ArtisticName: input.ArtisticName,
		}

		w.WriteHeader(http.StatusCreated)

		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

func getIdol(service idol.UseCase, socialNetworksService socialnetworks.UseCase, videoService video.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		errorMessage := "Error reading idol"

		vars := mux.Vars(r)

		id, err := entity.StringToID(vars["id"])

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		dataIdol, err := service.GetIdol(id)

		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}

		videos, err := videoService.GetVideos(dataIdol.ID)

		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}

		socialUUID, err := uuid.Parse(dataIdol.SocialNetworksID)
		dataSocialNetworks, err := socialNetworksService.GetSocialNetworks(socialUUID)

		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}

		w.Header().Set("Content-Type", "application/json")

		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}

		toJ := &presenter.IdolInformation{
			ID:           dataIdol.ID,
			ArtisticName: dataIdol.ArtisticName,
			Description:  dataIdol.Description,
			Profession:   dataIdol.Profession,
			Value:        dataIdol.Value,
			Deadline:     dataIdol.Deadline,
			DenyContent:  dataIdol.DenyContent,
			Youtube:      dataSocialNetworks.Youtube,
			Instagram:    dataSocialNetworks.Instagram,
			Twitter:      dataSocialNetworks.Twitter,
			Tiktok:       dataSocialNetworks.TikTok,
			VideoURL:     videos,
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
	})
}

func getIdols(service idol.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		errorMessage := "Error reading idols"

		var data []*entity.Idol
		var err error

		data, err = service.GetIdols()

		w.Header().Set("Content-Type", "application/json")

		if err != nil && err != entity.ErrNotFound {
			errorMessage = "erro generico"
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		if err != nil {
			errorMessage = "erro not found idols"
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}

		var toJ []*presenter.Idol

		for _, d := range data {
			toJ = append(toJ, &presenter.Idol{
				ID:           d.ID,
				ArtisticName: d.ArtisticName,
				Description:  d.Description,
				Value:        d.Value,
			})
		}

		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
	})
}

func endpointTest() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Funciona a API!")
	})
}

//MakeIdolHandlers make url handlers
func MakeIdolHandlers(
	r *mux.Router,
	n negroni.Negroni,
	IdolService idol.UseCase,
	UserService user.UseCase,
	SocialNetworksService socialnetworks.UseCase,
	VideoService video.UseCase,
	bankAccountService bankaccount.UseCase,
) {
	r.Handle("/", n.With(
		negroni.Wrap(endpointTest()),
	)).Methods("GET", "OPTIONS").Name("endpointTest")

	r.Handle("/v1/idol", n.With(
		negroni.Wrap(createIdol(IdolService, UserService, SocialNetworksService, bankAccountService)),
	)).Methods("POST", "OPTIONS").Name("createIdol")

	r.Handle("/v1/idol/{id}", n.With(
		negroni.Wrap(getIdol(IdolService, SocialNetworksService, VideoService)),
	)).Methods("GET", "OPTIONS").Name("getIdol")

	r.Handle("/v1/idols", n.With(
		negroni.Wrap(getIdols(IdolService)),
	)).Methods("GET", "OPTIONS").Name("getIdols")
}
