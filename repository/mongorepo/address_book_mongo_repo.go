package mongorepo

import (
	"context"
	"firstMS/repository"
	"firstMS/repository/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AddressBookMongoRepo struct {
	collection *mongo.Collection
}

func (a *AddressBookMongoRepo) GetAddressBook(ctx context.Context) (*models.AddressBook, error) {
	var persons []models.Person
	cursor, err := a.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &persons); err != nil {
		return nil, err
	}
	var ab models.AddressBook
	ab.Persons = persons
	return &ab, nil
}

func (a *AddressBookMongoRepo) StoreOnePerson(ctx context.Context, person models.Person) error {
	_, err := a.collection.InsertOne(ctx, person)
	return err
}

func GetAddressBookMongoRepo(collection *mongo.Collection) repository.AddressBookRepo {
	return &AddressBookMongoRepo{collection: collection}
}
