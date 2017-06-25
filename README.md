[![License](http://img.shields.io/:license-mit-blue.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/o1egl/ratelimit?status.svg)](https://godoc.org/github.com/o1egl/ratelimit)
[![Build Status](http://img.shields.io/travis/o1egl/ratelimit.svg?style=flat-square)](https://travis-ci.org/o1egl/ratelimit)
[![Coverage Status](http://img.shields.io/coveralls/o1egl/ratelimit.svg?style=flat-square)](https://coveralls.io/r/o1egl/ratelimit)
# RateLimit

Very simple rate limit implementation.

## Install

```
$ go get -u github.com/o1egl/ratelimit
```

## Usage

```go
    r = NewRatelimiter(limit, interval)

    if err := r.Get(context.Background()); err == nil {
        // do something
    }
       

```

## Copyright, License & Contributors

### Submitting a Pull Request

1. Fork it.
2. Open a [Pull Request](https://github.com/o1egl/ratelimit/pulls)
3. Enjoy a refreshing Diet Coke and wait

ratelimit is released under the MIT license. See [LICENSE](LICENSE)