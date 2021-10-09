package main

import (
	"crypto/md5"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func upload(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("TRIALS/upload.gtpl")
		t.Execute(w, token)

	} else {

		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}

		filename := handler.Filename
		filesize := handler.Size
		fmt.Println(filesize)
		data, _ := ioutil.ReadFile(filename)

		conn := AccessDB("dbUser" , "dbPass")

		bucket , err := gridfs.NewBucket(conn.Database("MAIN"))
		if err != nil{
			log.Fatal(err)
		}
		uploadStream, err := bucket.OpenUploadStream(
			filename,
		)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer func(uploadStream *gridfs.UploadStream) {
			err := uploadStream.Close()
			if err != nil {
				panic(err)
			}
		}(uploadStream)

		fileSize, err := uploadStream.Write(data)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Write file to DB was successful. File size: %d M\n", fileSize)


		defer file.Close()
		fmt.Println(handler.Filename)



		/*
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)*/
	}
}