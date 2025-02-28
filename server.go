// Copyright 2022 - MinIO, Inc. All rights reserved.
// Use of this source code is governed by the AGPLv3
// license that can be found in the LICENSE file.

package kes

import "time"

// State is a KES server status snapshot.
type State struct {
	Version string // The KES server version

	UpTime time.Duration // The time the KES server has been up and running
}

// API describes a KES server API.
type API struct {
	Method  string        // The HTTP method
	Path    string        // The API path without its arguments. For example: "/v1/status"
	MaxBody int64         // The max. size of request bodies accepted
	Timeout time.Duration // Amount of time after which request will time out
}
