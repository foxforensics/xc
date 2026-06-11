NAME
====

**xc** - experimental terminal chain-of-custody

SYNOPSIS
========

```console
$ xc COMMAND ...
```

DESCRIPTION
===========

xc is an experimental terminal chain-of-custody tool. It executes the given command, prints the combined output of STDOUT and STDERR, and hashes this output using the SHA256 hash algorithm. Any input from STDIN will be forwarded to the command. Any occurring error will be returned transparently and the commands exit code will be mirrored.

INSTALLATION
============

```console
$ go install go.foxforensics.eu/xc@latest
```

REFERENCES
==========

- [_Legal Guide for the Forensic Expert_](https://nij.ojp.gov/nij-hosted-online-training-courses/law-101-legal-guide-forensic-expert/pretrial/pretrial-motions/chain-custody) - National Institute of Justice

SEE ALSO
========

[**sha256sum(1)**](https://man7.org/linux/man-pages/man1/sha256sum.1.html)