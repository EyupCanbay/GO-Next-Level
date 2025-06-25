/*package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var (
	db    *sql.DB
	dbErr error
)

type City struct {
	Id   int
	Name string
	Code int
}

func main() {

	// preparing db connecting driver
	host := "localhost"
	port := 5432
	user := "postgres"
	password := "postgres"
	dbName := "godb"
	psgInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)
	db, dbErr = sql.Open("postgres", psgInfo)
	if dbErr != nil {
		panic(dbErr)
	}

	//insertCity()
	//selectCity()
	//selectOneCity(1)
	selectWithPreparedStatement("istanbul")
}

func insertCity(city City) {
	result, err := db.Exec("insert into cities(name, code) values('istanbul', 34)")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(result.RowsAffected())
	}
}

func selectCity() {
	//	prepare statement
	var cityList []City
	rows, err := db.Query("select * from cities")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		// arka palnda dbye connection acılıyor ve rows bu connectiona bind oluyor

		//rows.Next() -> elimde kaç tane rows var hala edevam eden rows var mı onun bilgisini bana veriyor
		for rows.Next() {
			var city City
			if err := rows.Scan(&city.Id, &city.Name, &city.Code); err != nil {
				fmt.Println(err)
			} else {
				cityList = append(cityList, city)
			}
		}

		rows.Close()

		fmt.Println(cityList)
	}

}

func selectOneCity(id int) {
	// carefuly, this is a vulneraility (sql injection)
	var city City
	formattedSql := fmt.Sprintf("select * from cities where id= %d", id)

	err := db.QueryRow(formattedSql).Scan(&city.Id, &city.Name, &city.Code)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(city)
	}
}

func selectWithPreparedStatement(cityName string) {
	stmt, err := db.Prepare("select * from cities where name=$1")
	if err != nil {
		fmt.Println(err)
	} else {
		var city City
		if err := stmt.QueryRow(cityName).Scan(&city.Id, &city.Name, &city.Code); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(city)
		}
	}
}
*/
