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
	router.HandleFunc("/registry", controller.CreateNewBirthCertificate).Methods("POST")
	router.HandleFunc("/doctorCreateUser", controller.DoctorCreateUser).Methods("POST")
	router.HandleFunc("/deleteUserID/{id}", controller.DeleteUserID).Methods("DELETE")
	//router.HandleFunc("/children/{jmbg}", controller.GetChildren).Methods("GET")
	//router.HandleFunc("/certificate/{jmbg}/{typeOfCertificate}", controller.GetCertificate).Methods("GET")
	//router.HandleFunc("/marriage", controller.Marriage).Methods("POST")
	//router.HandleFunc("/isParent/{jmbg}", controller.IsParent).Methods("GET")
	//router.HandleFunc("/died", controller.UpdateCertificate).Methods("POST")
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

func (controller *RegistrarController) CreateNewBirthCertificate(writer http.ResponseWriter, req *http.Request) {
	var user entity.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Problem to parsing JSON to entity!"))
		return
	}

	value, err := controller.service.CreateNewBirthCertificate(user)
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

//func (controller *RegistrarController) Marriage(writer http.ResponseWriter, req *http.Request) {
//
//	var marriage entity.ExcerptFromTheMarriageRegister
//	err := json.NewDecoder(req.Body).Decode(&marriage)
//	if err != nil {
//		writer.WriteHeader(http.StatusInternalServerError)
//		writer.Write([]byte("Problem to parsing JSON to entity!"))
//		return
//	}
//
//	//find Svedok1 i Svedok2
//	var svedok1 *entity.User
//	var svedok2 *entity.User
//	var mladozenja *entity.User
//	var mlada *entity.User
//
//	mladozenja = controller.service.FindOneUser(marriage.JMBGMladozenje)
//	mlada = controller.service.FindOneUser(marriage.JMBGMlade)
//	svedok1 = controller.service.FindOneUser(marriage.Svedok1.JMBG)
//	svedok2 = controller.service.FindOneUser(marriage.Svedok2.JMBG)
//
//	//kreiranje vencanja je moguce samo ukoliko postoje oba svedoka u bazi
//	if mladozenja == nil {
//		writer.WriteHeader(http.StatusAccepted)
//		writer.Write([]byte("Ne postoji mladozenja u sistemu"))
//		return
//	} else if mlada == nil {
//		writer.WriteHeader(http.StatusAccepted)
//		writer.Write([]byte("Ne postoji mlada u sistemu"))
//		return
//	} else if svedok1 == nil {
//		writer.WriteHeader(http.StatusAccepted)
//		writer.Write([]byte("Ne postoji prvi svedok u sistemu"))
//		return
//	} else if svedok2 == nil {
//		writer.WriteHeader(http.StatusAccepted)
//		writer.Write([]byte("Ne postoji drugi svedok u sistemu"))
//		return
//	}
//
//	marriage.Svedok1 = *svedok1
//	marriage.Svedok2 = *svedok2
//
//	controller.service.CreateNewMarriage(marriage)
//}
//
//func (controller *RegistrarController) UpdateCertificate(writer http.ResponseWriter, req *http.Request) {
//	var userDied entity.UserDied
//	err := json.NewDecoder(req.Body).Decode(&userDied)
//	if err != nil {
//		writer.WriteHeader(http.StatusInternalServerError)
//		writer.Write([]byte("Problem to parsing JSON to entity!"))
//		return
//	}
//
//	err = controller.service.UpdateCertificate(userDied)
//	if err != nil {
//		writer.WriteHeader(http.StatusInternalServerError)
//		writer.Write([]byte(err.Error()))
//		return
//	}
//
//	writer.WriteHeader(http.StatusOK)
//	//writer.Write([]byte("Okej"))
//	//jsonResponse(token, writer)
//}

//func (controller *RegistrarController) GetChildren(writer http.ResponseWriter, req *http.Request) {
//	vars := mux.Vars(req)
//	jmbg, _ := vars["jmbg"]
//
//	//children := controller.service.GetChildren(jmbg)
//	//fmt.Println(children)
//
//	jsonResponse(controller.service.GetChildren(jmbg), writer)
//
//	writer.WriteHeader(http.StatusOK)
//}

//func (controller *RegistrarController) GetCertificate(writer http.ResponseWriter, req *http.Request) {
//	vars := mux.Vars(req)
//	typeStr, _ := vars["typeOfCertificate"]
//	num, err := strconv.Atoi(typeStr)
//	if err != nil {
//		writer.WriteHeader(http.StatusInternalServerError)
//		writer.Write([]byte("Error in convert string to int"))
//	}
//	jmbg, _ := vars["jmbg"]
//
//	one, two, three := controller.service.FindOneCertificateByType(jmbg, num)
//
//	if num == 1 {
//		jsonResponse(one, writer)
//
//	} else if num == 2 {
//		jsonResponse(two, writer)
//
//	} else if num == 3 {
//		jsonResponse(three, writer)
//
//	} else {
//		writer.WriteHeader(http.StatusNotAcceptable)
//		writer.Write([]byte("That type of certificate not exist!"))
//		return
//	}
//
//	writer.WriteHeader(http.StatusOK)
//}

//func (controller *RegistrarController) IsParent(writer http.ResponseWriter, req *http.Request) {
//
//	authToken := req.Header.Get("Authorization")
//	splitted := strings.Split(authToken, " ")
//	claims := authorization.GetMapClaims([]byte(splitted[1]))
//
//	loggedInJMBG := claims["jmbg"]
//
//	vars := mux.Vars(req)
//	jmbgStr, _ := vars["jmbg"]
//
//	user := controller.service.FindOneUser(jmbgStr)
//
//	//dodati prvo proveru ussera
//
//	if user == nil {
//		jsonResponse(false, writer)
//	} else if user.JMBGOca == loggedInJMBG || user.JMBGMajke == loggedInJMBG {
//		jsonResponse(true, writer)
//	} else {
//		jsonResponse(false, writer)
//	}
//}
