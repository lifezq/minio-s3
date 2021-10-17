package handler

import (
	"fmt"
	"github.com/minio/minio-go/v6"
	"io/ioutil"
	"log"
	"net/http"

	"minio-s3/internal/svc"
)

func UploadHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//var req types.UploadReq
		fmt.Println("1111111111111111111")
		if err := r.ParseForm(); err != nil {
			fmt.Errorf("parse form error:%s\n", err.Error())
			return
		}
		fmt.Printf("-----------------")
		if err := r.ParseMultipartForm(32 << 20); err != nil {
			fmt.Errorf("ParseMultipartForm form error:%s\n", err.Error())
			return
		}

		formvalue := r.FormValue("filename")
		fmt.Printf("r.Form len:%d  formvalue:%v\n", r.Form, formvalue)
		formvalue = r.PostFormValue("filename")
		fmt.Printf("------formvalue:%s\n", formvalue)
		file, fileHeader, err := r.FormFile("filename")
		fmt.Printf("formfile err:%s\n", err)
		fmt.Printf("------file:%v  fileHeader.len:%v\n", file, fileHeader.Filename)

		fmt.Printf("-----------------")
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Errorf("Parse body error:%s\n", err.Error())
			return
		}
		fmt.Printf("body len:%d\n", len(body))

		endpoint := "192.168.0.110:9000"
		accessKeyID := "miniouser"
		secretAccessKey := "miniopwd"
		useSSL := false

		// 初使化 minio client对象。
		minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
		if err != nil {
			log.Fatalln("发生错误：", err)
		}

		log.Printf("minioClient初使化成功 %#v\n", minioClient) // minioClient初使化成功

		// 创建一个叫mymusic的存储桶。
		bucketName := "bucketprogram"
		location := "us-east-1"
		err = minioClient.MakeBucket(bucketName, location)
		if err != nil {
			// 检查存储桶是否已经存在。
			exists, err := minioClient.BucketExists(bucketName)
			if err == nil && exists {
				log.Printf("We already own %s\n", bucketName)
			} else {
				log.Fatalln(err)
			}
		}
		log.Printf("Successfully created %s\n", bucketName)

		// 上传一个zip文件。
		objectName := "zip/" + fileHeader.Filename
		n, err := minioClient.PutObject(bucketName, objectName, file, fileHeader.Size, minio.PutObjectOptions{})
		if err != nil {
			log.Fatalln(err)
		}

		log.Printf("Successfully uploaded %s of size %d\n", objectName, n)

		//object, err := minioClient.GetObject(bucketName, objectName, minio.GetObjectOptions{})
		//if err != nil {
		//	fmt.Printf("get object error:%s\n", err.Error())
		//	return
		//}
		//
		//statObject, err := minioClient.StatObject(bucketName, objectName, minio.StatObjectOptions{})
		//if err != nil {
		//	fmt.Printf("state object error:%s\n", err.Error())
		//	return
		//}

		//if err := httpx.Parse(r, &req); err != nil {
		//	fmt.Printf("post err:%s\n", err.Error())
		//	httpx.Error(w, err)
		//	return
		//}
		//fmt.Println("222222222222222222222")
		//l := logic.NewUploadLogic(r.Context(), ctx)
		//resp, err := l.Upload(req)
		//if err != nil {
		//	httpx.Error(w, err)
		//} else {
		//	httpx.OkJson(w, resp)
		//}
	}
}
