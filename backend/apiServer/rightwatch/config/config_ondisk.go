package config

import (
	"os"

	"github.com/spf13/viper"
	"rightwatch/services"
)

// GetOndiskConfig reads ondisk login config for the screenshotter.
// Credentials prefer env vars (ONDISK_USERNAME/ONDISK_PASSWORD) over the
// config.toml [ondisk] section so secrets stay out of committed config (M0 hygiene).
// Username이 비면 screenshotter는 로그인을 건너뛰고 공개 화면만 캡처한다.
func GetOndiskConfig() services.OndiskConfig {
	username := os.Getenv("ONDISK_USERNAME")
	if username == "" {
		username = viper.GetString("ondisk.username")
	}
	password := os.Getenv("ONDISK_PASSWORD")
	if password == "" {
		password = viper.GetString("ondisk.password")
	}
	loginURL := viper.GetString("ondisk.login_url")
	if loginURL == "" {
		loginURL = "https://ondisk.co.kr/index.php"
	}
	return services.OndiskConfig{
		LoginURL: loginURL,
		Username: username,
		Password: password,
	}
}
