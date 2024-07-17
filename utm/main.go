package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<ns:Documents xmlns:ce="http://fsrar.ru/WEGAIS/CommonV3" xmlns:ns="http://fsrar.ru/WEGAIS/WB_DOC_SINGLE_01" xmlns:rst="http://fsrar.ru/WEGAIS/ReplyRestBCode">
  <ns:Owner>
    <ns:FSRAR_ID>3463047</ns:FSRAR_ID>
</ns:Owner>
  <ns:Document>
    <ns:ReplyRestBCode>
      <rst:RestsDate>2024-07-17T18:21:47.2147</rst:RestsDate>
      <rst:Inform2RegId>FB-000006988452001</rst:Inform2RegId>
<rst:MarkInfo>
<ce:amc>188404477177461222001JUMLQHV766R5CM3M457AL2TYPMQGXJKDIKQ7HFJDFFUCG6UESJULPMH6WP6HZAHQ355VW5QKTJAJRTKPWCYY2G34FILWPVZDJEKSS7OQXQPCNHWAQOWCAOBOF4HFDAD51</ce:amc>
<ce:amc>188404477177461222001JUMLQHV766R5CM3M457AL2TYPMQGXJKDIKQ7HFJDFFUCG6UESJULPMH6WP6HZAHQ355VW5QKTJAJRTKPWCYY2G34FILWPVZDJEKSS7OQXQPCNHWAQOWCAOBOF4HFDAD52</ce:amc>
</rst:MarkInfo>
    </ns:ReplyRestBCode>
  </ns:Document>
</ns:Documents>`)
		c.Request.Response.Header.Add("Content-Type", "text/xml;charset=UTF-8")
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	/* example curl for /admin with basicauth header
	   Zm9vOmJhcg== is base64("foo:bar")

		curl -X POST \
	  	http://localhost:8080/admin \
	  	-H 'authorization: Basic Zm9vOmJhcg==' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'
	*/
	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
