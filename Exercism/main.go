package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://iqoption.com/api/quotes/EURUSD"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "YOUR_API_KEY")
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}
