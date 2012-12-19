// Copyright (C) 2012 Miquel Sabaté <mikisabate@gmail.com>
// This file is licensed under the GNU LGPL v3 or later.
// See the LICENSE file.

package user_agent;

import "testing";
import "github.com/bmizerany/assert";


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
