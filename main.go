package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

type Logindetails struct {
	Name  string
	Email string `gorm:"primary_key"`
	Pass  string
	Roll  string
	// No_of_tickets_booked int
	// Total_amount         int
}
type Busdetails struct {
	Id           uint `gorm:"AUTO_INCREMENT"`
	Start        string
	Destination  string
	Name         string
	Desc         string
	Type         string
	Wifi         string
	Water        string
	Refreshments string
	Capacity     int
	Fare         int
	// Img          string
	Count int
}
type Bookingdetails struct {
	Email                string
	Bus_id               int
	No_of_tickets_booked int
	Total_amount         int
}
type Result struct {
	Email string
	Total int
}

func connection() {
	db, err := gorm.Open("postgres", "port=5432 user=postgres dbname=user password=55021007 sslmode=disable")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect")
	} else {
		fmt.Println("Connected successfully")
	}
	defer db.Close()
	db.AutoMigrate(&Logindetails{})
	db.AutoMigrate(&Busdetails{})
	db.AutoMigrate(&Bookingdetails{})
	db.AutoMigrate(&Result{})
}

func request() {
	route := mux.NewRouter().StrictSlash(true)
	route.HandleFunc("/Logindetails", Allusers).Methods("GET")
	route.HandleFunc("/Busdetails", Allbus).Methods("GET")
	route.HandleFunc("/Busdetails/{Id}", Retbus).Methods("GET")
	route.HandleFunc("/Logindetails/{Name}/{Email}/{Pass}/{Roll}", Createuser).Methods("POST")
	route.HandleFunc("/Logindetails/{Email}/{Pass}", Returnuser).Methods("GET")
	route.HandleFunc("/Bookingdetails/{Email}/{Bus_id}/{No_of_tickets_booked}/{Total_amount}", Createbooking).Methods("POST")
	route.HandleFunc("/Busdetails/{Start}/{End}/{Name}/{Desc}/{Type}/{Wifi}/{Water}/{Refreshments}/{Capacity}/{Fare}/{Count}", Createbus).Methods("POST")
	route.HandleFunc("/Busdetails/{Start}/{End}", Returnbus).Methods("GET")
	route.HandleFunc("/Bookingdetails/{Email}", Returnbooking).Methods("GET")
	route.HandleFunc("/Busdetails/{Id}/{No}", Updateseatcount).Methods("POST")
	route.HandleFunc("/Result/{Email}/{Total}", Createtransaction).Methods("POST")
	route.HandleFunc("/Result/{Email}/{No}/{Char}", Updatetotal).Methods("POST")
	route.HandleFunc("/Bookingdetails", Returnusercount).Methods("GET")
	route.HandleFunc("/Result", Returnresult).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", cors.Default().Handler(route)))
}
func main() {
	connection()
	request()

}
