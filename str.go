package main

import (
	"strings"
	"fmt"
)

func main(){
	srt := "这是一段字符串"

	srt += "增加的第二段文字"


	//strings.HasPrefix 判断字符串是否以xxx开头
	isSrtFist := strings.HasPrefix(srt,"这")

	//strings.HasSuffix 判断字符串是否以xxx结尾
	isSrtLast := strings.HasSuffix(srt,"ha")

	//strings.Contains 判断字符串是否包含
	isSrt := strings.Contains(srt,"一")

	//返回对应字符串的索引
	isIndex := strings.Index(srt,"字")

	//返回字符串中最后出现的索引
	isLastIndex := strings.LastIndex(srt,"字")


	//替换字符串
	rep := strings.Replace(srt,"的","di",1)



	//统计字符串中出现的次数
	cun := strings.Count(srt,"字")

	fmt.Println(rep)
	fmt.Println(srt)


	var newSrt string
	//重复字符串
	newSrt = strings.Repeat(srt,3)

	fmt.Println("重复字符串:"+newSrt)

	fmt.Println(isSrtFist,isSrtLast,isSrt,isIndex,isLastIndex,cun)


	//ToLower 将字符串中的 Unicode 字符全部转换为相应的小写字符
	//ToUpper 将字符串中的 Unicode 字符全部转换为相应的大写字符
}


