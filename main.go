package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strings"
)

type Object = map[string]any
type Array = []any

// This application is to translate the json file to Golang struct
func main() {
	data := make(Object)
	_ = json.Unmarshal([]byte(large), &data)
	walk(data)
	//_ = "{ \"id\": 1, \"name\": \"wzy\", \"marry\": false, \"no\": null }"
	fmt.Println()
	s := ToStruct(large, "Data", nil)
	fmt.Println(s)

	//item := Item{}
	//err := json.Unmarshal([]byte(large), &item)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("%v", item)

}

func walk(data any) {
	switch data.(type) {
	case Object:
		fmt.Println("#JsonObject", "len: ", len(data.(Object)))
		for key, value := range data.(Object) {
			fmt.Printf("Key [\"%v\"]: ", key)
			walk(value)
		}
	case Array:
		fmt.Println("#JsonArray", "len: ", len(data.(Array)))
		for i, value := range data.(Array) {
			fmt.Printf("Index [%v]: ", i)
			walk(value)
		}
	default:
		switch data.(type) {
		case string:
			fmt.Printf(" '%v' ", data)
		case int:
			fmt.Printf(" %v ", data)
		case bool:
			fmt.Printf(" %v ", data)
		default:
			fmt.Printf(" %v ", data)
		}
	}
}

var path []string = make([]string, 0)

func ToStruct(s string, name string, ignore map[string]bool) string {
	obj := make(Object)
	err := json.Unmarshal([]byte(s), &obj)
	if err != nil {
		log.Fatal(err)
	}
	ss := SBuilder{}
	ss.writeString("type ").writeString(name).writeRune(' ')
	build(obj, &ss, 0)
	return ss.String()
}

type SBuilder struct {
	strings.Builder
}

func (br *SBuilder) writeString(s string) *SBuilder {
	if _, err := br.WriteString(s); err != nil {
		log.Fatal(err)
	}
	return br
}

func (br *SBuilder) writeRune(r rune) *SBuilder {
	if _, err := br.WriteRune(r); err != nil {
		log.Fatal(err)
	}
	return br
}

func build(elem any, builder *SBuilder, depth int) {
	if elem == nil {
		builder.writeString("any\n    ")
		return
	}
	switch elem.(type) {
	case Object:
		depth++
		builder.writeString("struct {\n")

		for key, value := range elem.(Object) {
			for i := 0; i < depth; i++ {
				builder.writeString("\t")
			}

			builder.writeString(strings.Title(key)).writeString(" ")
			build(value, builder, depth)
		}
		for i := 0; i < depth-1; i++ {
			builder.writeString("\t")
		}
		builder.writeString("}")
	case Array:
		builder.writeString("[]")
		if len(elem.(Array)) == 0 {
			builder.writeString("any")
		} else {
			build(elem.(Array)[0], builder, depth)
		}
	default:
		builder.writeString(reflect.TypeOf(elem).Kind().String())
	}
	builder.writeString("\n")
}

var str = `
{
	"paramz": {
		"feeds": [{
			"id": 299076,
			"oid": 288340,
			"category": "article",
			"data": {
				"subject": "荔枝新闻3.0：不止是阅读",
				"summary": "江苏广电旗下资讯类手机应用“荔枝新闻”于近期推出全新升级换代的3.0版。",
				"cover": "/Attachs/Article/288340/3e8e2c397c70469f8845fad73aa38165_padmini.JPG",
				"pic": "",
				"format": "txt",
				"changed": "2015-09-22 16:01:41"
			}
		}]
	},
	"Index": 1
}`

var large = `{
	"wordRank":1,
	"headWord":"consistent",
	"content":{"word":{"wordHead":"consistent","wordId":"CET6luan_1_1","content":{"sentence":{"sentences":[{"sContent":"She’s the team’s most consistent player.","sCn":"她是该队中表现最为稳定的选手。"}],"desc":"例句"},"usphone":"kən'sɪstənt","syno":{"synos":[{"pos":"adj","tran":"始终如一的，[数]一致的；坚持的","hwds":[{"w":"united"},{"w":"corresponding"},{"w":"uniform"},{"w":"matching"},{"w":"solid"}]}],"desc":"同近"},"ukphone":"kənˈsɪstənt","ukspeech":"consistent&type=1","star":0,"phrase":{"phrases":[{"pContent":"consistent with","pCn":"符合；与…一致"},{"pContent":"consistent quality","pCn":"始终如一的质量"},{"pContent":"consistent policy","pCn":"一贯的政策"},{"pContent":"consistent principle","pCn":"一致性原则"},{"pContent":"self consistent field","pCn":"自洽场"}],"desc":"短语"},"phone":"kən'sistənt","speech":"consistent","relWord":{"desc":"同根","rels":[{"pos":"adv","words":[{"hwd":"consistently","tran":"一贯地；一致地；坚实地"}]},{"pos":"n","words":[{"hwd":"consistency","tran":"[计] 一致性；稠度；相容性"}]}]},"usspeech":"consistent&type=2","trans":[{"tranCn":"一致的","descOther":"英释","pos":"adj","descCn":"中释","tranOther":"always behaving in the same way or having the same attitudes, standards etc – usually used to show approval"}]}}},"bookId":"CET6luan_1"}`

// Item test
type Item struct {
	HeadWord string
	Content  struct {
		Word struct {
			WordId  string
			Content struct {
				Phrase struct {
					Phrases []struct {
						PCn      string
						PContent string
					}

					Desc string
				}
				Phone string
				Trans []struct {
					TranCn    string
					DescOther string
					Pos       string
					DescCn    string
					TranOther string
				}

				Usphone string
				Syno    struct {
					Synos []struct {
						Pos  string
						Tran string
						Hwds []struct {
							W string
						}
					}

					Desc string
				}
				Ukspeech string
				Star     float64
				Speech   string
				RelWord  struct {
					Desc string
					Rels []struct {
						Pos   string
						Words []struct {
							Hwd  string
							Tran string
						}
					}
				}
				Usspeech string
				Sentence struct {
					Sentences []struct {
						SContent string
						SCn      string
					}

					Desc string
				}
				Ukphone string
			}
			WordHead string
		}
	}
	BookId   string
	WordRank float64
}
