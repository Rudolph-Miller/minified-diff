# Midified Diff

Get diff from minified files.


# Usage

```sh
go get github.com/Rudolph-Miller/minified-diff
$GOPATH/bin/minified-diff test/lodash_4_0_0.min.js test/lodash_4_0_1.min.js
```


# Example

```
$ cat test/sample1
sample
text1
```

```
$ cat test/sample2
sample
text2
```

```sh
$ minified-diff test/sample1 test/sample2
 sample
-text1
+text2
```

```sh
$ minified-diff test/sample1 test/sample2 | grep '^+'
text2
```
