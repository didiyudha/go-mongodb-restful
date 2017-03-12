package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/didiyudha/go-mongodb-restful/models"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2/bson"
)

// UsersController struc
type UsersController struct{}

var usr *models.User

func init() {
	usr = models.NewUser()
}

// NewUserController return pointer to UsersController struct
func NewUserController() *UsersController {
	return &UsersController{}
}

// SaveUser handler for saving user into database
func (uc *UsersController) SaveUser(wr http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	wr.Header().Set("Content-Type", "application/json")
	body := req.Body
	user := models.NewUser()
	err := json.NewDecoder(body).Decode(user)
	checkErr(wr, err)
	err = user.Save()
	checkErr(wr, err)
}

// GetUsers grab all users
func (uc *UsersController) GetUsers(wr http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	users, err := usr.FindAll()
	checkErr(wr, err)
	wr.Header().Set("Content-Type", "application/json")
	json.NewEncoder(wr).Encode(users)
}

// FindUser find a user by ID
func (uc *UsersController) FindUser(wr http.ResponseWriter, req *http.Request, prm httprouter.Params) {
	ID := prm.ByName("id")
	var user = models.User{}
	if !bson.IsObjectIdHex(ID) {
		errorResponse(wr, "Invalid Object ID", http.StatusBadRequest)
		return
	}
	oID := bson.ObjectIdHex(ID)
	user, er := usr.FindByID(oID)
	if user.ID == "" {
		errorResponse(wr, "Data Not Found", http.StatusNotFound)
		return
	}
	if er != nil {
		errorResponse(wr, er.Error(), http.StatusInternalServerError)
		return
	}
	uJSON, er := json.Marshal(user)
	if er != nil {
		errorResponse(wr, er.Error(), http.StatusInternalServerError)
		return
	}
	jsonResponse(wr, uJSON, http.StatusOK)
}

// UpdateUser update a user from mongodb
func (uc *UsersController) UpdateUser(wr http.ResponseWriter, req *http.Request, prm httprouter.Params) {
	ID := prm.ByName("id")
	if !bson.IsObjectIdHex(ID) {
		errorResponse(wr, "Invalid object ID", http.StatusBadRequest)
		return
	}
	oID := bson.ObjectIdHex(ID)
	usr := models.User{}
	body := req.Body
	err := json.NewDecoder(body).Decode(&usr)
	if err != nil {
		errorResponse(wr, err.Error(), http.StatusInternalServerError)
		return
	}
	body.Close()
	usr.Update(oID)
	wr.WriteHeader(http.StatusOK)
}

// DeleteUser delete a user from mongodb
func (uc *UsersController) DeleteUser(wr http.ResponseWriter, req *http.Request, prm httprouter.Params) {
	ID := prm.ByName("id")
	if !bson.IsObjectIdHex(ID) {
		errorResponse(wr, "Invalid Object ID", http.StatusBadRequest)
		return
	}
	oID := bson.ObjectIdHex(ID)
	user, er := usr.FindByID(oID)
	if user.ID == "" {
		errorResponse(wr, "Data not found", http.StatusNotFound)
		return
	}
	if er != nil {
		errorResponse(wr, er.Error(), http.StatusInternalServerError)
		return
	}
	er = user.Delete(oID)
	if er != nil {
		errorResponse(wr, er.Error(), http.StatusInternalServerError)
		return
	}
	wr.WriteHeader(http.StatusOK)
}
