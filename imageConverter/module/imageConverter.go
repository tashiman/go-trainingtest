package module

import (
	"errors"
	"fmt"
	"image"
	_ "image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

type ConvertOption struct {
	SrcExtention string
	DstExtention string
	Path         string
}

var ImageExtentions = [4]string{"jpg", "jpeg", "png", "gif"}

//指定パス以下の指定拡張の画像を指定の形式に変換
func ImageConvert(opt ConvertOption) error {
	//指定ディレクトリをwalkして都度convert呼ぶ
	err := filepath.Walk(opt.Path,
		func(path string, info os.FileInfo, err error) error {
			if filepath.Ext(path) == ("." + opt.SrcExtention) {
				err := convert(path, opt.DstExtention)
				if err != nil {
					return err
				}
			}
			return nil
		})
	if err != nil {
		return err
	}
	return nil
}

//指定の拡張子のバリデーション(小文字化前提)
func ValidateExtention(extention string) error {
	for i := 0; i < len(ImageExtentions); i++ {
		if extention == ImageExtentions[i] {
			return nil
		}
	}
	return errors.New("improper extention")
}

//指定拡張子のフォーマット(小文字化してるだけ)
func FormatExtention(extention string) string {
	return strings.ToLower(extention)
}

//ファイル単品を変換
func convert(path, dstExtention string) error {
	//.付ける
	extention := "." + dstExtention
	//変換元画像を開く
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("file open error : " + path)
		return err
	}
	//変換元の画像を閉じる
	defer file.Close()

	//画像オブジェクトに変換
	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("image decode error : " + path)
		return err
	}

	//出力用ファイル
	newPath := strings.Split(path, ".")[0] + extention
	outputFile, err := os.Create(newPath)
	if err != nil {
		fmt.Println("output file error : " + newPath)
		return err
	}
	//出力用ファイル閉じる
	defer outputFile.Close()

	//画像出力
	switch dstExtention {
	case "png":
		err := png.Encode(outputFile, img)
		if err != nil {
			fmt.Println("encode error : to png -> " + path)
			return err
		}
	case "jpeg", "jpg":
		err := jpeg.Encode(outputFile, img, nil)
		if err != nil {
			fmt.Println("encode error : to jpg -> " + path)
			return err
		}
	default:
		fmt.Println("unknown format -> " + path)
		return errors.New("image convert unknown format")
	}
	fmt.Println(path + " -> " + newPath)
	return nil
}
