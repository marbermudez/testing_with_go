package main

import "fmt"

func main(){
	fmt.Println(Hello("Martin",""))
}

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const portugueseHelloPrefix = "Olá, "
const germanHelloPrefix = "Hallo, "
const spanish = "Spanish"
const french = "French"
const portuguese = "Portuguese"
const german = "German"

func Hello(name string, language string) string {
	if name == ""{
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {

	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case portuguese:
		prefix = portugueseHelloPrefix
	case german:
		prefix = germanHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return prefix
}