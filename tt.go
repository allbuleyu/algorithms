package main

import (
	"github.com/allbuleyu/algorithms/algorithms/ch1"
	"fmt"

	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	fmt.Println(ch1.Mystery(2, 25), ch1.Mystery(3, 11))

	fmt.Println(ch1.Ln(8))



	//bd := strings.NewReader("client_id=HAPwktAi+8bwoEbH8E905BRpncj9M8UD&client_secret=02&grant_type=authorization_code&code=112311421315870920081181443233244&format=json&cb=")

	//client := http.DefaultClient
	//request, err := http.NewRequest("POST", "https://api.passport.pptv.com/accessToken", bd)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//
	//request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//resp, err := client.Do(request)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	data := url.Values{}
	data.Add("client_id", "HAPwktAi+8bwoEbH8E905BRpncj9M8UD")
	data.Add("client_secret", "02")
	data.Add("grant_type", "authorization_code")
	data.Add("code", "143342123215870934654621113211413")
	data.Add("format", "json")
	resp, _ := http.PostForm("https://api.passport.pptv.com/accessToken", data)
	fmt.Println(data.Encode())


	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
