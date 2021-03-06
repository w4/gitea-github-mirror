package mirror

import (
	"fmt"

	"encoding/json"
	"os"

	"github.com/namsral/flag"
)

type Config struct {
	DbUser         string
	DbPass         string
	DbUrl          string
	DbName         string
	IP             string
	Port           int
	Path           string
	UserConfigFile string
	BaseUrl        string
}

type ServiceConfig struct {
	Url string
}

type LocalUser struct {
	Ignores     []string
	RemoteUsers []string `json:"remote_users"`
}

type RemoteUser struct {
	Username string
	Service  string
	Token    string
}

type UserConfig struct {
	Users       map[string]LocalUser
	Services    map[string]ServiceConfig
	RemoteUsers map[string]RemoteUser `json:"remote_users"`
}

var cfg Config

func GetConfig() Config {
	if cfg.DbUser != "" {
		return cfg
	}
	user := flag.String("db_user", "gitea", "Gitea db user (defaults to gitea)")
	pass := flag.String("db_pass", "", "Gitea db password")
	db := flag.String("db", "gitea", "Gitea db (defaults to gitea)")
	url := flag.String("db_url", "127.0.0.1", "Mysql DB url (defaults to 127.0.0.1)")
	users := flag.String("user_conf", "user.json", "File containing mappings of gitea users to github credentials")
	apiIp := flag.String("ip", "127.0.0.1", "Webhook IP (network accessable by Gitea/external)")
	port := flag.Int("port", 9001, "Webhook port (defaults to 9001)")
	path := flag.String("repo_path", "", "Path to Gitea repositories")
	baseurl := flag.String("base_url", "", "Base URL for your Gitea instance")
	fmt.Println("Loading config...")
	flag.Parse()
	cfg = Config{*user, *pass, *url, *db, *apiIp, *port, *path, *users, *baseurl}
	return cfg
}

func GetUserConfig(cfg Config) UserConfig {
	f, err := os.Open(cfg.UserConfigFile)
	if err != nil {
		panic("Error opening user config: " + err.Error())
	}
	defer f.Close()
	decode := json.NewDecoder(f)
	ucfg := UserConfig{}
	err = decode.Decode(&ucfg)
	if err != nil {
		panic("Error reading user config: " + err.Error())
	}
	return ucfg
}
