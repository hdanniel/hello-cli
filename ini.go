package main

import (
	"github.com/go-ini/ini"
)

/*
[credentials]
client_id = XXXXXXXX
client_secret = YYYYYYYYYYY

*/


func SaveCredentialsToFile(client_id string, client_secret string) error {
	cfg, err := ini.Load(getHomeDir() + "/" + cfg_dir + "/" + cfg_file)
	if err != nil {
		return err
	}
	cfg.Section("credentials").Key("client_id").SetValue(client_id)
	cfg.Section("credentials").Key("client_secret").SetValue(client_secret)
	err = cfg.SaveTo(getHomeDir() + "/" + cfg_dir + "/" + cfg_file)

	if err != nil {
		return err
	}
	return nil

}