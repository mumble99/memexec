#! /bin/bash

memexec(){python3 -c 'from os import memfd_create, getpid, system; from sys import argv; a=" ".join(argv[1:]); fd = memfd_create("", 1); open(f"/proc/self/fd/{fd}", "wb").write(open(0, "rb").read()); system(f"exec /proc/{getpid()}/fd/{fd} {a}")' "$@"}
# Example: cat /usr/bin/id | memexec -u