package crypto

type Options struct {
	Type string
}

func (Crypto) Echo(options Options) string {
	return options.Type
}
