package util

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"net"
	"net/http"
)

var DB *gorm.DB

func InitDB() {
	sqlType := viper.GetString("db.sqlType")
	database := viper.GetString("db.database")
	switch sqlType {
	case "mysql":
		username := viper.GetString("db.username")
		password := viper.GetString("db.password")
		host := viper.GetString("db.host")
		port := viper.GetString("db.port")
		dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"
		var err error
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Panicln("mysql数据库连接失败。", err)
		}
	case "sqlite":
		var err error
		DB, err = gorm.Open(sqlite.Open(database+".db"), &gorm.Config{})
		if err != nil {
			log.Panicln("sqlite数据库连接失败。", err)
		}
	}

}

func GetIPs() (localIP, publicIP string) {

	resp, err := http.Get("https://myexternalip.com/raw")
	if err != nil {
		publicIP = "0.0.0.0"
	}
	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
	publicIP = string(content)
	localIP = "0.0.0.0"
	interfaces, _ := net.Interfaces()
	for _, iface := range interfaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			localIP = ip.To4().String()
			if len(localIP) >= 7 && len(localIP) <= 15 {
				return
			}
		}
	}
	return
}
