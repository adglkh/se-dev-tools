# SE Dev Tools

This repository contains tools that help the SE team in performing day to day
duties.

Tools included:

* Snap Revision Provider
* Flight Schedule

## Snap Revision Provider

This tool uses the snap store API to fetch and display version and revision of
a particular snap in all channels. It returns the results for all supported
architectures.

### Usage

```
$ ./snap-revision-provider [snap]
```

For example:

```
$ ./snap-revision-provider bluez
 bluez | stable    | 5.37-5     | 54 | armhf
 bluez | stable    | 5.37-5     | 52 | amd64
 bluez | stable    | 5.37-5     | 55 | arm64
 bluez | stable    | 5.37-5     | 53 | i386
 bluez | candidate | 5.44-1     | 65 | amd64
 bluez | candidate | 5.44-1     | 67 | arm64
 bluez | candidate | 5.44-1     | 66 | i386
 bluez | candidate | 5.44-1     | 69 | armhf
 bluez | beta      | 5.44-1     | 69 | armhf
 bluez | beta      | 5.44-1     | 65 | amd64
 bluez | beta      | 5.44-1     | 67 | arm64
 bluez | beta      | 5.44-1     | 66 | i386
 bluez | edge      | 5.44-2-dev | 83 | armhf
 bluez | edge      | 5.44-2-dev | 81 | i386
 bluez | edge      | 5.44-2-dev | 80 | amd64
 bluez | edge      | 5.44-2-dev | 82 | arm6
```

## Flight Schedule

This tool queries the Launchpad and returns the list of the merge proposals
currently in motion for the [snappy-hwe-team](https://launchpad.net/~snappy-hwe-team)

### Usage

```
$ ./flightschedule
```

## Dependencies

```
go get github.com/bergotorino/go-launchpad/launchpad
```
