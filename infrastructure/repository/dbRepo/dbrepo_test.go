package dbrepo

import (
	"backend/api/presenter"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func init() {

	//The response recorder used to record HTTP responses
	httptest.NewRecorder()
}

func TestGetUser(t *testing.T) {
	// apiUrl := config.DB_HOST + string(rune(config.API_PORT))
	// resource := "/user/"
	data := url.Values{}
	data.Set("email", "tom@example.com")

	resp, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/user", strings.NewReader(data.Encode())) // URL-encoded payload
	defer resp.Body.Close()

	var user presenter.User
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	fmt.Println("sb", sb)
	buf := bytes.NewBuffer(body)
	fmt.Println("clgt", json.NewDecoder(buf).Decode(&user))
	if err := json.NewDecoder(buf).Decode(&user); err != nil {
		t.Errorf("Error decoding response body: %v", err.Error())
	}
	result := user.Name
	expected := "Tom Nguyen"
	if result != expected {
		t.Errorf("Expected: %s. Got: %s.", expected, result)
	}
}
