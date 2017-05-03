# SE Dev Tools

This repository contains tools that help the SE team in performing day to day
duties.

## Tools

Tools included:

* Snap Revision Provider
* Flight Schedule
* Bug Surfer

### Snap Revision Provider

This tool uses the snap store API to fetch and display version and revision of
a particular snap in all channels. It returns the results for all supported
architectures.

#### Usage

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

### Flight Schedule

This tool queries the Launchpad and returns the list of the merge proposals
currently in motion for the [snappy-hwe-team](https://launchpad.net/~snappy-hwe-team)

#### Usage

```
$ ./flightschedule
```

### Bug Surfer

This tool queries the Launchpad and returns the list of 10 recently updated bugs
for a person who is logged in. Additionally it can read a config file named 
*bugsurfer.config* from:

* ~/.go-launchpad/
* $SNAP_DATA/.go-launchpad

for additional bug lists to provide.

#### Config file format

Each line shall contain either $DISTRIBUTION or $DISTRIBUTION/$PACKAGE. The
Distribution is for example 'ubuntu', a package can be 'bluez'. It can be easily
read from the bug url: https://bugs.launchpad.net/ubuntu/+source/bluez/+bug/1221524
The pattern is: $DISTRIBUTION/+source/$PACKAGE

For example a config file for including all bugs for ubuntu and for bluez
separately would look like:

```
$ cat ~/.go-launchpad/bugsurfer.config
ubuntu
ubuntu/bluez
$
```

### Dependencies

```
go get github.com/bergotorino/go-launchpad/launchpad
```
