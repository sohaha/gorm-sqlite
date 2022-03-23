//go:build !gcc
// +build !gcc

package sqlite

import (
	_ "github.com/glebarez/go-sqlite"
)

// DriverName is the default driver name for SQLite.
const DriverName = "sqlite"
