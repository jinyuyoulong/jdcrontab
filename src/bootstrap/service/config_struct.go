package service

import "time"

type Config struct {
	App struct {
		Name  string
		URL   string
		Port  string
		Debug bool
		AppOwner string
	}

	Database struct {
		Dirver string
	} `toml:"database"`

	Mysql struct {
		Dbname   string
		Username string
		Password string
	} `toml:"mysql"`

	Website struct {
		static_uri string
		site_title string
		copy_right string
	}

	Redis struct {
		// Network "tcp"
		Network string
		// Addr "127.0.0.1:6379"
		Addr string
		// Password string .If no password then no 'AUTH'. Default ""
		Password string
		// If Database is empty "" then no 'SELECT'. Default ""
		Database string
		// MaxIdle 0 no limit
		MaxIdle int
		// MaxActive 0 no limit
		MaxActive int
		// IdleTimeout  time.Duration(5) * time.Minute
		IdleTimeout time.Duration
		// Prefix "myprefix-for-this-website". Default ""
		Prefix string
	}

	Session struct {
		Cookie  string
		Expires time.Duration
		Dirver  string
	}
}