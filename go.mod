module github.com/rsteube/carapace-spec-kong

go 1.19

require (
	github.com/alecthomas/kong v0.7.1
	github.com/rsteube/carapace-spec v0.10.1
	gopkg.in/yaml.v3 v3.0.1
)

replace github.com/rsteube/carapace-spec => ../carapace-spec/
