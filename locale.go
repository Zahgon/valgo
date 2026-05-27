package valgo

const (
	LocaleCodeEn = "en"
	LocaleCodeEs = "es"
	LocaleCodeDe = "de"
	LocaleCodeHu = "hu"
)

const localeCodeDefault = LocaleCodeEn

// Locale is a type alias that represents a map of locale entries.
// The keys in the map are strings that represent the entry's identifier, and
// the values are strings that contain the corresponding localized text
// for that entry
type Locale map[string]string

func getLocaleWithSkipDefaultOption(code string, skipDefault bool, factoryLocales ...map[string]*Locale) *Locale {
	_ = "STUB: not implemented"
	return nil
}

func getLocaleAndSkipDefaultOption(code string, factoryLocales ...map[string]*Locale) *Locale {
	_ = "STUB: not implemented"
	return nil
}

func getLocale(code string, factoryLocales ...map[string]*Locale) *Locale {
	_ = "STUB: not implemented"
	return nil
}

func (_locale *Locale) merge(locale *Locale) *Locale { _ = "STUB: not implemented"; return nil }
