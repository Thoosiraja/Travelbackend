// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"

// 	"github.com/gorilla/mux"
// 	"github.com/jinzhu/gorm"
// 	_ "github.com/jinzhu/gorm/dialects/postgres"
// 	"github.com/rs/cors"
// )

// var db *gorm.DB
// var err error

// type Admindetails struct {
// 	gorm.Model
// 	Name  string
// 	Email string
// 	Pass  string
// 	Roll  string
// }
// type Busdetails struct {
// 	gorm.Model
// 	Start        string
// 	End          string
// 	Name         string
// 	Desc         string
// 	Type         string
// 	Wifi         string
// 	Water        string
// 	Refreshments string
// 	Capacity     string
// 	Fare         string
// 	Img          string
// }

// // type Account struct {
// // 	gorm.Model
// // 	Email string
// // 	Pass  string
// // }

// func InitialMigration() {

// 	db, err = gorm.Open("postgres", "port=5432 user=postgres dbname=user password=55021007 sslmode=disable")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		panic("Failed to connect")
// 	} else {
// 		fmt.Println("Connected successfully")
// 	}

// 	defer db.Close()

// 	db.AutoMigrate(&Admindetails{})
// 	db.AutoMigrate(&Busdetails{})
// 	// db.AutoMigrate(&Account{})
// }

// func helloworld(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Helloworld")
// }

// func handleRequests() {
// 	myRouter := mux.NewRouter().StrictSlash(true)
// 	myRouter.HandleFunc("/", helloworld).Methods("GET")
// 	myRouter.HandleFunc("/admindetails", Allusers).Methods("GET")
// 	myRouter.HandleFunc("/admindetails/{Name}/{Email}/{Pass}/{Roll}", Newuser).Methods("POST")
// 	// myRouter.HandleFunc("/admindetails/{Email}", Deleteuser).Methods("DELETE")
// 	// myRouter.HandleFunc("/admindetails/{Name}/{Email}/{Pass}/{roll}", Updateuser).Methods("PUT")
// 	myRouter.HandleFunc("/busdetails", Allbus).Methods("GET")
// 	myRouter.HandleFunc("/busdetails/{Start}/{End}/{Name}/{Desc}/{Type}/{Wifi}/{Water}/{Refreshments}/{Capacity}/{Fare}/{Img}", Newbus).Methods("POST")
// 	// myRouter.HandleFunc("/busdetails/{Name}", Deletebus).Methods("DELETE")
// 	// myRouter.HandleFunc("/busdetails/{Name}/{Start}/{End}", Updatebus).Methods("PUT")
// 	// myRouter.HandleFunc("/account", Alladmin).Methods("GET")
// 	// myRouter.HandleFunc("/account/{Email}/{Pass}", Newadmin).Methods("POST")
// 	log.Fatal(http.ListenAndServe(":8000", cors.Default().Handler(myRouter)))
// }

// func main1() {
// 	// fmt.Println("Started")
// 	InitialMigration()
// 	handleRequests()
// }
