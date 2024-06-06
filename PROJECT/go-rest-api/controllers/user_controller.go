package controllers

import (
    "encoding/json"
    "net/http"
    "go-rest-api/models"
    "go-rest-api/utils"
    "github.com/gorilla/mux"
    "strconv"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
    users, err := models.GetUsers()
    if err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }
    utils.RespondWithJSON(w, http.StatusOK, users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, err := strconv.Atoi(params["id"])
    if err != nil {
        utils.RespondWithError(w, http.StatusBadRequest, "Invalid user ID")
        return
    }

    user, err := models.GetUser(id)
    if err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }
    utils.RespondWithJSON(w, http.StatusOK, user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
    var user models.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
        return
    }

    if err := user.CreateUser(); err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }
    utils.RespondWithJSON(w, http.StatusCreated, user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, err := strconv.Atoi(params["id"])
    if err != nil {
        utils.RespondWithError(w, http.StatusBadRequest, "Invalid user ID")
        return
    }

    var user models.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
        return
    }

    user.ID = id

    if err := user.UpdateUser(); err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }
    utils.RespondWithJSON(w, http.StatusOK, user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, err := strconv.Atoi(params["id"])
    if err != nil {
        utils.RespondWithError(w, http.StatusBadRequest, "Invalid user ID")
        return
    }

    if err := models.DeleteUser(id); err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }
    utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
