|Date      |Issue |Description                                                                                              |
|----------|------|---------------------------------------------------------------------------------------------------------|
|2020/01/12|      |Release 0.10.0                                                                                           |
|2019/12/24|152   |Add a Rego policy engine                                                                                 |
|2019/12/25|134   |Support formatting aggregates for display                                                                |
|2019/12/24|143   |Support formatting RPC replies for console use                                                           |
|2019/12/07|      |Release 0.9.0                                                                                            |
|2019/12/03|134   |Support regular expressions in callerid matches in action policy                                         |
|2019/11/28|57    |Improve Ruby DDL generation                                                                              |
|2019/11/27|128   |Correctly track total time spend processing requests                                                     |
|2019/11/25|126   |Improve handling for actions without any inputs in DDL validation                                        |
|2019/09/22|121   |Expose the path to the agent configuration to external agents                                            |
|2019/09/19|      |Release 0.8.0                                                                                            |
|2019/09/17|112   |Create an `action_policy` like authorization framework                                                   |
|2019/09/16|23    |Remove old progress system                                                                               |
|2019/09/14|97    |Support `External Agents`                                                                                |
|2019/09/13|102   |Support generating a Ruby DDL                                                                            |
|2019/09/12|99    |Prevent the ruby provider from attempting to start agents it does not know                               |
|2019/09/08|      |Release 0.7.1                                                                                            |
|2019/09/08|94    |Support type hints on outputs                                                                            |
|2019/09/08|93    |Support setting defaults on inputs and outputs based on DDL hints                                        |
|2019/09/08|93    |Support `Hash`, `hash`, `Array` and `array` as types in DDL inputs and outputs                           |
|2019/09/07|      |Release 0.7.0                                                                                            |
|2019/09/07|88    |Support converting string from the CLI to the correct data types for any input                           |
|2019/09/07|88    |Support `shellsafe`, `ipv4address`, `ipv6address`, `ipaddress`, `shellsafe` and regex validators         |
|2019/09/07|88    |Support `input` in action DDLs                                                                           |
|2019/09/05|      |Release 0.6.1                                                                                            |
|2019/09/05|86    |Add a `chart` aggregator                                                                                 |
|2019/09/05|      |Release 0.6.0                                                                                            |
|2019/09/05|81    |Add callbacks before and after discovery to assist clients with progress indicators and feedback         |
|2019/09/05|72    |Add `summary` and `average` aggregation methods for reply data                                           |
|2019/09/03|68    |Add utilities for limiting the request targets - percentage or specific counts                           |
|2019/08/18|62    |Support checking `plugin.$agent.activate_agent` when loading ruby agents                                 |
|2019/07/24|      |Release 0.5.0                                                                                            |
|2019/06/27|58    |Allow go agents to determine at run time if they should activate                                         |
|2019/06/26|53    |Use extracted `go-config` and `go-srvcache`, improve use of interfaces                                   |
|2019/06/21|      |Release 0.4.1                                                                                            |
|2019/06/21|      |Support `go mod`                                                                                         |
|2018/05/23|      |Release 0.4.0                                                                                            |
|2018/05/23|47    |Log discovery requests in a similar way to rpc requests                                                  |
|2018/05/02|28,42 |Support Choria Autonomous Agents in `choria_util`                                                        |
|2018/12/27|      |Release 0.3.3                                                                                            |
|2018/12/26|36    |Increase `choria_util` agent timeout to facilitate slow facter runs                                      |
|2018/12/19|      |Release 0.3.2                                                                                            |
|2018/12/19|33    |Report protocol security and connector TLS in choria_util#info                                           |
|2018/09/16|      |Release 0.3.1                                                                                            |
|2018/09/16|29    |Adds an utility to generate Choria agent plugins                                                         |
|2018/09/15|      |Release 0.3.0                                                                                            |
|2018/09/15|26    |Support the new Choria plugin system                                                                     |
|2018/09/05|18    |When doing discovery ensure the agent is in the filter                                                   |
|2018/09/05|19    |Record `agent` and `action` in the stats recorded by the client                                          |
|2018/08/16|      |Release 0.2.0                                                                                            |
|2018/08/16|13-15 |Rework discovery helper to fit better with overall package design                                        |
|2018/08/09|      |Release 0.1.0                                                                                            |
|2018/08/03|6     |Move provisioning agent to own repository                                                                |
|2018/08/03|8     |Fix deny listing ruby agents like rpcutil and choria_util                                                |
|2018/08/02|4     |Support generating CSRs for the provisioner                                                              |
