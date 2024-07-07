package main

import (
	"WriteOff/db"
	"WriteOff/structures"
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	db.ConnectDatabase()

	router := gin.Default()

	router.Any("/", fufu)

	router.Run(":8080")

}

func fufu(c *gin.Context) {

	var documents structures.Documents
	today := time.Now()

	documents.Version = "1.0"
	documents.Xsi = "http://www.w3.org/2001/XMLSchema-instance"
	documents.Ns = "http://fsrar.ru/WEGAIS/WB_DOC_SINGLE_01"
	documents.Pref = "http://fsrar.ru/WEGAIS/ProductRef_v2"
	documents.Awr = "http://fsrar.ru/WEGAIS/ActWriteOff_v4"
	documents.Ce = "http://fsrar.ru/WEGAIS/CommonV3"

	documents.Owner.FSRAR_ID = "123456"
	loadIdentity()
	identityCounter2 = 0
	documents.Document.ActWriteOffv4.Identity = generateIdentity()

	documents.Document.ActWriteOffv4.Header.ActNumber = generateIdentity()
	documents.Document.ActWriteOffv4.Header.ActDate = today.Format("2006-01-02")
	documents.Document.ActWriteOffv4.Header.TypeWriteOff = "Реализация"
	documents.Document.ActWriteOffv4.Header.Note = "кря"

	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	var request struct {
		AmcValues []string `json:"marks"`
	}
	err = json.Unmarshal([]byte(bodyBytes), &request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	var records []db.Table
	result := db.DB.Where("amc IN (?)", request.AmcValues).Find(&records)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	for _, record := range records {
		position := structures.Position{
			Identity:       gen(),
			Writeoffvolume: &structures.Writeoffvolume{Volume: record.Volume},
			InformF1F2:     &structures.InformF1F2{InformF2: &structures.InformF2{F2RegId: record.F2RegId}},
			MarkCodeInfo:   &structures.MarkCodeInfo{Amc: record.Amc},
		}
		documents.Document.ActWriteOffv4.Content.Position = append(documents.Document.ActWriteOffv4.Content.Position, position)
	}

	xmlFile, err := os.Create("documents.xml")
	if err != nil {
		log.Fatal(err)
	}
	//defer delete()
	defer xmlFile.Close()

	encoder := xml.NewEncoder(xmlFile)
	err = encoder.Encode(documents)
	if err != nil {
		log.Fatal(err)
	}

	saveIdentity()
	send()

	fmt.Println("XML file created successfully!")

}

var identityCounter int
var identityCounter2 int
var identityMutex sync.Mutex

func generateIdentity() int {
	identityMutex.Lock()
	defer identityMutex.Unlock()

	identityCounter++
	return identityCounter
}

func gen() int {
	identityCounter2++
	return identityCounter2
}

func loadIdentity() {
	identityFile, err := os.ReadFile("identity.json")
	if err != nil {

		return
	}

	var counter struct {
		Identity int
	}
	err = json.Unmarshal(identityFile, &counter)
	if err != nil {

		return
	}

	identityCounter = counter.Identity
}

func saveIdentity() {
	identityFile, err := os.Create("identity.json")
	if err != nil {

		return
	}
	defer identityFile.Close()

	data, err := json.Marshal(struct {
		Identity int
	}{
		Identity: identityCounter,
	})
	if err != nil {

		return
	}

	_, err = identityFile.Write(data)
	if err != nil {

		return
	}
}

// func send() {

// 	xmlFile, err := os.Open("documents.xml")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer xmlFile.Close()

// 	url := "http://localhost:8085/xml"
// 	request, err := http.NewRequest("POST", url, xmlFile)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// request.Header.Set("documents", "application/xml")

// 	client := &http.Client{}
// 	resp, err := client.Do(request)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer resp.Body.Close()

// }

// func delete() {

// 	err := os.Remove("documents.xml")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// }

// package main

// import (
//     "bytes"
//     "fmt"
//     "io"
//     "mime/multipart"
//     "net/http"
//     "net/http/httptest"
//     "net/http/httputil"
//     "os"
//     "strings"
// )

// func main() {

//     var client *http.Client
//     var remoteURL string
//     {
//         //setup a mocked http client.
//         ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//             b, err := httputil.DumpRequest(r, true)
//             if err != nil {
//                 panic(err)
//             }
//             fmt.Printf("%s", b)
//         }))
//         defer ts.Close()
//         client = ts.Client()
//         remoteURL = ts.URL
//     }

//     //prepare the reader instances to encode
//     values := map[string]io.Reader{
//         "file":  mustOpen("main.go"), // lets assume its this file
//         "other": strings.NewReader("hello world!"),
//     }
//     err := Upload(client, remoteURL, values)
//     if err != nil {
//         panic(err)
//     }
// }

// func Upload(client *http.Client, url string, values map[string]io.Reader) (err error) {
//     // Prepare a form that you will submit to that URL.
//     var b bytes.Buffer
//     w := multipart.NewWriter(&b)
//     for key, r := range values {
//         var fw io.Writer
//         if x, ok := r.(io.Closer); ok {
//             defer x.Close()
//         }
//         // Add an image file
//         if x, ok := r.(*os.File); ok {
//             if fw, err = w.CreateFormFile(key, x.Name()); err != nil {
//                 return
//             }
//         } else {
//             // Add other fields
//             if fw, err = w.CreateFormField(key); err != nil {
//                 return
//             }
//         }
//         if _, err = io.Copy(fw, r); err != nil {
//             return err
//         }

//     }
//     // Don't forget to close the multipart writer.
//     // If you don't close it, your request will be missing the terminating boundary.

//     w.Close()

//     // Now that you have a form, you can submit it to your handler.
//     req, err := http.NewRequest("POST", url, &b)
//     if err != nil {
//         return
//     }
//     // Don't forget to set the content type, this will contain the boundary.
//     req.Header.Set("Content-Type", w.FormDataContentType())

//     // Submit the request
//     res, err := client.Do(req)
//     if err != nil {
//         return
//     }

//     // Check the response
//     if res.StatusCode != http.StatusOK {
//         err = fmt.Errorf("bad status: %s", res.Status)
//     }
//     return
// }

// func mustOpen(f string) *os.File {
//     r, err := os.Open(f)
//     if err != nil {
//         panic(err)
//     }
//     return r
// }

func send() {
	fileDir, _ := os.Getwd()
	fileName := "documents.xml"
	filePath := path.Join(fileDir, fileName)

	file, _ := os.Open(filePath)
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", filepath.Base(file.Name()))
	io.Copy(part, file)
	writer.Close()

	r, _ := http.NewRequest("POST", "http://localhost:8085/xml", body)
	r.Header.Add("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	client.Do(r)
}
