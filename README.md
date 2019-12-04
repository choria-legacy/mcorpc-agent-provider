# Choria `mcorpc` Agent Provider

[![Go Report Card](https://goreportcard.com/badge/github.com/choria-io/mcorpc-agent-provider)](https://goreportcard.com/report/github.com/choria-io/mcorpc-agent-provider)

[![GoDoc](https://godoc.org/github.com/choria-io-mcorpc-agent-provider?status.svg)](https://godoc.org/github.com/choria-io/mcorpc-agent-provider)

This is an Agent Provider for the [Choria Server](https://github.com/choria-io/go-choria) that provides compatability with The Marionette Collective (mcollective) Simple RPC system.

### Agent Features:

 * Agent DDLs represented in memory and loaded from JSON files
 * Agents, Actions and Agent Metadata that maps to the same terminology and behavior as MCollective
 * Auditing that is compatible with the Ruby based Choria auditing plugin
 * A framework for writing new Agents in Go that can be compiled into the Choria Server (see [the docs](https://choria.io/docs/development/server/) for more information on this.)
 * Agents written in Go and compiled into the Choria Server:
   * `choria_util` - Utilities used to interrogate version information about the server
   * `rpcutil` - General RPC utilities like extracting facts and statistics, compatible with MCollective as far as possible
   * `discovery` - Agent used to assist broadcast based discovery
 * A wrapper around the historical Ruby MCollective agent system that can run a MCollective agent in a sandbox without requiring the deprecated `mcollectived` daemon

### Client Features:

 * A fully featured Go client to the MCollective RPC system that is compatible with Ruby and Go nodes
   * Broadcast discovery
   * RPC Requests with the usual features like batches, direct, broadcast and more

### TODO

 * Authorization system that matches the action policy feature
 * Support limiting compiled in agents using the new Authorization feature above
 * Ability to execute ruby data plugins during discovery to enable compound filters
