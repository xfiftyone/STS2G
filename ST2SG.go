package main

import (
	"ST2G/cvemod/s001"
	"ST2G/cvemod/s005"
	"ST2G/cvemod/s007"
	"ST2G/cvemod/s008"
	"ST2G/cvemod/s009"
	"ST2G/cvemod/s012"
	"ST2G/cvemod/s013"
	"ST2G/cvemod/s015"
	"ST2G/cvemod/s016"
	"ST2G/cvemod/s045"
	"ST2G/cvemod/s046"
	"ST2G/cvemod/s048"
	"ST2G/cvemod/s053"
	"ST2G/cvemod/s057"
	"ST2G/cvemod/utils"
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	var mode string
	var url string
	var cmd string
	var vn int
	var data string
	app := &cli.App{
		Name:"ST2SG",
		Usage:"Struts2 Scanner Written in Golang",
		UsageText:"" +
			"ST2SG --url http://xxxxxxx.com/test.action --vn 15 --mode scan ",
		Authors: []*cli.Author{
			&cli.Author{
				Name:  "x51",
				Email: "x51enter@gmail.com",
			},
		},
		Flags: []cli.Flag {
			&cli.StringFlag{
				Name:        "mode",
				Usage:       "Specify work mode",
				Destination: &mode,
			},
			&cli.IntFlag{
					//指定漏洞编号
				Name:        "vn",
				Usage:       "Vulnerability number",
				Value:		  000,
				Destination: &vn,
			},
			&cli.StringFlag{
				Name:        "url",
				Usage:       "Set target url",
				Destination: &url,
			},
			&cli.StringFlag{
				Name:        "cmd",
				Usage:       "Exec command(Only works on mode exec.)",
				Destination: &cmd,
			},
			&cli.StringFlag{
				Name:        "data",
				Usage:       "Specific vulnerability packets",
				Destination: &data,
			},

		},
		
		Action: func(c *cli.Context) error {
			 if mode == "scan" {
				 if url == ""{
					log.Fatalln("url未指定")
				 }
				 switch vn {
						 case 1:
							 if  data !=""{
								 s001.Check(url,data)
							 }else {
								 fmt.Println("s001须指定POST数据包内容，并用<fuckit>标记出测试点，如: --data \"user=a&pass=fuckit\"")
							 }
						 case 5:
							s005.Check(url)
						 case 7:
							 if  data !=""{
								 s007.Check(url,data)
							 }else {
								 fmt.Println("s007需指定POST数据包内容，并用<fuckit>标记出测试点，如: --data \"user=a&pass=fuckit\"")
							 }
				 		 case 8:
					 		s008.Check(url)
				 		 case 9:
							 if  data !=""{
								 s009.Check(url,data)
							 }else {
								 fmt.Println("s009需指定要测试的GET参数，如: --data=\"name\"")
							 }
						 case 12:
							 if  data !=""{
								 s012.Check(url,data)
							 }else {
								 fmt.Println("s012需手动指定POST数据包内容，并用<fuckit>标记出测试点，如: --data \"user=a&pass=fuckit\"")
							 }
						 case 13:
							s013.Check(url)
						 case 15:
							 s015.Check(url)
						 case 16:
							 s016.Check(url)
						 case 45:
							s045.Check(url)
						 case 46:
							 s046.Check(url)
						 case 48:
							 if  data !=""{
								 s048.Check(url,data)
							 }else {
								 fmt.Println("s048需手动指定POST数据包内容，并用<fuckit>标记出测试点，如: --data \"user=a&pass=fuckit\"")
							 }
						 case 53:
							 if  data !=""{
								 s053.Check(url,data)
							 }else {
								 fmt.Println("s053需手动指定POST数据包内容，并用<fuckit>标记出测试点，如: --data \"user=a&pass=fuckit\"")
							 }
						 case 57:
							s057.Check(url)
						 case 000:
							 fmt.Println("未指定漏洞编号,默认全检测")
							 s001.Check(url,data)
							 s005.Check(url)
							 s007.Check(url,data)
							 s008.Check(url)
							 s009.Check(url,data)
							 s012.Check(url,data)
							 s013.Check(url)
							 s015.Check(url)
							 s016.Check(url)
							 s045.Check(url)
							 s046.Check(url)
							 s048.Check(url,data)
							 s053.Check(url,data)
							 s057.Check(url)
						 default:
							fmt.Println("漏洞编号设置错误，目前支持检测：")
							for _,vnn := range utils.Vnlist{
								fmt.Println(vnn)
							}
				 }
			} else if mode=="exec" && cmd != ""{
				 if url == ""{
					 log.Fatalln("url未指定")
				 }
				 switch vn {
				 case 1:
					 if  data !=""{
						 s001.ExecCommand(url,cmd,data)
					 }else {
						 fmt.Println("s001需手动指定POST数据包内容，并用fuckit标记出测试点，如: --data=\"user=a&pass=fuckit\"")
					 }
				 case 5:
				 	 s005.ExecCommand(url,cmd)
				 case 7:
					 if  data !=""{
						 s007.ExecCommand(url,cmd,data)
					 }else {
						 fmt.Println("s007需手动指定POST数据包内容，并用fuckit标记出测试点，如: --data=\"user=a&pass=fuckit\"")
					 }
				 case 8:
					 s008.ExecCommand(url,cmd)
				 case 9:
					 s009.ExecCommand(url,cmd,data)
				 case 12:
					 if  data!=""{
						 s012.ExecCommand(url,cmd,data)
					 }else {
						 fmt.Println("s012需手动指定POST数据包内容，并用<fuckit>标记出测试点，如: --data=\"user=a&pass=fuckit\"")
					 }
				 case 13:
					 s013.ExecCommand(url,cmd)
				 case 15:
					 s015.ExecCommand(url,cmd)
				 case 16:
					 s016.ExecCommand(url,cmd)
				 case 45:
					 s045.ExecCommand(url,cmd)
				 case 46:
					 s046.ExecCommand(url,cmd)
				 case 48:
					 if  data!=""{
						 s048.ExecCommand(url,cmd,data)
					 }else {
						 fmt.Println("s048需手动指定POST数据包内容，并用<fuckit>标记出测试点，如: --data=\"user=a&pass=fuckit\"")
					 }
				 case 53:
					 if  data !=""{
						 s053.ExecCommand(url,cmd,data)
					 }else {
						 fmt.Println("s053需手动指定POST数据包内容，并用<fuckit>标记出测试点，如: --data=\"user=a&pass=fuckit\"")
					 }
				 case 57:
				 	s057.ExecCommand(url,cmd)
				 case 000:
					 log.Fatalf("命令执行模式必须指定漏洞编号")
				 default:
					 log.Fatalf("命令执行模式必须指定漏洞编号")
				 }
			}else {
				fmt.Println("参数错误")
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}