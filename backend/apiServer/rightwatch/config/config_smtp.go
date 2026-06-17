package config

import (
	"github.com/spf13/viper"
	"rightwatch/services"
)

// GetSMTP config.toml의 [smtp] 섹션을 읽어 services.SMTPConfig를 반환한다.
// host가 비면 mailer는 dry-run으로 폴백한다.
func GetSMTP() services.SMTPConfig {
	return services.SMTPConfig{
		Host:     viper.GetString("smtp.host"),
		Port:     viper.GetInt("smtp.port"),
		Username: viper.GetString("smtp.username"),
		Password: viper.GetString("smtp.password"),
		From:     viper.GetString("smtp.from"),
	}
}
