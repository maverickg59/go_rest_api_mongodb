package main

//Represents patient data model
type Patient struct {
	ID                    int    `json:"id"`
	FirstName             string `json:"first_name"`
	LastName              string `json:"last_name"`
	Gender                string `json:"gender"`
	PhoneNumber           string `json:"phone_number"`
	Email                 string `json:"email"`
	Address               string `json:"address"`
	VisitDate             string `json:"visit_date"`
	Diagnosis             string `json:"diagnosis"`
	DrugCode              string `json:"drug_code"`
	AdditionalInformation []struct {
		Notes      string `json:"notes"`
		NewPatient bool   `json:"new_patient"`
		Race       string `json:"race"`
		Ssn        string `json:"ssn"`
	} `json:"additional_information"`
}

//For use when searching by term
type PatientID struct {
	ID int `json:"id"`
}
