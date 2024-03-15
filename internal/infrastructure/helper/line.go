package helper

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func LineNoti(token, message *string) {
	url := "https://notify-api.line.me/api/notify"

	payload := strings.NewReader("message=" + *message)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Bearer "+*token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer res.Body.Close()

	// Check the response
	if res.StatusCode != http.StatusOK {
		log.Fatalf("Error sending Line notification. Status: %s", res.Status)
	}

	log.Println("Line notification sent successfully!")
}
