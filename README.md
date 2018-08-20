# sshmidr

Generates list of IP globs usable in `Host` lines in your `~/.ssh/config` file from a given CIDR.

## Usage

    $ sshmidr 10.48.0.0/16
    10.48.*

    $ sshmidr 10.48.0.0/17
    10.48.?.* 10.48.??.* 10.48.10?.* 10.48.11?.* 10.48.120.* 10.48.121.* 10.48.122.* 10.48.123.* 10.48.124.* 10.48.125.* 10.48.126.* 10.48.127.*

    $ sshmidr 192.168.1.0/25
    192.168.1.?.* 192.168.1.??.* 192.168.1.10?.* 192.168.1.11?.* 192.168.1.120.* 192.168.1.121.* 192.168.1.122.* 192.168.1.123.* 192.168.1.124.* 192.168.1.125.* 192.168.1.126.* 192.168.1.127.*

## Installation

    go get github.com/daveadams/sshmidr

## License

This software is public domain. No rights are reserved. See LICENSE for more
information.
