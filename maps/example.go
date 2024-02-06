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

//go:generate mockgen -destination=mocks/mock_iPostgresDB.go -source example.go

type iPostgresDB interface {
	PostRecord(Home) (sql.Result, error)
	GetRecord() ([]Home, error)
	DelRecord() (sql.Result, error)
}

type PostgresDB struct {
	db *sql.DB
}

func newDB() (iPostgresDB, error) {
	connStr := "user=postgres dbname=postgres password=123456 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	defer db.Close()

	return &PostgresDB{db}, nil
}

func (p PostgresDB) PostRecord(h Home) (sql.Result, error) {

	row1, err := p.db.Exec(" insert into public.home (state, zip) values ($1, $2)", h.State, h.Zip)

	if err != nil {
		return nil, err
	}
	return row1, nil
}

func (p PostgresDB) GetRecord() ([]Home, error) {
	rows, err := p.db.Query("SELECT  * from public.home")

	if err != nil {
		return nil, err
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

	return home, nil

}

func (p PostgresDB) DelRecord() (sql.Result, error) {
	res1, err := p.db.Exec("delete from public.home")
	if err != nil {
		return nil, err
	}

	return res1, nil
}

func MJSON() {

	db, err := newDB()

	if err != nil {
		log.Fatal(err)
	}

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

	pos, err := db.PostRecord(slc[0])
	if err != nil {
		log.Print(err)
	}

	fmt.Println(pos.RowsAffected())

}

func WrapperFunc(db iPostgresDB, h Home) (sql.Result, error) {
	return db.PostRecord(h)
}
