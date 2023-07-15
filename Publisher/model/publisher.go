package model

type Publisher struct {
	PublisherID    string `json:"publisher_id"`
	PublisherName  string `json:"publisher_name"`
	RegisteredDate string `json:"registered_date"`
	Password       string `json:"password"`
}
