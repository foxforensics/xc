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

xr is an experimental terminal chain-of-custody tool. It targets to answer two main questions about event logs: WHAT and WHEN did it happen? Contrary to existing tools, it tries to answer these questions by analyzing the raw event record structure, rather than parsing whole event log chunks. By reading from any input stream, xr is capable of carving raw forensic disk images and memory dumps.

INSTALLATION
============

```console
$ go install go.foxforensics.dev/xc@latest
```

REFERENCES
==========

- [_Legal Guide for the Forensic Expert_](https://nij.ojp.gov/nij-hosted-online-training-courses/law-101-legal-guide-forensic-expert/pretrial/pretrial-motions/chain-custody) - National Institute of Justice

SEE ALSO
========

[**sha256sum(1)**](https://man7.org/linux/man-pages/man1/sha256sum.1.html)