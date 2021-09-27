package handler

import (
	"geocoder-v2/global"
	"geocoder-v2/model"
	"geocoder-v2/util"
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

func Upload(context *gin.Context) {
	file, err := context.FormFile("file")
	if err != nil {
		context.String(http.StatusBadRequest, "接收文件失败！")
		return
	}
	path := "upload/" + file.Filename
	if err := context.SaveUploadedFile(file, path); err != nil {
		context.String(http.StatusBadRequest, "保存文件失败！")
		return
	}

	//读取 CSV 文件
	csv := util.ReadCsv(path)

	locations := make([]model.Location, 0)

	for _, row := range csv {
		input := row[0]
		smallLocSlice := make([]string, 0)
		largeLocSlice := make([]string, 0)
		// 查找所有被包裹的大地点 %%
		var (
			smallLoc  string
			largeLoc  string
			queryText string
		)
		location := model.Location{
			InputText: input,
		}
		regLarge := regexp.MustCompile(`%[\s\S]*?%`)
		regSmall := regexp.MustCompile(`\$[\s\S]*?\$`)
		if allString := regLarge.FindAllString(input, -1); len(allString) > 0 {
			for _, s := range allString {
				largeLocSlice = append(largeLocSlice, strings.Trim(s, "%"))
			}
			largeLoc = strings.Replace(strings.Join(largeLocSlice, "%20"), " ", "%20", -1)
			//fmt.Println(largeLoc)
		}
		if allString := regSmall.FindAllString(input, -1); len(allString) > 0 {
			for _, s := range allString {
				smallLocSlice = append(smallLocSlice, strings.Trim(s, "$"))
			}
			smallLoc = strings.Replace(strings.Join(smallLocSlice, "%20"), " ", "%20", -1)
		}
		// 如果没有包裹词，直接将输入作为查询词语
		if len(smallLoc) == 0 && len(largeLoc) == 0 {
			queryText = strings.Replace(input, " ", "%20", -1)
		} else if len(smallLoc) > 0 {
			queryText = smallLoc
		} else {
			queryText = largeLoc
		}
		// 去除 queryText 的各种特殊符号
		queryText = strings.ReplaceAll(queryText, ".", "")
		queryText = strings.ReplaceAll(queryText, "&", "")
		queryText = strings.ReplaceAll(queryText, ":", "")
		queryText = strings.ReplaceAll(queryText, ";", "")
		queryText = strings.ReplaceAll(queryText, "<", "")
		location.QueryText = queryText
		// 将 location 存入数据库

		// 开始查询 暂时写在这里
		// 拼接查询 url
		//url := "http://dev.virtualearth.net/REST/v1/Locations/" + queryText + "?include=queryParse&maxRes=1&key=" + global.KEY

		global.DB.Create(&location)
		locations = append(locations, location)
	}
	/*for i := range locations {
		fmt.Println(locations[i].InputText, locations[i].QueryText)
	}*/
}
