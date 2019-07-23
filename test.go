package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/smtp"
	"strconv"

	"github.com/domodwyer/mailyak"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func Allbus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	db, err = gorm.Open("postgres", "port=5432 user=postgres dbname=user password=55021007 sslmode=disable")
	if err != nil {
		panic("Failed to connect")
	} else {
		fmt.Println("Connection Successfull")
	}
	var busdetails []Busdetails
	db.Find(&busdetails)
	json.NewEncoder(w).Encode(busdetails)
}
func mailSending(Email string) {
	mail := mailyak.New("smtp.gmail.com:587", smtp.PlainAuth("", "thoosiraja@codingmart.com", "********", "smtp.gmail.com"))

	mail.To(Email)
	mail.From("way@happen.com")
	mail.Subject("Bus Booking")
	mail.HTML().Set(
		"<html>" +
			"<body>" + "\r\n" +
			"<div class='row'>" + "\r\n" +
			"Tickets Booked" +
			"</div>" + "\r\n" +
			"</body>" + "\r\n" +
			"</html>")
	if err := mail.Send(); err != nil {
		panic(" 💣 ")
	}
}

func Retbus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	db, err = gorm.Open("postgres", "port=5432 user=postgres dbname=user password=55021007 sslmode=disable")
	if err != nil {
		panic("Failed to connect")
	} else {
		fmt.Println("Connection Successfull")
	}
	var busdetails []Busdetails
	db.Limit(5).Order("count desc").Find(&busdetails)
	json.NewEncoder(w).Encode(busdetails)
}
func Allusers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	db, err = gorm.Open("postgres", "port=5432 user=postgres dbname=user password=55021007 sslmode=disable")
	if err != nil {
		panic("Failed to connect")
	} else {
		fmt.Println("Connected Sucessfully")
	}
	var allusers []Logindetails
	db.Find(&allusers)
	json.NewEncoder(w).Encode(allusers)
}
func Createuser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	db, err = gorm.Open("postgres", "port=5432 user=postgres dbname=user password=55021007 sslmode=disable")
	if err != nil {
		panic("Failed to connect")
	} else {
		fmt.Println("Connection Successfull")
	}

	vars := mux.Vars(r)
	Name := vars["Name"]
	Email := vars["Email"]
	Pass := vars["Pass"]
	Roll := vars["Roll"]
	db.Create(&Logindetails{Name: Name, Email: Email, Pass: Pass, Roll: Roll})
	fmt.Fprintf(w, "New user created successfully")

}
func Createbooking(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	db, err = gorm.Open("postgres", "port=5432 user=postgres dbname=user password=55021007 sslmode=disable")
	if err != nil {
		panic("Failed to connect")
	} else {
		fmt.Println("Connected successfully")
	}
	vars := mux.Vars(r)
	Email := vars["Email"]
	Bus_id, err := strconv.Atoi(vars["Bus_id"])
	if err != nil {
		fmt.Println("Cannot able to convert")
	}
	No_of_tickets_booked, err := strconv.Atoi(vars["No_of_tickets_booked"])
	if err != nil {
		fmt.Println("Cannot able to convert")
	}
	Temp, err := strconv.Atoi(vars["Total_amount"])
	if err != nil {
		fmt.Println("Cannot able to convert")
	}
	Total_amount := Temp * No_of_tickets_booked
	db.Create(&Bookingdetails{Email: Email, Bus_id: Bus_id, No_of_tickets_booked: No_of_tickets_booked, Total_amount: Total_amount})
	fmt.Println("Tickets Booked")
	mailSending(Email)
}

func Createbus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	db, err = gorm.Open("postgres", "port=5432 user=postgres dbname=user password=55021007 sslmode=disable")
	if err != nil {
		panic("Failed to connect")
	} else {
		fmt.Println("Connection Successfull")
	}

	vars := mux.Vars(r)
	// ID := vars["Id"]
	Start := vars["Start"]
	Destination := vars["End"]
	Name := vars["Name"]
	Desc := vars["Desc"]
	Type := vars["Type"]
	Wifi := vars["Wifi"]
	Water := vars["Water"]
	Refreshments := vars["Refreshments"]
	Capacity, err := strconv.Atoi(vars["Capacity"])
	if err != nil {
		fmt.Println("cannot able to convert")
	}
	Fare, err := strconv.Atoi(vars["Fare"])
	if err != nil {
		fmt.Println("cannot able to convert")
	}
	// Img := vars["Img"]
	db.Create(&Busdetails{Start: Start, Destination: Destination, Name: Name, Desc: Desc, Type: Type, Wifi: Wifi, Water: Water, Refreshments: Refreshments, Capacity: Capacity, Fare: Fare /*Img: Img*/})
	fmt.Fprintf(w, "New bus created successfully")

}

func Returnuser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	db, err := gorm.Open("postgres", "port=5432 user=postgres dbname=user password=55021007 sslmode=disable")
	if err != nil {
		panic("Failed")
	} else {
		fmt.Println("Connected")
	}
	vars := mux.Vars(r)
	Email := vars["Email"]
	Pass := vars["Pass"]
	var logindetails []Logindetails
	db.Where("Email=? AND Pass=?", Email, Pass).Find(&logindetails)
	json.NewEncoder(w).Encode(logindetails)

}

func Returnbus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	db, err := gorm.Open("postgres", "port=5432 user=postgres dbname=user password=55021007 sslmode=disable")
	if err != nil {
		panic("Failed")
	} else {
		fmt.Println("Connected")
	}
	vars := mux.Vars(r)
	// ID := vars["Id"]
	Start := vars["Start"]
	End := vars["End"]
	var busdetails []Busdetails
	db.Where("Start = ? AND Destination = ?", Start, End).Order("count desc").Find(&busdetails)
	json.NewEncoder(w).Encode(busdetails)
}

func Updateseatcount(w http.ResponseWriter, r *http.Request) {

	db, err := gorm.Open("postgres", "port=5432 user=postgres dbname=user password=55021007 sslmode=disable")
	if err != nil {
		panic("Failed")
	} else {
		fmt.Println("Connected")
	}
	defer db.Close()
	vars := mux.Vars(r)
	Id, err := strconv.Atoi(vars["Id"])
	if err != nil {
		fmt.Println("cannot able to convert")
	}
	No, err := strconv.Atoi(vars["No"])
	if err != nil {
		fmt.Println("cannot able to convert")
	}
	var bus Busdetails
	db.Where("Id = ?", Id).Find(&bus)
	bus.Capacity = bus.Capacity - No
	bus.Count = bus.Count + No
	db.Save(&bus)
	json.NewEncoder(w).Encode(bus)
	fmt.Println("Seat Updated")
}

func Returnbooking(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	db, err := gorm.Open("postgres", "port=5432 user=postgres dbname=user password=55021007 sslmode=disable")
	if err != nil {
		panic("Failed")
	} else {
		fmt.Println("Connected")
	}
	vars := mux.Vars(r)
	Email := vars["Email"]
	var bookingdetails []Bookingdetails
	db.Where("Email=?", Email).Find(&bookingdetails)
	json.NewEncoder(w).Encode(bookingdetails)

}
func Returnusercount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	db, err := gorm.Open("postgres", "port=5432 user=postgres dbname=user password=55021007 sslmode=disable")
	if err != nil {
		panic("Failed")
	} else {
		fmt.Println("Connected Return")
	}
	var bookingdetails []Bookingdetails
	db.Order("total_amount desc").Find(&bookingdetails)
	// db.Table("bookingdetails").Select("bookingdetails.Email").Scan(&bookingdetails)
	json.NewEncoder(w).Encode(&bookingdetails)

}

func Createtransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	db, err = gorm.Open("postgres", "port=5432 user=postgres dbname=user password=55021007 sslmode=disable")
	if err != nil {
		panic("Failed to connect")
	} else {
		fmt.Println("Connection Successfull")
	}

	vars := mux.Vars(r)
	Email := vars["Email"]
	Total, err := strconv.Atoi(vars["Total"])
	if err != nil {
		fmt.Println("cannot able to convert")
	}
	db.Create(&Result{Email: Email, Total: Total})

}

func Updatetotal(w http.ResponseWriter, r *http.Request) {

	db, err := gorm.Open("postgres", "port=5432 user=postgres dbname=user password=55021007 sslmode=disable")
	if err != nil {
		panic("Failed")
	} else {
		fmt.Println("Connected")
	}
	defer db.Close()
	vars := mux.Vars(r)
	Email := vars["Email"]
	No, err := strconv.Atoi(vars["No"])
	if err != nil {
		fmt.Println("cannot able to convert")
	}
	var tot Result
	db.Where("Email = ?", Email).Find(&tot)
	tot.Total = tot.Total + No
	db.Save(&tot)
	json.NewEncoder(w).Encode(tot)
	fmt.Println("Seat Updated")
}

func Returnresult(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	db, err := gorm.Open("postgres", "port=5432 user=postgres dbname=user password=55021007 sslmode=disable")
	if err != nil {
		panic("Failed")
	} else {
		fmt.Println("Connected")
	}
	var result []Result
	db.Limit(5).Order("total desc").Find(&result)
	json.NewEncoder(w).Encode(result)
}
