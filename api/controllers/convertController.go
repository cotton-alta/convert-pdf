package controllers

import (
	"net/http"
	"io/ioutil"
	"log"
	"strings"
	"os"
	"github.com/russross/blackfriday/v2"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/labstack/echo"
)

func GetPdf() echo.HandlerFunc {
	return func(c echo.Context) error {
		url := c.FormValue("url") + ".md"

		top_html := `
			<html>
			<head>
					<meta charset="UTF-8">
					<style>
						img {
							width: 90%;
						}
						pre {
							background: #364549;
							color: #FFFFFF;
						}
					</style>
			</head>
			<body>
		`
		bottom_html := `
			</body>
			</html>
		`
	
		res, err := http.Get(url)
		checkErr(err, "get request failed")
		defer res.Body.Close()
		
		data, err := ioutil.ReadAll(res.Body)
		checkErr(err, "get read failed")
	
		html := blackfriday.Run(data)

		replace_html := strings.Replace(string(html), `<h2>slide: false</h2>`, ``, 1)
		replace_html = strings.Replace(replace_html, `<p>`, `<h1>`, 1)
		replace_html = strings.Replace(replace_html, `</p>`, `</h1>`, 1)
	
		modified_html := top_html + replace_html + bottom_html
	
		htmlConvert(modified_html)

		file, err := os.Open("test.pdf")
		checkErr(err, "file open failed")
		defer file.Close()

		fileInfo, err := file.Stat()
    checkErr(err, "fileInfo get failed")

    size := fileInfo.Size()
    buffer := make([]byte, size)

    file.Read(buffer)

		return c.Blob(http.StatusOK, "application/pdf", buffer)
	}
}

func htmlConvert(item string) {	
	
	pdf, err := wkhtmltopdf.NewPDFGenerator()
	checkErr(err, "generate failed")
	pdf.PageSize.Set("B5")
	pdf.AddPage(wkhtmltopdf.NewPageReader(strings.NewReader(item)))
	
	err = pdf.Create()
	checkErr(err, "create failed")
	
	err = pdf.WriteFile("./test.pdf")
	checkErr(err, "write failed")
}

func checkErr(err error, msg string) {
	if err != nil {
			log.Fatalln(msg, err)
	}
}