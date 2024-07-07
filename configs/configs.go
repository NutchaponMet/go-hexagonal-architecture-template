package configs

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	infisical "github.com/infisical/go-sdk"
	"github.com/spf13/viper"
)

var (
	cfgs map[string]string
)

func Init() {
	initConfigLoader()
	initTimeZone()
	loadConfigs(&cfgs)
}

func initConfigLoader() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.ReadInConfig()
}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	time.Local = ict
}

func loadConfigs(cfg *map[string]string) {
	mode := viper.GetString("infisical.mode")
	url := viper.GetString("infisical.url")
	clientID := viper.GetString("infisical.clientID")
	clientSecret := viper.GetString("infisical.clientSecret")
	projectID := viper.GetString("infisical.projectID")

	client := infisical.NewInfisicalClient(infisical.Config{
		SiteUrl: url,
	})

	_, err := client.Auth().UniversalAuthLogin(clientID, clientSecret)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	apiKeySecrets, err := client.Secrets().List(infisical.ListSecretsOptions{
		Environment: mode,
		SecretPath:  "/go-backend",
		ProjectID:   projectID,
	})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	conf := map[string]string{}
	for _, obj := range apiKeySecrets {
		conf[obj.SecretKey] = obj.SecretValue
	}
	*cfg = conf
}

func GetString(s string) string {
	return cfgs[s]
}

func GetInt(s string) int {
	number, _ := strconv.Atoi(cfgs[s])
	return number
}
