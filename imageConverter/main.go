package main

//画像変換ターゲットの拡張子は-src=で指定前提
//画像変換後の拡張子は-dst=で指定前提
//ディレクトリは-pathで指定前提

import (
	"flag"
	"fmt"
	cnv "imageConverter/module"
)

var convertOption cnv.ConvertOption
var src, dst, path string

func init() {
	//コマンドライン引数からオプション読み取り
	flag.StringVar(&src, "src", "jpg", "文字列の値を指定")
	flag.StringVar(&dst, "dst", "png", "文字列の値を指定")
	flag.StringVar(&path, "path", "", "文字列の値を指定")
}

func main() {
	flag.Parse()
	//拡張子は小文字化して代入
	convertOption.SrcExtention = cnv.FormatExtention(src)
	convertOption.DstExtention = cnv.FormatExtention(dst)
	convertOption.Path = path

	//バリデーションチェック 課題2でここの改良の話が出てきそう
	var err = cnv.ValidateExtention(convertOption.SrcExtention)
	if err != nil {
		fmt.Println("error at src : " + src)
		return
	}
	err = cnv.ValidateExtention(convertOption.DstExtention)
	if err != nil {
		fmt.Println("error at dst extention : " + dst)
		return
	}

	//変換
	err = cnv.ImageConvert(convertOption)
	if err != nil {
		fmt.Println("error")
	}
}
