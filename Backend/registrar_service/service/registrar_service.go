package service

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"os"
	"registrar_service/model/entity"
	"registrar_service/repository"
)

type RegistrarService struct {
	store          repository.RegistrarRepository
	natsConnection *nats.Conn
}

func NewRegistrarService(store repository.RegistrarRepository, natsConnection *nats.Conn) *RegistrarService {
	return &RegistrarService{
		store:          store,
		natsConnection: natsConnection,
	}
}

func (service *RegistrarService) GetAllUsers() ([]*entity.User, error) {
	return service.store.GetAllUsers()
}

func (service *RegistrarService) GetUserJMBG(jmbg string) (*entity.User, error) {
	return service.store.GetUserJMBG(jmbg)
}

func (service *RegistrarService) FindOneUserID(id primitive.ObjectID) (*entity.User, error) {
	return service.store.FindOneUserID(id)
}

func (service *RegistrarService) CreateNewUser(user entity.User) (int, error) {
	isExist := service.store.IsUserExist(user.JMBG)
	if isExist {
		return 1, nil
	}

	user.ID = primitive.NewObjectID()
	user.JMBGOca = ""
	user.JMBGMajke = ""
	err := service.store.CreateNewUser(user)
	if err != nil {
		return 0, err
	}
	return 0, nil
}

func (service *RegistrarService) DoctorCreateUser(user *entity.User) (int, error) {
	mother, err := service.store.GetUserJMBG(user.JMBGMajke)
	if mother == nil {
		return 1, nil
	} else if mother.Pol == "Muski" {
		return 2, nil
	}

	user.ID = primitive.NewObjectID()
	user.Ime = ""
	user.Prezime = ""
	user.JMBGOca = ""
	user.JMBG = ""

	//JMBGMajke, Pol, DatumRodjenja i MestoRodjenja se unose

	err = service.store.DoctorCreateUser(user)
	if err != nil {
		return 0, err
	}
	return 0, nil
}

func (service *RegistrarService) ParentCreateUser(user *entity.User) (int, error) {
	father, err := service.store.GetUserJMBG(user.JMBGOca)
	if father == nil {
		return 1, nil
	} else if father.Pol == "Zenski" {
		return 2, nil
	}

	err = service.store.ParentCreateUser(user)
	if err != nil {
		return 0, err
	}
	return 0, nil
}

func (service *RegistrarService) GetNewbornByMotherJMBG(jmbgMajke string) ([]*entity.User, error) {
	return service.store.GetNewbornByMotherJMBG(jmbgMajke)
}

func (service *RegistrarService) GetChildrenByParentJMBG(jmbg string, pol string) ([]*entity.User, error) {
	return service.store.GetChildrenByParentJMBG(jmbg, pol)
}

func (service *RegistrarService) DeleteUserID(id primitive.ObjectID) error {
	_, err := service.store.FindOneUserID(id)
	if err != nil {
		return err
	}

	return service.store.DeleteUserID(id)
}

func (service *RegistrarService) DeleteAllUsers() error {
	return service.store.DeleteAllUsers()
}

func (service *RegistrarService) DeleteAllPotvrde() error {
	return service.store.DeleteAllPotvrde()
}

func (service *RegistrarService) FindOneUser(jmbg string) *entity.User {
	return service.store.FindOneUser(jmbg)
}

func (service *RegistrarService) GetAllPotvrdeSmrti() ([]*entity.PotvrdaSmrti, error) {
	return service.store.GetAllPotvrdeSmrti()
}

func (service *RegistrarService) GetPotvrdaSmrtiJMBG(jmbg string) (*entity.PotvrdaSmrti, error) {
	return service.store.GetPotvrdaSmrtiJMBG(jmbg)
}

func (service *RegistrarService) IsPotvrdaExist(jmbg string) bool {
	return service.store.IsPotvrdaExist(jmbg)
}

func (service *RegistrarService) PostPotvrdaSmrti(potvrda entity.PotvrdaSmrti) (int, error) {
	potvrda.ID = primitive.NewObjectID()

	//existingPotvrda, err := service.store.GetPotvrdaSmrtiJMBG(potvrda.JMBG)
	//if err != nil {
	//	log.Println("Error in trying to get potvrda smrti")
	//	return 0, err
	//}
	//if existingPotvrda != nil {
	//	return 1, nil
	//}

	err := service.store.PostPotvrdaSmrti(potvrda)
	if err != nil {
		log.Println("Error in trying to save Potvrda")
		return 0, err
	}

	return 0, nil
}

func (service *RegistrarService) DeletePotvrdaSmrtiID(id primitive.ObjectID) error {
	return service.store.DeleteUserID(id)
}

func (service *RegistrarService) SubscribeToNats(natsConnection *nats.Conn) {

	_, err := natsConnection.QueueSubscribe(os.Getenv("CHECK_USER_JMBG"), "queue-registrar-group", func(message *nats.Msg) {
		var credentials entity.Credentials
		err := json.Unmarshal(message.Data, &credentials)
		if err != nil {
			log.Println("Error in unmarshal JSON!")
			return
		}

		isExist := service.store.IsUserExist(credentials.JMBG)

		dataToSend, err := json.Marshal(isExist)
		if err != nil {
			log.Println("Error in marshaling json")
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
		return
	}

	log.Printf("Subscribed to channel: %s", os.Getenv("CHECK_USER_JMBG"))

	_, err = natsConnection.QueueSubscribe(os.Getenv("CHECK_POTVRDA_SMRTI_JMBG"), "queue-registrar-group", func(message *nats.Msg) {
		var jmbg string
		err := json.Unmarshal(message.Data, &jmbg)
		if err != nil {
			log.Println("Error in unmarshal JSON")
			return
		}

		isExist := service.IsPotvrdaExist(jmbg)

		dataToSend, err := json.Marshal(isExist)
		if err != nil {
			log.Println("Error in marshaling json")
			return
		}
		reply := dataToSend
		err = natsConnection.Publish(message.Reply, reply)
		if err != nil {
			log.Printf("Error in publishing response: %s", err.Error())
			return
		}

	})
	if err != nil {
		log.Printf("Error in receiving message: %s", err.Error())
	}

	log.Printf("Subscribed to channel: %s", os.Getenv("CHECK_POTVRDA_SMRTI_JMBG"))

	_, err = natsConnection.QueueSubscribe(os.Getenv("GET_USER_BY_JMBG"), "queue-registrar-group", func(message *nats.Msg) {
		var jmbg string
		err := json.Unmarshal(message.Data, &jmbg)
		if err != nil {
			log.Println("Error in unmarshal JSON!")
			return
		}

		user := service.FindOneUser(jmbg)

		dataToSend, err := json.Marshal(user)
		if err != nil {
			log.Println("Error in marshaling json")
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

	log.Printf("Subscribed to channel: %s", os.Getenv("GET_USER_BY_JMBG"))
}
