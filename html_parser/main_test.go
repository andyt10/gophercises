package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBaseHtmlFile(t *testing.T) {
	fileName := "tests/basetest.html"
	expected := []link{{Href: "/other-page", Text: "A link to another page"}}

	file, _ := openSource(fileName)
	actual := loadSource(file)

	assert.Equal(t, expected, actual)
}

func TestFile1(t *testing.T) {
	fileName := "tests/test1.html"
	expected := []link{{Href: "/other-page", Text: "A link to another page"}}

	file, _ := openSource(fileName)
	actual := loadSource(file)

	assert.Equal(t, expected, actual)
}

func TestFile2(t *testing.T) {
	fileName := "tests/test2.html"
	expected := []link{{Href: "https://www.twitter.com/joncalhoun", Text: "Check me out on twitter"}, {Href: "https://github.com/gophercises", Text: "Gophercises is on Github!"}} // Github!

	file, _ := openSource(fileName)
	actual := loadSource(file)

	assert.Equal(t, expected, actual)
}

func TestFile3(t *testing.T) {
	fileName := "tests/test3.html"
	expected := []link{{Href: "#", Text: "Login"}, {Href: "/lost", Text: "Lost? Need help?"}, {Href: "https://twitter.com/marcusolsson", Text: "@marcusolsson"}}

	file, _ := openSource(fileName)
	actual := loadSource(file)

	assert.Equal(t, expected, actual)
}

func TestFile4(t *testing.T) {
	fileName := "tests/test4.html"
	expected := []link{{Href: "/dog-cat", Text: "dog cat"}}

	file, _ := openSource(fileName)
	actual := loadSource(file)

	assert.Equal(t, expected, actual)
}
