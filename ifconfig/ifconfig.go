package ifconfig

import (
	"net/http"
	"io/ioutil"
)

const Endpoint = "https://ifconfig.co"

func GetPublicIP() string {
	response, err := http.Get(Endpoint)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	return string(body)
}