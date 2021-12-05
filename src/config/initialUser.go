package config

type InitialUserConfig struct {
	Mail              string
	Password          string
	AccessToken       string
	AccessTokenSecret string
}

func LoadInitialUserConfiguration() *InitialUserConfig {
	requireUserConfiguration, _ := GetEnvBool("CONFIGURE_INITIAL_USER")

	if !requireUserConfiguration {
		return nil
	}

	mail := RequireEnvString("INITIAL_USER_MAIL")
	password := RequireEnvString("INITIAL_USER_PASSWORD")
	accessToken := RequireEnvString("INITIAL_USER_ACCESS_TOKEN")
	accessTokenSecret := RequireEnvString("INITIAL_USER_ACCESS_TOKEN_SECRET")

	return &InitialUserConfig{
		Mail:              mail,
		Password:          password,
		AccessToken:       accessToken,
		AccessTokenSecret: accessTokenSecret,
	}
}
