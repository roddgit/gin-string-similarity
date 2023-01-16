package configs

import (
	"database/sql"
	"fmt"
	"net/url"

	_ "github.com/sijms/go-ora/v2"
)

var env = EnvConfig()

var ParamDBOra = map[string]string{
	"service":  env["OR_SID"],
	"username": env["OR_USER_DB"],
	"server":   env["OR_HOST_DB"],
	"port":     env["OR_PORT_DB"],
	"password": env["OR_PASSWORD_DB"],
}

func connectORADB(dbParams map[string]string) *sql.DB {
	connectionString := "oracle://" + dbParams["username"] + ":" + dbParams["password"] + "@" + dbParams["server"] + ":" + dbParams["port"] + "/" + dbParams["service"]

	db, err := sql.Open("oracle", connectionString)
	if err != nil {
		panic(fmt.Errorf("error in sql.Open: %w", err))
	}

	err = db.Ping()
	if err != nil {
		panic(fmt.Errorf("error pinging db: %w", err))
	}

	return db
}

func DoDBThings(dbParams map[string]string, logs_id string, req_pmo_raw string, req_core_raw string, req_pmo string, req_core string, treshold float64) {
	connectionString := "oracle://" + dbParams["username"] + ":" + dbParams["password"] + "@" + dbParams["server"] + ":" + dbParams["port"] + "/" + dbParams["service"]
	if val, ok := dbParams["walletLocation"]; ok && val != "" {
		connectionString += "?TRACE FILE=trace.log&SSL=enable&SSL Verify=false&WALLET=" + url.QueryEscape(dbParams["walletLocation"])
	}
	db, err := sql.Open("oracle", connectionString)
	if err != nil {
		panic(fmt.Errorf("error in sql.Open: %w", err))
	}
	defer func() {
		err = db.Close()
		if err != nil {
			fmt.Println("Can't close connection: ", err)
		}
	}()

	err = db.Ping()
	if err != nil {
		panic(fmt.Errorf("error pinging db: %w", err))
	}

	someAdditionalActions(db, logs_id, req_pmo_raw, req_core_raw, req_pmo, req_core, treshold)
}

const insertStatement = `INSERT INTO COMPARE_NAME (LOGS_ID, REQ_PMO_RAW, REQ_CORE_RAW, REQ_PMO, REQ_CORE, TRESHOLD)
VALUES (:logs_id, :req_pmo_raw, :req_core_raw, :req_pmo, :req_core, :treshold )`

func someAdditionalActions(db *sql.DB, logs_id string, req_pmo_raw string, req_core_raw string, req_pmo string, req_core string, treshold float64) {
	stmt, err := db.Prepare(insertStatement)
	HandleError("prepare insert statement", err)

	sqlresult, err := stmt.Exec(logs_id, req_pmo_raw, req_core_raw, req_pmo, req_core, treshold)
	HandleError("execute insert statement", err)
	rowCount, _ := sqlresult.RowsAffected()
	fmt.Println("Inserted number of rows = ", rowCount)

}

func InsertTresholdORA(logs_id string, req_pmo_raw string, req_core_raw string, req_pmo string, req_core string, treshold float64) string {
	var insertStatementStr = `INSERT INTO COMPARE_NAME (LOGS_ID, REQ_PMO_RAW, REQ_CORE_RAW, REQ_PMO, REQ_CORE, TRESHOLD)
							VALUES (:logs_id, :req_pmo_raw, :req_core_raw, :req_pmo, :req_core, :treshold )`

	db := connectORADB(ParamDBOra)

	stmt, err := db.Prepare(insertStatementStr)

	HandleError("prepare insert statement", err)
	sqlresult, err := stmt.Exec(logs_id, req_pmo_raw, req_core_raw, req_pmo, req_core, treshold)
	HandleError("execute insert statement", err)
	rowCount, _ := sqlresult.RowsAffected()
	fmt.Println("Inserted number of rows = ", rowCount)

	// must close in the end
	defer func() {
		err = db.Close()
		if err != nil {
			fmt.Println("Can't close connection: ", err)
		}
	}()

	result := fmt.Sprint(rowCount)

	return result
}
