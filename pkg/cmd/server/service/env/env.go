package env

import (
	"os"
)

var ServerURL = os.Getenv("URL")

var Broker = &struct {
	URI string
}{
	URI: os.Getenv("BROKER_URI"),
}

var Database = &struct {
	URI string
}{
	URI: os.Getenv("DATABASE_URI"),
}

var SMTP = &struct {
	Host, Port, Hotelname, Password, ServerURL string
}{
	Host:      os.Getenv("SMTP_HOST"),
	Port:      os.Getenv("SMTP_PORT"),
	Hotelname: os.Getenv("SMTP_HotelNAME"),
	Password:  os.Getenv("SMTP_PASSWORD"),
	ServerURL: ServerURL,
}

var Security = &struct {
	AllowedHosts string
}{
	AllowedHosts: os.Getenv("ALLOWED_HOSTS"),
}

var JWT = &struct {
	SecretKey string
}{
	SecretKey: os.Getenv("JWT_SECRET_KEY"),
}

var Cookie = &struct {
	SecretKey, SessionName string
}{
	SecretKey:   os.Getenv("COOKIE_SECRET_KEY"),
	SessionName: os.Getenv("COOKIE_SESSION_NAME"),
}
