/*
 *  @Author : huangzj
 *  @Time : 2020/12/16 14:27
 *  @Description：
 */

package participle

import (
	"fmt"
	"strings"
	"testing"

	"github.com/yanyiwu/gojieba"
)

const (
	separator = "|" //字符串拼接，和该工具无关
)

func TestParticipleUseHMM(t *testing.T) {
	var seg = gojieba.NewJieba()
	defer seg.Free()

	var resWords []string
	var sentence = "万里长城万里长"

	resWords = seg.CutAll(sentence)
	fmt.Printf("%s\t全模式：%s \n", sentence, strings.Join(resWords, separator))

	resWords = seg.Cut(sentence, true)
	fmt.Printf("%s\t精确模式：%s \n", sentence, strings.Join(resWords, separator))
	var addWord = "万里长"
	seg.AddWord(addWord)
	fmt.Printf("添加新词：%s\n", addWord)

	resWords = seg.Cut(sentence, true)
	fmt.Printf("%s\t精确模式：%s \n", sentence, strings.Join(resWords, separator))

	sentence = "北京鲜花速递"
	resWords = seg.Cut(sentence, true)
	fmt.Printf("%s\t新词识别：%s \n", sentence, strings.Join(resWords, separator))

	sentence = "北京鲜花速递"
	resWords = seg.CutForSearch(sentence, true)
	fmt.Println(sentence, "\t搜索引擎模式：", strings.Join(resWords, separator))

	sentence = "北京市朝阳公园"
	resWords = seg.Tag(sentence)
	fmt.Println(sentence, "\t词性标注：", strings.Join(resWords, separator))

	sentence = "鲁迅先生"
	resWords = seg.CutForSearch(sentence, false)
	fmt.Println(sentence, "\t搜索引擎模式：", strings.Join(resWords, separator))

	words := seg.Tokenize(sentence, gojieba.SearchMode, false)
	fmt.Println(sentence, "\tTokenize Search Mode 搜索引擎模式：", words)

	words = seg.Tokenize(sentence, gojieba.DefaultMode, false)
	fmt.Println(sentence, "\tTokenize Default Mode搜索引擎模式：", words)

	word2 := seg.ExtractWithWeight(sentence, 5)
	fmt.Println(sentence, "\tExtract：", word2)
}
