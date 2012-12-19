// Copyright (C) 2012 Miquel Sabaté <mikisabate@gmail.com>
// This file is licensed under the GNU LGPL v3 or later.
// See the LICENSE file.

package user_agent;


import "testing";
import "github.com/bmizerany/assert";


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
