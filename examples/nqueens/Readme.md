
In main:
```
	f, _ := os.Create("cpu.prof")
    pprof.StartCPUProfile(f)
    defer pprof.StopCPUProfile()
```
then
`go tool pprof cpu.prof`

top
list functionName