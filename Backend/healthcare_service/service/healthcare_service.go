package service

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"healthcare_service/model"
	"healthcare_service/repository"
	"log"
	"os"
	"time"
)

type HealthcareService struct {
	repository     repository.HealthcareRepository
	natsConnection *nats.Conn
}

func NewHealthcareService(repository repository.HealthcareRepository, natsConnection *nats.Conn) *HealthcareService {
	return &HealthcareService{
		repository:     repository,
		natsConnection: natsConnection,
	}
}

func (service *HealthcareService) GetSviPregledi() ([]*model.Pregled, error) {
	return service.repository.GetSviPregledi()
}

//Pregled ------------------------------------------------------------------------------------------------------------------

func (service *HealthcareService) GetMojiPreglediLekar(jmbg string) ([]*model.Pregled, error) {
	dataToSend, err := json.Marshal(jmbg)
	if err != nil {
		log.Println("Error Marshaling JMBG")
	}

	response, err := service.natsConnection.Request(os.Getenv("GET_USER_BY_JMBG"), dataToSend, 5*time.Second)

	var doctor model.User
	err = json.Unmarshal(response.Data, &doctor)
	if err != nil {
		log.Println("Error in Unmarshalling json")
		return nil, err
	}

	doctorID := doctor.ID

	return service.repository.GetMojiPreglediLekar(doctorID)
}

func (service *HealthcareService) GetMojiSlobodniPreglediLekar(jmbg string) ([]*model.Pregled, error) {
	dataToSend, err := json.Marshal(jmbg)
	if err != nil {
		log.Println("Error Marshaling JMBG")
	}

	response, err := service.natsConnection.Request(os.Getenv("GET_USER_BY_JMBG"), dataToSend, 5*time.Second)

	var doctor model.User
	err = json.Unmarshal(response.Data, &doctor)
	if err != nil {
		log.Println("Error in Unmarshalling json")
		return nil, err
	}

	doctorID := doctor.ID

	return service.repository.GetMojiSlobodniPreglediLekar(doctorID)
}

func (service *HealthcareService) GetMojiZauzetiPreglediLekar(jmbg string) ([]*model.Pregled, error) {
	dataToSend, err := json.Marshal(jmbg)
	if err != nil {
		log.Println("Error Marshaling JMBG")
	}

	response, err := service.natsConnection.Request(os.Getenv("GET_USER_BY_JMBG"), dataToSend, 5*time.Second)

	var doctor model.User
	err = json.Unmarshal(response.Data, &doctor)
	if err != nil {
		log.Println("Error in Unmarshalling json")
		return nil, err
	}

	doctorID := doctor.ID

	return service.repository.GetMojiZauzetiPreglediLekar(doctorID)
}

func (service *HealthcareService) GetMojiPreglediGradjanin(jmbg string) ([]*model.Pregled, error) {
	dataToSend, err := json.Marshal(jmbg)
	if err != nil {
		log.Println("Error Marshaling JMBG")
	}

	response, err := service.natsConnection.Request(os.Getenv("GET_USER_BY_JMBG"), dataToSend, 5*time.Second)

	var gradjanin model.User
	err = json.Unmarshal(response.Data, &gradjanin)
	if err != nil {
		log.Println("Error in Unmarshalling json")
		return nil, err
	}

	gradjaninID := gradjanin.ID

	return service.repository.GetMojiPreglediGradjanin(gradjaninID)
}

func (service *HealthcareService) GetSviSlobodniPregledi() ([]*model.Pregled, error) {
	return service.repository.GetSviSlobodniPregledi()
}

func (service *HealthcareService) GetPregledID(id primitive.ObjectID) (*model.Pregled, error) {
	return service.repository.GetPregledID(id)
}

func (service *HealthcareService) GetPreglediByGradjaninID(id primitive.ObjectID) ([]*model.Pregled, error) {
	return service.repository.GetPreglediByGradjaninID(id)
}

//																		PRE konekcije sa nats da se stavi

func (service *HealthcareService) PostPregled(addPregled *model.AddPregled, jmbg string) (int, error) {
	var pregled model.Pregled
	var vakcina *model.Vakcina

	dataToSend, err := json.Marshal(jmbg)
	if err != nil {
		log.Println("Error Marshaling JMBG")
	}

	existingPregledi, err := service.GetMojiSlobodniPreglediLekar(jmbg)
	for _, existingPregled := range existingPregledi {
		if (existingPregled.PocetakPregleda >= addPregled.PocetakPregleda && existingPregled.PocetakPregleda <= addPregled.ZavrsetakPregleda) ||
			(existingPregled.ZavrsetakPregleda >= addPregled.PocetakPregleda && existingPregled.ZavrsetakPregleda <= addPregled.ZavrsetakPregleda) ||
			(existingPregled.PocetakPregleda >= addPregled.PocetakPregleda && existingPregled.ZavrsetakPregleda <= addPregled.ZavrsetakPregleda) {
			return 1, nil
		}
	}
	if err != nil {
		log.Println("Error getting all Appointments", err)
		return 0, err
	}

	response, err := service.natsConnection.Request(os.Getenv("GET_USER_BY_JMBG"), dataToSend, 5*time.Second)
	if err != nil {
		log.Println("Error requesting user by JMBG:", err)
		return 0, err
	}

	var lekar model.User
	if err := json.Unmarshal(response.Data, &lekar); err != nil {
		log.Println("Error in Unmarshalling json:", err)
		return 0, err
	}

	if addPregled.VakcinaID != "" {
		vakcinaID, err := primitive.ObjectIDFromHex(addPregled.VakcinaID)
		if err != nil {
			log.Println("Convert to Primitive error:", err)
			return 0, err
		}

		vakcina, err = service.repository.GetVakcinaID(vakcinaID)
		if err != nil {
			log.Println("Error getting vaccine by ID:", err)
			return 0, err
		}
	} else {
		vakcina = nil
	}

	pocetakTime := time.Unix(addPregled.PocetakPregleda, 0)
	zavrsetakTime := time.Unix(addPregled.ZavrsetakPregleda, 0)

	appointmentDuration := 20 * time.Minute

	for current := pocetakTime; current.Before(zavrsetakTime); current = current.Add(appointmentDuration) {
		pregled.ID = primitive.NewObjectID()
		pregled.Lekar = &lekar
		pregled.Gradjanin = nil
		pregled.PocetakPregleda = current.Unix()
		pregled.ZavrsetakPregleda = current.Add(appointmentDuration).Unix()
		pregled.TipPregleda = addPregled.TipPregleda
		if vakcina != nil {
			pregled.Vakcina = vakcina
		} else {
			pregled.Vakcina = nil
		}

		err := service.repository.PostPregled(&pregled)
		if err != nil {
			log.Println("Error in trying to save Pregled:", err)
			return 0, err
		}
	}

	return 0, nil
}

func (service *HealthcareService) PutPregled(id primitive.ObjectID, jmbg string) (*model.Pregled, error) {
	dataToSend, err := json.Marshal(jmbg)
	if err != nil {
		log.Println("Error Marshaling JMBG")
	}

	response, err := service.natsConnection.Request(os.Getenv("GET_USER_BY_JMBG"), dataToSend, 5*time.Second)

	var gradjanin model.User
	err = json.Unmarshal(response.Data, &gradjanin)
	if err != nil {
		log.Println("Error in Unmarshalling json")
		return nil, err
	}

	pregled, err := service.repository.GetPregledID(id)
	if err != nil {
		log.Println("Error in finding Pregled By ID")
		return nil, err
	}

	pregled.Gradjanin = &gradjanin

	err = service.repository.PutPregled(pregled)
	if err != nil {
		log.Println("Error in Updating Appointment")
		return nil, err
	}

	return pregled, nil
}

func (service *HealthcareService) DeletePregledID(id primitive.ObjectID) error {
	return service.repository.DeletePregledID(id)
}

//Vakcina ------------------------------------------------------------------------------------------------------------------

func (service *HealthcareService) GetSveVakcine() ([]*model.Vakcina, error) {
	return service.repository.GetSveVakcine()
}

func (service *HealthcareService) GetVakcinaID(id primitive.ObjectID) (*model.Vakcina, error) {
	return service.repository.GetVakcinaID(id)
}

func (service *HealthcareService) PostVakcina(vakcina *model.Vakcina) (int, error) {
	vakcina.ID = primitive.NewObjectID()

	existingVakcina, _ := service.repository.GetVakcinaNaziv(vakcina.Naziv)
	if existingVakcina != nil {
		return 1, nil
	}

	err := service.repository.PostVakcina(vakcina)
	if err != nil {
		log.Println("Error in trying to save Vakcina")
		return 0, err
	}

	return 0, nil
}

func (service *HealthcareService) PutVakcina(vakcina *model.Vakcina, id primitive.ObjectID) (int, error) {
	updateVakcina, err := service.repository.GetVakcinaID(id)
	if err != nil {
		log.Println("Error in trying to update Vakcina")
		return 0, err
	}

	existingVakcina, _ := service.repository.GetVakcinaNaziv(vakcina.Naziv)
	if existingVakcina != nil {
		return 1, nil
	}

	updateVakcina.Naziv = vakcina.Naziv
	updateVakcina.Kompanija = vakcina.Kompanija

	err = service.repository.PutVakcina(updateVakcina)
	if err != nil {
		log.Println("Error in trying to save Vakcina")
		return 0, err
	}

	return 0, nil
}

func (service *HealthcareService) DeleteVakcinaID(id primitive.ObjectID) error {
	return service.repository.DeleteVakcinaID(id)
}

//Alergija ------------------------------------------------------------------------------------------------------------------

func (service *HealthcareService) GetSveAlergije() ([]*model.Alergija, error) {
	return service.repository.GetSveAlergije()
}

func (service *HealthcareService) GetAlergijaID(id primitive.ObjectID) (*model.Alergija, error) {
	return service.repository.GetAlergijaID(id)
}

func (service *HealthcareService) PostAlergija(alergija *model.Alergija) (int, error) {
	alergija.ID = primitive.NewObjectID()

	existingAlergija, _ := service.repository.GetAlergijaNaziv(alergija.Naziv)
	if existingAlergija != nil {
		return 1, nil
	}

	err := service.repository.PostAlergija(alergija)
	if err != nil {
		log.Println("Error in trying to save Alergija")
		return 0, err
	}

	return 0, nil
}

func (service *HealthcareService) DeleteAlergijaID(id primitive.ObjectID) error {
	return service.repository.DeleteAlergijaID(id)
}

//Invaliditet ------------------------------------------------------------------------------------------------------------------

func (service *HealthcareService) GetSveInvaliditete() ([]*model.Invaliditet, error) {
	return service.repository.GetSveInvaliditete()
}

func (service *HealthcareService) GetInvaliditetID(id primitive.ObjectID) (*model.Invaliditet, error) {
	return service.repository.GetInvaliditetID(id)
}

func (service *HealthcareService) PostInvaliditet(invaliditet *model.Invaliditet) (int, error) {
	invaliditet.ID = primitive.NewObjectID()

	existingInvaliditet, _ := service.repository.GetInvaliditetNaziv(invaliditet.Naziv)
	if existingInvaliditet != nil {
		return 1, nil
	}

	err := service.repository.PostInvaliditet(invaliditet)
	if err != nil {
		log.Println("Error in trying to save Invaliditet")
		return 0, err
	}

	return 0, nil
}

func (service *HealthcareService) DeleteInvaliditetID(id primitive.ObjectID) error {
	return service.repository.DeleteInvaliditetID(id)
}

//Invaliditet ------------------------------------------------------------------------------------------------------------------

func (service *HealthcareService) GetSveKartone() ([]*model.Karton, error) {
	return service.repository.GetSveKartone()
}

func (service *HealthcareService) GetKartoneJMBG(jmbg string) ([]*model.Karton, error) {
	return service.repository.GetKartoneJMBG(jmbg)
}

func (service *HealthcareService) GetKartonJMBG(jmbg string) (*model.Karton, error) {
	return service.repository.GetKartonJMBG(jmbg)
}

func (service *HealthcareService) PutKarton(karton *model.Karton, jmbg string) (int, error) {
	updateKarton, err := service.repository.GetKartonJMBG(jmbg)
	if err != nil {
		log.Println("Error in trying to update Karton")
		return 0, err
	}

	updateKarton = karton

	err = service.repository.PutKarton(updateKarton)
	if err != nil {
		log.Println("Error in trying to save Karton")
		return 0, err
	}

	return 0, nil
}

func (service *HealthcareService) DeleteKartonID(id primitive.ObjectID) error {
	return service.repository.DeleteKartonID(id)
}

func (service *HealthcareService) DeleteKartonJMBG(jmbg string) error {
	return service.repository.DeleteKartonJMBG(jmbg)
}

//------------------------------------------------------------------------------------------------------------------------------------------

func (service *HealthcareService) GetMe(jmbg string) (*model.User, error) {
	dataToSend, err := json.Marshal(jmbg)
	if err != nil {
		log.Println("Error Marshaling JMBG")
	}

	response, err := service.natsConnection.Request(os.Getenv("GET_USER_BY_JMBG"), dataToSend, 5*time.Second)

	var user model.User
	err = json.Unmarshal(response.Data, &user)
	if err != nil {
		log.Println("Error in Unmarshalling json")
		return nil, err
	}

	return &user, nil
}

func (service *HealthcareService) SubscribeToNats(natsConnection *nats.Conn) {

	_, err := natsConnection.QueueSubscribe(os.Getenv("POST_KARTON"), "queue-healthcare-group", func(message *nats.Msg) {
		var karton model.Karton
		err := json.Unmarshal(message.Data, &karton)
		if err != nil {
			log.Println("Error in unmarshal JSON!")
			return
		}

		karton.ID = primitive.NewObjectID()
		err = service.repository.PostKarton(karton)
		if err != nil {
			log.Println("Error in Nats")
			return
		}

		dataToSend, err := json.Marshal(karton)
		if err != nil {
			log.Println("Error in marshaling JSON!")
			return
		}

		reply := dataToSend
		err = natsConnection.Publish(message.Reply, reply)
		if err != nil {
			log.Printf("Error in publish response: %s", err.Error())
			return
		}

	})
	if err != nil {
		log.Printf("Error in receiving message: %s", err.Error())
	}

	log.Printf("Subscribed to channel: %s", os.Getenv("POST_KARTON"))
}

//-------------------------------------------
func (service *HealthcareService) DeleteAllPregled() error {
	return service.repository.DeleteAllPregled()
}
func (service *HealthcareService) DeleteAllVakcina() error {
	return service.repository.DeleteAllVakcina()
}
func (service *HealthcareService) DeleteAllAlergija() error {
	return service.repository.DeleteAllAlergija()
}
func (service *HealthcareService) DeleteAllInvaliditet() error {
	return service.repository.DeleteAllInvaliditet()
}
func (service *HealthcareService) DeleteAllKarton() error {
	return service.repository.DeleteAllKarton()
}
