package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var config = Config{}
var dao = PatientsDAO{}

//Route handler create endpoint
func CreatePatientEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	//Parse input
	var patient Patient
	if err := json.NewDecoder(r.Body).Decode(&patient); err != nil {
		respondWithError(w, http.StatusBadRequest, "Patient information is not valid") //invalid json
		return
	}
	//Find last record ID to handle blank input
	if patient.ID == 0 {
		lastRecord, err := dao.FindsLastRecord()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Something went wrong with the server.") //database error
			return
		}
		patient.ID = lastRecord.ID + 1
	}
	//Create patient
	if err := dao.Insert(patient); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Something went wrong with the server.") //database error
		return
	}
	respondWithJson(w, http.StatusCreated, patient)
}

//Route handler read all endpoint
func AllPatientsEndPoint(w http.ResponseWriter, r *http.Request) {
	//Find all records
	patients, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Something went wrong with the server.") //database error
		return
	}
	respondWithJson(w, http.StatusOK, patients)
}

//Route handler read by id endpoint
func FindPatientEndpoint(w http.ResponseWriter, r *http.Request) {
	//PArse ID
	params := mux.Vars(r)
	id := params["id"]
	uid, err := strconv.Atoi(id)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid ID") //invalid input
		return
	}
	//Find record by ID
	patient, err := dao.FindById(uid)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Patient does not exist.") //no record exists
		return
	}
	respondWithJson(w, http.StatusOK, patient)
}

//Route handler search by term endpoint
func SearchPatientEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	//Parse query
	var term, key string
	err := r.ParseForm()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Something went wrong with the server.") //database error
		return
	}
	// range over map
	for local_key, values := range r.Form {
		term = local_key
		// range over []string
		for _, value := range values {
			key = value
		}
	}
	//Find records by search term
	ids, err := dao.FindByTerm(term, key)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Something went wrong with the server.") //database error
		return
	} else if ids != nil {
		respondWithJson(w, http.StatusOK, ids)
	} else {
		respondWithJson(w, http.StatusOK, "Records matching these search terms do not exist.") //no record exists
	}
}

//Route handler update by ID endpoint
func UpdatePatientEndPoint(w http.ResponseWriter, r *http.Request) {
	//Parse ID
	params := mux.Vars(r)
	id := params["id"]
	uid, err := strconv.Atoi(id)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Patient does not exist.") //no record exists
		return
	}
	//Parse input
	var patient Patient
	if err := json.NewDecoder(r.Body).Decode(&patient); err != nil {
		respondWithError(w, http.StatusBadRequest, "Patient information is not valid") //invalid json
		return
	}
	//Update data
	if err := dao.Update(uid, patient); err != nil {
		respondWithError(w, http.StatusBadRequest, "Something went wrong with the server") //database error
		return
	}
	//Get updated data
	patient, err = dao.FindById(uid)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Something went wrong with the server.") //database error
		return
	}
	respondWithJson(w, http.StatusOK, patient)
}

//Route handler delete by ID endpoint
func DeletePatientEndPoint(w http.ResponseWriter, r *http.Request) {
	//Parse ID
	params := mux.Vars(r)
	id := params["id"]
	uid, err := strconv.Atoi(id)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid ID") //invalid input
		return
	}
	//Delete record
	if err := dao.Delete(uid); err != nil {
		respondWithError(w, http.StatusBadRequest, "Something went wrong with the server") //database error
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{})
}

//Respond with error function
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

//Respond with JSON function
func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

//Initializes .toml
func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

//Creates router and sets endpoints
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/patients", CreatePatientEndPoint).Methods("POST")
	r.HandleFunc("/patients", AllPatientsEndPoint).Methods("GET")
	r.HandleFunc("/patients/{id}", FindPatientEndpoint).Methods("GET")
	r.HandleFunc("/search", SearchPatientEndPoint).Methods("GET")
	r.HandleFunc("/patients/{id}", UpdatePatientEndPoint).Methods("PUT")
	r.HandleFunc("/patients/{id}", DeletePatientEndPoint).Methods("DELETE")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
