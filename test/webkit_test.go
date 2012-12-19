// Copyright (C) 2012 Miquel Sabaté <mikisabate@gmail.com>
// This file is licensed under the GNU LGPL v3 or later.
// See the LICENSE file.

package user_agent;

import "testing"
import "github.com/bmizerany/assert"


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
