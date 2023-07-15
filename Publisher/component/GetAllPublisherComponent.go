package component

import "Publisher/model"

func GetAllPublisher() []model.Publisher {
	publishers := []model.Publisher{
		{PublisherID: "theroyakash", PublisherName: "Roy, Akash", RegisteredDate: "16/07/2023", Password: "3kn3noinet"},
		{PublisherID: "chennaipolice", PublisherName: "Chennai Police", RegisteredDate: "16/07/2023", Password: "dqwr32r5"},
	}

	return publishers
}
