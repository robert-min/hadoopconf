package crawling

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func createTxt(doc string, path string) {
	var _, err = os.Stat(path)
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			return
		}
		fmt.Fprintf(file, doc)
		defer file.Close()
	}
}

func getHtml(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return string(body), nil

}

func Crawler(url string, path string) (string, error) {
	html, err := getHtml(url)
	if err != nil {
		return html, err
	}
	createTxt(html, path)
	// doc, _ := htmlquery.Parse(strings.NewReader(html))

	// property -> name, value, description

	return html, err
}
