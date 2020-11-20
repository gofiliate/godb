package godb

func LoadConfigExternal(mode string, username string , password string, host string, port int, database string) {

	if mode == "read" {
		 readUser = username
		 readPass = password
		 readHost =  host
		 readPort = port
		 readDB = database
	}

	if mode == "write" {
		writeUser = username
		writePass = password
		writeHost =  host
		writePort = port
		writeDB = database
	}
}
