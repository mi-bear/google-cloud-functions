// Copyright 2017 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"./nodego"
)

// Token is Slash Command Token.
const Token = "YOUR TOKEN"

// Response xxx.
type Response struct {
	ResponseType string `json:"response_type"`
	Text         string `json:"text"`
}

func init() {
	nodego.OverrideLogger()
}

func main() {
	flag.Parse()

	http.HandleFunc(nodego.HTTPTrigger, nodego.WithLoggerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		if r.Method != "POST" {
			log.Println("Only POST! bear is angry!!!")
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		// application/x-www-form-urlencoded
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		body := buf.String()
		parsed, err := url.ParseQuery(body)
		if err != nil {
			log.Println("ParseQuery error.")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// verify
		if parsed.Get("token") != Token {
			log.Println("Invalid credentials! bear is angry!!!")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		response := Response{
			ResponseType: "in_channel",
			Text:         talk(parsed.Get("user_id"), parsed.Get("user_name"), parsed.Get("text")),
		}

		// response
		w.Header().Add("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&response)
	}))

	defer nodego.TakeOver()
}

func talk(uid, uname, text string) string {
	// TODO: add message valiations
	switch text {
	case "のどがかわいた":
		return fmt.Sprintf("<@%s|%s> :beers:", uid, uname)
	default:
		return text
	}
}
