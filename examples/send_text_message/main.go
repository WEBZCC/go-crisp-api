// Copyright 2018 Crisp IM SAS All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
  "github.com/crisp-im/go-crisp-api/crisp/v3"
  "fmt"
)

func main() {
  client := crisp.New()
  client.Authenticate("7c3ef21c-1e04-41ce-8c06-5605c346f73e", "cc29e1a5086e428fcc6a697d5837a66d82808e65c5cce006fbf2191ceea80a0a")

  // Build text message
  message := crisp.ConversationTextMessageNew{Type: "text", From: "operator", Origin: "chat", Content: "Hello :)", Fingerprint: 12345}

  // Send text message for `website_id` and `session_id`
  _, _, err := client.Website.SendTextMessageInConversation("8c842203-7ed8-4e29-a608-7cf78a7d2fcc", "session_19e5240f-0a8d-461e-a661-a3123fc6eec9", message)

  if err != nil {
    fmt.Printf("Error: %s", err)
  } else {
    fmt.Print("Message sent")
  }
}
