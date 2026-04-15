package main

func main() {
	config := config{}
	configPtr := &config
	startRepl(configPtr)
}
