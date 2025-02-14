package finance

import (
	"log"
)

type Finace struct {
	ID       string
	Action   string
	CreateAt string
	UpdateAt string
	Status   string
}

func Start() {
	log.Println("Start Finance Module")
}
