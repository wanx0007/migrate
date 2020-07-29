// +build clickhouse

package cli

import (
	_ "github.com/ClickHouse/clickhouse-go"
	_ "github.com/wanx0007/migrate/v4/database/clickhouse"
)
