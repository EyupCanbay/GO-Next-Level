package repository

import (
	"database/sql"
	"database_external/entity"
	"fmt"
)

type CityRepo struct {
	Db *sql.DB
}

func NewRepo(db *sql.DB) *CityRepo {
	return &CityRepo{
		Db: db,
	}
}

func (r *CityRepo) Insert(city entity.City) {
	stmt, err := r.Db.Prepare("insert into cities(name, code) values($1, $2)")

	result, err := stmt.Exec(city.Name, city.Code)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(result.RowsAffected())
	}
}

func (r *CityRepo) List() []entity.City {
	var listCity []entity.City

	rows, err := r.Db.Query("select * from cities")
	if err != nil {
		fmt.Println(err)
		return nil
	} else {
		for rows.Next() {
			var city entity.City
			err := rows.Scan(&city.Id, &city.Name, &city.Code)
			if err != nil {
				fmt.Println(err)
			} else {
				listCity = append(listCity, city)
			}
		}
		rows.Close()

		return listCity
	}
}

func (r *CityRepo) GetById(id int) *entity.City {
	var city entity.City

	rows, err := r.Db.Query("SELECT id, name, code FROM cities WHERE id = $1", id)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&city.Id, &city.Name, &city.Code)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		return &city
	}

	return nil
}
