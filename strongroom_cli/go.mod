module github.com/shunr/strongroom/strongroom_cli

go 1.14

replace github.com/shunr/strongroom_core => ../strongroom_core

require (
	github.com/google/uuid v1.1.2
	github.com/shunr/strongroom_core v0.0.0-00010101000000-000000000000
	github.com/urfave/cli/v2 v2.2.0
)
