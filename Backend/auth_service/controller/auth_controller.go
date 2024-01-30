package controller

import (
	domain "auth_service/model/entity"
	"auth_service/service"
	"authorization"
	"encoding/json"
	"fmt"
	"github.com/casbin/casbin"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
)

type AuthController struct {
	service *service.AuthService
}

func NewAuthController(service *service.AuthService) *AuthController {
	return &AuthController{
		service: service,
	}
}

func (controller *AuthController) Init(router *mux.Router) {
	authEnforcer, err := casbin.NewEnforcerSafe("./auth_model.conf", "./policy.csv")
	if err != nil {
		log.Fatal(err)
	}

	router.HandleFunc("/allCredentials", controller.GetAllCredentials).Methods("GET")
	router.HandleFunc("/registration", controller.Register).Methods("POST")
	router.HandleFunc("/login", controller.Login).Methods("POST")
	router.HandleFunc("/deleteCredentialsID/{id}", controller.DeleteCredentialsID).Methods("DELETE")
	router.HandleFunc("/deleteAllCredentials", controller.DeleteAllCredentials).Methods("DELETE")

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8002", authorization.Authorizer(authEnforcer)(router)))
}

func (controller *AuthController) GetAllCredentials(writer http.ResponseWriter, req *http.Request) {
	credentials, err := controller.service.GetAllCredentials()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	jsonResponse(credentials, writer)
	writer.WriteHeader(http.StatusOK)
}

func (controller *AuthController) Register(writer http.ResponseWriter, req *http.Request) {
	var credentials domain.Credentials
	err := json.NewDecoder(req.Body).Decode(&credentials)
	fmt.Println(credentials)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("There is problem in decoding JSON"))
		return
	}

	value, err := controller.service.Register(credentials)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println(value)
	if value == -1 {
		writer.WriteHeader(http.StatusAccepted)
		writer.Write([]byte("JMBG je vec registrovan!"))
		return
	} else if value == -2 {
		writer.WriteHeader(http.StatusCreated)
		writer.Write([]byte("JMBG nije pronadjen u izvodima rodjenih lica!"))
		return
	}
}

func (controller *AuthController) Login(writer http.ResponseWriter, req *http.Request) {
	var credentials domain.Credentials
	err := json.NewDecoder(req.Body).Decode(&credentials)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("There is problem in decoding JSON"))
		return
	}

	token, value := controller.service.Login(credentials.JMBG, credentials.Password)
	if value == 1 {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("JMBG does not exist"))
		return
	} else if value == 2 {
		writer.WriteHeader(http.StatusUnauthorized)
		writer.Write([]byte("Incorrect password"))
		return
	} else if value == 3 {
		writer.WriteHeader(http.StatusConflict)
		writer.Write([]byte("User died"))
		return
	} else if value == 4 {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Problem with generating token"))
		return
	} else if value == 5 {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("NATS Error"))
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte(token))
}

func (controller *AuthController) DeleteCredentialsID(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, _ := vars["id"]

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Convert to Primitive error")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = controller.service.DeleteCredentialsID(objectID)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}

	writer.WriteHeader(http.StatusOK)
}

func (controller *AuthController) DeleteAllCredentials(writer http.ResponseWriter, req *http.Request) {
	err := controller.service.DeleteAllCredentials()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}

	writer.WriteHeader(http.StatusOK)
}
