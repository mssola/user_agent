// Copyright (C) 2012-2013 Miquel Sabaté Solà
// This file is licensed under the MIT license.
// See the LICENSE file.


package user_agent

import "strings"


// Internal: normalize the name of the operating system. By now, this just
// affects to Windows.
//
// name - The original name of the Operating System.
//
// Returns a string containing the normalized name for the Operating System.
func normalizeOS(name string) string {
    windows := map[string]string{
        "Windows NT 6.2":  "Windows 8",
        "Windows NT 6.1":  "Windows 7",
        "Windows NT 6.0":  "Windows Vista",
        "Windows NT 5.2":  "Windows XP x64 Edition",
        "Windows NT 5.1":  "Windows XP",
        "Windows NT 5.01": "Windows 2000, Service Pack 1 (SP1)",
        "Windows NT 5.0":  "Windows 2000",
        "Windows NT 4.0":  "Windows NT 4.0",
        "Windows 98":      "Windows 98",
        "Windows 95":      "Windows 95",
        "Windows CE":      "Windows CE",
    }
    if val, ok := windows[name]; ok {
        return val
    }
    return name
}

// Internal: guess the OS, the localization and if this is a mobile device
// for a Webkit-powered browser.
//
// p       - a reference to the current UserAgent.
// comment - a slice of strings containing the comment.
func webkit(p *UserAgent, comment []string) {
    if p.platform == "webOS" {
        p.browser.name = p.platform
        p.os = "Palm"
        p.localization = comment[2]
        p.mobile = true
    } else if p.platform == "Symbian" {
        p.mobile = true
        p.browser.name = p.platform
        p.os = comment[0]
        if len(comment) > 3 {
            p.localization = comment[3]
        }
    } else if p.platform == "Linux" {
        p.mobile = true
        if p.browser.name == "Safari" {
            p.browser.name = "Android"
        }
        if comment[1] == "U" {
            if len(comment) > 2 {
                p.os = comment[2]
            } else {
                p.mobile = false
                p.os = comment[0]
            }
        } else {
            p.os = comment[1]
        }
        if len(comment) > 3 {
            p.localization = comment[3]
        }
    } else if len(comment) > 0 {
        if len(comment) > 3 {
            p.localization = comment[3]
        }
        if strings.HasPrefix(comment[0], "Windows NT") {
            p.os = normalizeOS(comment[0])
        } else if len(comment) < 2 {
            p.localization = comment[0]
        } else if len(comment) < 3 {
            p.os = normalizeOS(comment[1])
        } else {
            p.os = normalizeOS(comment[2])
        }
        if p.platform == "BlackBerry" {
            p.browser.name = p.platform
            if p.os == "Touch" {
                p.os = p.platform
            }
        }
    }
}

// Internal: guess the OS, the localization and if this is a mobile device
// for a Gecko-powered browser.
//
// p       - a reference to the current UserAgent.
// comment - a slice of strings containing the comment.
func gecko(p *UserAgent, comment []string) {
    if p.platform == "Android" {
        p.mobile = true
    }
    if len(comment) > 1 {
        if comment[1] == "U" {
            p.os = normalizeOS(comment[2])
        } else {
            p.os = normalizeOS(comment[1])
        }
        if len(comment) > 3 {
            p.localization = comment[3]
        }
    }
}

// Internal: given the comment of the first section of the UserAgent string,
// get the platform.
//
// comment - a slice of strings containing the comment itself.
//
// Returns a string that is the platform.
func getPlatform(comment []string) string {
    if len(comment) > 0 {
        if comment[0] != "compatible" {
            if strings.HasPrefix(comment[0], "Windows") {
                return "Windows"
            } else if strings.HasPrefix(comment[0], "Symbian") {
                return "Symbian"
            } else if strings.HasPrefix(comment[0], "webOS") {
                return "webOS"
            } else if comment[0] == "BB10" {
                return "BlackBerry"
            }
            return comment[0]
        }
    }
    return ""
}

// Internal: detect some properties of the OS.
//
// section - an UASection containing the first section of the User-Agent.
func (p *UserAgent) detectOS(section UASection) {
    if section.name == "Mozilla" {
        p.platform = getPlatform(section.comment)
        switch p.browser.engine {
        case "Gecko":
            gecko(p, section.comment)
        case "AppleWebKit":
            webkit(p, section.comment)
        case "Trident":
            p.platform = "Windows"
            p.os = normalizeOS(section.comment[2])
            for _, v := range section.comment {
                if strings.HasPrefix(v, "IEMobile") {
                    p.mobile = true
                    return
                }
            }
        }
    } else if section.name == "Opera" {
        if len(section.comment) > 0 {
            if strings.HasPrefix(section.comment[0], "Windows") {
                p.platform = "Windows"
                p.os = normalizeOS(section.comment[0])
                if len(section.comment) > 2 {
                    p.localization = section.comment[2]
                }
            } else {
                if strings.HasPrefix(section.comment[0], "Android") {
                    p.mobile = true
                }
                p.platform = section.comment[0]
                p.os = section.comment[1]
                if len(section.comment) > 3 {
                    p.localization = section.comment[3]
                }
            }
        }
    } else {
        // At this point, we assume that this is a bot.
        p.bot = true
        p.mobile = true
        p.mozilla = ""
        p.browser.name = section.name
        p.browser.version = section.version
    }
}

// Returns a string containing the platform..
func (p *UserAgent) Platform() string {
    return p.platform
}

// Returns a string containing the name of the Operating System.
func (p *UserAgent) OS() string {
    return p.os
}

// Returns a string containing the localization.
func (p *UserAgent) Localization() string {
    return p.localization
}
