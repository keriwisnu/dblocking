package models

type (
	Sales struct {
		ID        int       `json:"id"`
		Name      string    `name:"name"`
		Stock	  int       `json:"stock"`
	}
)
