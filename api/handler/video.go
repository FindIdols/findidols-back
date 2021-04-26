package handler

import (
	"encoding/json"
	"net/http"

	"github.com/FindIdols/findidols-back/api/presenter"
	"github.com/FindIdols/findidols-back/entity"
	"github.com/FindIdols/findidols-back/usecase/idol"
	"github.com/FindIdols/findidols-back/usecase/video"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func getVideo(videoService video.UseCase, idolService idol.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading video"

		vars := mux.Vars(r)

		id, err := entity.StringToID(vars["id"])

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		data, err := videoService.Get(id)

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

		toJ := &presenter.Video{
			ID:       data.ID,
			VideoURL: data.URL,
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
	})
}

// func uploadVideo(videoService video.UseCase, idolService idol.UseCase) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		errorMessage := "Error adding idol"

// 		r.ParseMultipartForm(10 << 20)

// 		file, handler, err := r.FormFile("myFile")
// 		fmt.Println(r.FormValue("palavra"))

// 		if err != nil {
// 			fmt.Printlnln("Erro receber file")
// 			return
// 		}

// 		defer file.Close()

// 		fmt.Printlnf("File name: %+v\n", handler.Filename)
// 		fmt.Printlnf("File zie: %+v\n", handler.Size)

// 		fileBytes, err := ioutil.ReadAll(file)

// 		if err != nil {
// 			fmt.Printlnln("Erro ler file")
// 			return
// 		}

// 		err = os.Mkdir("../storage/videos/joao_cleber", 0755)

// 		if err != nil {
// 			fmt.Printlnln(err)
// 		}

// 		err = ioutil.WriteFile("../storage/videos/s/teste.mp4", fileBytes, 0644)

// 		if err != nil {
// 			teste := err.Error() == "comer"
// 			fmt.Println(teste)
// 			return
// 		}

// 		var input struct {
// 			IdolID string `json:"idol_id"`
// 		}

// 		err = json.NewDecoder(r.Body).Decode(&input)

// 		if err != nil {
// 			fmt.Println("erro decode idol")
// 			log.Println(err.Error())
// 			w.WriteHeader(http.StatusInternalServerError)
// 			w.Write([]byte(errorMessage))
// 			return
// 		}

// 		videoID, err := videoService.UploadVideo(
// 			input.IdolID,
// 		)

// 		if err != nil {
// 			fmt.Println("erro criacao idol")
// 			log.Println(err.Error())
// 			w.WriteHeader(http.StatusInternalServerError)
// 			w.Write([]byte(errorMessage))
// 			return
// 		}

// 		toJ := &presenter.Video{
// 			ID: videoID,
// 		}

// 		w.WriteHeader(http.StatusCreated)

// 		if err := json.NewEncoder(w).Encode(toJ); err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			w.Write([]byte(errorMessage))
// 			return
// 		}
// 	})
// }

//MakeVideoHandlers make url handlers
func MakeVideoHandlers(r *mux.Router, n negroni.Negroni, VideoService video.UseCase, IdolService idol.UseCase) {
	r.Handle("/v1/video", n.With(
		negroni.Wrap(getVideo(VideoService, IdolService)),
	)).Methods("GET", "OPTIONS").Name("getVideo")

	// r.Handle("/v1/idol/{id}", n.With(
	// 	negroni.Wrap(getIdol(IdolService)),
	// )).Methods("GET", "OPTIONS").Name("getIdol")

}
