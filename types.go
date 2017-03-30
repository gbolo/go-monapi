package main

import "time"

type HealthCheck struct {
	Name      string    `json:"name"`
	Status    int       `json:"status"`
	Completed bool      `json:"completed"`
	Date      time.Time `json:"date"`
	Verbose   string    `json:"verbose"`
}
