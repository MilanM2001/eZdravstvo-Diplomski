package service

import (
	domain "auth_service/model/entity"
	"auth_service/repository"
	"encoding/json"
	"github.com/cristalhq/jwt/v4"
	"github.com/nats-io/nats.go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"time"
)

type AuthService struct {
	store          repository.AuthRepository
	natsConnection *nats.Conn
}

func NewAuthService(store repository.AuthRepository, natsConnection *nats.Conn) *AuthService {
	return &AuthService{
		store:          store,
		natsConnection: natsConnection,
	}
}

func (service *AuthService) GetAllCredentials() ([]*domain.Credentials, error) {
	return service.store.GetAllCredentials()
}

func (service *AuthService) IsJMBGUnique(jmbg string) bool {
	return service.store.IsJMBGUnique(jmbg)
}

func (service *AuthService) Register(credentials domain.Credentials) (int, error) {
	dataToSend, err := json.Marshal(credentials)

	response, err := service.natsConnection.Request(os.Getenv("CHECK_USER_JMBG"), dataToSend, 5*time.Second)

	var isJMBGExist bool
	err = json.Unmarshal(response.Data, &isJMBGExist)
	if err != nil {
		log.Println("Error in unmarshal json")
		return 0, err
	}

	if isJMBGExist {
		isExists := service.IsJMBGUnique(credentials.JMBG)
		if isExists == true {
			return -1, nil
		}

		credentials.ID = primitive.NewObjectID()                                                       //creating unique UUID for MongoDB
		password, err := bcrypt.GenerateFromPassword([]byte(credentials.Password), bcrypt.DefaultCost) //hashing password
		credentials.Password = string(password)
		if err != nil {
			return 0, err
		}
		service.store.Register(credentials)

		var karton domain.Karton
		karton.JMBG = credentials.JMBG
		karton.Alergije = []domain.Alergija{}
		karton.Invaliditeti = []domain.Invaliditet{}

		kartonToSend, err := json.Marshal(karton)
		if err != nil {
			log.Println("Error in Marshaling JSON!")
			return 0, err
		}

		_, err = service.natsConnection.Request(os.Getenv("POST_KARTON"), kartonToSend, 5*time.Second)
		if err != nil {
			log.Println(err)
			return 0, err
		}

		return 0, nil
	} else {
		return -2, nil
	}

}

func (service *AuthService) Login(jmbg string, password string) (string, int) {
	credentials, err := service.store.GetCredentials(jmbg)
	if err != nil {
		log.Println(err)
		return "", 1
	}

	err = bcrypt.CompareHashAndPassword([]byte(credentials.Password), []byte(password))
	if err != nil {
		log.Println(err)
		return "", 2
	}

	tokenString, err := GenerateJWT(credentials)
	if err != nil {
		return "", 3
	}

	return tokenString, 0
}

func (service *AuthService) DeleteCredentialsID(id primitive.ObjectID) error {
	return service.store.DeleteCredentialsID(id)
}

func (service *AuthService) DeleteAllCredentials() error {
	return service.store.DeleteAllCredentials()
}

func (service *AuthService) DeleteCredentialsJMBG(jmbg string) error {
	return service.store.DeleteCredentialsJMBG(jmbg)
}

func (service *AuthService) SubscribeToNats(natsConnection *nats.Conn) {
	_, err := natsConnection.QueueSubscribe(os.Getenv("DELETE_CREDENTIALS"), "queue-auth-group", func(message *nats.Msg) {
		var jmbg string
		err := json.Unmarshal(message.Data, &jmbg)
		if err != nil {
			log.Println("Error in unamrshal JSON")
			return
		}

		err = service.DeleteCredentialsJMBG(jmbg)
		if err != nil {
			log.Println(err)
			return
		}

		dataToSend, err := json.Marshal(jmbg)
		if err != nil {
			log.Println("Error in marshaling JSON!")
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
		log.Printf("Error in publishing response: %s", err.Error())
		return
	}

	log.Printf("Error in receiving message: %s", err.Error())
}

func GenerateJWT(credentials *domain.Credentials) (string, error) {
	key := []byte(os.Getenv("SECRET_KEY"))
	signer, err := jwt.NewSignerHS(jwt.HS256, key)
	if err != nil {
		log.Println(err)
	}

	builder := jwt.NewBuilder(signer)

	claims := &domain.Claims{
		UserID:    credentials.ID,
		JMBG:      credentials.JMBG,
		Role:      credentials.UserType,
		ExpiresAt: time.Now().Add(time.Minute * 60),
	}

	token, err := builder.Build(claims)
	if err != nil {
		log.Println(err)
	}

	return token.String(), nil
}
