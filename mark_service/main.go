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

	"github.com/gin-gonic/gin"
)

func proxy(c *gin.Context) {
	remote, err := url.Parse("http://localhost:8080")
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

	db.ConnectDatabase()

	routerProxy := gin.Default()

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
	body := ioutil.NopCloser(bytes.NewReader(b))
	resp.Body = body
	// resp.ContentLength = int64(len(b))
	// resp.Header.Set("Content-Length", strconv.Itoa(len(b)))

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
