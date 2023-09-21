# Problem-topK
My solution for Big size numbers with Small memory to select top K numbers. 

## How to use

1. generate big size random_numbers by /generate/generate.go. You can change source code file to generate any size you want. Default size is 512MB. Just type it on your terminal `go run generate.go`.
2. fix-up /heap/heap.go and /quick/quick.go source code, `file, err := os.Open("/path/to/your/random_numbers.bin")` means you should replace it with your random_numbers' local path.
3. Just run it. `go run quick.go` or `go run heap.go`

> Tips
> My solution is coding by golang, so make sure your machine is alreay have it.
> You can use pprof or other tools to analyze cpu-usage and mem-usage.
> Different level of hardware makes result very different.

## Result on my machine

### heap usage

![cpu.svg](/heap/heap-cpu.svg)

![mem.svg](heap/heap-mem.svg)

### quick usage

![cpu.svg](quick/quick-cpu.svg)

![mem.svg](quick/quick-mem.svg)

