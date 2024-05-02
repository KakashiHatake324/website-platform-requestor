package parser

import (
	"context"
	"errors"
	"strings"
	"website-platform-requestor/client"
	"website-platform-requestor/platforms"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func ParseUrlToPlatform(web string) (string, error) {
	proxy := ""
	req_client := client.CreateRequestClient(proxy)
	ua := client.SelectUserAgent()
	Request := client.DoRequest{
		Client:         req_client,
		CTX:            context.TODO(),
		AcceptedStatus: []int{200},
		Req:            map[string]string{"Method": "GET", "URL": web, "Data": "nil"},
		Headers: map[string][]string{
			"Accept":             {"*/*"},
			"User-Agent":         {ua[0]}, // useragent
			"Origin":             {web},
			"Sec-Ch-Ua-Mobile":   {"?0"},
			"Sec-Ch-Ua-Platform": {ua[1]}, // useragent's platform
			"Sec-Fetch-Site":     {"same-site"},
			"Sec-Fetch-Mode":     {"cors"},
			"Sec-Fetch-Dest":     {"empty"},
			"Accept-Language":    {"en-US,en;q=0.9"},
			"Accept-Encoding":    {"gzip, deflate, br"},
		},
	}

	response := Request.MakeRequest()
	if response.Error != nil {
		return "", response.Error
	}

	for platform, parse := range platforms.PlatformsMap {
		c := cases.Title(language.Und)
		if strings.Contains(strings.ToLower(response.ResponseBody), parse) {
			return c.String(platform), nil
		}
	}

	return "", errors.New("the web address' platform has not been added to the parser yet")
}
