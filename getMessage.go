package golocale

func GetMessage(lang string, label string) string {

	if Source[lang] == nil {
		return Source[Locale][label]
	}

	return Source[lang][label]
}
