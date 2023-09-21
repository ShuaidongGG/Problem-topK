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

- cpu-usage [cpu](/heap/heap-cpu.svg)
  - all cost: 16.26s
  - heap Push func: 2.90s
  - heap Pop func: 5.19s
  - read from disk: 0.27s
  - read from memory: 2.80s
  - read int64 one by one: 3.46s

- mem-usage [mem](/heap/heap-mem.svg)
  -  all cost: 20MB

### quick usage

- cpu-usage [cpu](/quick/quick-cpu.svg)
  - all cost: 15.02s
  - sort func: 13.61s
  - read from disk: 0.11s
  - read from memory: 1.23s

- mem-usage [mem](/quick/quick-mem.svg)
  -  all cost: 40MB
  -  main cost: 20MB

### Discussions

In the case of selecting the top K data scenarios from a huge amount of data, Heap sort does show better performance on my machine compared to quick sort. 

The possible reasons for this are as follows:

1. Adjust small size heap is `O(logK)`. heap sort is `O(NlogK)`, N is counts of numbers. quick sort is `O(N/n*nlogn) == O(Nlogn)`, N is counts of numbers, n is counts of every chunk(20MB) numbers.
2. Quick sort is implemented recursively and requires additional memory.