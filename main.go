package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func heure(w http.ResponseWriter) {
	now := time.Now()
	fmt.Fprintf(w, "%02dh%02d\n", now.Hour(), now.Minute())
}

func Handleheure(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		heure(w)
	}
}

func get_random_number(min int, max int) int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	random_number := r1.Intn(max-min) + min
	return (random_number)
}

func dice(w http.ResponseWriter, max int, format string) {
	min := 1

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	random_number := r1.Intn(max-min) + min
	fmt.Fprintf(w, format, random_number)
}

func Handledice(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		dice(w, 1000, "%04d\n")
	}
}

type d struct {
	max    int
	format string
}

func dices(w http.ResponseWriter) {
	array_dices := []d{{int(2), "%d"}, {int(4), "%d"}, {int(6), "%d"}, {int(8), "%d"}, {int(10), "%02d"}, {int(12), "%02d"}, {int(20), "%02d"}, {int(100), "%03d"}}

	for i := 0; i < 15; i++ {
		rand_dice := get_random_number(0, 8)
		dice(w, array_dices[rand_dice].max, array_dices[rand_dice].format)
		if i < 14 {
			fmt.Fprintf(w, " ")
		}
	}
	fmt.Fprintf(w, "\n")
}

func Handledices(w http.ResponseWriter, r *http.Request) {
	//array_dices := []d{{int(2), "%d"}, {int(4), "%d"}, {int(6), "%d"}, {int(8), "%d"}, {int(10), "%02d"}, {int(12), "%02d"}, {int(20), "%02d"}, {int(100), "%03d"}}
	//type_number := 0
	//type_dice := false
	switch r.Method {
	case http.MethodGet:
		//²values := r.URL.Query()
		//for k, v := range values {
		//	if k == "type" {
		//		type_dice = true
		//		type_number, err := strconv.Atoi(v[0])
		//		if err != nil {
		//		}
		//	}
		//}
		//if type_dice == true {
		//	des_number := 0
		//	for i := 0; i < len(array_dices); i++ {
		//		if type_number == array_dices[i].max {
		//			des_number = i
		//		}
		//	}
		//	for i := 0; i < 15; i++ {
		//		dice(w, array_dices[des_number].max, array_dices[des_number].format)
		//		if i < 14 {
		//			fmt.Fprintf(w, " ")
		//		}
		//	}
		//	fmt.Fprintf(w, "\n")
		//} else {
		dices(w)
		//}
	}
}

func main() {
	http.HandleFunc("/", Handleheure)
	http.HandleFunc("/dice", Handledice)
	http.HandleFunc("/dices", Handledices)
	http.ListenAndServe(":4567", nil)
}
