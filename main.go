package main

import (
	"fmt"
	"log"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/getallschedule", getAllSchedule).Methods("GET")
	router.HandleFunc("/getschedulebyrangedatetime", getScheduleByRangeDatetime).Methods("GET")
	router.HandleFunc("/getschedulebyjudul", getScheduleByJudul).Methods("GET")
	router.HandleFunc("/getschedulebyscheduleid", getScheduleByScheduleId).Methods("GET")
	router.HandleFunc("/addschedule", addSchedule).Methods("POST")
	router.HandleFunc("/updateschedule", updateSchedule).Methods("PUT")
	router.HandleFunc("/deleteschedulemultipart", deleteScheduleMultipart).Methods("DELETE")
	router.HandleFunc("/getactivitybyscheduleid", getActivityByScheduleId).Methods("GET")
	router.HandleFunc("/getactivitybyactivityid", getActivityByActivityId).Methods("GET")
	router.HandleFunc("/addactivityforschedule", addActivityForSchedule).Methods("POST")
	router.HandleFunc("/updateactivityforschedule", updateActivityForSchedule).Methods("PUT")
	router.HandleFunc("/deleteactivityforschedulemultipart", deleteActivityForScheduleMultipart).Methods("DELETE")

	http.Handle("/", router)
	fmt.Println("Listen to port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}