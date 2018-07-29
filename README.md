# go-hdfs
A fast HDFS CLI 

## What is this?
This is Cobra CLI for [http://github.com/colinmarc/hdfs]. colinmarc/hdfs is a great library, but the CLI left a lot to be desired. In the spirit of Go, this package has no vendoring, and will use the master branch of the upstream HDFS library (and Cobra, incidentally), so it's functionally equivalent.

The advantage of this CLI is both its usability and that it'll be easy to drop in as a static client into a container; HDFS control with no need for Java.

## What are the goals of this instead of upstream?

This package won't depend on an `hdfs-site.xml` for configuration -- while it will support one that exists by default (as does upstream), it will also be fully configured by environment variables and command line flags. This is super useful for containerized environments and Kubernetes.

Flags defined in the Java client will have the same functionality, and may have long flag versions. Flags added to this CLI will follow GNU coreutils usage, as opposed to strict POSIX, which is what upstream prefers.

##
Usage:

```
$ hdfs
GoHDFS is a very fast client for HDFS clusters

Usage:
  gohdfs [flags]
  gohdfs [command]

Available Commands:
  cat         concatenate HDFS files and print on the standard output
  checksum    calculate the checksum of the HDFS files
  chmod       change HDFS file mode bits
  chown       change HDFS file owner and group
  df          concatenate HDFS files and print on the standard output
  du          estimate HDFS space usage
  get         get files from HDFS to the local filesystem
  getmerge    get a directory from HDFS and merge the files into a single local file
  head        output the first part of HDFS files
  help        Help about any command
  ls          list HDFS directory contents
  mkdir       make HDFS directories
  mv          move HDFS files
  put         put files from local filesystem to HDFS
  rm          remove HDFS files or directories
  tail        output the last part of HDFS files
  touch       create HDFS files and modify their timestamps

Flags:
  -h, --help      help for gohdfs
      --version   version for gohdfs

Use "gohdfs [command] --help" for more information about a command.
```

Subcommands:
```
$ hdfs mv
```


Compared to upstream:

```
$ hdfs mv
Both a source and destination are required. 
Usage: ./hdfs COMMAND
The flags available are a subset of the POSIX ones, but should behave similarly.

Valid commands:
  ls [-lah] [FILE]...
  rm [-rf] FILE...
  mv [-nT] SOURCE... DEST
  mkdir [-p] FILE...
  touch [-amc] FILE...
  chmod [-R] OCTAL-MODE FILE...
  chown [-R] OWNER[:GROUP] FILE...
  cat SOURCE...
  head [-n LINES | -c BYTES] SOURCE...
  tail [-n LINES | -c BYTES] SOURCE...
  du [-sh] FILE...
  checksum FILE...
  get SOURCE [DEST]
  getmerge SOURCE DEST
  put SOURCE DEST
  df [-h]
```


