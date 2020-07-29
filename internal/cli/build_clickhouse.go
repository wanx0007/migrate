// +build clickhouse

package cli

import (
	_ "github.com/ClickHouse/clickhouse-go"
	_ "github.com/fun/golang-migrate/migrate/v4/database/clickhouse"
)
