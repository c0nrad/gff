Global File Finder
==================

This little golang program finds all globally writable files. It recursively finds file that have the "w" permission for all Users.

This is useful for finding entry points when doing penetration testing on an unknown system.

Usage: dff.go -h:

```
Global File Finder
Usage of ./gff:
  -depth=3: depth of directories to search
  -dir=".": directory to start searching from
```

Questions or comments? poptarts4liffe@gmail.com
