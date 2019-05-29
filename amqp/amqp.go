package amqp

import (
	"fmt"
	"github.com/inteleon/go-integration-testutil/env"
	"github.com/streadway/amqp"
	"log"
)

func SetupAmqp(fallback string) (*amqp.Connection, error){
	var err error
	amqpConn, err := amqp.Dial(env.GetEnv("AMQP_DSN", fallback))
	if err != nil {
		log.Fatal("no rabbitmq connection.")
		return nil, err
	}
	return amqpConn, nil
}

func ResetQueues(queues []string, amqpConn *amqp.Connection) {
	ch, err := amqpConn.Channel()
	if err != nil {
		log.Fatalf("unable to obtain channel: %v", err)
	}
	for _, queue := range queues {
		count, err := ch.QueuePurge(queue, false)
		if err != nil {
			fmt.Printf("Problem purging queue %v: %v\n", queue, err)
		} else {
			fmt.Printf("Successfully purged %v messages from queue %v\n", count, queue)
		}
	}
}

func UntilQueueFulfills(ch *amqp.Channel, expectedCount int, queueName string) func() error {
	numGotten := 0
	isExpected := false

	return func() error {
		q, err := ch.QueueInspect(queueName)
		if err != nil {
			return err
		}
		numGotten = q.Messages
		err = fmt.Errorf("wrong number of messages: expected %v, got %v", expectedCount, numGotten)
		if numGotten > expectedCount {
			return err
		} else if numGotten == expectedCount && isExpected {
			return nil
		} else if numGotten == expectedCount {
			isExpected = true
		}
		return err
	}
}