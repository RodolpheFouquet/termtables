// Copyright 2013 Apcera Inc. All rights reserved.

// +build !windows

package locale

import (
	"os"
	"strings"
)

// GetCharmap returns the character map (aka CODESET) of the current locale,
// according to the system libc implementation.  The value is returned as a
// string.  Common values which might be seen include "US-ASCII" and "UTF-8".
// check : http://www.opensource.apple.com/source/Libc/Libc-498/locale/nl_langinfo-fbsd.c
func GetCharmap() string {
	locale := strings.Split(os.Getenv("LC_CTYPE"), ".")
	if len(locale) != 2 || locale[1] == "C" || locale[1] == "POSIX" {
		return "US-ASCII"
	} else {
		return locale[1]

	}
}
