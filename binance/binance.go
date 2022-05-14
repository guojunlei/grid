package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

var host string = "https://fapi.binance.com"

type exchange struct {
	ApiKey  string
	Secret  string
	TimeOut int
	Host    string
}

func (e exchange) get(url string) string {
	resp, err := http.Get(e.Host + url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}

func (e exchange) post(Url string, s string, t string) {
	client := &http.Client{}
	url1 := e.Host + Url
	parseurl, err := url.Parse(url1)
	if err != nil {
		log.Fatal(err)
	}
	// t := time.Now()
	// T := t.UnixMicro()
	// T1 := strconv.FormatInt(T, 10)
	params := "timestamp=" + t + "&" + "signature=" + s
	parseurl.RawQuery = params
	url2 := parseurl.String()
	fmt.Println(url2)
	reqest, err := http.NewRequest("GET", url2, nil)
	reqest.Header.Add("X-MBX-APIKEY", e.ApiKey)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(reqest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(b))

}

func Sha265(s, t string) string {
	T1 := "timestamp=" + t
	T2 := []byte(T1)
	S := []byte(s)

	m := hmac.New(sha256.New, S)
	m.Write(T2)
	signature := hex.EncodeToString(m.Sum(nil))

	return signature
}

func main() {
	binance := exchange{"v6JS4TzWBIr9nz4diXkBu8is3frsfnfT1qVU3UVEyjkwP2whzjxK2DAtzLeIzxDE", "7UC38Ky2fOzHjpFPonS8UgAc3WbMXYR7syp71SipiR5Hb1OZQTJQCI7mjYyaIJxu", 3000, "https://fapi.binance.com"}
	conent := binance.get("/fapi/v1/time")
	timeMap := make(map[string]int64)
	err := json.Unmarshal([]byte(conent), &timeMap)
	if err != nil {
		log.Fatal(err)
	}
	t := strconv.FormatInt(timeMap["serverTime"], 10)
	signature := Sha265(binance.Secret, t)

	binance.post("/fapi/v2/account", signature, t)

}
