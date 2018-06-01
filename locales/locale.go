package locales

var Langs = map[string]map[int32]string{
	"zh-CN": zhCN,
	"en":    en,
}

func Get(lang string, code int32) string {
	if locale, ok := Langs[lang]; ok {
		if message, ok := locale[code]; ok {
			return message
		}
	}
	return ""
}
