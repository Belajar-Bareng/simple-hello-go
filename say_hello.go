package simplehellogo

import "fmt"

var lang string
var langs map[string]string = map[string]string{
	"ID": "Halo! Salam kenal",
	"EN": "Hello there",
	"ES": "Hola",
}

func init() {
	lang = "ID"
}

func Hello() string {
	return langs[lang]
}

func SayHello() {
	fmt.Println(Hello())
}

func GreetPerson(name string) string {
	return Hello() + ", " + name
}

func SayGreetPerson(name string) {
	fmt.Println(GreetPerson(name))
}

func GetCurrentLanguage() string {
	return lang
}

func GetAvailableLanguages() []string {
	var avl []string
	for cd := range langs {
		avl = append(avl, cd)
	}
	return avl
}

func SetLanguange(language string) bool {
	for cd := range langs {
		if language == cd {
			lang = language
			return true
		}
	}

	return false
}
