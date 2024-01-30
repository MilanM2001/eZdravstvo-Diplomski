package controller

import (
	"authorization"
	"encoding/json"
	"fmt"
	"github.com/casbin/casbin"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"registrar_service/model/entity"
	"registrar_service/service"
)

type RegistrarController struct {
	service *service.RegistrarService
}

func NewRegistrarController(service *service.RegistrarService) *RegistrarController {
	return &RegistrarController{
		service: service,
	}
}

func (controller *RegistrarController) Init(router *mux.Router) {
	authEnforcer, err := casbin.NewEnforcerSafe("./auth_model.conf", "./policy.csv")
	if err != nil {
		log.Fatal(err)
	}

	router.HandleFunc("/allUsers", controller.GetAllUsers).Methods("GET")
	router.HandleFunc("/getUserJMBG/{jmbg}", controller.GetUserJMBG).Methods("GET")
	router.HandleFunc("/getUserID/{id}", controller.GetUserID).Methods("GET")
	router.HandleFunc("/getNewbornsByMotherJMBG/{jmbg}", controller.GetNewbornByMotherJMBG).Methods("GET")
	router.HandleFunc("/registry", controller.CreateNewUser).Methods("POST")
	router.HandleFunc("/doctorCreateUser", controller.DoctorCreateUser).Methods("POST")
	router.HandleFunc("/parentCreateUser", controller.ParentCreateUser).Methods("POST")
	router.HandleFunc("/deleteUserID/{id}", controller.DeleteUserID).Methods("DELETE")
	router.HandleFunc("/deleteAllUsers", controller.DeleteAllUsers).Methods("DELETE")
	router.HandleFunc("/postPotvrdaSmrti", controller.PostPotvrdaSmrti).Methods("POST")
	router.HandleFunc("/allPotvrdeSmrti", controller.GetAllPotvrdeSmrti).Methods("GET")
	router.HandleFunc("/deletePotvrdaSmrtiID/{id}", controller.DeletePotvrdaSmrtiID).Methods("DELETE")
	router.HandleFunc("/isPotvrdaExist/{jmbg}", controller.IsPotvrdaExistJMBG).Methods("GET")
	router.HandleFunc("/getPotvrdaSmrtiJMBG/{jmbg}", controller.GetPotvrdaSmrtiJMBG).Methods("GET")
	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(":8001", authorization.Authorizer(authEnforcer)(router)))
}

func (controller *RegistrarController) GetAllUsers(writer http.ResponseWriter, req *http.Request) {
	users, err := controller.service.GetAllUsers()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	jsonResponse(users, writer)
	writer.WriteHeader(http.StatusOK)
}

func (controller *RegistrarController) GetUserJMBG(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	jmbg, _ := vars["jmbg"]

	user, err := controller.service.GetUserJMBG(jmbg)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	jsonResponse(user, writer)
	writer.WriteHeader(http.StatusOK)
}

func (controller *RegistrarController) GetUserID(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, _ := vars["id"]

	objectID, err := primitive.ObjectIDFromHex(id)

	user, err := controller.service.FindOneUserID(objectID)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	jsonResponse(user, writer)
	writer.WriteHeader(http.StatusOK)
}

func (controller *RegistrarController) CreateNewUser(writer http.ResponseWriter, req *http.Request) {
	var user entity.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Problem to parsing JSON to entity!"))
		return
	}

	value, err := controller.service.CreateNewUser(user)
	if value == 1 {
		writer.WriteHeader(http.StatusAccepted)
		writer.Write([]byte("JMBG already exist in system!"))
		return
	}
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}

	jsonResponse(user, writer)
	writer.WriteHeader(http.StatusOK)
}

func (controller *RegistrarController) DoctorCreateUser(writer http.ResponseWriter, req *http.Request) {
	var user entity.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		fmt.Printf("Error decoding JSON: %s\n", err)
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Problem to parsing JSON to entity!"))
		return
	}

	value, err := controller.service.DoctorCreateUser(&user)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	} else if value == 1 {
		writer.WriteHeader(http.StatusConflict)
		writer.Write([]byte("JMBG Majke ne postoji u sistemu"))
		return
	} else if value == 2 {
		writer.WriteHeader(http.StatusForbidden)
		writer.Write([]byte("JMBG ne pripada zenskoj osobi"))
		return
	}

	jsonResponse(user, writer)
	writer.WriteHeader(http.StatusOK)
}

func (controller *RegistrarController) ParentCreateUser(writer http.ResponseWriter, req *http.Request) {
	var user entity.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		fmt.Printf("Error decoding JSON: %s\n", err)
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Problem to parsing JSON to entity!"))
		return
	}

	value, err := controller.service.ParentCreateUser(&user)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	} else if value == 1 {
		writer.WriteHeader(http.StatusConflict)
		writer.Write([]byte("JMBG Oca ne postoji u sistemu"))
		return
	} else if value == 2 {
		writer.WriteHeader(http.StatusForbidden)
		writer.Write([]byte("JMBG ne pripada muskoj osobi"))
		return
	}

	jsonResponse(user, writer)
	writer.WriteHeader(http.StatusOK)
}

func (controller *RegistrarController) GetNewbornByMotherJMBG(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	jmbg, _ := vars["jmbg"]

	newborns, err := controller.service.GetNewbornByMotherJMBG(jmbg)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Errror in getting newborns"))
		return
	}

	jsonResponse(newborns, writer)
	writer.WriteHeader(http.StatusOK)
}

func (controller *RegistrarController) DeleteUserID(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, _ := vars["id"]

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Convert to Primitive error")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = controller.service.DeleteUserID(objectID)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}

	writer.WriteHeader(http.StatusOK)
}

func (controller *RegistrarController) DeleteAllUsers(writer http.ResponseWriter, req *http.Request) {
	err := controller.service.DeleteAllUsers()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}

	writer.WriteHeader(http.StatusOK)
}

func (controller *RegistrarController) GetAllPotvrdeSmrti(writer http.ResponseWriter, req *http.Request) {
	potvrde, err := controller.service.GetAllPotvrdeSmrti()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	jsonResponse(potvrde, writer)
	writer.WriteHeader(http.StatusOK)
}

func (controller *RegistrarController) GetPotvrdaSmrtiJMBG(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	jmbg, _ := vars["jmbg"]

	potvrda, err := controller.service.GetPotvrdaSmrtiJMBG(jmbg)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	jsonResponse(potvrda, writer)
	writer.WriteHeader(http.StatusOK)
}

func (controller *RegistrarController) IsPotvrdaExistJMBG(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	jmbg, _ := vars["jmbg"]
	isExist := controller.service.IsPotvrdaExist(jmbg)

	jsonResponse(isExist, writer)
	writer.WriteHeader(http.StatusOK)
}

func (controller *RegistrarController) PostPotvrdaSmrti(writer http.ResponseWriter, req *http.Request) {
	var potvrda entity.PotvrdaSmrti
	err := json.NewDecoder(req.Body).Decode(&potvrda)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Problem to parsing JSON to entity!"))
		return
	}

	value, err := controller.service.PostPotvrdaSmrti(potvrda)
	if value == 1 {
		writer.WriteHeader(http.StatusAccepted)
		writer.Write([]byte("Potvrda vec postoji u sistemu"))
		return
	}
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}

	jsonResponse(potvrda, writer)
	writer.WriteHeader(http.StatusOK)
}

func (controller *RegistrarController) DeletePotvrdaSmrtiID(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, _ := vars["id"]

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Convert to Primitive error")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = controller.service.DeletePotvrdaSmrtiID(objectID)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}

	writer.WriteHeader(http.StatusOK)
}
