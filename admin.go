// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"

// 	"github.com/gorilla/mux"
// 	"github.com/jinzhu/gorm"
// )

// func Allusers(w http.ResponseWriter, r *http.Request) {

// 	db, err = gorm.Open("postgres", "port=5432 user=postgres dbname=user password=55021007 sslmode=disable")
// 	if err != nil {
// 		panic("Failed to connect")
// 	} else {
// 		fmt.Println("Connection Successfull")
// 	}
// 	var admindetails []Admindetails
// 	db.Find(&admindetails)
// 	json.NewEncoder(w).Encode(admindetails)
// }
// func Allbus(w http.ResponseWriter, r *http.Request) {

// 	db, err = gorm.Open("postgres", "port=5432 user=postgres dbname=user password=55021007 sslmode=disable")
// 	if err != nil {
// 		panic("Failed to connect")
// 	} else {
// 		fmt.Println("Connection Successfull")
// 	}
// 	var busdetails []Busdetails
// 	db.Find(&busdetails)
// 	json.NewEncoder(w).Encode(busdetails)
// }

// // func Alladmin(w http.ResponseWriter, r *http.Request) {

// // 	db, err = gorm.Open("postgres", "port=5432 user=postgres dbname=user password=55021007 sslmode=disable")
// // 	if err != nil {
// // 		panic("Failed to connect")
// // 	} else {
// // 		fmt.Println("Connection Successfull")
// // 	}
// // 	var account []Account
// // 	db.Find(&account)
// // 	json.NewEncoder(w).Encode(account)
// // }

// func Newuser(w http.ResponseWriter, r *http.Request) {

// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Content-Type", "application/json")

// 	db, err = gorm.Open("postgres", "port=5432 user=postgres dbname=user password=55021007 sslmode=disable")
// 	if err != nil {
// 		panic("Failed to connect")
// 	} else {
// 		fmt.Println("Connection Successfull")
// 	}

// 	vars := mux.Vars(r)
// 	Name := vars["Name"]
// 	Email := vars["Email"]
// 	Pass := vars["Pass"]
// 	Roll := vars["Roll"]
// 	db.Create(&Admindetails{Name: Name, Email: Email, Pass: Pass, Roll: Roll})
// 	fmt.Fprintf(w, "New user created successfully")

// }

// // func Newadmin(w http.ResponseWriter, r *http.Request) {

// // 	w.Header().Set("Access-Control-Allow-Origin", "*")
// // 	w.Header().Set("Content-Type", "application/json")

// // 	db, err = gorm.Open("postgres", "port=5432 user=postgres dbname=user password=55021007 sslmode=disable")
// // 	if err != nil {
// // 		panic("Failed to connect")
// // 	} else {
// // 		fmt.Println("Connection Successfull")
// // 	}

// // 	vars := mux.Vars(r)
// // 	Email := vars["Email"]
// // 	Pass := vars["Pass"]
// // 	db.Create(&Account{Email: Email, Pass: Pass})
// // 	fmt.Fprintf(w, "New admin created successfully")

// // }

// func Newbus(w http.ResponseWriter, r *http.Request) {

// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Content-Type", "application/json")

// 	db, err = gorm.Open("postgres", "port=5432 user=postgres dbname=user password=55021007 sslmode=disable")
// 	if err != nil {
// 		panic("Failed to connect")
// 	} else {
// 		fmt.Println("Connection Successfull")
// 	}

// 	vars := mux.Vars(r)
// 	Start := vars["Start"]
// 	End := vars["End"]
// 	Name := vars["Name"]
// 	Desc := vars["Desc"]
// 	Type := vars["Type"]
// 	Wifi := vars["Wifi"]
// 	Water := vars["Water"]
// 	Refreshments := vars["Refreshments"]
// 	Capacity := vars["Capacity"]
// 	Fare := vars["Fare"]
// 	Img := vars["Img"]
// 	db.Create(&Busdetails{Start: Start, End: End, Name: Name, Desc: Desc, Type: Type, Wifi: Wifi, Water: Water, Refreshments: Refreshments, Capacity: Capacity, Fare: Fare, Img: Img})
// 	fmt.Fprintf(w, "New bus created successfully")

// }

// // func Deleteuser(w http.ResponseWriter, r *http.Request) {

// // 	db, err = gorm.Open("postgres", "port=5432 user=postgres dbname=user password=55021007 sslmode=disable")
// // 	if err != nil {
// // 		panic("Failed to connect")
// // 	} else {
// // 		fmt.Println("Connection Successfull")
// // 	}

// // 	vars := mux.Vars(r)
// // 	Name := vars["Name"]

// // 	var admindetails Admindetails
// // 	db.Where("Name=?", Name).Find(&admindetails)
// // 	db.Delete(admindetails)

// // 	fmt.Fprintf(w, "User deleted successfully")

// // }
// // func Deletebus(w http.ResponseWriter, r *http.Request) {

// // 	db, err = gorm.Open("postgres", "port=5432 user=postgres dbname=user password=55021007 sslmode=disable")
// // 	if err != nil {
// // 		panic("Failed to connect")
// // 	} else {
// // 		fmt.Println("Connection Successfull")
// // 	}

// // 	vars := mux.Vars(r)
// // 	Name := vars["Name"]

// // 	var busdetails Busdetails
// // 	db.Where("Name=?", Name).Find(&busdetails)
// // 	db.Delete(busdetails)

// // 	fmt.Fprintf(w, "Bus deleted successfully")

// // }
// // func Updateuser(w http.ResponseWriter, r *http.Request) {

// // 	db, err = gorm.Open("postgres", "port=5432 user=postgres dbname=user password=55021007 sslmode=disable")
// // 	if err != nil {
// // 		panic("Failed to connect")
// // 	} else {
// // 		fmt.Println("Connection Successfull")
// // 	}

// // 	vars := mux.Vars(r)
// // 	Name := vars["Name"]
// // 	Email := vars["Email"]

// // 	var admindetails Admindetails
// // 	db.Where("Name=?", Name).Find(&admindetails)
// // 	admindetails.Email = Email
// // 	db.Save(&admindetails)

// // 	fmt.Fprintf(w, "User Updated successfully")

// // }
// func Login(w http.ResponseWriter, r *http.Request) {

// 	db, err = gorm.Open("postgres", "port=5432 user=postgres dbname=user password=55021007 sslmode=disable")
// 	if err != nil {
// 		panic("Failed to connect")
// 	} else {
// 		fmt.Println("Connection Successfull")
// 	}

// 	vars := mux.Vars(r)
// 	Name := vars["Name"]
// 	Email := vars["Email"]

// 	var admindetails Admindetails
// 	db.Where("Name=? AND Email=?", Name, Email).Find(&admindetails)
// 	json.NewEncoder(w).Encode(admindetails)
// 	// fmt.Fprintf(w, "User Updated successfully")

// }
// func Updatebus(w http.ResponseWriter, r *http.Request) {

// 	db, err = gorm.Open("postgres", "port=5432 user=postgres dbname=user password=55021007 sslmode=disable")
// 	if err != nil {
// 		panic("Failed to connect")
// 	} else {
// 		fmt.Println("Connection Successfull")
// 	}

// 	vars := mux.Vars(r)
// 	Name := vars["Name"]
// 	Start := vars["Start"]
// 	End := vars["End"]

// 	var busdetails Busdetails
// 	db.Where("Name=?", Name).Find(&busdetails)
// 	busdetails.Start = Start
// 	busdetails.End = End
// 	db.Save(&busdetails)

// 	fmt.Fprintf(w, "Bus Updated successfully")

// }
