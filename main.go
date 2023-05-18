package main

import "github.com/Hamster601/fastweb/cmd"

// @title swagger 接口文档
// @version 2.0
// @description

// @contact.name
// @contact.url
// @contact.email

// @license.name MIT
// @license.url https://github.com/Hamster601/fastweb/blob/master/LICENSE

// @securityDefinitions.apikey  LoginToken
// @in                          header
// @name                        token
// @BasePath /
func main() {
	cmd.Execute()
}
