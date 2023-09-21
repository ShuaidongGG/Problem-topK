package main

import (
	"bufio"
	"bytes"
	"container/heap"
	"encoding/binary"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

const (
	K = 100
)

type priorityqueue []int64

func (pq priorityqueue) Len() int           { return len(pq) }
func (pq priorityqueue) Less(i, j int) bool { return pq[i] < pq[j] }
func (pq priorityqueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq *priorityqueue) Push(x interface{}) {
	*pq = append(*pq, x.(int64))
}
func (pq *priorityqueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[:n-1]
	return x
}

func main() {
	var cpuProfile = flag.String("cpuprofile", "", "write cpu profile to file")
	var memProfile = flag.String("memprofile", "", "write mem profile to file")
	flag.Parse()
	//采样cpu运行状态
	if *cpuProfile != "" {
		f, err := os.Create(*cpuProfile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	start := time.Now()
	// 替换为你生成的文件路径
	file, err := os.Open("/path/to/your/random_numbers.bin")
	if err != nil {
		println(err)
		return
	}
	defer file.Close()

	blockSize := 20 * 1024 * 1024
	buffer := make([]byte, blockSize)
	reader := bufio.NewReader(file)
	topK := new(priorityqueue)
	for {
		// 读取块数据
		n, err := reader.Read(buffer)
		if err != nil && err != io.EOF {
			fmt.Println("Error reading file:", err)
			break
		}

		// 处理块数据
		if n > 0 {
			// 在这里处理读取的数据，例如将其写入其他地方或进行其他操作
			// 这里只是示例，打印块大小
			// fmt.Printf("Read %d bytes\n", n)
			for i := 0; i < blockSize; i += 8 {
				var data int64
				err := binary.Read(bytes.NewReader(buffer[i:i+8]), binary.LittleEndian, &data)
				if err != nil {
					fmt.Println("Error parsing int64:", err)
					break
				}
				heap.Push(topK, data)
				if topK.Len() > K {
					heap.Pop(topK)
				}
			}
		}

		// 检查是否已经到达文件末尾
		if err == io.EOF {
			break
		}
	}
	// 打印显示topK
	for i := 0; i < K; i++ {
		println(heap.Pop(topK).(int64))
	}
	println("cost time:")
	println(time.Since(start).Seconds())
	if *memProfile != "" {
		f, err := os.Create(*memProfile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.WriteHeapProfile(f)
		f.Close()
	}
}
