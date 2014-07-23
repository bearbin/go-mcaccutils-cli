go-mcaccutils-cli
=================

A CLI interface for [go-mcaccutils](https://github.com/bearbin/go-mcaccutils),
allowing easy access to player name and UUID resolution.

Installation
------------

    go get github.com/bearbin/go-mcaccutils-cli

Usage
-----

    $GOPATH/bin/go-mcaccutils-cli uuid notch
    $GOPATH/bin/go-mcaccutils-cli name 104d79d60df34908b7b748432185a63b

Warnings
--------

 * This is not guaranteed to work, and will probably break down at some point, 
   so use with caution!
 * The Mojang APIs that this uses have a rate limit of about 1 a second. If this
   is exceeded for a reasonable period of time it might kick you out. 