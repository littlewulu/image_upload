package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"image_upload/config"
	"image_upload/database/models"
	"image_upload/utils"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

type ErrorResponse struct {
	Msg string `json:"message"`
	Code int `json:"code"`
}

// /api/tomato/image/upload_image
// 上传图片
func UploadImageFileFunc(c *gin.Context){
	// check pw
	pw, _:= c.GetPostForm("password")
	if pw == "" || pw != config.ConfigInstance.OperationPassword{
		c.JSON(400, ErrorResponse{Msg:"params error", Code:10})
		return
	}
	img, err := c.FormFile("img")
	if err != nil{
		c.JSON(400, ErrorResponse{Msg:"params error", Code:11})
		return
	}
	// 大小限制
	if img.Size > config.ConfigInstance.ImageSizeLimit{
		c.JSON(400, ErrorResponse{Msg:"too large", Code:15})
		return
	}

	// 读取文件
	f, err := img.Open()
	if err != nil{
		c.JSON(400, ErrorResponse{Msg:"open file fail", Code:12})
		return
	}
	defer f.Close()
	data, err := ioutil.ReadAll(f)
	if err != nil{
		c.JSON(400, ErrorResponse{Msg:"read file fail", Code:13})
		return
	}

	// 计算md5
	md5Str := utils.CalMd5(data)
	// 是否已存在
	tmpImgModel, _ := models.GetImageByMd5(md5Str)
	if tmpImgModel != nil{
		c.JSON(200, map[string]interface{}{
			"url": tmpImgModel.GetUrl(),
		})
		return
	}

	// 保存文件
	fileType, isValid := getFileType(data)
	if ! isValid{
		c.JSON(400, ErrorResponse{Msg:"error file type", Code:14})
		return
	}
	fileKey := generateFileKey(fileType)
	err = saveImageFile(data, fileKey)
	if err != nil{
		c.JSON(400, ErrorResponse{Msg:err.Error(), Code:16})
		return
	}
	// 保存到数据库
	imgModel, err := models.CreateImageModel(&models.ImageModelIn{
		FileSize: len(data),
		FileType:fileType,
		FileName: img.Filename,
		FileKey:fileKey,
		FileMd5: md5Str,
	})
	if err != nil{
		c.JSON(400, ErrorResponse{Msg:err.Error(), Code:17})
		return
	}
	c.JSON(200, map[string]interface{}{
		"url": imgModel.GetUrl(),
	})
	return
}

// 生成file key
func generateFileKey(fileType string)string{
	// 按年月拆分文件夹
	ym := time.Now().Format("0601")
	// 日期
	d := time.Now().Format("02")
	// 生成随机字符串
	r := utils.RanString(30)

	return fmt.Sprintf("%s/%s_%s.%s", ym, d, r, fileType)
}

// 保存文件
func saveImageFile(data []byte, fileKey string)error{
	filePath := config.ConfigInstance.ImageSourceDirectory + "/" + fileKey
	// 创建文件夹
	fileDir := filepath.Dir(filePath)
	err := os.MkdirAll(fileDir, 0755)
	if err != nil{
		return err
	}
	f, err := os.Create(filePath)
	if err != nil{
		return err
	}
	f.Write(data)
	f.Close()
	return nil
}

// get file type
func getFileType(data []byte)(string, bool){
	t, err := utils.DetectContentType(data)
	if err != nil{
		return "", false
	}
	return utils.ValidateImageType(t)
}

