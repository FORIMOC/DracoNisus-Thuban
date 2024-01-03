package detect

import (
	"regexp"

	"forimoc.DracoNisus-Thuban/model/ModelKeyword"
	"github.com/tidwall/gjson"
	"gorm.io/gorm"
)

// GetKeywordRegexp 对消息使用关键词词表进行正则表达式匹配
// => 匹配的关键词列表
func GetKeywordRegexp(DB *gorm.DB, data string) []string {
	msg := gjson.Get(data, "message").String()
	var keywordList []ModelKeyword.KeywordList
	var keywords []string
	var record []string

	DB.Find(&keywordList)
	for i := 0; i < len(keywordList); i++ {
		keywords = append(keywords, keywordList[i].Keyword)
	}

	// 对于每一个关键词进行正则表达式匹配
	for i := 0; i < len(keywords); i++ {
		re := regexp.MustCompile("(?i)" + keywords[i])
		if len(re.FindStringIndex(msg)) > 0 {
			record = append(record, keywords[i])
		}
	}
	return record
}
