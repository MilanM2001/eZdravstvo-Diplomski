package controller

import (
	"authorization"
	"encoding/json"
	"github.com/casbin/casbin"
	"github.com/cristalhq/jwt/v4"
	"github.com/gorilla/mux"
	"healthcare_service/model"
	"healthcare_service/service"
	"log"
	"net/http"
	"os"
)

type HealthcareController struct {
	service *service.HealthcareService
}

var jwtKey = []byte(os.Getenv("SECRET_KEY"))
var verifier, _ = jwt.NewVerifierHS(jwt.HS256, jwtKey)

func NewHealthcareController(service *service.HealthcareService) *HealthcareController {
	return &HealthcareController{
		service: service,
	}
}

func (controller *HealthcareController) Init(router *mux.Router) {
	authEnforcer, err := casbin.NewEnforcerSafe("./auth_model.conf", "./policy.csv")
	if err != nil {
		log.Fatal(err)
	}

	//Pregled
	router.HandleFunc("/getSviPregledi", controller.GetSviPregledi).Methods("GET")
	router.HandleFunc("/getMojiPreglediLekar", controller.GetMojiPreglediLekar).Methods("GET")
	router.HandleFunc("/getMojiSlobodniPreglediLekar", controller.GetMojiSlobodniPreglediLekar).Methods("GET")
	router.HandleFunc("/getMojiZauzetiPreglediLekar", controller.GetMojiZauzetiPreglediLekar).Methods("GET")
	router.HandleFunc("/getMojiPreglediGradjanin", controller.GetMojiPreglediGradjanin).Methods("GET")
	router.HandleFunc("/getSviSlobodniPregledi", controller.GetSviSlobodniPregledi).Methods("GET")
	router.HandleFunc("/getPreglediByGradjaninID/{id}", controller.GetPreglediByGradjaninID).Methods("GET")
	router.HandleFunc("/getPregledID/{id}", controller.GetPregledID).Methods("GET")
	router.HandleFunc("/postPregled", controller.PostPregled).Methods("POST")
	router.HandleFunc("/zakaziPregled/{id}", controller.ZakaziPregled).Methods("PUT")
	router.HandleFunc("/deletePregledID/{id}", controller.DeletePregledID).Methods("DELETE")
	router.HandleFunc("/deleteAllPregled", controller.DeleteAllPregled).Methods("DELETE")

	//Vakcina
	router.HandleFunc("/getSveVakcine", controller.GetSveVakcine).Methods("GET")
	router.HandleFunc("/getVakcinaID/{id}", controller.GetVakcinaID).Methods("GET")
	router.HandleFunc("/postVakcina", controller.PostVakcina).Methods("POST")
	router.HandleFunc("/putVakcina/{id}", controller.PutVakcina).Methods("PUT")
	router.HandleFunc("/deleteVakcinaID/{id}", controller.DeleteVakcinaID).Methods("DELETE")
	router.HandleFunc("/deleteAllVakcina", controller.DeleteAllVakcina).Methods("DELETE")

	//Alergija
	router.HandleFunc("/getSveAlergije", controller.GetSveAlergije).Methods("GET")
	router.HandleFunc("/getAlergijaID/{id}", controller.GetAlergijaID).Methods("GET")
	router.HandleFunc("/postAlergija", controller.PostAlergija).Methods("POST")
	router.HandleFunc("/deleteAlergijaID/{id}", controller.DeleteAlergijaID).Methods("DELETE")
	router.HandleFunc("/deleteAllAlergija", controller.DeleteAllAlergija).Methods("DELETE")

	//Invaliditet
	router.HandleFunc("/getSveInvaliditete", controller.GetSveInvaliditete).Methods("GET")
	router.HandleFunc("/getInvaliditetID/{id}", controller.GetInvaliditetID).Methods("GET")
	router.HandleFunc("/postInvaliditet", controller.PostInvaliditet).Methods("POST")
	router.HandleFunc("/deleteInvaliditetID/{id}", controller.DeleteInvaliditetID).Methods("DELETE")
	router.HandleFunc("/deleteAllInvaliditet", controller.DeleteAllInvaliditet).Methods("DELETE")

	//Karton
	router.HandleFunc("/getSveKartone", controller.GetSveKartone).Methods("GET")
	router.HandleFunc("/getKartoneJMBG/{jmbg}", controller.GetKartoneJMBG).Methods("GET")
	router.HandleFunc("/getKartonJMBG/{jmbg}", controller.GetKartonJMBG).Methods("GET")
	router.HandleFunc("/putKarton/{jmbg}", controller.PutKarton).Methods("PUT")
	router.HandleFunc("/deleteKartonID/{id}", controller.DeleteKartonID).Methods("DELETE")
	router.HandleFunc("/deleteAllKarton", controller.DeleteAllKarton).Methods("DELETE")

	router.HandleFunc("/getMe", controller.GetMe).Methods("GET")

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8003", authorization.Authorizer(authEnforcer)(router)))
}

//Pregled ------------------------------------------------------------------------------------------------------------------

func (controller *HealthcareController) GetSviPregledi(writer http.ResponseWriter, _ *http.Request) {
	pregledi, err := controller.service.GetSviPregledi()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	writer.WriteHeader(http.StatusOK)
	jsonResponse(pregledi, writer)
}

func (controller *HealthcareController) GetMojiPreglediLekar(writer http.ResponseWriter, req *http.Request) {
	jmbg, err := extractJMBGFromClaims(writer, req)

	pregledi, err := controller.service.GetMojiPreglediLekar(jmbg)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	writer.WriteHeader(http.StatusOK)
	jsonResponse(pregledi, writer)
}

func (controller *HealthcareController) GetMojiSlobodniPreglediLekar(writer http.ResponseWriter, req *http.Request) {
	jmbg, err := extractJMBGFromClaims(writer, req)

	pregledi, err := controller.service.GetMojiSlobodniPreglediLekar(jmbg)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	writer.WriteHeader(http.StatusOK)
	jsonResponse(pregledi, writer)
}

func (controller *HealthcareController) GetMojiZauzetiPreglediLekar(writer http.ResponseWriter, req *http.Request) {
	jmbg, err := extractJMBGFromClaims(writer, req)

	pregledi, err := controller.service.GetMojiZauzetiPreglediLekar(jmbg)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	writer.WriteHeader(http.StatusOK)
	jsonResponse(pregledi, writer)
}

func (controller *HealthcareController) GetMojiPreglediGradjanin(writer http.ResponseWriter, req *http.Request) {
	jmbg, err := extractJMBGFromClaims(writer, req)

	pregledi, err := controller.service.GetMojiPreglediGradjanin(jmbg)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	writer.WriteHeader(http.StatusOK)
	jsonResponse(pregledi, writer)
}

func (controller *HealthcareController) GetSviSlobodniPregledi(writer http.ResponseWriter, _ *http.Request) {
	pregledi, err := controller.service.GetSviSlobodniPregledi()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	jsonResponse(pregledi, writer)
}

func (controller *HealthcareController) GetPreglediByGradjaninID(writer http.ResponseWriter, req *http.Request) {
	objectID, err := getIDFromReqAsPrimitive(writer, req)

	pregledi, err := controller.service.GetPreglediByGradjaninID(objectID)
	if err != nil {
		log.Println("Error finding Appointment By ID")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	writer.WriteHeader(http.StatusOK)
	jsonResponse(pregledi, writer)
}

func (controller *HealthcareController) GetPregledID(writer http.ResponseWriter, req *http.Request) {
	objectID, err := getIDFromReqAsPrimitive(writer, req)

	pregled, err := controller.service.GetPregledID(objectID)
	if err != nil {
		log.Println("Error finding Appointment By ID")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	writer.WriteHeader(http.StatusOK)
	jsonResponse(pregled, writer)
}

func (controller *HealthcareController) PostPregled(writer http.ResponseWriter, req *http.Request) {
	var pregled model.AddPregled
	err := json.NewDecoder(req.Body).Decode(&pregled)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("There is a problem in decoding JSON"))
		log.Println(err)
		return
	}

	jmbg, err := extractJMBGFromClaims(writer, req)

	_, err = controller.service.PostPregled(&pregled, jmbg)
	//if value == 1 {
	//	writer.WriteHeader(http.StatusNotAcceptable)
	//	writer.Write([]byte("Appointment already exists in that time"))
	//	return
	//}
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	writer.WriteHeader(http.StatusOK)
	jsonResponse(pregled, writer)
}

func (controller *HealthcareController) ZakaziPregled(writer http.ResponseWriter, req *http.Request) {
	objectID, err := getIDFromReqAsPrimitive(writer, req)
	jmbg, err := extractJMBGFromClaims(writer, req)

	appointment, err := controller.service.PutPregled(objectID, jmbg)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}

	writer.WriteHeader(http.StatusOK)
	jsonResponse(appointment, writer)
}

func (controller *HealthcareController) DeletePregledID(writer http.ResponseWriter, req *http.Request) {
	objectID, err := getIDFromReqAsPrimitive(writer, req)

	err = controller.service.DeletePregledID(objectID)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}

	writer.WriteHeader(http.StatusOK)
}

//Vakcina ------------------------------------------------------------------------------------------------------------------

func (controller *HealthcareController) GetSveVakcine(writer http.ResponseWriter, req *http.Request) {
	vakcine, err := controller.service.GetSveVakcine()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	writer.WriteHeader(http.StatusOK)
	jsonResponse(vakcine, writer)
}

func (controller *HealthcareController) GetVakcinaID(writer http.ResponseWriter, req *http.Request) {
	objectID, err := getIDFromReqAsPrimitive(writer, req)

	vakcina, err := controller.service.GetVakcinaID(objectID)
	if err != nil {
		log.Println("Error finding Vakcina By ID")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	writer.WriteHeader(http.StatusOK)
	jsonResponse(vakcina, writer)
}

func (controller *HealthcareController) PostVakcina(writer http.ResponseWriter, req *http.Request) {
	var vakcina model.Vakcina
	err := json.NewDecoder(req.Body).Decode(&vakcina)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("There is a problem in decoding JSON"))
		return
	}

	value, err := controller.service.PostVakcina(&vakcina)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	} else if value == 1 {
		writer.WriteHeader(http.StatusNotAcceptable)
		writer.Write([]byte("Vaccine already exists"))
		return
	}

	writer.WriteHeader(http.StatusOK)
	jsonResponse(vakcina, writer)
}

func (controller *HealthcareController) PutVakcina(writer http.ResponseWriter, req *http.Request) {
	var vakcina model.Vakcina
	err := json.NewDecoder(req.Body).Decode(&vakcina)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("There is a problem in decoding JSON"))
		return
	}

	objectID, err := getIDFromReqAsPrimitive(writer, req)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}

	value, err := controller.service.PutVakcina(&vakcina, objectID)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	} else if value == 1 {
		writer.WriteHeader(http.StatusNotAcceptable)
		writer.Write([]byte("Vaccine already exists"))
		return
	}

	writer.WriteHeader(http.StatusOK)
	jsonResponse(vakcina, writer)
}

func (controller *HealthcareController) DeleteVakcinaID(writer http.ResponseWriter, req *http.Request) {
	objectID, err := getIDFromReqAsPrimitive(writer, req)

	err = controller.service.DeleteVakcinaID(objectID)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}

	writer.WriteHeader(http.StatusOK)
}

//Alergija ------------------------------------------------------------------------------------------------------------------

func (controller *HealthcareController) GetSveAlergije(writer http.ResponseWriter, _ *http.Request) {
	alergije, err := controller.service.GetSveAlergije()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	writer.WriteHeader(http.StatusOK)
	jsonResponse(alergije, writer)
}

func (controller *HealthcareController) GetAlergijaID(writer http.ResponseWriter, req *http.Request) {
	objectID, err := getIDFromReqAsPrimitive(writer, req)

	alergija, err := controller.service.GetAlergijaID(objectID)
	if err != nil {
		log.Println("Error finding Appointment By ID")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	writer.WriteHeader(http.StatusOK)
	jsonResponse(alergija, writer)
}

func (controller *HealthcareController) PostAlergija(writer http.ResponseWriter, req *http.Request) {
	var alergija model.Alergija
	err := json.NewDecoder(req.Body).Decode(&alergija)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("There is a problem in decoding JSON"))
		log.Println(err)
		return
	}

	value, err := controller.service.PostAlergija(&alergija)
	if value == 1 {
		writer.WriteHeader(http.StatusNotAcceptable)
		writer.Write([]byte("Appointment already exists in that time"))
		return
	}
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	writer.WriteHeader(http.StatusOK)
	jsonResponse(alergija, writer)
}

func (controller *HealthcareController) DeleteAlergijaID(writer http.ResponseWriter, req *http.Request) {
	objectID, err := getIDFromReqAsPrimitive(writer, req)

	err = controller.service.DeleteAlergijaID(objectID)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}

	writer.WriteHeader(http.StatusOK)
}

//Invaliditet ------------------------------------------------------------------------------------------------------------------

func (controller *HealthcareController) GetSveInvaliditete(writer http.ResponseWriter, _ *http.Request) {
	invaliditeti, err := controller.service.GetSveInvaliditete()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	writer.WriteHeader(http.StatusOK)
	jsonResponse(invaliditeti, writer)
}

func (controller *HealthcareController) GetInvaliditetID(writer http.ResponseWriter, req *http.Request) {
	objectID, err := getIDFromReqAsPrimitive(writer, req)

	invaliditet, err := controller.service.GetInvaliditetID(objectID)
	if err != nil {
		log.Println("Error finding Appointment By ID")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	writer.WriteHeader(http.StatusOK)
	jsonResponse(invaliditet, writer)
}

func (controller *HealthcareController) PostInvaliditet(writer http.ResponseWriter, req *http.Request) {
	var invaliditet model.Invaliditet
	err := json.NewDecoder(req.Body).Decode(&invaliditet)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("There is a problem in decoding JSON"))
		log.Println(err)
		return
	}

	value, err := controller.service.PostInvaliditet(&invaliditet)
	if value == 1 {
		writer.WriteHeader(http.StatusNotAcceptable)
		writer.Write([]byte("Appointment already exists in that time"))
		return
	}
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	writer.WriteHeader(http.StatusOK)
	jsonResponse(invaliditet, writer)
}

func (controller *HealthcareController) DeleteInvaliditetID(writer http.ResponseWriter, req *http.Request) {
	objectID, err := getIDFromReqAsPrimitive(writer, req)

	err = controller.service.DeleteInvaliditetID(objectID)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}

	writer.WriteHeader(http.StatusOK)
}

//Karton ------------------------------------------------------------------------------------------------------------------

func (controller *HealthcareController) GetSveKartone(writer http.ResponseWriter, _ *http.Request) {
	kartoni, err := controller.service.GetSveKartone()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	writer.WriteHeader(http.StatusOK)
	jsonResponse(kartoni, writer)
}

func (controller *HealthcareController) GetKartoneJMBG(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	jmbg, _ := vars["jmbg"]

	kartoni, err := controller.service.GetKartoneJMBG(jmbg)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	writer.WriteHeader(http.StatusOK)
	jsonResponse(kartoni, writer)
}

func (controller *HealthcareController) GetKartonJMBG(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	jmbg, _ := vars["jmbg"]

	karton, err := controller.service.GetKartonJMBG(jmbg)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Print(err)
		return
	}

	writer.WriteHeader(http.StatusOK)
	jsonResponse(karton, writer)
}

func (controller *HealthcareController) PutKarton(writer http.ResponseWriter, req *http.Request) {
	var karton model.Karton
	err := json.NewDecoder(req.Body).Decode(&karton)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("There is a problem in decoding JSON"))
		return
	}
	vars := mux.Vars(req)
	jmbg, _ := vars["jmbg"]

	_, err = controller.service.PutKarton(&karton, jmbg)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	jsonResponse(karton, writer)
}

func (controller *HealthcareController) DeleteKartonID(writer http.ResponseWriter, req *http.Request) {
	objectID, err := getIDFromReqAsPrimitive(writer, req)

	err = controller.service.DeleteKartonID(objectID)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}

	writer.WriteHeader(http.StatusOK)
}

func (controller *HealthcareController) GetMe(writer http.ResponseWriter, req *http.Request) {
	jmbg, err := extractJMBGFromClaims(writer, req)

	user, err := controller.service.GetMe(jmbg)
	if err != nil {
		log.Println("Error getting User")
	}

	writer.WriteHeader(http.StatusOK)
	jsonResponse(user, writer)
}

//---------------------------------------------------------------------------

func (controller *HealthcareController) DeleteAllPregled(writer http.ResponseWriter, req *http.Request) {
	err := controller.service.DeleteAllPregled()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}

	writer.WriteHeader(http.StatusOK)
}

func (controller *HealthcareController) DeleteAllVakcina(writer http.ResponseWriter, req *http.Request) {
	err := controller.service.DeleteAllVakcina()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}

	writer.WriteHeader(http.StatusOK)
}

func (controller *HealthcareController) DeleteAllAlergija(writer http.ResponseWriter, req *http.Request) {
	err := controller.service.DeleteAllAlergija()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}

	writer.WriteHeader(http.StatusOK)
}

func (controller *HealthcareController) DeleteAllInvaliditet(writer http.ResponseWriter, req *http.Request) {
	err := controller.service.DeleteAllInvaliditet()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}

	writer.WriteHeader(http.StatusOK)
}

func (controller *HealthcareController) DeleteAllKarton(writer http.ResponseWriter, req *http.Request) {
	err := controller.service.DeleteAllKarton()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}

	writer.WriteHeader(http.StatusOK)
}
