# Choria `mcorpc` Agent Provider

This is an Agent Provider for the [Choria Server](https://github.com/choria-io/go-choria) that provides compatability with The Marionette Collective (mcollective) Simple RPC system.

Features:

 * Agent DDLs represented in memory and loaded from JSON files
 * Agents, Actions and Agent Metadata that maps to the same terminology and behaviour as MCollective
 * Auditing that is compatible with the Ruby based Choria auditing plugin
 * A framework for writing new Agents in Go that can be compiled into the Choria Server
 * Agents written in Go and compiled into the Choria Server:
   * `choria_util` - Utilities used to interrogate version information about the server
   * `rpcutil` - General RPC utilities like extracting facts and statistics, compatible with MCollective as far as possible
   * `discovery` - Agent used to assist broadcast based discovery
   * `choria_provision` - Beginnings of a Choria self provisioning system via the [provisioning-agent](https://github.com/choria-io/provisioning-agent) repo
 * A wrapper around the historical Ruby MCollective agent system that can run a MCollective agent in a sandbox without requiring the deprecated `mcollectived` daemon
 * A full featured Go client to the MCollective RPC system that is compatible with Ruby and Go nodes
   * Broadcast discovery
   * RPC Requests with the usual features like batches, direct, broadcast and more

TODO:

 * Authorization system that matches the action policy feature
 * Support limiting compiled in agents using the new Authorization feature above
 * Ability to execute ruby data plugins during discovery to enable compound filters