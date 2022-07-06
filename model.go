package main

type Activity struct {
	Id			int `form:"id" json:"id"`
	Schedule_id	int `form:"schedule_id" json:"schedule_id"`
	Activity	string `form:"activity" json:"activity"`
}

type Schedule struct {
	Id			int `form:"id" json:"id"`
	Judul		string `form:"judul" json:"judul"`
	Start_at	string `form:"start_at" json:"start_at"`
	End_at		string `form:"end_at" json:"end_at"`
	Activities	[]Activity
}

type Response struct {
	Code	int `json:"code"`
}

type ResponseSchedule struct {
	Code	int `json:"code"`
	Data    []Schedule
}

type ResponseActivity struct {
	Code	int `json:"code"`
	Data    []Activity
}