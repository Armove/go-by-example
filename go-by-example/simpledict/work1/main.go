package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type DictRequestVolc struct{
	Language string `json:"language"` 
	Text string `json:"text"`
	//UserID string `json:"user_id"`
}

type DictRequest struct {
	TransType string `json:"trans_type"`
	Source    string `json:"source"`
	UserID    string `json:"user_id"`
}
type DictResponseVolc struct {
	Words []struct {
		Source int `json:"source"`
		Text string `json:"text"`
		PosList []struct {
			Type int `json:"type"`
			Phonetics []struct {
				Type int `json:"type"`
				Text string `json:"text"`
			} `json:"phonetics"`
			Explanations []struct {
				Text string `json:"text"`
				Examples []struct {
					Type int `json:"type"`
					Sentences []struct {
						Text string `json:"text"`
						TransText string `json:"trans_text"`
					} `json:"sentences"`
				} `json:"examples"`
				Synonyms []interface{} `json:"synonyms"`
			} `json:"explanations"`
			Relevancys []interface{} `json:"relevancys"`
		} `json:"pos_list"`
	} `json:"words"`
	Phrases []interface{} `json:"phrases"`
	BaseResp struct {
		StatusCode int `json:"status_code"`
		StatusMessage string `json:"status_message"`
	} `json:"base_resp"`
}


type DictResponse struct {
	Rc   int `json:"rc"`
	Wiki struct {
		KnownInLaguages int `json:"known_in_laguages"`
		Description     struct {
			Source string      `json:"source"`
			Target interface{} `json:"target"`
		} `json:"description"`
		ID   string `json:"id"`
		Item struct {
			Source string `json:"source"`
			Target string `json:"target"`
		} `json:"item"`
		ImageURL  string `json:"image_url"`
		IsSubject string `json:"is_subject"`
		Sitelink  string `json:"sitelink"`
	} `json:"wiki"`
	Dictionary struct {
		Prons struct {
			EnUs string `json:"en-us"`
			En   string `json:"en"`
		} `json:"prons"`
		Explanations []string      `json:"explanations"`
		Synonym      []string      `json:"synonym"`
		Antonym      []string      `json:"antonym"`
		WqxExample   [][]string    `json:"wqx_example"`
		Entry        string        `json:"entry"`
		Type         string        `json:"type"`
		Related      []interface{} `json:"related"`
		Source       string        `json:"source"`
	} `json:"dictionary"`
}

func querryVolc(word string) {
	client := &http.Client{}
	//var data = strings.NewReader(`{"text":"god","language":"en"}`)
	request := DictRequestVolc{Language: "en", Text: word}
	buf, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}
	var data = bytes.NewReader(buf)
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
	//防御性编程
	if resp.StatusCode != 200 {
		log.Fatal("bad StatusCode:", resp.StatusCode, "body", string(bodyText))
	}
	var dictResponseVolc DictResponseVolc
	err = json.Unmarshal(bodyText, &dictResponseVolc)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(word, "UK:", dictResponseVolc.Dictionary.Prons.En, "US:", dictResponseVolc.Dictionary.Prons.EnUs)
	for _, word := range dictResponseVolc.Words {
		for _,poslist := range word.PosList {
			for _,explanation := range poslist.Explanations {
				fmt.Println(explanation.Text)
			}
			
		}
	}

}


func query(word string) {
	client := &http.Client{}
	request := DictRequest{TransType: "en2zh", Source: word}
	buf, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}
	var data = bytes.NewReader(buf)
	req, err := http.NewRequest("POST", "https://api.interpreter.caiyunai.com/v1/dict", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("DNT", "1")
	req.Header.Set("os-version", "")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36")
	req.Header.Set("app-name", "xy")
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("device-id", "")
	req.Header.Set("os-type", "web")
	req.Header.Set("X-Authorization", "token:qgemv4jr1y38jyq6vhvi")
	req.Header.Set("Origin", "https://fanyi.caiyunapp.com")
	req.Header.Set("Sec-Fetch-Site", "cross-site")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", "https://fanyi.caiyunapp.com/")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Cookie", "_ym_uid=16456948721020430059; _ym_d=1645694872")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	//防御性编程
	if resp.StatusCode != 200 {
		log.Fatal("bad StatusCode:", resp.StatusCode, "body", string(bodyText))
	}
	
	var dictResponse DictResponse
	err = json.Unmarshal(bodyText, &dictResponse)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(word, "UK:", dictResponse.Dictionary.Prons.En, "US:", dictResponse.Dictionary.Prons.EnUs)
	for _, item := range dictResponse.Dictionary.Explanations {
		fmt.Println(item)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, `usage: simpleDict WORD
example: simpleDict hello
		`)
		os.Exit(1)
	}
	word := os.Args[1]
	querryVolc(word)
}
