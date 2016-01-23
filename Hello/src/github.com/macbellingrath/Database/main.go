package main 

import (
    "database/sql"
    _ "github.com/lib/pq"
	"log"

)

func main()  {
    db, err := sql.Open("postgres","user=MacBellingrath dbname=macdb sslmode=disable")
    if err != nil {
        log.Fatal("Error: The data source arguments are not valid")
    }
    err = db.Ping()
    
    if err != nil {
  
    log.Fatal("Error: Could not establish a connection with the database")
    }
    queryStmt, err := db.Prepare("SELECT name FROM jobs WHERE id = $1")
       
       if err != nil {
    log.Fatal(err)
    }
    var name string
    err = queryStmt.QueryRow(1).Scan(&name)
    // r,e := queryStmt.Exec(1)
    // if e != nil {
    //     log.Fatal(e)
    // } else {
    //     // print(r.RowsAffected)
    //     print(name)
    // }
  print(name)
    
    if err == sql.ErrNoRows {
  log.Fatal("No Results Found")
}
if err != nil {
  log.Fatal(err)
}
}



