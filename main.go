package main

import "fmt"

func main() {
	fmt.Println("SendEmail")
	SendEmail("Novo Título", "Mensagem", "fluis@cpqd.com.br", "fabiolf@gmail.com")
}

func SendEmail(subject string, message string, fromEmail string, toEmail string) {
	fmt.Println("Subject: "+subject, "Mensagem: "+message, "De:"+fromEmail, "Para"+toEmail)
}
