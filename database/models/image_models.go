package models

import (
	"errors"
	"image_upload/config"
	"image_upload/database/mysql"
	"time"
)

type ImageModel struct {
	ID uint32 `gorm:"column:id; primary_key"`
	FileName string `gorm:"column:filename"`
	FileMd5 string `gorm:"column:file_md5"`
	FileKey string `gorm:"column:file_key"`
	FileType string `gorm:"column:file_type"`
	FileSize int `gorm:"column:file_size"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (i *ImageModel)TableName()string{
	return "tmt_upload_image"
}

func (i *ImageModel)GetUrl()string{
	return config.ConfigInstance.ImageServerHost + "/" + i.FileKey
}

type ImageModelIn struct {
	FileName string
	FileMd5 string
	FileKey string
	FileType string
	FileSize int
}

func CreateImageModel(imgIn *ImageModelIn)(*ImageModel, error){
	tmp, _ := GetImageByMd5(imgIn.FileMd5)
	if tmp != nil{
		return tmp, nil
	}
	i := ImageModel{
		FileName: imgIn.FileName,
		FileMd5: imgIn.FileMd5,
		FileKey: imgIn.FileKey,
		FileType: imgIn.FileType,
		FileSize: imgIn.FileSize,
	}
	if err := mysql.DB.Create(&i).Error; err != nil{
		return nil, err
	}
	return &i, nil
}

func GetImageByMd5(m string)(*ImageModel, error){
	if m == ""{
		return nil, errors.New("md5 is empty")
	}
	i := ImageModel{}
	if err := mysql.DB.Where("file_md5 = ?", m).First(&i).Error; err != nil{
		return nil, err
	}
	return &i, nil
}



