# Yet Another CLI, for multiple proxy platforms

## Usage
- [Help](#help)
- [Initialize](#initialize)
- [Provider](#provider)
    * [Switch](#switch)
- [Subscription](#subscription)
    * [Add](#add)
    * [List](#list)
    * [Edit](#edit)
    * [Remove](#remove)
    * [Switch](#switch-1)
    * [Sync](#sync)
- [Daemon](#daemon)
    * [Run](#run)
    * [Kill](#kill)
    * [Enable](#enable)
    * [Disable](#disable)
    * [Restart](#restart)
- [Node](#node)
    * [List](#list-1)
    * [Switch](#switch-2)
- [Conf](#conf)
    * [Mode](#mode)
- [Status](#status)
- [Version](#version)

### Help
Use "ya [command] --help" for more information about a command.

```shell
Usage:
  ya [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  conf        For manage application settings and display program information.
  daemon      Used to manage the daemon processes.
  help        Help about any command
  init        Initialise the application.
  node        To view and switch subscription nodeswitch (Proxies).
  provider    Used to manage the proxy platform utilised.
  status      Print the status of the application.
  subscribe   Subscription management.
  version     Print the version.

Flags:
  -h, --help     help for ya
  -t, --toggle   Help message for toggle
```

---
### Initialize
Before use, you must initialise the application to make it operational.

```shell
Usage:
  ya init --provider <path> [flags]

Aliases:
  init, i

Flags:
  -h, --help              help for init
  -p, --provider string   path to the provider.
```

Choose and initialize the proxy platform to use, either by specifying the provider's path with `ya i` or `ya i -p <path>`. Decide on the provider (proxy platform) during initialization.

---
### Provider
It allows for the display or alteration of the
proxy platform that ya operates on, among othe
r functionalities, for provider management.

```shell
Usage:
  ya provider [command]

Aliases:
  provider, p

Available Commands:
  switch      Switch between different providers.

Flags:
  -h, --help   help for provider
```
#### Switch
Switch the proxy platform used by ya with `ya <provider/p> <switch/s>`. You can also add a proxy program through init, which then becomes your current selected proxy platform by default.

Switch provider: `ya p s`

### Subscription
This command is used to manage subscriptions.

```shell
Usage:
  ya subscribe [command]

Aliases:
  subscribe, s, sub

Available Commands:
  add         Add a new subscription.
  edit        Edit a subscription.
  list        List subscriptions
  remove      Remove a subscription.
  switch      Switch subscription.
  sync        Synchronize the current subscription.

Flags:
  -h, --help   help for subscribe
```

#### Add
Add a new subscription for the current proxy platform with a remote subscription or specify a subscription file using `ya <subscribe/s/sub> <add/a>`, with the optional `-l` or `-p` flags for Link and Path, respectively.
```shell
ya subscribe add <name> --link <link> --path <path>
```
Add a remote subscription with `ya s a foo -l https://foo.bar`, or add a local config file with `ya s a foo -p /path/to/bar.json`.

#### List
List subscriptions with `ya <subscribe/s/sub> <list/ls/l>`, or view subscription details by ID.
```shell
ya subscribe list --id <id>
```
List subscriptions with `ya s l` or use a subscription ID with `ya s l -i <id>` to view subscription details directly.

#### Edit
Edit a specified subscription with `ya <subscribe/s/sub> <edit/e>`, specifying optional `-l` or `-n` flags.
```shell
ya subscribe edit <id> --name <name> --link <link>
```
Edit the name of the subscription with ID 1 with `ya s e 1 -n foo`, or edit the remote subscription with `ya s e 1 -l https://foo.bar`.

#### Remove
Remove a specified subscription with `ya <subscribe/sub/s> <remove/r/rm/delete/d>`.
```shell
ya subscribe remove <id>
```
Remove the subscription with ID 1 with `ya s r 1`.

#### Switch
Switch subscriptions with `ya <subscribe/sub/s> <switch/s>`, or specify an ID to switch subscriptions.
```shell
ya subscribe switch --id <id>
```
Switch subscriptions with `ya s s` or specify an ID with `ya s s -i 1`.

#### Sync
Synchronize the current subscription with `ya <subscribe/sub/s> sync`, or specify an ID to sync.
```shell
ya subscribe sync
```
Synchronize subscriptions with `ya s sync`.

---
### Daemon
Used to manage the daemon processes.

```shell
Usage:
  ya daemon [command]

Aliases:
  daemon d

Available Commands:
  disable     Disable daemon autostart on boot
  enable      Enable autostart on boot.
  kill        Kill the daemon.
  restart     Restart the daemon.
  run         Run the daemon.

Flags:
  -h, --help   help for daemon
```
#### Run
Start the daemon process with `ya <daemon/d> <run/start>`.
```shell
ya d run
```

#### Kill
Terminate the daemon process with `ya <daemon/d> <kill/k/stop>`.
```shell
ya d k
```

#### Enable
Set the daemon to automatically start on boot with `ya <daemon/d> <enable/e>`.
```shell
ya d e
```

#### Disable
Prevent the daemon from starting automatically on boot with `ya <daemon/d> <disable/d>`.
```shell
ya d d
```

#### Restart
Restart the daemon with `ya <daemon/d> restart`.
```shell
ya d restart
```

---
### Node
View and switch subscription nodes (Proxies).

```shell
Usage:
  ya node [command]

Available Commands:
  list        Used to display nodeswitch (Proxies) information in subscription.
  switch      Switch switch.

Flags:
  -h, --help   help for node
```
#### List
List the current strategy groups or view nodes in a group by name with `ya <node/n> <list/proxies/ls/l>`.
```shell
ya node list --group <name>
```
List nodes with `ya n l` or view nodes in a specific group with `ya n l -g foo`.

#### Switch
Switch nodes within a group with `ya <node/n> <switch/s>`.

```shell
ya node switch --group <name>
```
Switch nodes with `ya n s` or specify a group name with `ya n s -g foo`.

---
### Conf
For manage application settings and display program information.

```shell
Usage:
  ya conf [command]
  
Aliases:
  mode m
  
Available Commands:
  mode        Switch the proxy mode (Global | Rule | Direct).
```

#### Mode
Switch the proxy mode between Global, Rule, or Direct with `ya <conf/c> <mode/m>`.

```shell
ya conf mode --mode <mode>
```
Switch proxy modes with `ya c m`, choosing between global for global mode, rule for rule-based mode, or direct for direct connection mode. Use `-m` to specify the mode directly. `-m` accepts only one of these three strings: `Global`, `Rule`, `Direct`.

---
### Status
Print the status of the application.

```shell
Usage:
  ya status [flags]

Aliases:
  status, state
```

---
### Version
Print version of the application.


## Build
Linux 64-bit AMD64 architecture

To compile, ensure CGO `CGO_ENABLED=1` is enabled:
- Use aarch64-linux-gnu-gcc for ARM architecture
- Use x86_64-w64-mingw32-gcc for Windows architecture