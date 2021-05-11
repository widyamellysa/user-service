package main

import (
	"digileaps/user/config"
	"digileaps/user/entities"
	"digileaps/user/model"

	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/user/findall", FindAllAPI).Methods("GET")
	r.HandleFunc("/api/user/find/{id}", FindAPI).Methods("GET")
	r.HandleFunc("/api/user/create", CreateAPI).Methods("POST")
	r.HandleFunc("/api/user/update", UpdateAPI).Methods("PUT")
	r.HandleFunc("/api/user/delete/{id}", FindAPI).Methods("DELETE")
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		fmt.Println(err)
	}
}

func FindAllAPI(w http.ResponseWriter, r *http.Request) {
	db, err := config.GetMongoDB()
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	} else {
		userModel := model.UserModel{
			Db:         db,
			Collection: "users",
		}
		user, err2 := userModel.FindAll()
		if err2 != nil {
			respondWithError(w, http.StatusBadRequest, err2.Error())
			return
		} else {
			respondWithJson(w, http.StatusOK, user)
		}
	}
}

func FindAPI(w http.ResponseWriter, r *http.Request) {
	db, err := config.GetMongoDB()
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	} else {
		userModel := model.UserModel{
			Db:         db,
			Collection: "users",
		}
		vars := mux.Vars(r)
		id := vars["_id"]
		user, err2 := userModel.Find(id)
		if err2 != nil {
			respondWithError(w, http.StatusBadRequest, err2.Error())
			return
		} else {
			respondWithJson(w, http.StatusOK, user)
		}
	}
}

func CreateAPI(w http.ResponseWriter, r *http.Request) {
	db, err := config.GetMongoDB()
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	} else {
		userModel := model.UserModel{
			Db:         db,
			Collection: "users",
		}
		var user entities.User
		user.ID = bson.NewObjectId()
		err2 := json.NewDecoder(r.Body).Decode(&user)
		if err2 != nil {
			respondWithError(w, http.StatusBadRequest, err2.Error())
			return
		} else {
			err3 := userModel.Create(&user)
			if err3 != nil {
				respondWithError(w, http.StatusBadRequest, err3.Error())
				return
			} else {
				respondWithJson(w, http.StatusOK, user)
			}
		}
	}
}

func UpdateAPI(w http.ResponseWriter, r *http.Request) {
	db, err := config.GetMongoDB()
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	} else {
		userModel := model.UserModel{
			Db:         db,
			Collection: "users",
		}
		var user entities.User
		err2 := json.NewDecoder(r.Body).Decode(&user)
		if err2 != nil {
			respondWithError(w, http.StatusBadRequest, err2.Error())
			return
		} else {
			err3 := userModel.Update(&user)
			if err3 != nil {
				respondWithError(w, http.StatusBadRequest, err3.Error())
				return
			} else {
				respondWithJson(w, http.StatusOK, user)
			}
		}
	}
}

func DeleteAPI(w http.ResponseWriter, r *http.Request) {
	db, err := config.GetMongoDB()
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	} else {
		userModel := model.UserModel{
			Db:         db,
			Collection: "users",
		}
		vars := mux.Vars(r)
		id := vars["_id"]
		user, _ := userModel.Find(id)
		err2 := userModel.Delete(user)
		if err2 != nil {
			respondWithError(w, http.StatusBadRequest, err2.Error())
			return
		} else {
			respondWithJson(w, http.StatusOK, user)
		}
	}
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
