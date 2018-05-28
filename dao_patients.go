package main

import (
	"log"
	"strings"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type PatientsDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "patients"
)

//Establishes connection to database
func (p *PatientsDAO) Connect() {
	session, err := mgo.Dial(p.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(p.Database)
}

//Creates database record
func (p *PatientsDAO) Insert(patient Patient) error {
	err := db.C(COLLECTION).Insert(&patient)
	return err
}

//Finds last record in database
func (p *PatientsDAO) FindsLastRecord() (Patient, error) {
	var patient Patient
	err := db.C(COLLECTION).Find(nil).Select(bson.M{"id": "$gt:0"}).Sort("-$natural").Limit(1).One(&patient) //Natural sort was chosen because timestamps are not overwritten in MongoDB and id's should increment by one
	return patient, err
}

//Finds all database records
func (p *PatientsDAO) FindAll() ([]Patient, error) {
	var patients []Patient
	err := db.C(COLLECTION).Find(bson.M{}).All(&patients)
	return patients, err
}

//Finds database record by ID
func (p *PatientsDAO) FindById(id int) (Patient, error) {
	var patient Patient
	err := db.C(COLLECTION).Find(bson.M{"id": id}).One(&patient)
	return patient, err
}

//Searches database record by term
func (p *PatientsDAO) FindByTerm(term string, key string) ([]PatientID, error) {
	var patients []PatientID
	term = strings.Replace(term, "_", "", -1) //Unable to access field names without removing underscore
	err := db.C(COLLECTION).Find(bson.M{term: key}).All(&patients)
	return patients, err
}

//Updates database record
func (p *PatientsDAO) Update(id int, patient Patient) error {
	err := db.C(COLLECTION).Update(bson.M{"id": id}, &patient)
	return err
}

//Deletes database record
func (p *PatientsDAO) Delete(id int) error {
	err := db.C(COLLECTION).Remove(bson.M{"id": id})
	return err
}
