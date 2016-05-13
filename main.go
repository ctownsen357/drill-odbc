package main

import (
	"fmt"
	"database/sql"
	_ "github.com/alexbrainman/odbc"
	)


func main() {
	connStr := `Driver=/opt/mapr/drillodbc/lib/64/libmaprdrillodbc64.so;
                        Catalog=DRILL;
                        Schema=hivestg;
                        ConnectionType=Direct;
                        Host=localhost;
                        Port=31010;
                        AuthenticationType=No Authentication`

	db, err := sql.Open("odbc", connStr)
	if err != nil {
		panic(err)	
	}

	sql := "select * from dfs.`/home/ctownsend/projects/drill-odbc/test_data/`"

	type rslt struct{
		PolID int
		LocID int
		SomeVal int
	}

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	
	for rows.Next() {
		var r rslt
		err = rows.Scan(&r.PolID, &r.LocID, &r.SomeVal)
		if err != nil {
			panic(err)
		}
		fmt.Println(r)
	} 
		

}
