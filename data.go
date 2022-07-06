package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func getAllSchedule(w http.ResponseWriter, r *http.Request) {
	var schedule Schedule
	var activity Activity
	var array_schedule []Schedule
	var array_activity []Activity
	var response ResponseSchedule

	db := connect()
	defer db.Close()

	rows, err := db.Query("SELECT schedule_id, judul, start_at, end_at FROM schedule")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&schedule.Id, &schedule.Judul, &schedule.Start_at, &schedule.End_at); err != nil {
			log.Fatal(err.Error())
		}else {
			array_activity = nil

			rows, err := db.Query("SELECT activity_id, schedule_id, activity FROM activity WHERE schedule_id=?", schedule.Id)
			if err != nil {
				log.Print(err)
			}

			for rows.Next() {
				if err := rows.Scan(&activity.Id, &activity.Schedule_id, &activity.Activity); err != nil {
					log.Fatal(err.Error())
				}else {
					array_activity = append(array_activity, activity)
				}
			}

			schedule.Activities = array_activity
			array_schedule 		= append(array_schedule, schedule)
		}
	}

	response.Code = 200
	response.Data = array_schedule

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func getScheduleByRangeDatetime(w http.ResponseWriter, r *http.Request) {
	var schedule Schedule
	var activity Activity
	var array_schedule []Schedule
	var array_activity []Activity
	var response ResponseSchedule

	start_at	:= r.URL.Query().Get("start_at")
	end_at		:= r.URL.Query().Get("end_at")
	
	db := connect()
	defer db.Close()

	rows, err := db.Query("SELECT schedule_id, judul, start_at, end_at FROM schedule WHERE start_at>=? AND end_at<=?", start_at, end_at)
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&schedule.Id, &schedule.Judul, &schedule.Start_at, &schedule.End_at); err != nil {
			log.Fatal(err.Error())
		}else {
			array_activity = nil

			rows, err := db.Query("SELECT activity_id, schedule_id, activity FROM activity WHERE schedule_id=?", schedule.Id)
			if err != nil {
				log.Print(err)
			}

			for rows.Next() {
				if err := rows.Scan(&activity.Id, &activity.Schedule_id, &activity.Activity); err != nil {
					log.Fatal(err.Error())
				}else {
					array_activity = append(array_activity, activity)
				}
			}

			schedule.Activities = array_activity
			array_schedule 		= append(array_schedule, schedule)
		}
	}

	response.Code = 200
	response.Data = array_schedule

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func getScheduleByJudul(w http.ResponseWriter, r *http.Request) {
	var schedule Schedule
	var activity Activity
	var array_schedule []Schedule
	var array_activity []Activity
	var response ResponseSchedule

	judul := "%" + r.URL.Query().Get("judul") + "%"
	
	db := connect()
	defer db.Close()

	rows, err := db.Query("SELECT schedule_id, judul, start_at, end_at FROM schedule WHERE judul LIKE ?", judul)
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&schedule.Id, &schedule.Judul, &schedule.Start_at, &schedule.End_at); err != nil {
			log.Fatal(err.Error())
		}else {
			array_activity = nil

			rows, err := db.Query("SELECT activity_id, schedule_id, activity FROM activity WHERE schedule_id=?", schedule.Id)
			if err != nil {
				log.Print(err)
			}

			for rows.Next() {
				if err := rows.Scan(&activity.Id, &activity.Schedule_id, &activity.Activity); err != nil {
					log.Fatal(err.Error())
				}else {
					array_activity = append(array_activity, activity)
				}
			}

			schedule.Activities = array_activity
			array_schedule 		= append(array_schedule, schedule)
		}
	}

	response.Code = 200
	response.Data = array_schedule

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func getScheduleByScheduleId(w http.ResponseWriter, r *http.Request) {
	var schedule Schedule
	var activity Activity
	var array_schedule []Schedule
	var array_activity []Activity
	var response ResponseSchedule

	id := r.URL.Query().Get("id")
	
	db := connect()
	defer db.Close()

	rows, err := db.Query("SELECT schedule_id, judul, start_at, end_at FROM schedule WHERE schedule_id=?", id)
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&schedule.Id, &schedule.Judul, &schedule.Start_at, &schedule.End_at); err != nil {
			log.Fatal(err.Error())
		}else {
			array_activity = nil

			rows, err := db.Query("SELECT activity_id, schedule_id, activity FROM activity WHERE schedule_id=?", schedule.Id)
			if err != nil {
				log.Print(err)
			}

			for rows.Next() {
				if err := rows.Scan(&activity.Id, &activity.Schedule_id, &activity.Activity); err != nil {
					log.Fatal(err.Error())
				}else {
					array_activity = append(array_activity, activity)
				}
			}

			schedule.Activities = array_activity
			array_schedule 		= append(array_schedule, schedule)
		}
	}

	response.Code = 200
	response.Data = array_schedule

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func addSchedule(w http.ResponseWriter, r *http.Request) {
	var schedule Schedule
	var activity Activity
	var array_schedule []Schedule
	var array_activity []Activity
	var response ResponseSchedule
	
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	
	schedule_id, err	:= strconv.Atoi(r.Form.Get("schedule_id"))
	if err != nil {
		log.Print(err)
	}
	judul 				:= r.Form.Get("judul")
	start_at			:= r.Form.Get("start_at")
	end_at				:= r.Form.Get("end_at")
	activity_id, err	:= strconv.Atoi(r.Form.Get("activity_id"))
	if err != nil {
		log.Print(err)
	}
	activities			:= r.Form.Get("activities")
	
	_, err = db.Exec("INSERT INTO schedule (schedule_id, judul, start_at, end_at) VALUES (?,?,?,?)", schedule_id, judul, start_at, end_at)
	if err != nil {
		log.Print(err)
	}

	_, err = db.Exec("INSERT INTO activity (activity_id, schedule_id, activity) VALUES (?,?,?)", activity_id, schedule_id, activities)
	if err != nil {
		log.Print(err)
	}
	
	activity.Id 			= activity_id
	activity.Schedule_id	= schedule_id
	activity.Activity		= activities
	array_activity 			= append(array_activity, activity)
	
	schedule.Id				= schedule_id
	schedule.Judul			= judul
	schedule.Start_at		= start_at
	schedule.End_at			= end_at
	schedule.Activities		= array_activity
	array_schedule 			= append(array_schedule, schedule)

	response.Code = 201
	response.Data = array_schedule

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func updateSchedule(w http.ResponseWriter, r *http.Request) {
	var schedule Schedule
	var activity Activity
	var array_schedule []Schedule
	var array_activity []Activity
	var response ResponseSchedule

	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	schedule_id, err	:= strconv.Atoi(r.Form.Get("schedule_id"))
	if err != nil {
		log.Print(err)
	}
	judul 				:= r.Form.Get("judul")
	start_at			:= r.Form.Get("start_at")
	end_at				:= r.Form.Get("end_at")
	activity_id, err	:= strconv.Atoi(r.Form.Get("activity_id"))
	if err != nil {
		log.Print(err)
	}
	activities			:= r.Form.Get("activities")

	_, err = db.Exec("UPDATE schedule SET judul=?, start_at=?, end_at=? WHERE schedule_id=?", judul, start_at, end_at, schedule_id)
	if err != nil {
		log.Print(err)
	}

	_, err = db.Exec("UPDATE activity SET activity=? WHERE activity_id=?", activities, activity_id)
	if err != nil {
		log.Print(err)
	}

	activity.Id 			= activity_id
	activity.Schedule_id	= schedule_id
	activity.Activity		= activities
	array_activity 			= append(array_activity, activity)
	
	schedule.Id				= schedule_id
	schedule.Judul			= judul
	schedule.Start_at		= start_at
	schedule.End_at			= end_at
	schedule.Activities		= array_activity
	array_schedule 			= append(array_schedule, schedule)

	response.Code = 200
	response.Data = array_schedule

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func deleteScheduleMultipart(w http.ResponseWriter, r *http.Request) {
	var response Response

	db := connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}
	
	schedule_id := r.FormValue("schedule_id")

	_, err = db.Exec("DELETE FROM schedule WHERE schedule_id=?", schedule_id)
	if err != nil {
		log.Print(err)
	}

	response.Code = 200

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func getActivityByScheduleId(w http.ResponseWriter, r *http.Request) {
	var activity Activity
	var array_activity []Activity
	var response ResponseActivity

	id := r.URL.Query().Get("id")
	
	db := connect()
	defer db.Close()

	rows, err := db.Query("SELECT activity_id, schedule_id, activity FROM activity WHERE schedule_id=?", id)
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&activity.Id, &activity.Schedule_id, &activity.Activity); err != nil {
			log.Fatal(err.Error())
		}else {
			array_activity = append(array_activity, activity)
		}
	}

	response.Code = 200
	response.Data = array_activity

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func getActivityByActivityId(w http.ResponseWriter, r *http.Request) {
	var activity Activity
	var array_activity []Activity
	var response ResponseActivity

	id := r.URL.Query().Get("id")
	
	db := connect()
	defer db.Close()

	rows, err := db.Query("SELECT activity_id, schedule_id, activity FROM activity WHERE activity_id=?", id)
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&activity.Id, &activity.Schedule_id, &activity.Activity); err != nil {
			log.Fatal(err.Error())
		}else {
			array_activity = append(array_activity, activity)
		}
	}

	response.Code = 200
	response.Data = array_activity

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func addActivityForSchedule(w http.ResponseWriter, r *http.Request) {
	var schedule Schedule
	var activity Activity
	var array_schedule []Schedule
	var array_activity []Activity
	var response ResponseSchedule
	
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	
	schedule_id, err	:= strconv.Atoi(r.Form.Get("schedule_id"))
	if err != nil {
		log.Print(err)
	}
	activity_id, err	:= strconv.Atoi(r.Form.Get("activity_id"))
	if err != nil {
		log.Print(err)
	}
	activities			:= r.Form.Get("activities")
	
	_, err = db.Exec("INSERT INTO activity (activity_id, schedule_id, activity) VALUES (?,?,?)", activity_id, schedule_id, activities)
	if err != nil {
		log.Print(err)
	}

	rows, err := db.Query("SELECT schedule_id, judul, start_at, end_at FROM schedule WHERE schedule_id=?", schedule_id)
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&schedule.Id, &schedule.Judul, &schedule.Start_at, &schedule.End_at); err != nil {
			log.Fatal(err.Error())
		}else {
			array_activity = nil

			rows, err := db.Query("SELECT activity_id, schedule_id, activity FROM activity WHERE schedule_id=?", schedule.Id)
			if err != nil {
				log.Print(err)
			}

			for rows.Next() {
				if err := rows.Scan(&activity.Id, &activity.Schedule_id, &activity.Activity); err != nil {
					log.Fatal(err.Error())
				}else {
					array_activity = append(array_activity, activity)
				}
			}

			schedule.Activities = array_activity
			array_schedule 		= append(array_schedule, schedule)
		}
	}

	response.Code = 200
	response.Data = array_schedule

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func updateActivityForSchedule(w http.ResponseWriter, r *http.Request) {
	var schedule Schedule
	var activity Activity
	var array_schedule []Schedule
	var array_activity []Activity
	var response ResponseSchedule
	
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	
	schedule_id, err	:= strconv.Atoi(r.Form.Get("schedule_id"))
	if err != nil {
		log.Print(err)
	}
	activity_id, err	:= strconv.Atoi(r.Form.Get("activity_id"))
	if err != nil {
		log.Print(err)
	}
	activities			:= r.Form.Get("activities")
	
	_, err = db.Exec("UPDATE activity SET activity=? WHERE activity_id=? AND schedule_id=?", activities, activity_id, schedule_id)
	if err != nil {
		log.Print(err)
	}

	rows, err := db.Query("SELECT schedule_id, judul, start_at, end_at FROM schedule WHERE schedule_id=?", schedule_id)
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&schedule.Id, &schedule.Judul, &schedule.Start_at, &schedule.End_at); err != nil {
			log.Fatal(err.Error())
		}else {
			array_activity = nil

			rows, err := db.Query("SELECT activity_id, schedule_id, activity FROM activity WHERE schedule_id=?", schedule.Id)
			if err != nil {
				log.Print(err)
			}

			for rows.Next() {
				if err := rows.Scan(&activity.Id, &activity.Schedule_id, &activity.Activity); err != nil {
					log.Fatal(err.Error())
				}else {
					array_activity = append(array_activity, activity)
				}
			}

			schedule.Activities = array_activity
			array_schedule 		= append(array_schedule, schedule)
		}
	}

	response.Code = 200
	response.Data = array_schedule

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func deleteActivityForScheduleMultipart(w http.ResponseWriter, r *http.Request) {
	var schedule Schedule
	var activity Activity
	var array_schedule []Schedule
	var array_activity []Activity
	var response ResponseSchedule
	
	db := connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}
	
	schedule_id := r.FormValue("schedule_id")
	activity_id := r.FormValue("activity_id")
		
	_, err = db.Exec("DELETE FROM activity WHERE activity_id=? AND schedule_id=?", activity_id, schedule_id)
	if err != nil {
		log.Print(err)
	}

	rows, err := db.Query("SELECT schedule_id, judul, start_at, end_at FROM schedule WHERE schedule_id=?", schedule_id)
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&schedule.Id, &schedule.Judul, &schedule.Start_at, &schedule.End_at); err != nil {
			log.Fatal(err.Error())
		}else {
			array_activity = nil

			rows, err := db.Query("SELECT activity_id, schedule_id, activity FROM activity WHERE schedule_id=?", schedule.Id)
			if err != nil {
				log.Print(err)
			}

			for rows.Next() {
				if err := rows.Scan(&activity.Id, &activity.Schedule_id, &activity.Activity); err != nil {
					log.Fatal(err.Error())
				}else {
					array_activity = append(array_activity, activity)
				}
			}

			schedule.Activities = array_activity
			array_schedule 		= append(array_schedule, schedule)
		}
	}

	response.Code = 200
	response.Data = array_schedule

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}