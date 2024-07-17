package main

import (
	"UtmStealer/db"
	"UtmStealer/structures"
	"bytes"
	"encoding/xml"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func getHost() string {
	host, ok := os.LookupEnv("UTM_HOST")
	if !ok {
		host = "http://localhost:8080"
	}

	return host
}

func proxy(c *gin.Context) {

	remote, err := url.Parse(getHost())
	if err != nil {
		panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.Director = func(req *http.Request) {
		req.Header = c.Request.Header
		req.Host = remote.Host
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
		req.URL.Path = c.Param("proxyPath")
	}

	proxy.ModifyResponse = Grub
	proxy.ServeHTTP(c.Writer, c.Request)

}

func main() {

	log.Println("Going to steal from UTM:", getHost())

	db.ConnectDatabase()

	f, _ := os.OpenFile("request.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	routerProxy := gin.Default()

	routerProxy.Use(ginBodyLogMiddleware, Logger())

	routerProxy.Any("/*proxyPath", proxy)

	routerProxy.Run(":8085")

	// routerMarks := gin.Default()

	// // routerMarks.GET("/marks", Grub)

	// routerMarks.Run(":8086")
}

func Grub(resp *http.Response) (err error) {
	var Documents structures.TTN

	// if err := c.BindXML(&Documents); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// c.JSON(http.StatusOK, gin.H{"message": "ok"})

	b, err := io.ReadAll(resp.Body) //Read html
	if err != nil {
		println("err")
	}

	newbody := strings.ReplaceAll(string(b), "http://localhost:8080", "http://localhost:8085")

	xmlerr := xml.Unmarshal(b, &Documents)
	if xmlerr != nil {
		print(xmlerr)
	} else {
		for _, position := range Documents.Document.WayBillV4.Content.Position {

			for _, informF2Item := range position.InformF2 {

				for _, amcItem := range informF2Item.MarkInfo.Boxpos.Amclist.Amc {

					if errbd := db.DB.Create(&db.Table{
						Volume:  position.Product.AlcVolume,
						F2RegId: informF2Item.F2RegId,
						Amc:     amcItem,
					}).Error; errbd != nil {
						log.Println("Error writing amclist:", err)
						return
					}
				}
			}

		}
	}
	body := ioutil.NopCloser(bytes.NewReader([]byte(newbody)))
	resp.Body = body
	resp.ContentLength = int64(len(b))
	resp.Header.Set("Content-Length", strconv.Itoa(len(b)))

	// resp.Header.Set("Content-Length", strconv.Itoa(int(resp.ContentLength)))
	// err = resp.Body.Close()
	// if err != nil {
	// 	println("close")
	// }

	return nil
}

// func Grub(c *gin.Context) {
// 	var Documents structures.TTN

// 	if err := c.BindXML(&Documents); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	for _, position := range Documents.Document.WayBillV4.Content.Position {

// 		for _, informF2 := range position.InformF2 {

// 			for _, amclistItem := range informF2.MarkInfo.Boxpos.Amclist {
// 				err := db.DB.Create(&db.Table{
// 					Volume:  position.Product.AlcVolume,
// 					Amc:     amclistItem.Amc,
// 					F2RegId: informF2.F2RegId,
// 				}).Error

// 				if err != nil {
// 					log.Println("Error writing amclist:", err)
// 				}
// 			}
// 		}
// 	}
// 	// c.JSON(http.StatusOK, gin.H{"message": "ok"})
// }

// for _, position := range ttn.Document.WayBillV4.Content.Position {

//     for _, informF2 := range position.InformF2 {

//         for _, amclistItem := range informF2.MarkInfo.Boxpos.Amclist {
//             err := db.Create(&AmcData{
//                 AlcVolume: position.Product.AlcVolume,
//                 Amc:       amclistItem.Amc,
//                 F2RegId:   informF2.F2RegId,
//             }).Error

//             if err != nil {
//                 log.Println("Error writing amclist:", err)
//             }
//         }
//     }
// }

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		buf, _ := io.ReadAll(c.Request.Body)
		rdr1 := io.NopCloser(bytes.NewBuffer(buf))

		f, _ := os.OpenFile("request.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		f.WriteString("REQUEST BODY: " + (string(buf)) + "\n")

		c.Request.Body = rdr1
		c.Next()
	}
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func ginBodyLogMiddleware(c *gin.Context) {
	blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
	c.Writer = blw
	c.Next()
	f, _ := os.OpenFile("request.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	f.WriteString("RESPONSE BODY: " + blw.body.String() + "\n")
}
