package settings

type Settings struct {
	PrivateKeyPath     string
	PublicKeyPath      string
	JWTExpirationDelta int
}

func Get() Settings {
	if &settings == nil {
		Init()
	}
	return settings
}
