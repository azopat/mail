// User authentication and CRUD
//

package main

import (
	"../common/const/mailconst"
	"encoding/json"
	"errors"
	"github.com/bitly/go-nsq"
	"github.com/codegangsta/cli"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

const (
	SERVICE_NAME          = "mail"
	ENV                   = "dev"
	API_VERSION           = "0.1"
	SERVICE_TOPIC         = "mail"
	FLAG_ENV              = "env"
	FLAG_PORT             = "port"
	ENV_FROM_ADDRESS      = "FROM_ADDRESS"
	ENV_NSQ_HOST          = "NSQ_HOST"
	ENV_MAIL_FROM_ADDRESS = "MAIL_FROM_ADDRESS"
)

var (
	BUILD       = "not defined"
	fromAddress string
	nsqHost     string
)

func main() {

	app := cli.NewApp()

	app.Name = SERVICE_NAME
	app.Usage = "User authentication"
	app.Action = runIt
	app.Version = BUILD + "\nAPI Version: " + API_VERSION + "\nENV: " + ENV
	app.Before = checkOptions
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  FLAG_ENV,
			Value: "dev",
			Usage: "live or dev mode",
		},
		cli.StringFlag{
			Name:  FLAG_PORT,
			Value: "80",
			Usage: "Port to listen on",
		},
	}
	app.Run(os.Args)
}

func checkOptions(c *cli.Context) error {

	e := c.String(FLAG_ENV)

	if (e != "dev") && (e != "live") {
		err := errors.New("ENV must be 'dev' or 'live'")
		return err
	}
	return nil
}

func runIt(c *cli.Context) {

	fromAddress = os.Getenv(ENV_FROM_ADDRESS)
	nsqHost = os.Getenv(ENV_NSQ_HOST)

	log.Printf("fromAddress :  %s", fromAddress)
	log.Printf("nsqHost :  %s", nsqHost)

	setupNSQ()

}

func setupNSQ() {

	log.Printf("setupNSQ")
	config := nsq.NewConfig()
	q, _ := nsq.NewConsumer(mailconst.NSQ_TOPIC, "process-request", config)
	q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Printf("Got a message: %v", message)

		// Unmarshal JSON to Result struct.
		var queueMailDoc mailconst.QueueMailDoc
		json.Unmarshal(message.Body, &queueMailDoc)

		sendEmail(queueMailDoc)
		return nil
	}))

	log.Printf(nsqHost)
	err := q.ConnectToNSQD(nsqHost)
	if err != nil {
		log.Fatal(err)
	}
	<-q.StopChan

}
