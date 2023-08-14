package model

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type DBConnection struct {
	ListenAddr string
}

func NewDBConnection(ListenAddr string) *DBConnection {
	return &DBConnection{
		ListenAddr: ListenAddr,
	}
}

func (db *DBConnection) StartRouter() {
	router := mux.NewRouter()

	router.HandleFunc("/store", db.StorePublisher)
	router.HandleFunc("/delete", db.DeletePublisher)

}

func (db *DBConnection) StorePublisher(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "hello")
}

func (db *DBConnection) DeletePublisher(writer http.ResponseWriter, request *http.Request) {
}
