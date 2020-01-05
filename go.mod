module github.com/Win-Man/tidb-healthy-check

go 1.13

require (
	github.com/bndr/gotabulate v1.1.2
	github.com/go-ini/ini v1.51.1 // indirect
	github.com/stretchr/testify v1.4.0 // indirect
)

replace (
	github.com/Win-Man/tidb-healthy-check/collector => ./tidb-healthy-check/collector
	github.com/Win-Man/tidb-healthy-check/pkg/config => ./tidb-healthy-check/pkg/config
	github.com/Win-Man/tidb-healthy-check/pkg/requests => ./tidb-healthy-check/pkg/requests
)
