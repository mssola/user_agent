// Copyright (C) 2012 Miquel Sabaté <mikisabate@gmail.com>
// This file is licensed under the GNU LGPL v3 or later.
// See the LICENSE file.

package user_agent;

import "user_agent";


/*
 * Internal: used by all tests to parse a User-Agent string.
 *
 * ua - the User-Agent string.
 *
 * Returns a reference to a newly created UserAgent.
 */
func parse(ua string) *user_agent.UserAgent {
    parser := new(user_agent.UserAgent);
    parser.Parse(ua);
    return parser;
}
