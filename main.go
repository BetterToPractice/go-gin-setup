package main

import "github.com/BetterToPractice/go-gin-setup/cmd"

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @schemes http https
// @basePath /
func main() {
	cmd.Execute()
}
