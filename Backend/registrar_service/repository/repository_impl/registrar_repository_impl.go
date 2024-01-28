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
	userRegistry *mongo.Collection
	potvrdaSmrti *mongo.Collection
}

const (
	DATABASE               = "users"
	CollectionUserRegistry = "user_registry"
	CollectionPotvrdaSmrti = "potvrda_smrti"
)

func NewRegistrarRepositoryImpl(client *mongo.Client) repository.RegistrarRepository {
	registrar := client.Database(DATABASE).Collection(CollectionUserRegistry)
	potvrdaOSmrti := client.Database(DATABASE).Collection(CollectionPotvrdaSmrti)

	return &RegistrarRepositoryImpl{
		userRegistry: registrar,
		potvrdaSmrti: potvrdaOSmrti,
	}
}

func (store *RegistrarRepositoryImpl) GetAllUsers() ([]*domain.User, error) {
	filter := bson.M{}
	return store.filterUserRegistries(filter)
}

func (store *RegistrarRepositoryImpl) GetUserJMBG(jmbg string) (*domain.User, error) {
	filter := bson.M{"jmbg": jmbg}
	return store.filterOneUserRegistry(filter)
}

func (store *RegistrarRepositoryImpl) GetNewbornByMotherJMBG(jmbgMajke string) ([]*domain.User, error) {
	filter := bson.M{"jmbgMajke": jmbgMajke, "ime": "", "prezime": ""}
	return store.filterUserRegistries(filter)
}

func (store *RegistrarRepositoryImpl) CreateNewUser(user domain.User) error {
	if !store.IsUserExist(user.JMBG) {
		_, err := store.userRegistry.InsertOne(context.Background(), user)
		if err != nil {
			log.Println("Error in saving User")
			return err
		}
		return nil
	}

	return fmt.Errorf("user already exists")
}

func (store *RegistrarRepositoryImpl) DoctorCreateUser(user *domain.User) error {
	_, err := store.userRegistry.InsertOne(context.Background(), user)
	if err != nil {
		log.Println("Error in saving User")
		return err
	}

	return nil
}
func (store *RegistrarRepositoryImpl) ParentCreateUser(user *domain.User) error {
	filter := bson.M{"_id": user.ID}
	update := bson.D{{"$set", user}}
	_, err := store.userRegistry.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Println("Updating User Error MongoDB", err.Error())
		return err
	}

	return nil
}

func (store *RegistrarRepositoryImpl) DeleteUserID(id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	_, err := store.userRegistry.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	return nil
}

func (store *RegistrarRepositoryImpl) DeleteAllUsers() error {
	filter := bson.D{}
	_, err := store.userRegistry.DeleteMany(context.Background(), filter)
	if err != nil {
		return err
	}

	return nil
}

func (store *RegistrarRepositoryImpl) IsUserExist(jmbg string) bool {
	user, err := store.filterOneUserRegistry(bson.M{"jmbg": jmbg})
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

func (store *RegistrarRepositoryImpl) FindOneUser(jmbg string) *domain.User {
	user, err := store.filterOneUserRegistry(bson.M{"jmbg": jmbg})
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	return user
}

func (store *RegistrarRepositoryImpl) FindOneUserID(id primitive.ObjectID) (*domain.User, error) {
	user, err := store.filterOneUserRegistry(bson.M{"_id": id})
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return user, nil
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

	_, err := store.userRegistry.UpdateOne(context.Background(), filter, update)
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

	users, err := store.userRegistry.Find(context.Background(), filter)
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

func (store *RegistrarRepositoryImpl) filterUserRegistries(filter interface{}) ([]*domain.User, error) {
	cursor, err := store.userRegistry.Find(context.Background(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}

	return decodeUserRegistry(cursor)
}

func (store *RegistrarRepositoryImpl) filterOneUserRegistry(filter interface{}) (user *domain.User, err error) {
	result := store.userRegistry.FindOne(context.TODO(), filter)
	err = result.Decode(&user)
	return
}

//------------------------------------------------------------------------------------------------------------------------------------------------

func (store *RegistrarRepositoryImpl) PostPotvrdaSmrti(potvrda domain.PotvrdaSmrti) error {
	_, err := store.potvrdaSmrti.InsertOne(context.Background(), potvrda)
	if err != nil {
		return err
	}
	return nil
}

func (store *RegistrarRepositoryImpl) GetPotvrdaSmrtiJMBG(jmbg string) *domain.PotvrdaSmrti {
	potvrdaOSmrti, err := store.filterOnePotvrdaSmrti(bson.M{"jmbg": jmbg})
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	return potvrdaOSmrti
}

func (store *RegistrarRepositoryImpl) GetAllPotvrdeSmrti() ([]*domain.PotvrdaSmrti, error) {
	filter := bson.M{}
	return store.filterPotvrdeSmrti(filter)
}

func (store *RegistrarRepositoryImpl) DeletePotvrdaSmrtiID(id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	_, err := store.potvrdaSmrti.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	return nil
}

func (store *RegistrarRepositoryImpl) filterPotvrdeSmrti(filter interface{}) ([]*domain.PotvrdaSmrti, error) {
	cursor, err := store.potvrdaSmrti.Find(context.Background(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}

	return decodePotvrdaOSmrti(cursor)
}

func (store *RegistrarRepositoryImpl) filterOnePotvrdaSmrti(filter interface{}) (user *domain.PotvrdaSmrti, err error) {
	result := store.potvrdaSmrti.FindOne(context.TODO(), filter)
	err = result.Decode(&user)
	return
}

func decodeUserRegistry(cursor *mongo.Cursor) (users []*domain.User, err error) {
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

func decodePotvrdaOSmrti(cursor *mongo.Cursor) (potvrdeOSmrti []*domain.PotvrdaSmrti, err error) {
	for cursor.Next(context.Background()) {
		var potvrdaSmrti domain.PotvrdaSmrti
		err = cursor.Decode(&potvrdaSmrti)
		if err != nil {
			return
		}
		potvrdeOSmrti = append(potvrdeOSmrti, &potvrdaSmrti)
	}
	err = cursor.Err()
	return
}
