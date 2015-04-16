package main

import (
	"net/http"
	"strconv"
	"time"
)

const listenAddress = ":5508"
const sleepPath = "/sleep/{SleepFor:[0-9]+}"
const sleepForParameterName = "SleepFor"
const respondWithStatusPath = "/status/{statuscode:[0-9]+}"
const respondWithDefaultStatusPath = "/status"
const statusCodeParameterName = "statuscode"

func StartRouter() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", router)
	http.ListenAndServe(listenAddress, mux)
}

func router(w http.ResponseWriter, r *http.Request) {
	rawQuery := r.URL.RawQuery
	queryString, _, _ := SplitRawQuery(rawQuery)
	//queryValue := queryFields[1]
	switch {
	case queryString == "sleep":
		//sleepFor(w, r, queryValue)
	case queryString == "status":
		//respondWithStatus(w, r, queryValue)
	}
}

func respondWithStatus(w http.ResponseWriter, r *http.Request, queryValue string) {
	status, err := strconv.Atoi(queryValue)

	if err != nil {
		w.WriteHeader(200)
	} else {

		w.WriteHeader(status)
	}
	w.Write([]byte("done!"))
}

func sleepFor(w http.ResponseWriter, request *http.Request, queryValue string) {
	sleepFor, _ := strconv.Atoi(queryValue)
	time.Sleep(time.Duration(sleepFor) * time.Second)
	w.Write([]byte("done!"))
}

func SplitRawQuery(rawQuery string) (string, string, error) {

	return "sleep", "", nil
}
