// Copyright (C) 2012 Miquel Sabaté <mikisabate@gmail.com>
// This file is licensed under the GNU LGPL v3 or later.
// See the LICENSE file.

package user_agent;

import "testing"
import "github.com/bmizerany/assert"


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
