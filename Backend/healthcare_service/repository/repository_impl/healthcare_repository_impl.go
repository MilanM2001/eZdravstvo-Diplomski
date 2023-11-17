package repository_impl

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"healthcare_service/model"
	"healthcare_service/repository"
	"log"
)

type HealthcareRepositoryImpl struct {
	pregled *mongo.Collection
	vakcina *mongo.Collection
}

const (
	DATABASE           = "healthcare"
	COLLECTION_PREGLED = "pregled"
	COLLECTION_VAKCINA = "vakcina"
)

func NewAuthRepositoryImpl(client *mongo.Client) repository.HealthcareRepository {
	pregled := client.Database(DATABASE).Collection(COLLECTION_PREGLED)
	vakcina := client.Database(DATABASE).Collection(COLLECTION_VAKCINA)

	return &HealthcareRepositoryImpl{
		pregled: pregled,
		vakcina: vakcina,
	}
}

//Pregled

func (repository *HealthcareRepositoryImpl) GetSviPregledi() ([]*model.Pregled, error) {
	filter := bson.M{}
	return repository.filterPregledi(filter)
}

func (repository *HealthcareRepositoryImpl) GetMojiPreglediLekar(id primitive.ObjectID) ([]*model.Pregled, error) {
	filter := bson.M{"lekar._id": id}
	return repository.filterPregledi(filter)
}

func (repository *HealthcareRepositoryImpl) GetMojiSlobodniPreglediLekar(id primitive.ObjectID) ([]*model.Pregled, error) {
	filter := bson.M{"lekar._id": id, "gradjanin": nil}
	return repository.filterPregledi(filter)
}

func (repository *HealthcareRepositoryImpl) GetMojiZauzetiPreglediLekar(id primitive.ObjectID) ([]*model.Pregled, error) {
	filter := bson.M{"lekar._id": id, "gradjanin": bson.M{"$ne": nil}}
	return repository.filterPregledi(filter)
}

func (repository *HealthcareRepositoryImpl) GetSviSlobodniPregledi() ([]*model.Pregled, error) {
	filter := bson.M{"gradjanin": nil}
	return repository.filterPregledi(filter)
}

func (repository *HealthcareRepositoryImpl) GetPregledID(id primitive.ObjectID) (*model.Pregled, error) {
	filter := bson.M{"_id": id}
	return repository.filterOnePregled(filter)
}

func (repository *HealthcareRepositoryImpl) PostPregled(pregled *model.Pregled) error {
	_, err := repository.pregled.InsertOne(context.Background(), pregled)
	if err != nil {
		return err
	}
	return nil
}

func (repository *HealthcareRepositoryImpl) PutPregled(pregled *model.Pregled) error {
	filter := bson.M{"_id": pregled.ID}
	update := bson.D{{"$set", pregled}}
	_, err := repository.pregled.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Println("Updating Pregled Error MongoDB", err.Error())
		return err
	}

	return nil
}

func (repository *HealthcareRepositoryImpl) DeletePregledID(id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	_, err := repository.pregled.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	return nil
}

func (repository *HealthcareRepositoryImpl) filterPregledi(filter interface{}) ([]*model.Pregled, error) {
	cursor, err := repository.pregled.Find(context.Background(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}

	return decodePregled(cursor)
}

func (repository *HealthcareRepositoryImpl) filterOnePregled(filter interface{}) (pregled *model.Pregled, err error) {
	result := repository.pregled.FindOne(context.Background(), filter)
	err = result.Decode(&pregled)
	return
}

//Vakcine

func (repository *HealthcareRepositoryImpl) GetSveVakcine() ([]*model.Vakcina, error) {
	filter := bson.M{}
	return repository.filterVakcine(filter)
}

func (repository *HealthcareRepositoryImpl) GetVakcinaID(id primitive.ObjectID) (*model.Vakcina, error) {
	filter := bson.M{"_id": id}
	return repository.filterOneTipVakcine(filter)
}

func (repository *HealthcareRepositoryImpl) GetVakcinaNaziv(naziv string) (*model.Vakcina, error) {
	filter := bson.M{"naziv": naziv}
	return repository.filterOneTipVakcine(filter)
}

func (repository *HealthcareRepositoryImpl) PostVakcina(vakcina *model.Vakcina) error {
	_, err := repository.vakcina.InsertOne(context.Background(), vakcina)
	if err != nil {
		return err
	}
	return nil
}

func (repository *HealthcareRepositoryImpl) PutVakcina(vakcina *model.Vakcina) error {
	filter := bson.M{"_id": vakcina.ID}
	update := bson.D{{"$set", vakcina}}
	_, err := repository.vakcina.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Println("Updating Pregled Error MongoDB", err.Error())
		return err
	}

	return nil
}

func (repository *HealthcareRepositoryImpl) DeleteVakcinaID(id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	_, err := repository.vakcina.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	return nil
}

func (repository *HealthcareRepositoryImpl) filterVakcine(filter interface{}) ([]*model.Vakcina, error) {
	cursor, err := repository.vakcina.Find(context.Background(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}

	return decodeVakcina(cursor)
}

func (repository *HealthcareRepositoryImpl) filterOneTipVakcine(filter interface{}) (tipVakcine *model.Vakcina, err error) {
	result := repository.vakcina.FindOne(context.Background(), filter)
	err = result.Decode(&tipVakcine)
	return
}

//func (repository *HealthcareRepositoryImpl) GetMyTakenVaccinationsRegular(id primitive.ObjectID) ([]*model.Vaccination, error) {
//	filter := bson.M{"user._id": id}
//	return repository.filterVaccinations(filter)
//}

//

//func (repository *HealthcareRepositoryImpl) GetZdravstvenoStanjeByJMBG(jmbg string) (*model.ZdravstvenoStanje, error) {
//	filter := bson.M{"jmbg": jmbg}
//	return repository.filterOneZdravstvenoStanje(filter)
//}
//
//func (repository *HealthcareRepositoryImpl) CreateNewZdravstvenoStanje(zdravstvenoStanje *model.ZdravstvenoStanje) error {
//	_, err := repository.zdravstvenoStanje.InsertOne(context.Background(), zdravstvenoStanje)
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//func (repository *HealthcareRepositoryImpl) DeleteZdravstvenoStanjeByJMBG(jmbg string) error {
//	filter := bson.M{"jmbg": jmbg}
//	_, err := repository.zdravstvenoStanje.DeleteOne(context.Background(), filter)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//func (repository *HealthcareRepositoryImpl) filterZdravstvenaStanja(filter interface{}) ([]*model.ZdravstvenoStanje, error) {
//	cursor, err := repository.zdravstvenoStanje.Find(context.Background(), filter)
//	defer cursor.Close(context.TODO())
//
//	if err != nil {
//		return nil, err
//	}
//
//	return decodeZdravstvenoStanje(cursor)
//}
//
//func (repository *HealthcareRepositoryImpl) filterOneZdravstvenoStanje(filter interface{}) (zdravstvenoStanje *model.ZdravstvenoStanje, err error) {
//	result := repository.zdravstvenoStanje.FindOne(context.Background(), filter)
//	err = result.Decode(&zdravstvenoStanje)
//	return
//}
//
func decodePregled(cursor *mongo.Cursor) (pregledi []*model.Pregled, err error) {
	for cursor.Next(context.Background()) {
		var pregled model.Pregled
		err = cursor.Decode(&pregled)
		if err != nil {
			return
		}
		pregledi = append(pregledi, &pregled)
	}
	err = cursor.Err()
	return
}

func decodeVakcina(cursor *mongo.Cursor) (vakcine []*model.Vakcina, err error) {
	for cursor.Next(context.Background()) {
		var vakcina model.Vakcina
		err = cursor.Decode(&vakcina)
		if err != nil {
			return
		}
		vakcine = append(vakcine, &vakcina)
	}
	err = cursor.Err()
	return
}
