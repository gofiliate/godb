package godb

type ConnectionDetails struct {

	DbrUser string `json:"DbrUser"`
	DbrPass string `json:"DbrPass"`
	DbrHost string `json:"DbrHost"`
	DbrPort string `json:"DbrPort"`
	DbrName string `json:"DbrName"`
	DbwUser string `json:"DbwUser"`
	DbwPass string `json:"DbwPass"`
	DbwHost string `json:"DbwHost"`
	DbwPort string `json:"DbwPort"`
	DbwName string `json:"DbwName"`
	SslMode string `json:"SslMode"`
}
