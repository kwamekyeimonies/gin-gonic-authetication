package utils

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func VerifyPassword(userPassword string, providedPassword string)(bool,string){
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword),[]byte(userPassword))
	check := true
	msg := ""

	if err != nil{
		msg = fmt.Sprintf("Email or Password Incorrect")
		check = false
	}

	return check, msg
}


func HashPassword(password string) string{
	bytes,err :=bcrypt.GenerateFromPassword([]byte(password),14)
	if err != nil{
		log.Panic(err)
		
	}
	return string(bytes)
}