package service

import (
	"fmt"
	"github.com/JiangInk/market_monitor/config"
	"github.com/JiangInk/market_monitor/extend/utils"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

type UploadService struct {}

// GetImgPath 获取图片相对目录
func (us *UploadService) GetImgPath() string {
	return config.ServerConf.StaticRootPath
}

// GetImgFullPath 获取图片完整目录
func (us *UploadService) GetImgFullPath() string {
	return config.ServerConf.StaticRootPath + config.ServerConf.UploadImagePath
}

// GetImgName 获取图片名称
func (us *UploadService) GetImgName(name string) string {
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = utils.MakeSha1(fileName)
	return fileName + ext
}

// GetImgFullUrl 获取图片完整URL
func (us *UploadService) GetImgFullUrl(name string) string {
	return config.ServerConf.PrefixUrl + config.ServerConf.UploadImagePath + name
}

// CheckImgExt 检查图片后缀是否满足要求
func (us *UploadService) CheckImgExt(fileName string) bool {
	ext := path.Ext(fileName)
	for _, allowExt := range config.ServerConf.ImageFormats {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}
	return false
}

// CheckImgSize 检查图片大小是否超出
func (us *UploadService) CheckImgSize(f multipart.File) bool {
	content, err := ioutil.ReadAll(f)
	if err != nil {
		log.Error().Msg(err.Error())
		return false
	}
	// 单位转换 bytes to Megabyte
	const converRatio float64 = 1024 * 1024
	fileSize := float64(len(content)) / converRatio
	// 文件大小不得超出上传限制：5M
	return fileSize <= config.ServerConf.UploadLimit
}

// CheckImgPath 检测图片路径是否创建及权限是否满足
func (us *UploadService) CheckImgPath(path string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}
	isExist, err := utils.IsExist(dir + "/" + path)
	if err != nil {
		return fmt.Errorf("utils.IsExist err: %v", err)
	}
	if isExist == false {
		// 若路径不存在，则创建
		err := os.MkdirAll(dir + "/" + path, os.ModePerm)
		if err != nil {
			return fmt.Errorf("os.MkdirAll err: %v", err)
		}
	}
	isPerm := utils.IsPerm(path)
	if isPerm {
		return fmt.Errorf("utils.IsPerm Permission denied src: %s", path)
	}
	return nil
}
