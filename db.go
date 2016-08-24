package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"os"
	"time"
)

type Expense struct {
	Amount   float64   `bson:"Amount"`
	Date     time.Time `bson:"Date"`
	Category string    `bson:"Category"`
}

type DayTotal struct {
	Date   time.Time
	Amount float64
}

type CategoryTotal struct {
	Category string
	Amount   float64
}

func connect() (session *mgo.Session) {
	connectURL := "localhost"
	session, err := mgo.Dial(connectURL)
	if err != nil {
		log.Fatalf("Can't connect to mongo, go error %v\n", err)
	}
	session.SetSafe(&mgo.Safe{})
	return session
}

func getDayTotals() (dayTotals []DayTotal) {
	session := connect()
	defer session.Close()

	collection := session.DB("expenseManager").C("expenses")

	//Group by day, the amount spent on a given day
	pipe := collection.Pipe([]bson.M{{"$group": bson.M{"_id": "$Date",
		"TotalAmount": bson.M{"$sum": "$Amount"}}}, {"$sort": bson.M{"_id": 1}}})
	iter := pipe.Iter()

	var x map[string]interface{}
	for iter.Next(&x) {
		dayTotals = append(dayTotals, DayTotal{Date: x["_id"].(time.Time),
			Amount: x["TotalAmount"].(float64)})
	}
	return dayTotals
}
func getCategoryTotals() (categoryTotals []CategoryTotal) {
	session := connect()
	defer session.Close()

	collection := session.DB("expenseManager").C("expenses")

	//Group by category, the amount spent on a category
	pipe := collection.Pipe([]bson.M{{"$group": bson.M{"_id": "$Category",
		"TotalAmount": bson.M{"$sum": "$Amount"}}}})
	iter := pipe.Iter()
	var x map[string]interface{}
	for iter.Next(&x) {
		categoryTotals = append(categoryTotals,
			CategoryTotal{Category: x["_id"].(string),
				Amount: x["TotalAmount"].(float64)})
	}
	return categoryTotals
}
func (exp *Expense) Write() (err error) {
	session := connect()
	defer session.Close()

	collection := session.DB("expenseManager").C("expenses")
	err = collection.Insert(exp)
	if err != nil {
		log.Fatal(err)
	}
	return err
}
