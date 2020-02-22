package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// *******************
// TEST URL FORMATTING
// *******************
func TestUrl1(t *testing.T) {
	url := "https://www.google.com"
	expected := urlParts{proto: "https", domain: "google.com", resource: "/"}

	actual := noramliseAddress(url)
	assert.Equal(t, expected, actual)
}

func TestUrl2(t *testing.T) {
	url := "https://google.com"
	expected := urlParts{proto: "https", domain: "google.com", resource: "/"}

	actual := noramliseAddress(url)
	assert.Equal(t, expected, actual)
}

func TestUrl3(t *testing.T) {
	url := "http://google.com"
	expected := urlParts{proto: "http", domain: "google.com", resource: "/"}

	actual := noramliseAddress(url)
	assert.Equal(t, expected, actual)
}

func TestUrl4(t *testing.T) {
	url := "http://www.google.com"
	expected := urlParts{proto: "http", domain: "google.com", resource: "/"}

	actual := noramliseAddress(url)
	assert.Equal(t, expected, actual)
}

func TestUrl5(t *testing.T) {
	url := "https://google.com"
	expected := urlParts{proto: "https", domain: "google.com", resource: "/"}

	actual := noramliseAddress(url)
	assert.Equal(t, expected, actual)
}

func TestUrl6(t *testing.T) {
	url := "http://google.com"
	expected := urlParts{proto: "http", domain: "google.com", resource: "/"}

	actual := noramliseAddress(url)
	assert.Equal(t, expected, actual)
}

func TestUrl7(t *testing.T) {
	url := "google.com"
	expected := urlParts{proto: "", domain: "google.com", resource: "/"}

	actual := noramliseAddress(url)
	assert.Equal(t, expected, actual)
}

func TestUrl8(t *testing.T) {
	url := "www.google.com"
	expected := urlParts{proto: "", domain: "google.com", resource: "/"}

	actual := noramliseAddress(url)
	assert.Equal(t, expected, actual)
}

func TestUrl9(t *testing.T) {
	url := "subdomain.google.com"
	expected := urlParts{proto: "", domain: "subdomain.google.com", resource: "/"}

	actual := noramliseAddress(url)
	assert.Equal(t, expected, actual)
}

func TestUrl10(t *testing.T) {
	url := "https://subdomain.google.com"
	expected := urlParts{proto: "https", domain: "subdomain.google.com", resource: "/"}

	actual := noramliseAddress(url)
	assert.Equal(t, expected, actual)
}

func TestUrl11(t *testing.T) {
	url := "https://subdomain.google.com/some/other/resource?val=key"
	expected := urlParts{proto: "https", domain: "subdomain.google.com", resource: "/some/other/resource?val=key"}

	actual := noramliseAddress(url)
	assert.Equal(t, expected, actual)
}

func TestUrl12(t *testing.T) {
	url := "/some/other/resource?val=key"
	expected := urlParts{proto: "", domain: "", resource: "/some/other/resource?val=key"}

	actual := noramliseAddress(url)
	assert.Equal(t, expected, actual)
}

// *******************
// TEST SAME SITE TEST
// *******************

func TestIsSame1(t *testing.T) {
	link := urlParts{proto: "https", domain: "google.com", resource: "/some/other/resource?val=key"}
	site := urlParts{proto: "https", domain: "google.com", resource: "/some/oasdsa?adsdsds=dasd"}

	assert.True(t, isLinkSameWebsite(link, site))

}

func TestIsSame2(t *testing.T) {
	link := urlParts{proto: "https", domain: "", resource: "/some/other/resource?val=key"}
	site := urlParts{proto: "https", domain: "google.com", resource: "/some/oasdsa?adsdsds=dasd"}

	assert.True(t, isLinkSameWebsite(link, site))
}
