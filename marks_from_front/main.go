package main

import (
	"bytes"
	"fmt"
	"front/structures"
	"io"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// {
//     "reason" : "Реализация"
//     "marks": [
//         "marka1",
//         "marka2",
//         "marka3"
//     ]
// }

func main() {

	f, _ := os.OpenFile("logs.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := gin.Default()

	router.Use(cors.Default())

	router.Use(Logger())

	router.Any("/marki", fufu)

	router.Run(":8081")
}

func fufu(c *gin.Context) {

	var request structures.Request

	if err := c.BindJSON(&request); err != nil {
		fmt.Println("ploho")
		return
	}

	for _, mark := range request.Marks {
		fmt.Println(mark)
	}

	// go Logger()

}

// func Logger() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		buf, _ := io.ReadAll(c.Request.Body)
// 		rdr1 := io.NopCloser(bytes.NewBuffer(buf))

// 		f, _ := os.OpenFile("logs.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
// 		f.WriteString("\n" + "REQUEST BODY: " + (string(buf)) + "\n")

// 		c.Request.Body = rdr1
// 		c.Next()
// 	}
// }

// func sendEmptyRequest() {
// 	resp, err := http.Get("http://localhost:8082/")
// 	if err != nil {
// 		fmt.Println("Error sending empty request:", err)
// 		return
// 	}

// 	defer resp.Body.Close()
// }

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {

		buf, _ := io.ReadAll(c.Request.Body)
		rdr1 := io.NopCloser(bytes.NewBuffer(buf))

		f, _ := os.OpenFile("request.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		f.WriteString("REQUEST BODY: " + string(buf) + "\n")
		f.Close()

		req, err := http.NewRequest("POST", "http://localhost:8082/", rdr1)
		if err != nil {
			c.AbortWithError(500, err)
			return
		}

		// req.Header.Set("documents", c.Request.Header.Get("application/xml"))
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			c.AbortWithError(500, err)
			return
		}

		defer resp.Body.Close()
		c.Next()
	}
}

// func Logger() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		buf, _ := io.ReadAll(c.Request.Body)
// 		rdr1 := io.NopCloser(bytes.NewBuffer(buf))

// 		f, _ := os.OpenFile("request.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
// 		f.WriteString("REQUEST BODY: " + string(buf) + "\n")
// 		f.Close()

// 		req, err := http.NewRequest("POST", "http://localhost:8082/", rdr1)
// 		if err != nil {
// 			c.AbortWithError(500, err)
// 			return
// 		}

// 		req.Header.Set("Content-Type", c.Request.Header.Get("Content-Type"))

// 		client := &http.Client{}
// 		resp, err := client.Do(req)
// 		if err != nil {
// 			c.AbortWithError(500, err)
// 			return
// 		}
// 		defer resp.Body.Close()

// 		c.Next()
// 	}
// }

// func zpr(c *gin.Context) {
// 	url := "http://localhost:8087"

// 	resp, err := http.Get(url)
// 	if err != nil {
// 		fmt.Println("Ошибка при отправке запроса", err)
// 		return
// 	}

// 	defer resp.Body.Close()

// 	fmt.Println("Статус ответа", resp.Status)
// }

// func zzpr(c *gin.Context) {

// 	url := "http://localhost:8087"

// 	resp, err := http.Get(url)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	defer resp.Body.Close()

// 	c.JSON(http.StatusOK, gin.H{"status": resp.Status})

// 	// if err != nil {
// 	// 	fmt.Println("Ошибка при запуске сервера:", err)
// 	// 	return
// 	// }
// }

// func handleRequest(c *gin.Context) {
//

//
// 	for _, mark := range request.Marks {
// 	  if len(mark) < 5 {
// 		fmt.Println("Марка", mark, "слишком короткая")
// 		continue
// 	  }

// 	  err := saveMarkToDB(mark)
// 	  if err != nil {
// 		fmt.Println("Ошибка при сохранении марки", mark, ":", err)
// 		continue
// 	  }

// 	  fmt.Println("Марка", mark, "успешно сохранена")
// 	}

// 	// ...
//   }

//   func saveMarkToDB(mark string) error {
//   }

// for _, mark := range request.Marks {
//     newMark := Mark{Value: mark}
//     err := db.Create(&newMark)
//     if err != nil {
//       fmt.Println("Ошибка при сохранении марки", mark, ":", err)
//       continue
//     }

//     fmt.Println("Марка", mark, "успешно сохранена в базу данных")
//   }
