package controllers

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
	"strings"
	"os"
	"github.com/russross/blackfriday"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/labstack/echo"
)

func GetPdf() echo.HandlerFunc {
	return func(c echo.Context) error {
		// url := c.FormValue("url")

		url := "https://qiita.com/cotton_alta_/items/6e2ca3d429c20a05d814.md"

		top_html := `
			<html>
			<head>
					<meta charset="UTF-8">
					<style>
						p {
							display: inline-block;
							font-size: 25px;
							line-height: 40px;
						}
						img {
							width: 90%;
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
	
		html := blackfriday.MarkdownBasic(data)
	
		modified_html := top_html + string(html) + bottom_html
	
		htmlConvert(modified_html)
	
		file, err := os.Open("test.pdf")
		checkErr(err, "file open failed")
		defer file.Close()

		fileInfo, err := file.Stat()
    checkErr(err, "fileInfo get failed")

    size := fileInfo.Size()
    buffer := make([]byte, size)

    file.Read(buffer)

    fmt.Printf("%T", buffer)

		return c.Blob(http.StatusOK, "application/pdf", buffer)
	}
}

func htmlConvert(item string) {	
	
	pdf, err := wkhtmltopdf.NewPDFGenerator()
	checkErr(err, "generate failed")
	
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