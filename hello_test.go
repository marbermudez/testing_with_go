package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string){
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q",got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Martin","")
		want := "Hello, Martin"
		assertCorrectMessage(t,got,want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("","")
		want := "Hello, World"
		assertCorrectMessage(t,got,want)
	})
	t.Run("saying hello to people in Spanish", func(t *testing.T) {
		got := Hello("Martin", "Spanish")
		want := "Hola, Martin"
		assertCorrectMessage(t,got,want)
	})
	t.Run("saying hello to people in French", func(t *testing.T) {
		got := Hello("Martin", "French")
		want := "Bonjour, Martin"
		assertCorrectMessage(t,got,want)
	})
	t.Run("saying hello to people in Portuguese", func(t *testing.T) {
		got := Hello("Martin", "Portuguese")
		want := "Olá, Martin"
		assertCorrectMessage(t,got,want)
	})
	t.Run("saying hello to people in German", func(t *testing.T) {
		got := Hello("Martin", "German")
		want := "Hallo, Martin"
		assertCorrectMessage(t,got,want)
	})

}