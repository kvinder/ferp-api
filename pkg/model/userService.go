package model

import (
	"time"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(data map[string]string) (map[string]string, int) {
	db := getConnection()
	defer db.Close()
	t := time.Now()
	now := t.Format("2006-01-02 15:04:05")

	resp := map[string]string{}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data["password"]), bcrypt.DefaultCost)
	checkErr(err)
	sqlQuery := `INSERT INTO APP_USER (employee_id,name,username,email,sex,
		telephone,department,password,status,createDate,updateDate) 
		VALUES (?,?,?,?,?,?,?,?,?,?,?)`
	_, err = db.Exec(
		sqlQuery,
		data["empnumber"], 
		data["name"], 
		data["username"], 
		data["email"], 
		data["sex"], 
		data["telephone"],
		data["department"],
		hashedPassword,
		"REGISTER",
		now, 
		now,
	)
	if err != nil {
		resp["error"] = err.Error()
		return resp, 500
	}
	
	resp["message"] = "successfully"
	return resp, 200
}