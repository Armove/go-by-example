

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	client := &http.Client{}
	var data = strings.NewReader(`{"text":"god","language":"en"}`)
	req, err := http.NewRequest("POST", "https://translate.volcengine.com/web/dict/match/v1/?msToken=&X-Bogus=DFSzKIVLQDGo/qZGSW81nTIbA69M&_signature=_02B4Z6wo00001buGQggAAIDAm2AFzDwon9W7l0aAAAyXCXZBWoQ6t9MPA9Jpqp77R0TGyuDsJwm6OZcjP829a9vcyusjRlV2aNzhVcffnRkf1ff1I6XSEtq0QgX11Zw02kbaN3FI7DkqnXa0f3", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:100.0) Gecko/20100101 Firefox/100.0")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2")
	//req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Origin", "https://translate.volcengine.com")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Referer", "https://translate.volcengine.com/translate?category=&home_language=zh&source_language=detect&target_language=zh&text=god")
	req.Header.Set("Cookie", "x-jupiter-uuid=16520193473474480; i18next=translate; ttcid=1668facb92474386aa57efed4bab926721; tt_scid=g7wTwbxU-Wiqkmc1rpXYKjiiaGhhouujBwcLnehQ45V24kPZv36TSK8MDKtuB3G99328")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("TE", "trailers")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
}

