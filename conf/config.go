package config

import (
	"github.com/joho/godotenv"
	"os"
)

var SI string
var SK string
var CosRegion string
var PictureStyle string
var VideoSavePath string
var ScreenshotSavePath string
var ScreenshotFrame string

//Initialize COS Settings - COS Region - COS Bucket - Workflow Format - Video Storage Path - Screenshot Storage Path
func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	SI = os.Getenv("SI")
	SK = os.Getenv("SK")
	CosRegion = os.Getenv("COS_REGION")
	PictureStyle = os.Getenv("PictureStyle")
	VideoSavePath = os.Getenv("VideoSavePath")
	ScreenshotSavePath = os.Getenv("ScreenshotSavePath")
	ScreenshotFrame = "_" + os.Getenv("ScreenshotFrame")
}
