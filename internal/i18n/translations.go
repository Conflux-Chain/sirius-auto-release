package i18n

import (
	"fmt"
	"reflect"
)

type Messages map[string]string

var translations = map[string]Messages{}

func RegisterTranslation(lang, key, translation string) {
	if _, ok := translations[lang]; !ok {
		translations[lang] = Messages{}
	}
	translations[lang][key] = translation
}

func RegisterMessages(lang string, messages Messages) {
	if _, ok := translations[lang]; !ok {
		translations[lang] = Messages{}
	}

	for key, translation := range messages {
		translations[lang][key] = translation
	}
}
func T(lang, key string) string {
	if messages, ok := translations[lang]; ok {
		if message, ok := messages[key]; ok {
			return message
		}
	}
	return key
}

func ExtractAllTranslations(obj any) {
	objType := reflect.TypeOf(obj)

	if objType.Kind() == reflect.Ptr {
		objType = objType.Elem()
	}

	if objType.Kind() != reflect.Struct {
		return
	}

	prefix := objType.Name()
	extractStructTranslations(objType, prefix)

}

func extractStructTranslations(objType reflect.Type, prefix string) {
	for i := range objType.NumField() {
		field := objType.Field(i)

		key := fmt.Sprintf("%s.%s", prefix, field.Name)

		if field.Type.Kind() == reflect.Struct {
			newPrefix := fmt.Sprintf("%s.%s", prefix, field.Name)
			extractStructTranslations(field.Type, newPrefix)
			continue
		}

		for _, lang := range []string{"en", "zh"} {
			tagName := fmt.Sprintf("huh_%s", lang)
			if translation, ok := field.Tag.Lookup(tagName); ok {
				RegisterTranslation(lang, key, translation)
			}
		}
	}
}

func DumpTranslations() {
	fmt.Println("=== Current Translations ===")
	for lang, messages := range translations {
		fmt.Printf("Language: %s\n", lang)
		for k, v := range messages {
			fmt.Printf("  %s: %s\n", k, v)
		}
	}
	fmt.Println("===========================")
}
