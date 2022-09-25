package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type URL struct {
	domain   string
	path     string
	endpoint string
}

type Badges struct {
	Links struct {
		Self string `json:"self"`
		Next string `json:"next"`
	} `json:"links"`
	Data []struct {
		ID         string `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
			ImageURL string `json:"imageUrl"`
			Name     string `json:"name"`
			State    string `json:"state"`
		} `json:"attributes"`
	} `json:"data"`
}

func (u *URL) get_url() string {
	return u.domain + "/" + u.path + "/" + u.endpoint
}

func get_request(token string) string {
	url := URL{"https://learningmanager.adobe.com", "primeapi/v2", "badges"}
	client := http.Client{}
	req, err := http.NewRequest("GET", url.get_url(), nil)
	if err != nil {
		fmt.Println("Error occured")
		return ""
	}
	req.Header = http.Header{
		"Accept":        {"application/vnd.api+json"},
		"Content-Type":  {"application/json"},
		"Authorization": {"oauth " + token},
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error Occured")
		return ""
	}
	resp_body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error Occured")
		return ""
	}
	return string(resp_body)
}

func main() {
	badges := Badges{}
	body := get_request("f76d0ae52b188e3d274b95c159649c6f")
	data := json.RawMessage(body)
	mydata, err := data.MarshalJSON()
	if err != nil {
		fmt.Println("Error occured")
		return
	}
	err = json.Unmarshal(mydata, &badges)
	if err != nil {
		fmt.Println("Error Ocuured")
		return
	}
	fmt.Println(badges.Links.Self)
	fmt.Println(badges.Links.Next)
	for _, value := range badges.Data {
		fmt.Println(value.ID)
	}
}
