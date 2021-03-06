package roquettor

import (
	// Posgresql driver

	"fmt"
	"log"

	"github.com/estebgonza/roquette/roqclient"
	_ "github.com/lib/pq"
)

// Database - Database settings connection
type Database struct {
	Driver     string `json:"driver"`
	Connection struct {
		Host string `json:"host"`
		Port int    `json:"port"`
		Db	string `json:"db"`
		User string `json:"user"`
		Pass string `json:"pass"`
	} `json:"connection"`
}

// Plan - Execution plan for roquettor
type Plan struct {
	Name            string `json:"name"`
	ConcurrentLevel int    `json:"concurrent-level"`
	Queries         []struct {
		SQL    string `json:"sql"`
		Repeat int    `json:"repeat"`
	} `json:"queries"`
}

// Row - Abstract row
type Row struct {
}

// Execute - test
func Execute(d *Database, p *Plan) error {
	rclient, err := roqclient.NewRClient(d.Driver)
	if err != nil {
		return fmt.Errorf("Error while instanciating RoqClient: %s", err.Error())
	}
	err = rclient.Connect(d.Connection.Host, d.Connection.Port, d.Connection.User, d.Connection.Pass, d.Connection.Db)
	if err != nil {
		return fmt.Errorf("Error while connection: %s", err.Error())
	}
	for _, query := range p.Queries {
		for i := 0; i < query.Repeat; i++ {
			_, err = rclient.Execute(query.SQL)
			if err != nil {
				log.Println(err)
			}
		}
	}
	// TODO: Temp reporting queries loading
	log.Println("Queries finished")
	return nil
}
