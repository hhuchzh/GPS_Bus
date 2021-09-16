package main

func init() {
	PriceConfig.Init()
}

func main() {
	app := NewApp()
	app.Run()
}