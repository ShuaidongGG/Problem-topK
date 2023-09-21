package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"runtime/pprof"
	"sort"
	"time"
)

const (
	K = 100
)

func main() {
	var cpuProfile = flag.String("cpuprofile", "", "write cpu profile to file")
	var memProfile = flag.String("memprofile", "", "write mem profile to file")
	// 替换为你生成的文件路径
	file, err := os.Open("/path/to/your/random_numbers.bin")
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
	if err != nil {
		println(err)
		return
	}
	defer file.Close()
	topK := []int64{}
	blockSize := 20 * 1024 * 1024
	buffer := make([]byte, blockSize)
	reader := bufio.NewReader(file)
	start := time.Now()
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
			temp := make([]int64, blockSize/8)
			err := binary.Read(bytes.NewReader(buffer), binary.LittleEndian, &temp)
			if err != nil {
				fmt.Println("Error parsing int64:", err)
				break
			}
			sort.Slice(temp, func(i, j int) bool { return temp[i] > temp[j] })
			topK = append(topK, temp[:100]...)
		}

		// 检查是否已经到达文件末尾
		if err == io.EOF {
			break
		}
	}
	sort.Slice(topK, func(i, j int) bool { return topK[i] > topK[j] })
	for i := 0; i < K; i++ {
		println(topK[i])
	}
	println("cost time:")
	println(time.Since(start).Seconds())
	//采样 memory 状态
	if *memProfile != "" {
		f, err := os.Create(*memProfile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.WriteHeapProfile(f)
		f.Close()
	}
}
