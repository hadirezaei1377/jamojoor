package saga

import (
	"fmt"
	"log"

	"github.com/Shopify/sarama"
)

func main() {
	// kafka config for using as a coordinator in SAGA pattern
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalf("error while connecting to KAFKA: %v", err)
	}
	defer producer.Close()

	// service one, for payment
	if err := performPayment(producer); err != nil {
		log.Printf("error while payment: %v", err)
		rollback("Service A", producer) // rolle back in error time
		return
	}

	// service two, reduce the balance
	if err := reduceInventory(producer); err != nil {
		log.Printf("error while reduce the balance: %v", err)
		rollback("Service A", producer) // rolle back in error time
		rollback("Service B", producer) // rolle back in error time
		return
	}

	// service three, register the order
	if err := placeOrder(producer); err != nil {
		log.Printf("error while register the order: %v", err)
		rollback("Service A", producer) // rolle back in error time
		rollback("Service B", producer) // rolle back in error time
		rollback("Service C", producer) // rolle back in error time
		return
	}

	log.Println("transaction done Successfully")
}

func performPayment(producer sarama.SyncProducer) error {
	msg := &sarama.ProducerMessage{
		Topic: "payment",
		Value: sarama.StringEncoder("payment done!"),
	}
	_, _, err := producer.SendMessage(msg)
	return err
}

func reduceInventory(producer sarama.SyncProducer) error {
	msg := &sarama.ProducerMessage{
		Topic: "inventory",
		Value: sarama.StringEncoder("balance reduced!"),
	}
	_, _, err := producer.SendMessage(msg)
	return err
}

func placeOrder(producer sarama.SyncProducer) error {
	msg := &sarama.ProducerMessage{
		Topic: "order",
		Value: sarama.StringEncoder("order registered!"),
	}
	_, _, err := producer.SendMessage(msg)
	return err
}

func rollback(serviceName string, producer sarama.SyncProducer) {
	msg := &sarama.ProducerMessage{
		Topic: "rollback",
		Value: sarama.StringEncoder(fmt.Sprintf("rolle back %s done.", serviceName)),
	}
	_, _, _ = producer.SendMessage(msg)
	log.Printf("rolle back %s done.", serviceName)
}
