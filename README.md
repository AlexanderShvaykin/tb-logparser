Remote grep on many servers

```shell script
make
cat hosts | tb-logparser grep 'some pattern' -f path_to_file/log.log
```

Async grep in local path

```shell script
make
cd my_log_dir/
tb-logparser grep 'some pattern' --local
```

It read all files and return matches lines