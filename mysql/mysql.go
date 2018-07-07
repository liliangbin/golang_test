package main
import (
	"fmt"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"log"
)

func main()  {
	db,err := sql.Open("mysql","root:123456@tcp(localhost:3306)/liliangbin")
	err = db.Ping()
	if err!=nil{
		log.Println(err)
		fmt.Println("lihai l ")
	}
	
	defer db.Close()
}