package config

type Mongo struct {
	Address     string       `json:"address"`
	DBName      string       `json:"dbName"`
	Credentials *Credentials `json:"credentials"`
}
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
