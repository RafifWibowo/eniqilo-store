package helper

import (
	"database/sql"
	"log"
	"time"
)

func CheckPhoneExist(db *sql.DB, tableName string, phoneNumber string) (bool, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM public."+tableName+" WHERE \"phoneNumber\" = $1", phoneNumber).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func ValidatePhoneNumber(phone string, errList map[string]string) {
	if phone == "" {
		errList["phoneNumber"] = "Phone number can't be null."
		return
	}
	if phone[0:1] != "+" {
		errList["phoneNumber"] = "Not in phone number format."
		return
	}
	if len(phone) < 10 || len(phone) > 16 {
		errList["phoneNumber"] = "Phone number must be between 10 and 16 characters in length."
		return
	}
}

func ValidateName(name string, errList map[string]string) {
	if name == "" {
		errList["name"] = "Name can't be null."
		return
	}
	if len(name) < 5 || len(name) > 50 {
		errList["name"] = "Name must be between 5 and 50 characters in length."
		return
	}
}

func ConvertToISO860(str string) (string, error){
	t, err := time.Parse(time.RFC3339, str)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return t.Format("2006-01-02T15:04:05Z07:00"), nil
}