// Copyright (C) 2012-2013 Miquel Sabaté Solà
// This file is licensed under the MIT license.
// See the LICENSE file.


package user_agent

import "strings"


// A struct containing all the information that we might be
// interested from the browser.
type Browser struct {
    engine         string
    engine_version string
    name           string
    version        string
}

// Internal: extract all the information that we can get from the User-Agent
// string about the browser and update the receiver with this information.
//
// sections - a slice of UASection's containing all the sections from the
//            User-Agent string after being parsed.
func (p *UserAgent) detectBrowser(sections []UASection) {
    slen := len(sections)

    if sections[0].name == "Opera" {
        p.mozilla = ""
        p.browser.name = "Opera"
        p.browser.version = sections[0].version
        p.browser.engine = "Presto"
        if slen > 1 {
            p.browser.engine_version = sections[1].version
        }
    } else if slen > 1 {
        engine := sections[1]
        p.browser.engine = engine.name
        p.browser.engine_version = engine.version
        if slen > 2 {
            p.browser.version = sections[2].version
            if engine.name == "AppleWebKit" {
                if sections[slen - 1].name == "OPR" {
                    p.browser.name = "Opera"
                    p.browser.version = sections[slen - 1].version
                } else if sections[2].name == "Chrome" {
                    p.browser.name = "Chrome"
                } else {
                    p.browser.name = "Safari"
                }
            } else if engine.name == "Gecko" {
                p.browser.name = sections[2].name
            }
        }
    } else if slen == 1 && len(sections[0].comment) > 1 {
        comment := sections[0].comment
        if comment[0] == "compatible" && strings.HasPrefix(comment[1], "MSIE") {
            p.browser.engine = "Trident"
            p.browser.version = strings.TrimSpace(comment[1][4:])
            p.browser.name = "Internet Explorer"
        }
    }
}

// Public: get the info from the browser's rendering engine.
// Returns:
//      - a string containing the name of the Engine.
//      - a string containing the version of the Engine.
func (p *UserAgent) Engine() (string, string) {
    return p.browser.engine, p.browser.engine_version
}

// Public: get the info from the browser itself.
// Returns:
//      - a string containing the name of the browser.
//      - a string containing the version of the browser.
func (p *UserAgent) Browser() (string, string) {
    return p.browser.name, p.browser.version
}
