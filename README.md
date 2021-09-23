# collectd-plugin-timesync

collectd plugin

# Behavior

1. Request send to NTP host(ex: 169.254.169.123, time.google.com, ...)
2. Retrive offset
3. Put offset to collectd

# Usage

```
  -host string
        destination host. (default "169.254.169.123")
  -identifier string
        collectd identifier. first tier is replaced to hostname. respect COLLECTD_HOSTNAME environment variable. (default "$COLLECTD_HOSTNAME/exec-timesync/gauge-time_offset")
  -interval int
        interval(sec). respect COLLECTD_INTERVAL environment variable. (default 60)
  -v    show version.
  -version
        show version.
```

# Example

## collectd.conf

```
LoadPlugin exec
<Plugin exec>
  Exec ubuntu "/usr/local/bin/collectd-plugin-timesync"
</Plugin>
```

with arguments

```
LoadPlugin exec
<Plugin exec>
  Exec ubuntu "/usr/local/bin/collectd-plugin-timesync" "-host=time.google.com" "-interval=3"
</Plugin>
```

## Output

Syntax

`PUTVAL {HOSTNAME}/exec-timesync/gauge-time_offset interval={N} {UNIXTIME}:{OFFSET_USEC}`

Example

```
PUTVAL "admiin01/exec-timesync/gauge-time_offset" interval=60 1632440516:19252.930000
PUTVAL "admiin01/exec-timesync/gauge-time_offset" interval=60 1632440576:16246.657000
...
```

# Install

1. Download and extract release artifacts
2. `chmod +x /path/to/collectd-plugin-timesync`
3. `mv /path/to/collectd-plugin-timesync /usr/local/bin/.`
