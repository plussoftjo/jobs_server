// Package config ..
package config

// Server ..
type Server struct {
	PublicPath string
	Dev        bool
	ServerName string
	ServerURI  string
}

// ServerInfo ..
var ServerInfo Server

// ServerInformations ..
func ServerInformations() Server {
	dev := false
	serverName := "jobs"

	var publicPath string
	var serverURI string
	if dev {
		publicPath = "./"
		serverURI = "http://192.168.1.108:8082/"

	} else {
		publicPath = "/var/www/" + serverName + "/"
		serverURI = "http://13.36.230.155:3002/"
	}

	ServerInfo = Server{
		Dev:        dev,
		PublicPath: publicPath,
		ServerName: serverName,
		ServerURI:  serverURI,
	}
	return Server{
		Dev:        dev,
		PublicPath: publicPath,
		ServerName: serverName,
		ServerURI:  serverURI,
	}
}
