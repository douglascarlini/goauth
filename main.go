package main

var CONFIG Config

func main() {
	url := CONFIG.OAuthURL
	go Open(url)
	Start()
}

