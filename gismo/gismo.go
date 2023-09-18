package gismo

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

const streamName = "GISMO"

type Config struct {
	URL string
	// NKey     string
	// NKeySeed string
}

func Listen(c Config) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	opts := nats.Options{
		Url:                  c.URL,
		MaxReconnect:         100, // try reconnect n times, and then give up
		RetryOnFailedConnect: true,
		ReconnectWait:        10 * time.Second,
		Timeout:              10 * time.Second, // connection timeout
		AllowReconnect:       true,
	}

	// Set NKeys authentication

	// Event callbacks

	nc, err := opts.Connect()
	if err != nil {
		return err
	}
	defer nc.Close()

	js, err := jetstream.New(nc)
	if err != nil {
		return err
	}

	s, err := js.Stream(ctx, "GISMO")

	if err != nil {
		return err
	}

	cons, err := s.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
		Durable:       "GISMO",
		AckPolicy:     jetstream.AckExplicitPolicy,
		FilterSubject: "project.*",
	})

	if err != nil {
		return err
	}

	consCtx, err := cons.Consume(func(msg jetstream.Msg) {
		fmt.Printf("Received a JetStream message: %s\n", string(msg.Data()))
		log.Println("test")
	})
	defer consCtx.Stop()

	// if err != nil {
	// 	return err
	// }

	return nil
}
