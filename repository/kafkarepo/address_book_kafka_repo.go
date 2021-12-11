package kafkarepo

import (
	"context"
	"encoding/json"
	"firstMS/repository"
	"firstMS/repository/models"
	"github.com/segmentio/kafka-go"
	"log"
)

type AddressBookKafkaRepo struct {
	Topic     string
	Partition int
	Address   string
}

func (a *AddressBookKafkaRepo) GetAddressBook(ctx context.Context) (*models.AddressBook, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AddressBookKafkaRepo) StoreOnePerson(ctx context.Context, person models.Person) error {
	w := &kafka.Writer{
		Addr:     kafka.TCP(a.Address),
		Topic:    a.Topic,
		Balancer: &kafka.LeastBytes{},
	}
	personJson, err := json.Marshal(person)
	if err != nil {
		log.Fatal("failed to marshal person:", err)
	}
	err = w.WriteMessages(ctx,
		kafka.Message{
			Key:   []byte("Person"),
			Value: personJson,
		},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := w.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
	return err
}

func getKafkaConnection(address string, topic string, partition int) (*kafka.Conn, error) {
	conn, err := kafka.DialLeader(context.Background(), "tcp", address, topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}
	return conn, err
}

func GetAddressBookKafkaRepo(address string, topic string, partition int) repository.AddressBookRepo {
	return &AddressBookKafkaRepo{Address: address, Topic: topic, Partition: partition}
}
