package main

import (
	"log"
	"time"

	"github.com/DataDog/datadog-go/statsd"
)

const (
	agentAddr = "127.0.0.1:8125"
)

type Status int32

const (
	OK       Status = 0
	CRITICAL Status = 1
)

var (
	cr = &statsd.ServiceCheck{
		Name: "playground.ping",
	}
)

// write playground.ping.* metric
func write(client *statsd.Client, status Status, tags []string) error {
	if status == OK {
		if err := client.Gauge("playground.ping.ok", 1, tags, 1); err != nil {
			return err
		}

		if err := client.Gauge("playground.ping.critical", 0, tags, 1); err != nil {
			return err
		}

		cr.Tags = tags
		cr.Status = statsd.Ok
		if err := client.ServiceCheck(cr); err != nil {
			return err
		}
	} else if status == CRITICAL {
		if err := client.Gauge("playground.ping.ok", 0, tags, 1); err != nil {
			return err
		}

		if err := client.Gauge("playground.ping.critical", 1, tags, 1); err != nil {
			return err
		}

		cr.Tags = tags
		cr.Status = statsd.Critical
		if err := client.ServiceCheck(cr); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	client, err := statsd.New(agentAddr)
	if err != nil {
		log.Fatal(err)
	}

	counter := 0
	for {
		if counter%12 <= 2 {
			if err := write(client, CRITICAL, []string{"service:A", "region:NORTH"}); err != nil {
				log.Fatal(err)
			}

			if err := write(client, CRITICAL, []string{"service:B", "region:NORTH"}); err != nil {
				log.Fatal(err)
			}

			if err := write(client, CRITICAL, []string{"service:C", "region:SOUTH"}); err != nil {
				log.Fatal(err)
			}

			if err := write(client, OK, []string{"service:D", "region:SOUTH"}); err != nil {
				log.Fatal(err)
			}
		} else if counter%12 <= 5 {
			if err := write(client, OK, []string{"service:A", "region:NORTH"}); err != nil {
				log.Fatal(err)
			}

			if err := write(client, OK, []string{"service:B", "region:NORTH"}); err != nil {
				log.Fatal(err)
			}

			if err := write(client, OK, []string{"service:C", "region:SOUTH"}); err != nil {
				log.Fatal(err)
			}

			if err := write(client, CRITICAL, []string{"service:D", "region:SOUTH"}); err != nil {
				log.Fatal(err)
			}
		}

		counter++
		time.Sleep(1 * time.Minute)
	}
}
