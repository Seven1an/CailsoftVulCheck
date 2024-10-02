package main

import (
	"flag"
	"fmt"
	"github.com/Seven1an/cailsoftVulCheck/vulList"
)

func main() {
	info := `
<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
____ ____ _ _    ____ ____ ____ ___ _  _ _  _ _    ____ _  _ ____ ____ _  _
|    |__| | |    [__  |  | |___  |  |  | |  | |    |    |__| |___ |    |_/
|___ |  | | |___ ___] |__| |     |   \/  |__| |___ |___ |  | |___ |___ | \_
							By:Seven1an    v0.1
>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
`
	fmt.Println(info)
	fmt.Println("Vulnerability List(As of 17:31:51 on October 2, 2024):")
	fmt.Println("\t 1. AuthToken interface arbitrary account login")
	fmt.Println("\t 2. DownloadBuilder arbitrary file read")
	fmt.Println("\t 3. GetCssFile arbitrary file read")
	fmt.Println("\t 4. GetExcellTemperature SQL injection")
	fmt.Println("\t 5. GetImportDetailJson SQL injection")
	fmt.Println("\t 6. GetJSFile arbitrary file read")
	fmt.Println("\t 7. ReadTxtLog arbitrary log read")
	fmt.Println("\t 8. EHRSubmitUploadify arbitrary file upload")
	fmt.Println("\t 9. SysSubmitUploadify arbitrary file upload")
	fmt.Println()
	target := flag.String("u", "", "URL")
	vul := flag.String("vul", "any", "Vulnerability")
	flag.Parse()

	if *target != "" {
		var err error
		switch *vul {
		case "1":
			err = vulList.AuthTokenCheck(*target)
		case "2":
			err = vulList.DownloadBuilderCheck(*target)
		case "3":
			err = vulList.GetCssFileCheck(*target)
		case "4":
			err = vulList.GetExecllTemperatureCheck(*target)
		case "5":
			err = vulList.GetImportDetailJsonCheck(*target)
		case "6":
			err = vulList.GetJSFileCheck(*target)
		case "7":
			err = vulList.ReadTxtLogCheck(*target)
		case "8":
			err = vulList.EHRSubmitUploadifyCheck(*target)
		case "9":
			err = vulList.SysSubmitUploadifyCheck(*target)
		case "any":
			err = vulList.AuthTokenCheck(*target)
			if err != nil {
				fmt.Println("Error in AuthTokenCheck:", err)
			}

			err = vulList.DownloadBuilderCheck(*target)
			if err != nil {
				fmt.Println("Error in DownloadBuilderCheck:", err)
			}

			err = vulList.GetCssFileCheck(*target)
			if err != nil {
				fmt.Println("Error in GetCssFileCheck:", err)
			}

			err = vulList.GetExecllTemperatureCheck(*target)
			if err != nil {
				fmt.Println("Error in GetExecllTemperatureCheck:", err)
			}

			err = vulList.GetImportDetailJsonCheck(*target)
			if err != nil {
				fmt.Println("Error in GetImportDetailJsonCheck:", err)
			}

			err = vulList.GetJSFileCheck(*target)
			if err != nil {
				fmt.Println("Error in GetJSFileCheck:", err)
			}

			err = vulList.ReadTxtLogCheck(*target)
			if err != nil {
				fmt.Println("Error in ReadTxtLogCheck:", err)
			}

			err = vulList.EHRSubmitUploadifyCheck(*target)
			if err != nil {
				fmt.Println("Error in EHRSubmitUploadifyCheck:", err)
			}

			err = vulList.SysSubmitUploadifyCheck(*target)
			if err != nil {
				fmt.Println("Error in SysSubmitUploadifyCheck:", err)
			}
		default:
			fmt.Println("What did you input??")
		}
		if err != nil {
			fmt.Println("An error occurred during the vulnerability check:", err)
		} else {
			fmt.Println("END.")
			fmt.Println()
		}
	} else {
		fmt.Println("Usage: \n\tCailsoftVulCheck.exe -u http://example.com/ -vul [1-9](default: any)\n\t!!!!Don't forget it!!! [/]")
		fmt.Println("\tCailsoftVulCheck.exe -u http://example.com/ OK")
		fmt.Println("\tCailsoftVulCheck.exe -u http://example.com  NO")
	}
}
