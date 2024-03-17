package main

import (
	"fmt"
	"time"
)

type feedback struct {
	text     string
	rateStar int // 1-5
}

type University struct {
	name string
	numberOfStudents int
	feedbacks        []feedback
	createdAt        time.Time
	acceptanceRate   float64 // 1-100%
}

func main() {
	var pdp = University{
		name: "pdp",
		numberOfStudents: 1000,
		feedbacks: []feedback{
			{
				text:     "yeah its good but not my dream uni",
				rateStar: 3,
			},
			{
				text:     "wow its so good, i wanna throw up",
				rateStar: 1,
			},
		},
		createdAt:      time.Now().UTC(),
		acceptanceRate: 100,
	}

	fmt.Printf("Name: %v\nStudents: %d\nAcceptance rate: %.2f\nCreated time: %v\nFeedbacks: %v", pdp.name, pdp.numberOfStudents, pdp.acceptanceRate, pdp.createdAt.Year(), pdp.feedbacks)
}


