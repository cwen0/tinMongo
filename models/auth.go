package models

type AuthConfig struct {
	HostName string `json:"hostname"`
	Port     int    `json:"port"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}
