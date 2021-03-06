package pirsch

import "testing"

// this can be used to manually test a User-Agent string
func TestParseUserAgentManually(t *testing.T) {
	ua := ParseUserAgent("Mozilla/5.0 (Windows; U; Windows NT 5.2; en-US) AppleWebKit/533.17.8 (KHTML, like Gecko) Version/5.0.1 Safari/533.17.8")
	t.Log(ua.OS)
	t.Log(ua.OSVersion)
	t.Log(ua.Browser)
	t.Log(ua.BrowserVersion)
}

func TestParseUserAgent(t *testing.T) {
	// just a simple test to check ParseUserAgent returns something for a clean User-Agent
	ua := ParseUserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:79.0) Gecko/20100101 Firefox/79.0")

	if ua.OS != OSMac || ua.OSVersion != "10.15" {
		t.Fatalf("Operating system information not as expected: %v %v", ua.OS, ua.OSVersion)
	}

	if ua.Browser != BrowserFirefox || ua.BrowserVersion != "79.0" {
		t.Fatalf("Browser information not as expected: %v %v", ua.Browser, ua.BrowserVersion)
	}
}

func TestGetBrowser(t *testing.T) {
	for _, ua := range userAgentsAll {
		system, products := parseUserAgent(ua.ua)
		browser, version := getBrowser(products, system, ua.os)

		if browser != ua.browser {
			t.Fatalf("Expected browser '%v' for user agent '%v', but was: %v", ua.browser, ua.ua, browser)
		}

		if version != ua.browserVersion {
			t.Fatalf("Expected version '%v' for user agent '%v', but was: %v", ua.browserVersion, ua.ua, version)
		}
	}
}

func TestGetBrowserChromeSafari(t *testing.T) {
	chrome := "AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36"
	system, products := parseUserAgent(chrome)
	browser, version := getBrowser(products, system, OSMac)

	if browser != BrowserChrome || version != "87.0" {
		t.Fatalf("UA must have been detected as Chrome, but was: %v %v", browser, version)
	}

	safari := "AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0.1 Safari/605.1.15"
	system, products = parseUserAgent(safari)
	browser, version = getBrowser(products, system, OSMac)

	if browser != BrowserSafari || version != "14.0" {
		t.Fatalf("UA must have been detected as Safari, but was: %v %v", browser, version)
	}
}

func TestGetOS(t *testing.T) {
	for _, ua := range userAgentsAll {
		system, _ := parseUserAgent(ua.ua)
		os, version := getOS(system)

		if os != ua.os {
			t.Fatalf("Expected OS '%v' for user agent '%v', but was: %v", ua.os, ua.ua, os)
		}

		if version != ua.osVersion {
			t.Fatalf("Expected version '%v' for user agent '%v', but was: %v", ua.osVersion, ua.ua, version)
		}
	}
}

func TestGetProductVersion(t *testing.T) {
	input := []struct {
		product string
		n       int
	}{
		{"", 0},
		{"", 1},
		{"", 42},
		{"   ", 0},
		{"Edg", 0},
		{"Edg/", 0},
		{"Edg/   ", 0},
		{"Safari/537.36", 0},
		{"Edg/79.0.309.43", 1},
		{"Chrome/79.0.3945.74", 2},
		{"Chrome/79.0.3945.74", 10},
	}
	expected := []string{
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"537",
		"79.0",
		"79.0.3945",
		"79.0.3945.74",
	}

	for i, in := range input {
		if version := getProductVersion(in.product, in.n); version != expected[i] {
			t.Fatalf("Expected version '%v' for string '%v' and n %v, but was: %v", expected[i], in.product, in.n, version)
		}
	}
}

func TestGetOSVersion(t *testing.T) {
	input := []struct {
		version string
		n       int
	}{
		{"", 0},
		{"", 1},
		{"", 42},
		{"   ", 0},
		{"10.0", 0},
		{"10.0", 1},
		{"10.15.6", 2},
		{"10.15.6", 42},
	}
	expected := []string{
		"",
		"",
		"",
		"",
		"10",
		"10.0",
		"10.15.6",
		"10.15.6",
	}

	for i, in := range input {
		if version := getOSVersion(in.version, in.n); version != expected[i] {
			t.Fatalf("Expected version '%v' for string '%v' and n %v, but was: %v", expected[i], in.version, in.n, version)
		}
	}
}

func TestParse(t *testing.T) {
	input := []string{
		// empty
		"",
		"  ",
		"'  '",
		` "   "`,

		// clean and simple
		"(system)",
		"version",
		"(system) version",

		// whitespace
		"   (system)   ",
		"   version    ",
		"   (   system   )   version   ",
		"   (  ;  system    ;  )   version   ",

		// multiple system entries and versions
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36",
	}
	expected := [][][]string{
		{{}, {}},
		{{}, {}},
		{{}, {}},
		{{}, {}},
		{{"system"}, {}},
		{{}, {"version"}},
		{{"system"}, {"version"}},
		{{"system"}, {}},
		{{}, {"version"}},
		{{"system"}, {"version"}},
		{{"system"}, {"version"}},
		{{"Macintosh", "Intel Mac OS X 10_10_5"}, {"Chrome/63.0.3239.132", "Safari/537.36"}},
	}

	for i, in := range input {
		system, products := parseUserAgent(in)

		if !testStringSlicesEqual(system, expected[i][0]) || !testStringSlicesEqual(products, expected[i][1]) {
			t.Fatalf("%v, expected: %v %v, was: %v %v", in, expected[i][0], expected[i][1], system, products)
		}
	}
}

func testStringSlicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
