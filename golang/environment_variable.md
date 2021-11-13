


### GOGC
The GOGC variable sets the initial garbage collection target percentage. A collection is triggered when the ratio of freshly allocated data to live data remaining after the previous collection reaches this percentage.

The default is GOGC=100, which means garbage collection will not be triggered until the heap has grown by 100% since the previous collection. Effectively GOGC=100 (the default) means the garbage collector will run each time the live heap doubles.

- Setting this value higher, say GOGC=200, will delay the start of a garbage collection cycle until the live heap has grown to 200% of the previous size.
- Setting the value lower, say GOGC=20 will cause the garbage collector to be triggered more often as less new data can be allocated on the heap before triggering a collection.


Setting GOGC=off disables the garbage collector entirely. The runtime/debug package’s SetGCPercent function allows changing this percentage at run time. See https://golang.org/pkg/runtime/debug/#SetGCPercent.

```
$ GOGC=200 ./main
```

https://swsmile.info/post/golang-gogc/