/*
	Copyright 2021 Alexander Vollschwitz <xelalex@gmx.net>

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	  http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package util

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
)

//
func SplitRef(ref string) (repo, path, tag string) {

	ix := strings.Index(ref, "/")

	if ix == -1 {
		repo = ""
		path = ref
	} else {
		repo = ref[:ix]
		path = ref[ix+1:]
	}

	ix = strings.Index(path, ":")

	if ix > -1 {
		tag = path[ix+1:]
		path = path[:ix]
	}

	return
}

//
func SplitPlatform(p string) (os, arch, variant string) {

	ix := strings.Index(p, "/")

	if ix == -1 {
		os = p
		arch = ""
	} else {
		os = p[:ix]
		arch = p[ix+1:]
	}

	ix = strings.Index(arch, "/")

	if ix > -1 {
		variant = arch[ix+1:]
		arch = arch[:ix]
	}

	return
}

//
type creds struct {
	Username string
	Password string
}

//
func DecodeJSONAuth(authBase64 string) string {

	if authBase64 == "" {
		return ""
	}

	decoded, err := base64.StdEncoding.DecodeString(authBase64)
	if err != nil {
		log.Error(err)
		return ""
	}

	var ret creds
	if err := json.Unmarshal([]byte(decoded), &ret); err != nil {
		log.Error(err)
		return ""
	}

	return fmt.Sprintf("%s:%s", ret.Username, ret.Password)
}
