package sendmail

import (
	env "docs/app/Env"
	"fmt"
	"io/ioutil"
	"gopkg.in/gomail.v2"
)




func SendGomail(templatePath string, subject string , email  string) error {
	htmlData, err := ioutil.ReadFile(templatePath)
	if err != nil {
		return fmt.Errorf("failed to read template file: %v", err)
	}
	m := gomail.NewMessage()
	m.SetHeader("From", "murtazobroimshoevm4@gmail.com")
	m.SetHeader("To",email) 
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", string(htmlData))

	d := gomail.NewDialer("smtp.gmail.com", 587, "murtazobroimshoevm4@gmail.com", env.Email_code)

	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}
	fmt.Println("Email sent successfully!")

	return nil
}
