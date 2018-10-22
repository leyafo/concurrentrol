# concurrentrol
Limit the maximum number of goroutines running at the same time.

## Usage
`go get -u github.com/leyafo/concurrentrol` or you can copy the limit.go for anyway.

The simplest code.
```go
	var count int32
	Run(10, 100, func(i int) error { //Running 100 tasks and concurrent task is 10.
		atomic.AddInt32(&count, int32(i))
		return nil
	})
```

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE.md) file for details.

## Author
李亚夫 - [leyafo](http://www.leyafo.com)