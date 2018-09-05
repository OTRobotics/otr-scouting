package main

import (
	"cloud.google.com/go/storage"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
	"google.golang.org/appengine/file"
	"google.golang.org/appengine/urlfetch"
	"html/template"
	"io/ioutil"
)
func UploadFile(fileName string, data []byte, c *gin.Context) {
	ctx := appengine.NewContext(c.Request)
	bucket, err := file.DefaultBucketName(ctx)

	if err != nil {
		bucket = "staging.otr-scouting.appspot.com"
	}
	client, err := storage.NewClient(ctx)
	if err != nil {
		fmt.Printf("failed to create client: %v", err)
		return
	}
	defer client.Close()

	wc := client.Bucket(bucket).Object(fileName).NewWriter(ctx)

	wc.ContentType = "text/plain"
	wc.Write(data)
	wc.Close()
}

func dumpStats(obj *storage.ObjectAttrs) {
	fmt.Printf( "(filename: /%v/%v, ", obj.Bucket, obj.Name)
	fmt.Printf( "ContentType: %q, ", obj.ContentType)
	fmt.Printf( "ACL: %#v, ", obj.ACL)
	fmt.Printf("Owner: %v, ", obj.Owner)
	fmt.Printf( "ContentEncoding: %q, ", obj.ContentEncoding)
	fmt.Printf("Size: %v, ", obj.Size)
	fmt.Printf( "MD5: %q, ", obj.MD5)
	fmt.Printf( "CRC32C: %q, ", obj.CRC32C)
	fmt.Printf("Metadata: %#v, ", obj.Metadata)
	fmt.Printf("MediaLink: %q, ", obj.MediaLink)
	fmt.Printf( "StorageClass: %q, ", obj.StorageClass)
	if !obj.Deleted.IsZero() {
		fmt.Printf( "Deleted: %v, ", obj.Deleted)
	}
	fmt.Printf( "Updated: %v)\n", obj.Updated)
}

func GetPageTemplate(page string, c *gin.Context) *template.Template {
	ctx := appengine.NewContext(c.Request)
	url := "https://storage.googleapis.com/staging.otr-scouting.appspot.com/web/"+page
	client := urlfetch.Client(ctx)
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	// do this now so it won't be forgotten
	defer resp.Body.Close()
	// reads html as a slice of bytes
	html, err := ioutil.ReadAll(resp.Body)
	// show the HTML code as a string %s
	fmt.Printf("%s\n", html)


	tmpl, err := template.New(page).Parse(string(html))
	return tmpl
}