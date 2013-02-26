// Copyright (C) 2012 Miquel Sabaté Solà
// This file is licensed under the MIT license.
// See the LICENSE file.

package user_agent

import "testing"
import "fmt"

// Slice that contains all the tests. Each test is contained in a struct
// that groups the name of the test and the User-Agent string to be tested.
var uastrings = []struct {
    name string
    ua   string
}{
    // Bots
    {"GoogleBot", "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"},
    {"BingBot", "Mozilla/5.0 (compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm)"},

    // Internet Explorer
    {"IE10", "Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; Trident/6.0)"},
    {"Tablet", "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.2; ARM; Trident/6.0; Touch; .NET4.0E; .NET4.0C; Tablet PC 2.0)"},
    {"Touch", "Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; ARM; Trident/6.0; Touch)"},
    {"Phone", "Mozilla/4.0 (compatible; MSIE 7.0; Windows Phone OS 7.0; Trident/3.1; IEMobile/7.0; SAMSUNG; SGH-i917)"},

    // Gecko
    {"FirefoxMac", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.6; rv:2.0b8) Gecko/20100101 Firefox/4.0b8"},
    {"FirefoxMacLoc", "Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.6; en-US; rv:1.9.2.13) Gecko/20101203 Firefox/3.6.13"},
    {"FirefoxLinux", "Mozilla/5.0 (X11; Linux x86_64; rv:17.0) Gecko/20100101 Firefox/17.0"},
    {"FirefoxWin", "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US; rv:1.8.1.14) Gecko/20080404 Firefox/2.0.0.14"},
    {"CaminoMac", "Mozilla/5.0 (Macintosh; U; Intel Mac OS X; en; rv:1.8.1.14) Gecko/20080409 Camino/1.6 (like Firefox/2.0.0.14)"},
    {"Iceweasel", "Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.8.1) Gecko/20061024 Iceweasel/2.0 (Debian-2.0+dfsg-1)"},
    {"SeaMonkey", "Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.6; en-US; rv:1.9.1.4) Gecko/20091017 SeaMonkey/2.0"},
    {"AndroidFirefox", "Mozilla/5.0 (Android; Mobile; rv:17.0) Gecko/17.0 Firefox/17.0"},

    // Opera
    {"OperaMac", "Opera/9.27 (Macintosh; Intel Mac OS X; U; en)"},
    {"OperaWin", "Opera/9.27 (Windows NT 5.1; U; en)"},
    {"OperaMinimal", "Opera/9.80"},
    {"OperaFull", "Opera/9.80 (Windows NT 6.0; U; en) Presto/2.2.15 Version/10.10"},
    {"OperaLinux", "Opera/9.80 (X11; Linux x86_64) Presto/2.12.388 Version/12.10"},
    {"OperaAndroid", "Opera/9.80 (Android 4.2.1; Linux; Opera Mobi/ADR-1212030829) Presto/2.11.355 Version/12.10"},

    // Other
    {"Nil", "nil"},
    {"Compatible", "Mozilla/4.0 (compatible)"},
    {"Mozilla", "Mozilla/5.0"},
    {"Amaya", "amaya/9.51 libwww/5.4.0"},
    {"Rails", "Rails Testing"},
    {"Python", "Python-urllib/2.7"},
    {"Curl", "curl/7.28.1"},

    // WebKit
    {"ChromeLinux", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.97 Safari/537.11"},
    {"ChromeWin7", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.168 Safari/535.19"},
    {"ChromeMinimal", "Mozilla/5.0 AppleWebKit/534.10 Chrome/8.0.552.215 Safari/534.10"},
    {"ChromeMac", "Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_6_5; en-US) AppleWebKit/534.10 (KHTML, like Gecko) Chrome/8.0.552.231 Safari/534.10"},
    {"SafariMac", "Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_6_3; en-us) AppleWebKit/533.16 (KHTML, like Gecko) Version/5.0 Safari/533.16"},
    {"SafariWin", "Mozilla/5.0 (Windows; U; Windows NT 5.1; en) AppleWebKit/526.9 (KHTML, like Gecko) Version/4.0dp1 Safari/526.8"},
    {"iPhone", "Mozilla/5.0 (iPhone; U; CPU like Mac OS X; en) AppleWebKit/420.1 (KHTML, like Gecko) Version/3.0 Mobile/4A102 Safari/419"},
    {"iPod", "Mozilla/5.0 (iPod; U; CPU like Mac OS X; en) AppleWebKit/420.1 (KHTML, like Gecko) Version/3.0 Mobile/4A102 Safari/419"},
    {"iPad", "Mozilla/5.0 (iPad; U; CPU OS 3_2 like Mac OS X; en-us) AppleWebKit/531.21.10 (KHTML, like Gecko) Version/4.0.4 Mobile/7B367 Safari/531.21.10"},
    {"webOS", "Mozilla/5.0 (webOS/1.4.0; U; en-US) AppleWebKit/532.2 (KHTML, like Gecko) Version/1.0 Safari/532.2 Pre/1.1"},
    {"Android", "Mozilla/5.0 (Linux; U; Android 1.5; de-; HTC Magic Build/PLAT-RC33) AppleWebKit/528.5+ (KHTML, like Gecko) Version/3.1.2 Mobile Safari/525.20.1"},
    {"BlackBerry", "Mozilla/5.0 (BlackBerry; U; BlackBerry 9800; en) AppleWebKit/534.1+ (KHTML, Like Gecko) Version/6.0.0.141 Mobile Safari/534.1+"},
    {"BB10", "Mozilla/5.0 (BB10; Touch) AppleWebKit/537.3+ (KHTML, like Gecko) Version/10.0.9.388 Mobile Safari/537.3+"},
    {"Ericsson", "Mozilla/5.0 (SymbianOS/9.4; U; Series60/5.0 Profile/MIDP-2.1 Configuration/CLDC-1.1) AppleWebKit/525 (KHTML, like Gecko) Version/3.0 Safari/525"},
    {"ChromeAndroid", "Mozilla/5.0 (Linux; Android 4.2.1; Galaxy Nexus Build/JOP40D) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166 Mobile Safari/535.19"},
}

// Slice of the expected results from the previous slice.
var expected = []string{
    // Bots
    "Mozilla:5.0 Bot:true Mobile:false",
    "Mozilla:5.0 Bot:true Mobile:false",

    // Internet Explorer
    "Mozilla:5.0 Platform:Windows OS:Windows 8 Browser:Internet Explorer-10.0 Engine:Trident Bot:false Mobile:false",
    "Mozilla:4.0 Platform:Windows OS:Windows 8 Browser:Internet Explorer-7.0 Engine:Trident Bot:false Mobile:false",
    "Mozilla:5.0 Platform:Windows OS:Windows 8 Browser:Internet Explorer-10.0 Engine:Trident Bot:false Mobile:false",
    "Mozilla:4.0 Platform:Windows OS:Windows Phone OS 7.0 Browser:Internet Explorer-7.0 Engine:Trident Bot:false Mobile:true",

    // Gecko
    "Mozilla:5.0 Platform:Macintosh OS:Intel Mac OS X 10.6 Browser:Firefox-4.0b8 Engine:Gecko-20100101 Bot:false Mobile:false",
    "Mozilla:5.0 Platform:Macintosh OS:Intel Mac OS X 10.6 Localization:en-US Browser:Firefox-3.6.13 Engine:Gecko-20101203 Bot:false Mobile:false",
    "Mozilla:5.0 Platform:X11 OS:Linux x86_64 Browser:Firefox-17.0 Engine:Gecko-20100101 Bot:false Mobile:false",
    "Mozilla:5.0 Platform:Windows OS:Windows XP Localization:en-US Browser:Firefox-2.0.0.14 Engine:Gecko-20080404 Bot:false Mobile:false",
    "Mozilla:5.0 Platform:Macintosh OS:Intel Mac OS X Localization:en Browser:Camino-1.6 Engine:Gecko-20080409 Bot:false Mobile:false",
    "Mozilla:5.0 Platform:X11 OS:Linux i686 Localization:en-US Browser:Iceweasel-2.0 Engine:Gecko-20061024 Bot:false Mobile:false",
    "Mozilla:5.0 Platform:Macintosh OS:Intel Mac OS X 10.6 Localization:en-US Browser:SeaMonkey-2.0 Engine:Gecko-20091017 Bot:false Mobile:false",
    "Mozilla:5.0 Platform:Android OS:Mobile Browser:Firefox-17.0 Engine:Gecko-17.0 Bot:false Mobile:true",

    // Opera
    "Platform:Macintosh OS:Intel Mac OS X Localization:en Browser:Opera-9.27 Engine:Presto Bot:false Mobile:false",
    "Platform:Windows OS:Windows XP Localization:en Browser:Opera-9.27 Engine:Presto Bot:false Mobile:false",
    "Browser:Opera-9.80 Engine:Presto Bot:false Mobile:false",
    "Platform:Windows OS:Windows Vista Localization:en Browser:Opera-9.80 Engine:Presto-2.2.15 Bot:false Mobile:false",
    "Platform:X11 OS:Linux x86_64 Browser:Opera-9.80 Engine:Presto-2.12.388 Bot:false Mobile:false",
    "Platform:Android 4.2.1 OS:Linux Browser:Opera-9.80 Engine:Presto-2.11.355 Bot:false Mobile:true",

    // Other
    "Browser:nil Bot:true Mobile:true",
    "Mozilla:4.0 Bot:false Mobile:false",
    "Mozilla:5.0 Bot:false Mobile:false",
    "Browser:amaya-9.51 Engine:libwww-5.4.0 Bot:true Mobile:true",
    "Browser:Rails Engine:Testing Bot:true Mobile:true",
    "Browser:Python-urllib-2.7 Bot:true Mobile:true",
    "Browser:curl-7.28.1 Bot:true Mobile:true",

    // WebKit
    "Mozilla:5.0 Platform:X11 OS:Linux x86_64 Browser:Chrome-23.0.1271.97 Engine:AppleWebKit-537.11 Bot:false Mobile:false",
    "Mozilla:5.0 Platform:Windows OS:Windows 7 Browser:Chrome-18.0.1025.168 Engine:AppleWebKit-535.19 Bot:false Mobile:false",
    "Mozilla:5.0 Browser:Chrome-8.0.552.215 Engine:AppleWebKit-534.10 Bot:false Mobile:false",
    "Mozilla:5.0 Platform:Macintosh OS:Intel Mac OS X 10_6_5 Localization:en-US Browser:Chrome-8.0.552.231 Engine:AppleWebKit-534.10 Bot:false Mobile:false",
    "Mozilla:5.0 Platform:Macintosh OS:Intel Mac OS X 10_6_3 Localization:en-us Browser:Safari-5.0 Engine:AppleWebKit-533.16 Bot:false Mobile:false",
    "Mozilla:5.0 Platform:Windows OS:Windows XP Localization:en Browser:Safari-4.0dp1 Engine:AppleWebKit-526.9 Bot:false Mobile:false",
    "Mozilla:5.0 Platform:iPhone OS:CPU like Mac OS X Localization:en Browser:Safari-3.0 Engine:AppleWebKit-420.1 Bot:false Mobile:true",
    "Mozilla:5.0 Platform:iPod OS:CPU like Mac OS X Localization:en Browser:Safari-3.0 Engine:AppleWebKit-420.1 Bot:false Mobile:true",
    "Mozilla:5.0 Platform:iPad OS:CPU OS 3_2 like Mac OS X Localization:en-us Browser:Safari-4.0.4 Engine:AppleWebKit-531.21.10 Bot:false Mobile:true",
    "Mozilla:5.0 Platform:webOS OS:Palm Localization:en-US Browser:webOS-1.0 Engine:AppleWebKit-532.2 Bot:false Mobile:true",
    "Mozilla:5.0 Platform:Linux OS:Android 1.5 Localization:de- Browser:Android-3.1.2 Engine:AppleWebKit-528.5+ Bot:false Mobile:true",
    "Mozilla:5.0 Platform:BlackBerry OS:BlackBerry 9800 Localization:en Browser:BlackBerry-6.0.0.141 Engine:AppleWebKit-534.1+ Bot:false Mobile:true",
    "Mozilla:5.0 Platform:BlackBerry OS:BlackBerry Browser:BlackBerry-10.0.9.388 Engine:AppleWebKit-537.3+ Bot:false Mobile:true",
    "Mozilla:5.0 Platform:Symbian OS:SymbianOS/9.4 Browser:Symbian-3.0 Engine:AppleWebKit-525 Bot:false Mobile:true",
    "Mozilla:5.0 Platform:Linux OS:Android 4.2.1 Browser:Chrome-18.0.1025.166 Engine:AppleWebKit-535.19 Bot:false Mobile:true",
}

// Internal: beautify the UserAgent reference into a string so it can be
// tested later on.
//
// ua - a UserAgent reference.
//
// Returns a string that contains the beautified representation.
func beautify(ua *UserAgent) (s string) {
    if len(ua.Mozilla()) > 0 {
        s += "Mozilla:" + ua.Mozilla() + " "
    }
    if len(ua.Platform()) > 0 {
        s += "Platform:" + ua.Platform() + " "
    }
    if len(ua.OS()) > 0 {
        s += "OS:" + ua.OS() + " "
    }
    if len(ua.Localization()) > 0 {
        s += "Localization:" + ua.Localization() + " "
    }
    str1, str2 := ua.Browser()
    if len(str1) > 0 {
        s += "Browser:" + str1
        if len(str2) > 0 {
            s += "-" + str2 + " "
        } else {
            s += " "
        }
    }
    str1, str2 = ua.Engine()
    if len(str1) > 0 {
        s += "Engine:" + str1
        if len(str2) > 0 {
            s += "-" + str2 + " "
        } else {
            s += " "
        }
    }
    s += "Bot:" + fmt.Sprintf("%v", ua.Bot()) + " "
    s += "Mobile:" + fmt.Sprintf("%v", ua.Mobile())
    return s
}

// The test suite.
func TestUserAgent(t *testing.T) {
    for i, tt := range uastrings {
        ua := new(UserAgent)
        ua.Parse(tt.ua)
        got := beautify(ua)
        if expected[i] != got {
            t.Errorf("Test %v => %q, expected %q", tt.name, got, expected[i])
        }
    }
}

// Benchmark: it parses each User-Agent string on the uastrings slice b.N times.
func BenchmarkUserAgent(b *testing.B) {
    for i := 0; i < b.N; i++ {
        b.StopTimer()
        for _, tt := range uastrings {
            ua := new(UserAgent)
            b.StartTimer()
            ua.Parse(tt.ua)
        }
    }
}
