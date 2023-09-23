package cli

import (
	"context"
	"encoding/json"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/spf13/cobra"
	"github.com/ugent-library/projects/gismo"
	"github.com/ugent-library/projects/models"
	"github.com/ugent-library/projects/repositories"
)

func init() {
	rootCmd.AddCommand(importCmd)
}

var importCmd = &cobra.Command{
	Use: "import",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		// setup repo
		repo, err := repositories.New(repositories.Config{
			Conn:   config.Repo.Conn,
			Secret: []byte(config.Repo.Secret),
		})
		if err != nil {
			return err
		}

		// setup nats
		opts := nats.Options{
			Url:                  config.Nats.URL,
			MaxReconnect:         100, // try reconnect n times, and then give up
			RetryOnFailedConnect: true,
			ReconnectWait:        10 * time.Second,
			Timeout:              10 * time.Second, // connection timeout
			AllowReconnect:       true,
		}

		// set NKeys authentication

		// nats event callbacks
		opts.DisconnectedErrCB = func(c *nats.Conn, err error) {
			if err != nil {
				logger.Errorf("connection got disconnected, and was unable to reconnect (num reconnections: %d): %s", c.Reconnects, err)
			}
		}

		opts.ReconnectedCB = func(c *nats.Conn) {
			logger.Infof("connection has been restored")
		}

		opts.ClosedCB = func(c *nats.Conn) {
			logger.Infof("connection has been closed")
		}

		nc, err := opts.Connect()
		if err != nil {
			return err
		}

		// setup jetstream
		js, err := jetstream.New(nc)

		if err != nil {
			return err
		}

		s, err := js.Stream(ctx, config.Nats.Stream)

		if err != nil {
			return err
		}

		cons, err := s.Consumer(ctx, config.Nats.Consumer)

		if err != nil {
			return err
		}

		logger.Info("Waiting for incoming messages...")

		consCtx, err := cons.Consume(func(msg jetstream.Msg) {
			ctx := context.Background()

			gp := &gismo.Project{}
			if err := json.Unmarshal(msg.Data(), gp); err != nil {
				logger.Warn(err)
			}

			ids := make([]*models.Identifier, 0, len(gp.Identifier))
			for _, id := range gp.Identifier {
				ids = append(ids, &models.Identifier{
					PropertyID: id.PropertyID,
					Value:      id.Value,
				})
			}

			if err := repo.AddProject(ctx, &models.Project{
				Name:            gp.Name,
				Description:     gp.Description,
				Identifier:      ids,
				FoundingDate:    gp.FoundingDate,
				DissolutionDate: gp.DissolutionDate,
			}); err != nil {
				logger.Warn(err)
			}

			msg.Ack()
		}, jetstream.ConsumeErrHandler(func(consumeCtx jetstream.ConsumeContext, err error) {
			logger.Error(err)
		}))

		if err != nil {
			return err
		}

		// Mimicking ory/http_graceful
		// See also: https://stackoverflow.com/a/28629623
		var (
			stopChan = make(chan os.Signal, 1)
			errChan  = make(chan error, 1)
		)

		go func() {
			signal.Notify(stopChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
			<-stopChan

			logger.Info("stop listening...")

			// Give nats 1 second to cancel the subscription, and close the connection
			// before exiting
			timer, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel()

			consCtx.Stop()
			nc.Close()

			<-timer.Done()

			errChan <- nil
		}()

		return <-errChan
	},
}
