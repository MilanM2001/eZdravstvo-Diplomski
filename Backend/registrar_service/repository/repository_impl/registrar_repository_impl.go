package repository_impl

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	domain "registrar_service/model/entity"
	"registrar_service/repository"
)

type RegistrarRepositoryImpl struct {
	user_registry *mongo.Collection //izvodi
	marriage      *mongo.Collection //vencani
}

const (
	DATABASE    = "users"
	COLLECTION  = "user_registry"
	COLLECTION1 = "marriage"
)

func NewRegistrarRepositoryImpl(client *mongo.Client) repository.RegistrarRepository {
	registrar := client.Database(DATABASE).Collection(COLLECTION)
	marriage := client.Database(DATABASE).Collection(COLLECTION1)

	return &RegistrarRepositoryImpl{
		user_registry: registrar,
		marriage:      marriage,
	}
}

func (store *RegistrarRepositoryImpl) GetAllUsers() ([]*domain.User, error) {
	filter := bson.M{}
	return store.filter(filter)
}

func (store *RegistrarRepositoryImpl) GetUserJMBG(jmbg string) (*domain.User, error) {
	filter := bson.M{"jmbg": jmbg}
	return store.filterOne(filter)
}

func (store *RegistrarRepositoryImpl) GetNewbornByMotherJMBG(jmbgMajke string) ([]*domain.User, error) {
	filter := bson.M{"jmbgMajke": jmbgMajke}
	return store.filter(filter)
}

func (store *RegistrarRepositoryImpl) CreateNewBirthCertificate(user domain.User) error {
	if !store.IsUserExist(user.JMBG) {
		_, err := store.user_registry.InsertOne(context.Background(), user)
		if err != nil {
			log.Println("Error in saving User")
			return err
		}
		return nil
	}

	return fmt.Errorf("user already exists")
}

func (store *RegistrarRepositoryImpl) DoctorCreateUser(user domain.User) error {
	if !store.IsUserExist(user.JMBG) {
		_, err := store.user_registry.InsertOne(context.Background(), user)
		if err != nil {
			log.Println("Error in saving User")
			return err
		}
		return nil
	}

	return fmt.Errorf("user already exists")
}

func (store *RegistrarRepositoryImpl) IsUserExist(jmbg string) bool {

	user, err := store.filterOne(bson.M{"jmbg": jmbg})
	if err != nil {
		log.Println(err.Error())
		return false
	}

	if user != nil {
		return true
	} else {
		return false
	}

}

func (store *RegistrarRepositoryImpl) filterOne(filter interface{}) (user *domain.User, err error) {
	result := store.user_registry.FindOne(context.TODO(), filter)
	err = result.Decode(&user)
	return
}

func (store *RegistrarRepositoryImpl) filter(filter interface{}) ([]*domain.User, error) {
	cursor, err := store.user_registry.Find(context.Background(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decodeUser(cursor)
}

func (store *RegistrarRepositoryImpl) FindOneUser(jmbg string) *domain.User {
	user, err := store.filterOne(bson.M{"jmbg": jmbg})
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	return user
}

func (store *RegistrarRepositoryImpl) FindOneUserID(id primitive.ObjectID) (*domain.User, error) {
	user, err := store.filterOne(bson.M{"_id": id})
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return user, nil
}

func (store *RegistrarRepositoryImpl) CreateNewMarriage(marriage domain.ExcerptFromTheMarriageRegister) {
	_, err := store.marriage.InsertOne(context.Background(), marriage)
	if err != nil {
		log.Printf("Error in RegistrarRepositoryImpl CreateNewMarriage(): %s", err.Error())
		return
	}
}

func (store *RegistrarRepositoryImpl) UpdateCertificate(user domain.User) error {

	log.Println(user.ID)

	update := bson.M{
		"$set": bson.M{
			//"Preminuo":   user.Preminuo,
			//"DatimSmrti": user.DatimSmrti,
			//"MestoSmrti": user.MestoSmrti,
			"Preminuo":   "",
			"DatimSmrti": "",
			"MestoSmrti": "",
		},
	}

	filter := bson.M{"_id": user.ID}

	_, err := store.user_registry.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Printf("Error in RegistrarRepositoryImpl UpdateOne(): %s", err.Error())
		return err
	}
	return nil
}

func (store *RegistrarRepositoryImpl) GetChildren(jmbg string, pol domain.Pol) []domain.User {

	var filter interface{}

	if pol == "Muski" {
		filter = bson.M{"JMBGOca": jmbg}
	} else if pol == "Zenski" {
		filter = bson.M{"JMBGMajke": jmbg}

	}
	fmt.Printf("JMBG: %s\nPol: %s\nfilter: %s\n", jmbg, pol, filter)

	users, err := store.user_registry.Find(context.Background(), filter)
	if err != nil {
		return nil
	}

	var children []domain.User

	// loop through the documents
	for users.Next(context.Background()) {
		var result domain.User
		err := users.Decode(&result)
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}

		children = append(children, result)
		// do something with the result
		fmt.Println(result)

	}

	return children
}

func decodeUser(cursor *mongo.Cursor) (users []*domain.User, err error) {
	for cursor.Next(context.Background()) {
		var user domain.User
		err = cursor.Decode(&user)
		if err != nil {
			return
		}
		users = append(users, &user)
	}
	err = cursor.Err()
	return
}
