
## Leetcode Problems in Go

### Live Testing

* On Mac: Install fswatch : `brew install fswatch`

```
while true; do; fswatch -o ../ | go test -v .; done
```

This will run tests on any changes to problems or the libs it depends on.
