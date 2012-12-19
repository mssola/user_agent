// Copyright (C) 2012 Miquel Sabaté <mikisabate@gmail.com>
// This file is licensed under the GNU GPL v3 or later.
// See the LICENSE file.

package user_agent;

import "testing";
import "github.com/bmizerany/assert";


// Internal: used by all tests to parse a User-Agent string.
//
// ua - the User-Agent string.
//
// Returns a reference to a newly created UserAgent.
func parse(ua string) *UserAgent {
    parser := new(UserAgent);
    parser.Parse(ua);
    return parser;
}

//BEGIN: BOT

const google_bot = "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)";

func TestGoogleBot(t *testing.T) {
    var str1, str2 string;
    ua := parse(google_bot);

    assert.Equal(t, ua.Mobile(), false);
    assert.Equal(t, ua.Bot(), true);
    assert.Equal(t, ua.Mozilla(), "5.0");

    assert.Equal(t, ua.Platform(), "");
    assert.Equal(t, ua.OS(), "");

    str1, str2 = ua.Engine();
    assert.Equal(t, str1, "");
    assert.Equal(t, str2, "");

    str1, str2 = ua.Browser();
    assert.Equal(t, str1, "");
    assert.Equal(t, str2, "");
}

const bing_bot = "Mozilla/5.0 (compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm)";

func TestBingBot(t *testing.T) {
    var str1, str2 string;
    ua := parse(bing_bot);

    assert.Equal(t, ua.Mobile(), false);
    assert.Equal(t, ua.Bot(), true);
    assert.Equal(t, ua.Mozilla(), "5.0");

    assert.Equal(t, ua.Platform(), "");
    assert.Equal(t, ua.OS(), "");

    str1, str2 = ua.Engine();
    assert.Equal(t, str1, "");
    assert.Equal(t, str2, "");

    str1, str2 = ua.Browser();
    assert.Equal(t, str1, "");
    assert.Equal(t, str2, "");
}

//END: BOT

//BEGIN: INTERNET EXPLORER

const ie10 = "Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; Trident/6.0)";

func TestExplorer10(t *testing.T) {
    var str1, str2 string;
    ua := parse(ie10);

    assert.Equal(t, ua.Mobile(), false);
    assert.Equal(t, ua.Bot(), false);
    assert.Equal(t, ua.Mozilla(), "5.0");

    assert.Equal(t, ua.Platform(), "Windows");
    assert.Equal(t, ua.OS(), "Windows 8");

    str1, str2 = ua.Engine();
    assert.Equal(t, str1, "Trident");
    assert.Equal(t, str2, "");

    str1, str2 = ua.Browser();
    assert.Equal(t, str1, "Internet Explorer");
    assert.Equal(t, str2, "10.0");
}

const tablet = "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.2; ARM; Trident/6.0; Touch; .NET4.0E; .NET4.0C; Tablet PC 2.0)";

func TestExplorerTable(t *testing.T) {
    var str1, str2 string;
    ua := parse(tablet);

    assert.Equal(t, ua.Mobile(), false);
    assert.Equal(t, ua.Bot(), false);
    assert.Equal(t, ua.Mozilla(), "4.0");

    assert.Equal(t, ua.Platform(), "Windows");
    assert.Equal(t, ua.OS(), "Windows 8");

    str1, str2 = ua.Engine();
    assert.Equal(t, str1, "Trident");
    assert.Equal(t, str2, "");

    str1, str2 = ua.Browser();
    assert.Equal(t, str1, "Internet Explorer");
    assert.Equal(t, str2, "7.0");
}

const touch = "Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; ARM; Trident/6.0; Touch)";

func TestTouch(t *testing.T) {
    var str1, str2 string;
    ua := parse(touch);

    assert.Equal(t, ua.Mobile(), false);
    assert.Equal(t, ua.Bot(), false);
    assert.Equal(t, ua.Mozilla(), "5.0");

    assert.Equal(t, ua.Platform(), "Windows");
    assert.Equal(t, ua.OS(), "Windows 8");

    str1, str2 = ua.Engine();
    assert.Equal(t, str1, "Trident");
    assert.Equal(t, str2, "");

    str1, str2 = ua.Browser();
    assert.Equal(t, str1, "Internet Explorer");
    assert.Equal(t, str2, "10.0");
}

const phone = "Mozilla/4.0 (compatible; MSIE 7.0; Windows Phone OS 7.0; Trident/3.1; IEMobile/7.0; SAMSUNG; SGH-i917)";

func TestWindowsPhone(t *testing.T) {
    var str1, str2 string;
    ua := parse(phone);

    assert.Equal(t, ua.Mobile(), true);
    assert.Equal(t, ua.Bot(), false);
    assert.Equal(t, ua.Mozilla(), "4.0");

    assert.Equal(t, ua.Platform(), "Windows");
    assert.Equal(t, ua.OS(), "Windows Phone OS 7.0");

    str1, str2 = ua.Engine();
    assert.Equal(t, str1, "Trident");
    assert.Equal(t, str2, "");

    str1, str2 = ua.Browser();
    assert.Equal(t, str1, "Internet Explorer");
    assert.Equal(t, str2, "7.0");
}

//END: INTERNET EXPLORER

//BEGIN: GECKO

const firefox_mac = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.6; rv:2.0b8) Gecko/20100101 Firefox/4.0b8";

func TestFirefoxMac(t *testing.T) {
    var str1, str2 string;
    ua := parse(firefox_mac);

    assert.Equal(t, ua.Mobile(), false);
    assert.Equal(t, ua.Bot(), false);
    assert.Equal(t, ua.Mozilla(), "5.0");

    assert.Equal(t, ua.Platform(), "Macintosh");
    assert.Equal(t, ua.OS(), "Intel Mac OS X 10.6");

    str1, str2 = ua.Engine();
    assert.Equal(t, str1, "Gecko");
    assert.Equal(t, str2, "20100101");

    str1, str2 = ua.Browser();
    assert.Equal(t, str1, "Firefox");
    assert.Equal(t, str2, "4.0b8");
}

const firefox_mac_loc = "Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.6; en-US; rv:1.9.2.13) Gecko/20101203 Firefox/3.6.13";

func TestFirefoxMacLoc(t *testing.T) {
    var str1, str2 string;
    ua := parse(firefox_mac_loc);

    assert.Equal(t, ua.Mobile(), false);
    assert.Equal(t, ua.Bot(), false);
    assert.Equal(t, ua.Mozilla(), "5.0");
    assert.Equal(t, ua.Localization(), "en-US");

    assert.Equal(t, ua.Platform(), "Macintosh");
    assert.Equal(t, ua.OS(), "Intel Mac OS X 10.6");

    str1, str2 = ua.Engine();
    assert.Equal(t, str1, "Gecko");
    assert.Equal(t, str2, "20101203");

    str1, str2 = ua.Browser();
    assert.Equal(t, str1, "Firefox");
    assert.Equal(t, str2, "3.6.13");
}

const firefox_linux = "Mozilla/5.0 (X11; Linux x86_64; rv:17.0) Gecko/20100101 Firefox/17.0";

func TestFirefoxLinux(t *testing.T) {
    var str1, str2 string;
    ua := parse(firefox_linux);

    assert.Equal(t, ua.Mobile(), false);
    assert.Equal(t, ua.Bot(), false);
    assert.Equal(t, ua.Mozilla(), "5.0");
    assert.Equal(t, ua.Localization(), "");

    assert.Equal(t, ua.Platform(), "X11");
    assert.Equal(t, ua.OS(), "Linux x86_64");

    str1, str2 = ua.Engine();
    assert.Equal(t, str1, "Gecko");
    assert.Equal(t, str2, "20100101");

    str1, str2 = ua.Browser();
    assert.Equal(t, str1, "Firefox");
    assert.Equal(t, str2, "17.0");
}

const firefox_win = "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US; rv:1.8.1.14) Gecko/20080404 Firefox/2.0.0.14";

func TestFirefoxWin(t *testing.T) {
    var str1, str2 string;
    ua := parse(firefox_win);

    assert.Equal(t, ua.Mobile(), false);
    assert.Equal(t, ua.Bot(), false);
    assert.Equal(t, ua.Mozilla(), "5.0");
    assert.Equal(t, ua.Localization(), "en-US");

    assert.Equal(t, ua.Platform(), "Windows");
    assert.Equal(t, ua.OS(), "Windows XP");

    str1, str2 = ua.Engine();
    assert.Equal(t, str1, "Gecko");
    assert.Equal(t, str2, "20080404");

    str1, str2 = ua.Browser();
    assert.Equal(t, str1, "Firefox");
    assert.Equal(t, str2, "2.0.0.14");
}

const camino_mac = "Mozilla/5.0 (Macintosh; U; Intel Mac OS X; en; rv:1.8.1.14) Gecko/20080409 Camino/1.6 (like Firefox/2.0.0.14)";

func TestCaminoMac(t *testing.T) {
    var str1, str2 string;
    ua := parse(camino_mac);

    assert.Equal(t, ua.Mobile(), false);
    assert.Equal(t, ua.Bot(), false);
    assert.Equal(t, ua.Mozilla(), "5.0");
    assert.Equal(t, ua.Localization(), "en");

    assert.Equal(t, ua.Platform(), "Macintosh");
    assert.Equal(t, ua.OS(), "Intel Mac OS X");

    str1, str2 = ua.Engine();
    assert.Equal(t, str1, "Gecko");
    assert.Equal(t, str2, "20080409");

    str1, str2 = ua.Browser();
    assert.Equal(t, str1, "Camino");
    assert.Equal(t, str2, "1.6");
}

const iceweasel_debian = "Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.8.1) Gecko/20061024 Iceweasel/2.0 (Debian-2.0+dfsg-1)";

func TestIceweasel(t *testing.T) {
    var str1, str2 string;
    ua := parse(iceweasel_debian);

    assert.Equal(t, ua.Mobile(), false);
    assert.Equal(t, ua.Bot(), false);
    assert.Equal(t, ua.Mozilla(), "5.0");
    assert.Equal(t, ua.Localization(), "en-US");

    assert.Equal(t, ua.Platform(), "X11");
    assert.Equal(t, ua.OS(), "Linux i686");

    str1, str2 = ua.Engine();
    assert.Equal(t, str1, "Gecko");
    assert.Equal(t, str2, "20061024");

    str1, str2 = ua.Browser();
    assert.Equal(t, str1, "Iceweasel");
    assert.Equal(t, str2, "2.0");
}

const seamonkey = "Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.6; en-US; rv:1.9.1.4) Gecko/20091017 SeaMonkey/2.0";

func TestSeaMonkey(t *testing.T) {
    var str1, str2 string;
    ua := parse(seamonkey);

    assert.Equal(t, ua.Mobile(), false);
    assert.Equal(t, ua.Bot(), false);
    assert.Equal(t, ua.Mozilla(), "5.0");
    assert.Equal(t, ua.Localization(), "en-US");

    assert.Equal(t, ua.Platform(), "Macintosh");
    assert.Equal(t, ua.OS(), "Intel Mac OS X 10.6");

    str1, str2 = ua.Engine();
    assert.Equal(t, str1, "Gecko");
    assert.Equal(t, str2, "20091017");

    str1, str2 = ua.Browser();
    assert.Equal(t, str1, "SeaMonkey");
    assert.Equal(t, str2, "2.0");
}

const android_firefox = "Mozilla/5.0 (Android; Mobile; rv:17.0) Gecko/17.0 Firefox/17.0";

func TestFirefoxAndroid(t *testing.T) {
    var str1, str2 string;
    ua := parse(android_firefox);

    assert.Equal(t, ua.Mobile(), true);
    assert.Equal(t, ua.Bot(), false);
    assert.Equal(t, ua.Mozilla(), "5.0");
    assert.Equal(t, ua.Localization(), "");

    assert.Equal(t, ua.Platform(), "Android");
    assert.Equal(t, ua.OS(), "Mobile");

    str1, str2 = ua.Engine();
    assert.Equal(t, str1, "Gecko");
    assert.Equal(t, str2, "17.0");

    str1, str2 = ua.Browser();
    assert.Equal(t, str1, "Firefox");
    assert.Equal(t, str2, "17.0");
}

//END: GECKO

//START: OPERA

const opera_mac = "Opera/9.27 (Macintosh; Intel Mac OS X; U; en)";

func TestOperaMac(t *testing.T) {
    var str1, str2 string;
    ua := parse(opera_mac);

    assert.Equal(t, ua.Mobile(), false);
    assert.Equal(t, ua.Bot(), false);
    assert.Equal(t, ua.Mozilla(), "");
    assert.Equal(t, ua.Localization(), "en");

    assert.Equal(t, ua.Platform(), "Macintosh");
    assert.Equal(t, ua.OS(), "Intel Mac OS X");

    str1, str2 = ua.Engine();
    assert.Equal(t, str1, "Presto");
    assert.Equal(t, str2, "");

    str1, str2 = ua.Browser();
    assert.Equal(t, str1, "Opera");
    assert.Equal(t, str2, "9.27");
}

const opera_win = "Opera/9.27 (Windows NT 5.1; U; en)";

func TestOperaWin(t *testing.T) {
    var str1, str2 string;
    ua := parse(opera_win);

    assert.Equal(t, ua.Mobile(), false);
    assert.Equal(t, ua.Bot(), false);
    assert.Equal(t, ua.Mozilla(), "");
    assert.Equal(t, ua.Localization(), "en");

    assert.Equal(t, ua.Platform(), "Windows");
    assert.Equal(t, ua.OS(), "Windows XP");

    str1, str2 = ua.Engine();
    assert.Equal(t, str1, "Presto");
    assert.Equal(t, str2, "");

    str1, str2 = ua.Browser();
    assert.Equal(t, str1, "Opera");
    assert.Equal(t, str2, "9.27");
}

const opera_minimal = "Opera/9.80";

func TestOperaMinimal(t *testing.T) {
    var str1, str2 string;
    ua := parse(opera_minimal);

    assert.Equal(t, ua.Mobile(), false);
    assert.Equal(t, ua.Bot(), false);
    assert.Equal(t, ua.Mozilla(), "");
    assert.Equal(t, ua.Localization(), "");

    assert.Equal(t, ua.Platform(), "");
    assert.Equal(t, ua.OS(), "");

    str1, str2 = ua.Engine();
    assert.Equal(t, str1, "Presto");
    assert.Equal(t, str2, "");

    str1, str2 = ua.Browser();
    assert.Equal(t, str1, "Opera");
    assert.Equal(t, str2, "9.80");
}

const opera_full = "Opera/9.80 (Windows NT 6.0; U; en) Presto/2.2.15 Version/10.10";

func TestOperaFull(t *testing.T) {
    var str1, str2 string;
    ua := parse(opera_full);

    assert.Equal(t, ua.Mobile(), false);
    assert.Equal(t, ua.Bot(), false);
    assert.Equal(t, ua.Mozilla(), "");
    assert.Equal(t, ua.Localization(), "en");

    assert.Equal(t, ua.Platform(), "Windows");
    assert.Equal(t, ua.OS(), "Windows Vista");

    str1, str2 = ua.Engine();
    assert.Equal(t, str1, "Presto");
    assert.Equal(t, str2, "2.2.15");

    str1, str2 = ua.Browser();
    assert.Equal(t, str1, "Opera");
    assert.Equal(t, str2, "9.80");
}

const opera_linux = "Opera/9.80 (X11; Linux x86_64) Presto/2.12.388 Version/12.10";

func TestOperaLinux(t *testing.T) {
    var str1, str2 string;
    ua := parse(opera_linux);

    assert.Equal(t, ua.Mobile(), false);
    assert.Equal(t, ua.Bot(), false);
    assert.Equal(t, ua.Mozilla(), "");
    assert.Equal(t, ua.Localization(), "");

    assert.Equal(t, ua.Platform(), "X11");
    assert.Equal(t, ua.OS(), "Linux x86_64");

    str1, str2 = ua.Engine();
    assert.Equal(t, str1, "Presto");
    assert.Equal(t, str2, "2.12.388");

    str1, str2 = ua.Browser();
    assert.Equal(t, str1, "Opera");
    assert.Equal(t, str2, "9.80");
}

const opera_android = "Opera/9.80 (Android 4.2.1; Linux; Opera Mobi/ADR-1212030829) Presto/2.11.355 Version/12.10";

func TestOperaAndroid(t *testing.T) {
    var str1, str2 string;
    ua := parse(opera_android);

    assert.Equal(t, ua.Mobile(), true);
    assert.Equal(t, ua.Bot(), false);
    assert.Equal(t, ua.Mozilla(), "");
    assert.Equal(t, ua.Localization(), "");

    assert.Equal(t, ua.Platform(), "Android 4.2.1");
    assert.Equal(t, ua.OS(), "Linux");

    str1, str2 = ua.Engine();
    assert.Equal(t, str1, "Presto");
    assert.Equal(t, str2, "2.11.355");

    str1, str2 = ua.Browser();
    assert.Equal(t, str1, "Opera");
    assert.Equal(t, str2, "9.80");
}

//END: OPERA

//START: OTHER

const null = "nil";

func TestNil(t *testing.T) {
    var str1, str2 string;
    ua := parse(null);

    assert.Equal(t, ua.Mobile(), true);
    assert.Equal(t, ua.Bot(), true);
    assert.Equal(t, ua.Mozilla(), "");

    assert.Equal(t, ua.Platform(), "");
    assert.Equal(t, ua.OS(), "");

    str1, str2 = ua.Engine();
    assert.Equal(t, str1, "");
    assert.Equal(t, str2, "");

    str1, str2 = ua.Browser();
    assert.Equal(t, str1, "");
    assert.Equal(t, str2, "");
}

const compatible = "Mozilla/4.0 (compatible)";

func TestCompatible(t *testing.T) {
    var str1, str2 string;
    ua := parse(compatible);

    assert.Equal(t, ua.Mobile(), false);
    assert.Equal(t, ua.Bot(), false);
    assert.Equal(t, ua.Mozilla(), "4.0");

    assert.Equal(t, ua.Platform(), "");
    assert.Equal(t, ua.OS(), "");

    str1, str2 = ua.Engine();
    assert.Equal(t, str1, "");
    assert.Equal(t, str2, "");

    str1, str2 = ua.Browser();
    assert.Equal(t, str1, "");
    assert.Equal(t, str2, "");
}

const mozilla = "Mozilla/5.0";

func TestMozilla(t *testing.T) {
    var str1, str2 string;
    ua := parse(mozilla);

    assert.Equal(t, ua.Mobile(), false);
    assert.Equal(t, ua.Bot(), false);
    assert.Equal(t, ua.Mozilla(), "5.0");

    assert.Equal(t, ua.Platform(), "");
    assert.Equal(t, ua.OS(), "");

    str1, str2 = ua.Engine();
    assert.Equal(t, str1, "");
    assert.Equal(t, str2, "");

    str1, str2 = ua.Browser();
    assert.Equal(t, str1, "");
    assert.Equal(t, str2, "");
}

const amaya = "amaya/9.51 libwww/5.4.0";

func TestAmaya(t *testing.T) {
    var str1, str2 string;
    ua := parse(amaya);

    assert.Equal(t, ua.Mobile(), true);
    assert.Equal(t, ua.Bot(), true);
    assert.Equal(t, ua.Mozilla(), "");

    assert.Equal(t, ua.Platform(), "");
    assert.Equal(t, ua.OS(), "");

    str1, str2 = ua.Engine();
    assert.Equal(t, str1, "libwww");
    assert.Equal(t, str2, "5.4.0");

    str1, str2 = ua.Browser();
    assert.Equal(t, str1, "");
    assert.Equal(t, str2, "");
}

const rails = "Rails Testing";

func TestRails(t *testing.T) {
    var str1, str2 string;
    ua := parse(rails);

    assert.Equal(t, ua.Mobile(), true);
    assert.Equal(t, ua.Bot(), true);
    assert.Equal(t, ua.Mozilla(), "");

    assert.Equal(t, ua.Platform(), "");
    assert.Equal(t, ua.OS(), "");

    str1, str2 = ua.Engine();
    assert.Equal(t, str1, "Testing"); // Weird, but doesn't hurt :)
    assert.Equal(t, str2, "");

    str1, str2 = ua.Browser();
    assert.Equal(t, str1, "");
    assert.Equal(t, str2, "");
}

const python = "Python-urllib/2.7";

func TestPython(t *testing.T) {
    var str1, str2 string;
    ua := parse(python);

    assert.Equal(t, ua.Mobile(), true);
    assert.Equal(t, ua.Bot(), true);
    assert.Equal(t, ua.Mozilla(), "");

    assert.Equal(t, ua.Platform(), "");
    assert.Equal(t, ua.OS(), "");

    str1, str2 = ua.Engine();
    assert.Equal(t, str1, "");
    assert.Equal(t, str2, "");

    str1, str2 = ua.Browser();
    assert.Equal(t, str1, "");
    assert.Equal(t, str2, "");
}

//END: OTHER

//START: WEBKIT

const chrome_linux = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.97 Safari/537.11";

func TestChromeLinux(t *testing.T) {
    var str1, str2 string;
    ua := parse(chrome_linux);

    assert.Equal(t, ua.Mobile(), false);
    assert.Equal(t, ua.Bot(), false);
    assert.Equal(t, ua.Mozilla(), "5.0");

    assert.Equal(t, ua.Platform(), "X11");
    assert.Equal(t, ua.OS(), "Linux x86_64");

    str1, str2 = ua.Engine();
    assert.Equal(t, str1, "AppleWebKit");
    assert.Equal(t, str2, "537.11");

    str1, str2 = ua.Browser();
    assert.Equal(t, str1, "Chrome");
    assert.Equal(t, str2, "23.0.1271.97");
}

const chrome_win7 = "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.168 Safari/535.19";

func TestChromeWin7(t *testing.T) {
    var str1, str2 string;
    ua := parse(chrome_win7);

    assert.Equal(t, ua.Mobile(), false);
    assert.Equal(t, ua.Bot(), false);
    assert.Equal(t, ua.Mozilla(), "5.0");

    assert.Equal(t, ua.Platform(), "Windows");
    assert.Equal(t, ua.OS(), "Windows 7");

    str1, str2 = ua.Engine();
    assert.Equal(t, str1, "AppleWebKit");
    assert.Equal(t, str2, "535.19");

    str1, str2 = ua.Browser();
    assert.Equal(t, str1, "Chrome");
    assert.Equal(t, str2, "18.0.1025.168");
}

const chrome_minimal = "Mozilla/5.0 AppleWebKit/534.10 Chrome/8.0.552.215 Safari/534.10";

func TestChromeMinimal(t *testing.T) {
    var str1, str2 string;
    ua := parse(chrome_minimal);

    assert.Equal(t, ua.Mobile(), false);
    assert.Equal(t, ua.Bot(), false);
    assert.Equal(t, ua.Mozilla(), "5.0");

    assert.Equal(t, ua.Platform(), "");
    assert.Equal(t, ua.OS(), "");

    str1, str2 = ua.Engine();
    assert.Equal(t, str1, "AppleWebKit");
    assert.Equal(t, str2, "534.10");

    str1, str2 = ua.Browser();
    assert.Equal(t, str1, "Chrome");
    assert.Equal(t, str2, "8.0.552.215");
}

const chrome_mac = "Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_6_5; en-US) AppleWebKit/534.10 (KHTML, like Gecko) Chrome/8.0.552.231 Safari/534.10";

func TestChromeMac(t *testing.T) {
    var str1, str2 string;
    ua := parse(chrome_mac);

    assert.Equal(t, ua.Mobile(), false);
    assert.Equal(t, ua.Bot(), false);
    assert.Equal(t, ua.Mozilla(), "5.0");

    assert.Equal(t, ua.Platform(), "Macintosh");
    assert.Equal(t, ua.OS(), "Intel Mac OS X 10_6_5");

    str1, str2 = ua.Engine();
    assert.Equal(t, str1, "AppleWebKit");
    assert.Equal(t, str2, "534.10");

    str1, str2 = ua.Browser();
    assert.Equal(t, str1, "Chrome");
    assert.Equal(t, str2, "8.0.552.231");
}

const safari_mac = "Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_6_3; en-us) AppleWebKit/533.16 (KHTML, like Gecko) Version/5.0 Safari/533.16";

func TestSafariMac(t *testing.T) {
    var str1, str2 string;
    ua := parse(safari_mac);

    assert.Equal(t, ua.Mobile(), false);
    assert.Equal(t, ua.Bot(), false);
    assert.Equal(t, ua.Mozilla(), "5.0");
    assert.Equal(t, ua.Localization(), "en-us");

    assert.Equal(t, ua.Platform(), "Macintosh");
    assert.Equal(t, ua.OS(), "Intel Mac OS X 10_6_3");

    str1, str2 = ua.Engine();
    assert.Equal(t, str1, "AppleWebKit");
    assert.Equal(t, str2, "533.16");

    str1, str2 = ua.Browser();
    assert.Equal(t, str1, "Safari");
    assert.Equal(t, str2, "5.0");
}

const safari_windows = "Mozilla/5.0 (Windows; U; Windows NT 5.1; en) AppleWebKit/526.9 (KHTML, like Gecko) Version/4.0dp1 Safari/526.8";

func TestSafariWindows(t *testing.T) {
    var str1, str2 string;
    ua := parse(safari_windows);

    assert.Equal(t, ua.Mobile(), false);
    assert.Equal(t, ua.Bot(), false);
    assert.Equal(t, ua.Mozilla(), "5.0");
    assert.Equal(t, ua.Localization(), "en");

    assert.Equal(t, ua.Platform(), "Windows");
    assert.Equal(t, ua.OS(), "Windows XP");

    str1, str2 = ua.Engine();
    assert.Equal(t, str1, "AppleWebKit");
    assert.Equal(t, str2, "526.9");

    str1, str2 = ua.Browser();
    assert.Equal(t, str1, "Safari");
    assert.Equal(t, str2, "4.0dp1");
}

const iPhone = "Mozilla/5.0 (iPhone; U; CPU like Mac OS X; en) AppleWebKit/420.1 (KHTML, like Gecko) Version/3.0 Mobile/4A102 Safari/419";

func TestIPhone(t *testing.T) {
    var str1, str2 string;
    ua := parse(iPhone);

    assert.Equal(t, ua.Mobile(), true);
    assert.Equal(t, ua.Bot(), false);
    assert.Equal(t, ua.Mozilla(), "5.0");
    assert.Equal(t, ua.Localization(), "en");

    assert.Equal(t, ua.Platform(), "iPhone");
    assert.Equal(t, ua.OS(), "CPU like Mac OS X");

    str1, str2 = ua.Engine();
    assert.Equal(t, str1, "AppleWebKit");
    assert.Equal(t, str2, "420.1");

    str1, str2 = ua.Browser();
    assert.Equal(t, str1, "Safari");
    assert.Equal(t, str2, "3.0");
}

const iPod = "Mozilla/5.0 (iPod; U; CPU like Mac OS X; en) AppleWebKit/420.1 (KHTML, like Gecko) Version/3.0 Mobile/4A102 Safari/419";

func TestIPod(t *testing.T) {
    var str1, str2 string;
    ua := parse(iPod);

    assert.Equal(t, ua.Mobile(), true);
    assert.Equal(t, ua.Bot(), false);
    assert.Equal(t, ua.Mozilla(), "5.0");
    assert.Equal(t, ua.Localization(), "en");

    assert.Equal(t, ua.Platform(), "iPod");
    assert.Equal(t, ua.OS(), "CPU like Mac OS X");

    str1, str2 = ua.Engine();
    assert.Equal(t, str1, "AppleWebKit");
    assert.Equal(t, str2, "420.1");

    str1, str2 = ua.Browser();
    assert.Equal(t, str1, "Safari");
    assert.Equal(t, str2, "3.0");
}

const iPad = "Mozilla/5.0 (iPad; U; CPU OS 3_2 like Mac OS X; en-us) AppleWebKit/531.21.10 (KHTML, like Gecko) Version/4.0.4 Mobile/7B367 Safari/531.21.10";

func TestIPad(t *testing.T) {
    var str1, str2 string;
    ua := parse(iPad);

    assert.Equal(t, ua.Mobile(), true);
    assert.Equal(t, ua.Bot(), false);
    assert.Equal(t, ua.Mozilla(), "5.0");
    assert.Equal(t, ua.Localization(), "en-us");

    assert.Equal(t, ua.Platform(), "iPad");
    assert.Equal(t, ua.OS(), "CPU OS 3_2 like Mac OS X");

    str1, str2 = ua.Engine();
    assert.Equal(t, str1, "AppleWebKit");
    assert.Equal(t, str2, "531.21.10");

    str1, str2 = ua.Browser();
    assert.Equal(t, str1, "Safari");
    assert.Equal(t, str2, "4.0.4");
}

const webOS = "Mozilla/5.0 (webOS/1.4.0; U; en-US) AppleWebKit/532.2 (KHTML, like Gecko) Version/1.0 Safari/532.2 Pre/1.1";

func TestWebOS(t *testing.T) {
    var str1, str2 string;
    ua := parse(webOS);

    assert.Equal(t, ua.Mobile(), true);
    assert.Equal(t, ua.Bot(), false);
    assert.Equal(t, ua.Mozilla(), "5.0");
    assert.Equal(t, ua.Localization(), "en-US");

    assert.Equal(t, ua.Platform(), "webOS");
    assert.Equal(t, ua.OS(), "Palm");

    str1, str2 = ua.Engine();
    assert.Equal(t, str1, "AppleWebKit");
    assert.Equal(t, str2, "532.2");

    str1, str2 = ua.Browser();
    assert.Equal(t, str1, "webOS");
    assert.Equal(t, str2, "1.0");
}

const android = "Mozilla/5.0 (Linux; U; Android 1.5; de-; HTC Magic Build/PLAT-RC33) AppleWebKit/528.5+ (KHTML, like Gecko) Version/3.1.2 Mobile Safari/525.20.1";

func TestAndroid(t *testing.T) {
    var str1, str2 string;
    ua := parse(android);

    assert.Equal(t, ua.Mobile(), true);
    assert.Equal(t, ua.Bot(), false);
    assert.Equal(t, ua.Mozilla(), "5.0");
    assert.Equal(t, ua.Localization(), "de-");

    assert.Equal(t, ua.Platform(), "Linux");
    assert.Equal(t, ua.OS(), "Android 1.5");

    str1, str2 = ua.Engine();
    assert.Equal(t, str1, "AppleWebKit");
    assert.Equal(t, str2, "528.5+");

    str1, str2 = ua.Browser();
    assert.Equal(t, str1, "Android");
    assert.Equal(t, str2, "3.1.2");
}

const blackberry = "Mozilla/5.0 (BlackBerry; U; BlackBerry 9800; en) AppleWebKit/534.1+ (KHTML, Like Gecko) Version/6.0.0.141 Mobile Safari/534.1+";

func TestBlackBerry(t *testing.T) {
    var str1, str2 string;
    ua := parse(blackberry);

    assert.Equal(t, ua.Mobile(), true);
    assert.Equal(t, ua.Bot(), false);
    assert.Equal(t, ua.Mozilla(), "5.0");
    assert.Equal(t, ua.Localization(), "en");

    assert.Equal(t, ua.Platform(), "BlackBerry");
    assert.Equal(t, ua.OS(), "BlackBerry 9800");

    str1, str2 = ua.Engine();
    assert.Equal(t, str1, "AppleWebKit");
    assert.Equal(t, str2, "534.1+");

    str1, str2 = ua.Browser();
    assert.Equal(t, str1, "BlackBerry");
    assert.Equal(t, str2, "6.0.0.141");
}

const bb10 = "Mozilla/5.0 (BB10; Touch) AppleWebKit/537.3+ (KHTML, like Gecko) Version/10.0.9.388 Mobile Safari/537.3+";

func TestBlackBerry10(t *testing.T) {
    var str1, str2 string;
    ua := parse(bb10);

    assert.Equal(t, ua.Mobile(), true);
    assert.Equal(t, ua.Bot(), false);
    assert.Equal(t, ua.Mozilla(), "5.0");
    assert.Equal(t, ua.Localization(), "");

    assert.Equal(t, ua.Platform(), "BlackBerry");
    assert.Equal(t, ua.OS(), "BlackBerry");

    str1, str2 = ua.Engine();
    assert.Equal(t, str1, "AppleWebKit");
    assert.Equal(t, str2, "537.3+");

    str1, str2 = ua.Browser();
    assert.Equal(t, str1, "BlackBerry");
    assert.Equal(t, str2, "10.0.9.388");
}

const ericsson = "Mozilla/5.0 (SymbianOS/9.4; U; Series60/5.0 Profile/MIDP-2.1 Configuration/CLDC-1.1) AppleWebKit/525 (KHTML, like Gecko) Version/3.0 Safari/525";

func TestEricsson(t *testing.T) {
    var str1, str2 string;
    ua := parse(ericsson);

    assert.Equal(t, ua.Mobile(), true);
    assert.Equal(t, ua.Bot(), false);
    assert.Equal(t, ua.Mozilla(), "5.0");
    assert.Equal(t, ua.Localization(), "");

    assert.Equal(t, ua.Platform(), "Symbian");
    assert.Equal(t, ua.OS(), "SymbianOS/9.4");

    str1, str2 = ua.Engine();
    assert.Equal(t, str1, "AppleWebKit");
    assert.Equal(t, str2, "525");

    str1, str2 = ua.Browser();
    assert.Equal(t, str1, "Symbian");
    assert.Equal(t, str2, "3.0");
}

const chrome_android = "Mozilla/5.0 (Linux; Android 4.2.1; Galaxy Nexus Build/JOP40D) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166 Mobile Safari/535.19";

func TestChromeAndroid(t *testing.T) {
    var str1, str2 string;
    ua := parse(chrome_android);

    assert.Equal(t, ua.Mobile(), true);
    assert.Equal(t, ua.Bot(), false);
    assert.Equal(t, ua.Mozilla(), "5.0");
    assert.Equal(t, ua.Localization(), "");

    assert.Equal(t, ua.Platform(), "Linux");
    assert.Equal(t, ua.OS(), "Android 4.2.1");

    str1, str2 = ua.Engine();
    assert.Equal(t, str1, "AppleWebKit");
    assert.Equal(t, str2, "535.19");

    str1, str2 = ua.Browser();
    assert.Equal(t, str1, "Chrome");
    assert.Equal(t, str2, "18.0.1025.166");
}

//END: WEBKIT
