package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"strings"
)

// 为dest在source中存在的字段自动赋值
func CopyStruct(source interface{}, dest interface{}) {
	//var (
	//	fields  []string
	//	convert bool
	//)
	//cp := Copy{
	//	Dest:    dest,
	//	Source:  source,
	//	Fields:  fields,
	//	Convert: convert,
	//}
	//if err := cp.CopySf(); nil != err {
	//	panic(err)
	//}
	sByte, _ := json.Marshal(source)
	_ = json.Unmarshal(sByte, dest)
}

// CameCase
func ToCamelCase(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}

	if strings.HasSuffix(string(data[:]), "Id") {
		data = append(data[0:len(data)-1], []byte("D")[0])
	}

	return string(data[:])
}

// Snake
func ToSnake(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data[:]))
}

func Md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func GetCookieDomain() string {
	return ".medlinker.com"
}
