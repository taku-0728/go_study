package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
)

func main4() {
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	writer.WriteField("name", "Michael Jackson")

	/*
	   fileWriter, err := writer.CreateFormFile("thumbnail", "photo.jpg")
	   if err != nil {
	     panic(err)
	   }
	   readFile, err := os.Open("photo.jpg")
	   if err != nil {
	     // ファイル読み込み失敗
	     panic(err)
	   }
	   defer readFile.Close()
	*/

	part := make(textproto.MIMEHeader)
	part.Set("Content-Type", "image/jpeg")
	part.Set("Content-Disposition", `form-data; name="thumbnail"; filename="photo.jpg"`)

	fileWriter, err := writer.CreatePart(part)

	if err != nil {
		panic(err)
	}

	readFile, err := os.Open("photo.jpg")

	if err != nil {
		panic(err)
	}

	io.Copy(fileWriter, readFile)
	writer.Close()

	resp, err := http.Post("http://localhost:18888", writer.FormDataContentType(), &buffer)

	if err != nil {
		// 送信失敗
		panic(err)
	}
	log.Println("Status:", resp.Status)
}
