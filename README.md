This is a small app which uses
 - golang's net/http
 - mgo driver, which connects to MongoDB

This webservice helps in expense management.

Following are the REST APIs it provides

1. GET request
-----------
a. "/totals/day"

    -- this returns the total amount spent on that day

    e.g. Request-> curl -G http://localhost:8080/totals/day
         Response-> {[{"Date":"2015-11-30T05:30:00+05:30","Amount":600}]}

b. "/totals/category"
    -- this returns the category wise total amount spent

    e.g. Request-> curl -G http://localhost:8080/totals/category
         Response-> [{"Category":"Food","Amount":600}]

2. POST request
---------------
 a. "/expense"

   -- this is used to add expenses in the database

   e.g. curl -d '{"Amount":300,"Category":"Food","Date":"2015-11-30T00:00:00Z"}' http://localhost:8080/expense

  will add the expense 300 to db with other details


The code can be built using
 go install expenseManager (given that your GOPATH is set correctly). This
creates a single binary expenseManager

The code can be run using
 expenseManager
This starts the server


Prerequisites
 - MongoDB should be installed
 - mongod service should be running
 - go get gopkg.in/mgo.v2
 - go get github.com/gorilla/mux
