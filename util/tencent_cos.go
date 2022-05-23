package util

import (
	"context"
	"douyin/conf"
	"github.com/google/uuid"
	"github.com/tencentyun/cos-go-sdk-v5"
	"mime"
	"mime/multipart"
	"net/http"
	"net/url"
	"path/filepath"
	"time"
)

var Ctx = context.Background()
var Client *cos.Client

//init TencentCOS
func init() {
	u, _ := url.Parse(config.CosRegion)
	b := &cos.BaseURL{BucketURL: u}
	Client = cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  config.SI,
			SecretKey: config.SK,
		},
	})
}

// GetPreSignedURL Get pre-signed with operation type
func GetPreSignedURL(name, httpMethod string) (*url.URL, error) {

	presignedURL, err := Client.Object.GetPresignedURL(Ctx, httpMethod, name, config.SI, config.SK, time.Hour, nil)
	if err != nil {
		return nil, err
	}
	return presignedURL, err
}

// ObjGetURL Get url by pre-signing
func ObjGetURL(filepath string) string {
	signedURL, err := GetPreSignedURL(filepath, http.MethodGet)
	if err != nil {
		return ""
	}
	return signedURL.String()
}

// ObjPost Upload file to bucket by pre-signing
func ObjPost(multfile *multipart.FileHeader) (string, string, error) {
	//OPEN FILE
	open, err := multfile.Open()
	if err != nil {
		return "", "", err
	}
	//extract file suffix
	ext := filepath.Ext(multfile.Filename)
	//To prevent file conflicts, use uuid
	uid := uuid.Must(uuid.NewRandom()).String()
	//Splicing bucket keys
	key_url := config.VideoSavePath + uid + ext
	//Get PreSigned
	putpresignedURL, err := Client.Object.GetPresignedURL(Ctx, http.MethodPut, key_url, config.SI, config.SK, time.Hour, nil)
	if err != nil {
		return "", "", err
	}
	//build request
	req, err := http.NewRequest(http.MethodPut, putpresignedURL.String(), open)
	if err != nil {
		return "", "", err
	}
	//Set Content-Type ， the file can be played directly
	req.Header.Set("Content-Type", mime.TypeByExtension(ext))
	_, err = http.DefaultClient.Do(req)
	if err != nil {
		return "", "", err
	}
	return key_url, config.ScreenshotSavePath + uid + config.ScreenshotFrame + config.PictureStyle, nil
}
