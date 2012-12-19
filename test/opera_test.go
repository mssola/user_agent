// Copyright (C) 2012 Miquel Sabaté <mikisabate@gmail.com>
// This file is licensed under the GNU LGPL v3 or later.
// See the LICENSE file.

package user_agent;

import "testing";
import "github.com/bmizerany/assert";


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
