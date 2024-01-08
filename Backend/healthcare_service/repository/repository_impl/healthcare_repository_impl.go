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
	pregled     *mongo.Collection
	vakcina     *mongo.Collection
	alergija    *mongo.Collection
	invaliditet *mongo.Collection
	karton      *mongo.Collection
}

const (
	DATABASE               = "healthcare"
	COLLECTION_PREGLED     = "pregled"
	COLLECTION_VAKCINA     = "vakcina"
	COLLECTION_ALERGIJA    = "alergija"
	COLLECTION_INVALIDITET = "invaliditet"
	COLLECTION_KARTON      = "karton"
)

func NewAuthRepositoryImpl(client *mongo.Client) repository.HealthcareRepository {
	pregled := client.Database(DATABASE).Collection(COLLECTION_PREGLED)
	vakcina := client.Database(DATABASE).Collection(COLLECTION_VAKCINA)
	alergija := client.Database(DATABASE).Collection(COLLECTION_ALERGIJA)
	invaliditet := client.Database(DATABASE).Collection(COLLECTION_INVALIDITET)
	karton := client.Database(DATABASE).Collection(COLLECTION_KARTON)

	return &HealthcareRepositoryImpl{
		pregled:     pregled,
		vakcina:     vakcina,
		alergija:    alergija,
		invaliditet: invaliditet,
		karton:      karton,
	}
}

//Pregled ------------------------------------------------------------------------------------------------------------------

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

func (repository *HealthcareRepositoryImpl) GetMojiPreglediGradjanin(id primitive.ObjectID) ([]*model.Pregled, error) {
	filter := bson.M{"gradjanin._id": id}
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

//Vakcina ------------------------------------------------------------------------------------------------------------------

func (repository *HealthcareRepositoryImpl) GetSveVakcine() ([]*model.Vakcina, error) {
	filter := bson.M{}
	return repository.filterVakcine(filter)
}

func (repository *HealthcareRepositoryImpl) GetVakcinaID(id primitive.ObjectID) (*model.Vakcina, error) {
	filter := bson.M{"_id": id}
	return repository.filterOneVakcina(filter)
}

func (repository *HealthcareRepositoryImpl) GetVakcinaNaziv(naziv string) (*model.Vakcina, error) {
	filter := bson.M{"naziv": naziv}
	return repository.filterOneVakcina(filter)
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
		log.Println("Updating Vakcina Error MongoDB", err.Error())
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

func (repository *HealthcareRepositoryImpl) filterOneVakcina(filter interface{}) (tipVakcine *model.Vakcina, err error) {
	result := repository.vakcina.FindOne(context.Background(), filter)
	err = result.Decode(&tipVakcine)
	return
}

//Alergija ------------------------------------------------------------------------------------------------------------------

func (repository *HealthcareRepositoryImpl) GetSveAlergije() ([]*model.Alergija, error) {
	filter := bson.M{}
	return repository.filterAlergije(filter)
}

func (repository *HealthcareRepositoryImpl) GetAlergijaID(id primitive.ObjectID) (*model.Alergija, error) {
	filter := bson.M{"_id": id}
	return repository.filterOneAlergija(filter)
}

func (repository *HealthcareRepositoryImpl) GetAlergijaNaziv(naziv string) (*model.Alergija, error) {
	filter := bson.M{"naziv": naziv}
	return repository.filterOneAlergija(filter)
}

func (repository *HealthcareRepositoryImpl) PostAlergija(alergija *model.Alergija) error {
	_, err := repository.alergija.InsertOne(context.Background(), alergija)
	if err != nil {
		return err
	}
	return nil
}

func (repository *HealthcareRepositoryImpl) DeleteAlergijaID(id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	_, err := repository.alergija.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	return nil
}

func (repository *HealthcareRepositoryImpl) filterAlergije(filter interface{}) ([]*model.Alergija, error) {
	cursor, err := repository.alergija.Find(context.Background(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}

	return decodeAlergija(cursor)
}

func (repository *HealthcareRepositoryImpl) filterOneAlergija(filter interface{}) (alergija *model.Alergija, err error) {
	result := repository.alergija.FindOne(context.Background(), filter)
	err = result.Decode(&alergija)
	return
}

//Invaliditet ------------------------------------------------------------------------------------------------------------------

func (repository *HealthcareRepositoryImpl) GetSveInvaliditete() ([]*model.Invaliditet, error) {
	filter := bson.M{}
	return repository.filterInvaliditeti(filter)
}

func (repository *HealthcareRepositoryImpl) GetInvaliditetID(id primitive.ObjectID) (*model.Invaliditet, error) {
	filter := bson.M{"_id": id}
	return repository.filterOneInvaliditet(filter)
}

func (repository *HealthcareRepositoryImpl) GetInvaliditetNaziv(naziv string) (*model.Invaliditet, error) {
	filter := bson.M{"naziv": naziv}
	return repository.filterOneInvaliditet(filter)
}

func (repository *HealthcareRepositoryImpl) PostInvaliditet(invaliditet *model.Invaliditet) error {
	_, err := repository.invaliditet.InsertOne(context.Background(), invaliditet)
	if err != nil {
		return err
	}
	return nil
}

func (repository *HealthcareRepositoryImpl) DeleteInvaliditetID(id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	_, err := repository.invaliditet.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	return nil
}

func (repository *HealthcareRepositoryImpl) filterInvaliditeti(filter interface{}) ([]*model.Invaliditet, error) {
	cursor, err := repository.invaliditet.Find(context.Background(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}

	return decodeInvaliditet(cursor)
}

func (repository *HealthcareRepositoryImpl) filterOneInvaliditet(filter interface{}) (invaliditet *model.Invaliditet, err error) {
	result := repository.invaliditet.FindOne(context.Background(), filter)
	err = result.Decode(&invaliditet)
	return
}

//Karton ------------------------------------------------------------------------------------------------------------------

func (repository *HealthcareRepositoryImpl) GetSveKartone() ([]*model.Karton, error) {
	filter := bson.M{}
	return repository.filterKartone(filter)
}

func (repository *HealthcareRepositoryImpl) GetKartoneJMBG(jmbg string) ([]*model.Karton, error) {
	filter := bson.M{"jmbg": jmbg}
	return repository.filterKartone(filter)
}

func (repository *HealthcareRepositoryImpl) GetKartonJMBG(jmbg string) (*model.Karton, error) {
	filter := bson.M{"jmbg": jmbg}
	return repository.filterOneKarton(filter)
}

func (repository *HealthcareRepositoryImpl) PostKarton(karton model.Karton) error {
	_, err := repository.karton.InsertOne(context.Background(), karton)
	if err != nil {
		return err
	}
	return nil
}

func (repository *HealthcareRepositoryImpl) PutKarton(karton *model.Karton) error {
	filter := bson.M{"_id": karton.ID}
	update := bson.D{{"$set", karton}}
	_, err := repository.karton.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Println("Updating Karton Error MongoDB", err.Error())
		return err
	}

	return nil
}

func (repository *HealthcareRepositoryImpl) DeleteKartonID(id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	_, err := repository.karton.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	return nil
}

func (repository *HealthcareRepositoryImpl) filterKartone(filter interface{}) ([]*model.Karton, error) {
	cursor, err := repository.karton.Find(context.Background(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}

	return decodeKarton(cursor)
}

func (repository *HealthcareRepositoryImpl) filterOneKarton(filter interface{}) (karton *model.Karton, err error) {
	result := repository.karton.FindOne(context.Background(), filter)
	err = result.Decode(&karton)
	return
}

//-----------------------------------------------------------------------------------------------------------------------

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

func decodeAlergija(cursor *mongo.Cursor) (alergije []*model.Alergija, err error) {
	for cursor.Next(context.Background()) {
		var alergija model.Alergija
		err = cursor.Decode(&alergija)
		if err != nil {
			return
		}
		alergije = append(alergije, &alergija)
	}
	err = cursor.Err()
	return
}

func decodeInvaliditet(cursor *mongo.Cursor) (invaliditeti []*model.Invaliditet, err error) {
	for cursor.Next(context.Background()) {
		var invaliditet model.Invaliditet
		err = cursor.Decode(&invaliditet)
		if err != nil {
			return
		}
		invaliditeti = append(invaliditeti, &invaliditet)
	}
	err = cursor.Err()
	return
}

func decodeKarton(cursor *mongo.Cursor) (kartoni []*model.Karton, err error) {
	for cursor.Next(context.Background()) {
		var karton model.Karton
		err = cursor.Decode(&karton)
		if err != nil {
			return
		}
		kartoni = append(kartoni, &karton)
	}
	err = cursor.Err()
	return
}
