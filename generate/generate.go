package main

import (
	"encoding/binary"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	// 设置随机数种子，以确保每次运行生成的随机数都不同

	// 计算要生成的int64数字的数量
	numInt64s := 512 * 1024 * 1024 / 8 // 512MB / 8字节

	// 生成随机的int64数字并存储在切片中
	randomInt64s := make([]int64, numInt64s)
	for i := 0; i < numInt64s; i++ {
		randomInt64s[i] = rand.Int63()
	}

	// 打开文件以写入随机数
	file, err := os.Create("random_numbers.bin")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// 将int64数字写入文件
	for _, num := range randomInt64s {
		err := binary.Write(file, binary.LittleEndian, num)
		if err != nil {
			println(err)
			return
		}
	}
	println("ok")
	// fmt.Printf("Generated and wrote %d random int64 numbers (512MB) to 'random_numbers.bin'.\n", numInt64s)
}
