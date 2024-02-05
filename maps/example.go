package maps

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Home struct {
	State string `db:"state"`
	Zip   int    `db:"zip"`
	//Detail map[string]string `json:"det"`
}

func MJSON() {
	slc := make([]Home, 0)

	// m1 := make(map[string]string, 0)
	// m1["City"] = "FV"
	// m1["Street"] = "Red Ced"
	// m2 := make(map[string]string, 0)
	// m2["City"] = "MO"
	// m2["Street"] = "wal"
	a := Home{State: "NC", Zip: 27527}
	b := Home{State: "NC", Zip: 27561}

	slc = append(slc, a, b)

	slc2 := make([]Home, 0)
	c := Home{State: "NC", Zip: 27524}

	slc2 = append(slc2, c)

	slc = append(slc, slc2...)

	connStr := "user=postgres dbname=postgres password=123456 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	row1, err := db.Exec(" insert into public.home (state, zip) values ($1, $2)", slc[2].State, slc[2].Zip)

	if err != nil {
		log.Print(err)
	}
	fmt.Println(row1.LastInsertId())

	// slcjson, err := json.Marshal(slc)
	// if err != nil {
	// 	log.Print(err)
	// }

	// if err = os.Remove("123.json"); err != nil {
	// 	log.Print(err)
	// }

	// err = os.WriteFile("123.json", slcjson, fs.FileMode(0644))
	// if err != nil {
	// 	log.Print(err)
	// }

	// sbyte, err := os.ReadFile("123.json")
	// if err != nil {
	// 	log.Print(err)
	// }

	// home := make([]Home, 0)

	// err = json.Unmarshal(sbyte, &home)
	// if err != nil {
	// 	log.Print(err)
	// }

	// fmt.Println(home)

	rows, err := db.Query("SELECT  * from public.home")

	if err != nil {
		log.Print(err)
	}

	defer rows.Close()

	home := make([]Home, 0)

	for rows.Next() {
		var r Home
		err := rows.Scan(&r.State, &r.Zip)
		if err != nil {
			log.Print(err)
		}
		home = append(home, r)
	}

	fmt.Println(home)

	res1, err := db.Exec("delete from public.home")
	if err != nil {
		log.Print(err)
	}

	fmt.Println(res1.RowsAffected())

}
