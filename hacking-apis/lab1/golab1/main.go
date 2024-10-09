package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	client := http.Client{}
	// userID := 1
	// user, err := getUser(client, userID)
	// if err != nil {
	// 	fmt.Println("Error getting user:", err)
	// 	os.Exit(1)
	// }
	// fmt.Printf("User %d : %+v\n", userID, user)
	lastUserID, err := testUsers(client, 1, 9999)
	if err != nil {
		fmt.Println("Last user requested:", lastUserID)
		fmt.Println("Failed with error:", err)
		fmt.Println(">> The answer is:", lastUserID-1)
	}
}

func testUsers(client http.Client, start, limit int) (int, error) {
	for i := start; i <= limit; i++ {
		_, err := getUser(client, i)
		if err != nil {
			return i, err
		}
		fmt.Println("Got user with ID =", i)
	}
	return limit, nil
}

func getUser(client http.Client, userID int) (*UserResponse, error) {
	URL := fmt.Sprintf("https://reqres.in/api/users/%d", userID)
	response, err := client.Get(URL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(response.Body)
		return nil, fmt.Errorf("response came back with status code %d and body %s", response.StatusCode, string(body))
	}
	userData, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	user := UserResponse{}
	err = json.Unmarshal(userData, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

type UserResponse struct {
	Data    UserData    `json:"data"`
	Support SupportData `json:"support"`
}

type UserData struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Avatar    string `json:"avatar"`
}

type SupportData struct {
	URL  string `json:"url"`
	Text string `json:"text"`
}
