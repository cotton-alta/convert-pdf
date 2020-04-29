package controllers

import (
	"fmt"
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
		// url := c.FormValue("url")

		url := "https://qiita.com/cotton_alta_/items/6e2ca3d429c20a05d814.md"

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
	
		modified_html := top_html + string(html) + bottom_html
	
		fmt.Printf("%v", modified_html)
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