// Copyright (C) 2012 Miquel Sabaté <mikisabate@gmail.com>
// This file is licensed under the GNU LGPL v3 or later.
// See the LICENSE file.

package user_agent;


import "testing";
import "github.com/bmizerany/assert";


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
