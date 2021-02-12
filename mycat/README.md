# mycat
Output the given file to standard output

## Usage
```sh
$ go run main.go testdata/hoge.txt testdata/fuga.txt
hoge
hoge hoge
fuga
fuga fuga
```

`-n` Can be added to display line numbers.
```
$ go run main.go -n testdata/hoge.txt testdata/fuga.txt
1: hoge
2: hoge hoge
3: fuga
4: fuga fuga
```
