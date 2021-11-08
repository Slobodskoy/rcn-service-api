package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ozonmp/rcn-service-api/internal/app/repo"
	"github.com/ozonmp/rcn-service-api/internal/app/retranslator"
	"github.com/ozonmp/rcn-service-api/internal/app/sender"
)

func main() {

	sigs := make(chan os.Signal, 1)

	cfg := retranslator.Config{
		ChannelSize:    512,
		ConsumerCount:  2,
		ConsumeSize:    10,
		ConsumeTimeout: time.Microsecond * 100,
		ProducerCount:  28,
		WorkerCount:    2,
		Repo:           repo.NewDbEventRepo(nil),
		Sender:         sender.NewKafkaEventSender(),
	}

	retranslator := retranslator.NewRetranslator(cfg)
	retranslator.Start()

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs

	retranslator.Close()
}
