package main

import (
	"fmt"
	"net"
	"sort"
	"time"
)

// 扫描 1~65535 以内的端口
const MIN_PORT = 1
const MAX_PORT = 65535

// 单线程版本 tcp 端口扫描器
// func main() {
// 	start := time.Now()

// 	for i := MIN_PORT; i < MAX_PORT; i++ {
// 		address := fmt.Sprintf("127.0.0.1:%d", i)
// 		conn, err := net.Dial("tcp", address)
// 		if err != nil {
// 			fmt.Printf("%s 关闭了\n", address)
// 			continue
// 		}

// 		conn.Close()
// 		fmt.Printf("%s 打开了\n", address)
// 	}

// 	elapsed := time.Since(start)
// 	fmt.Printf("\n\nSince: %d\n", elapsed)
// }

// 并发版本 tcp 端口扫描器
// func main() {
// 	var wg sync.WaitGroup
// 	start := time.Now()

// 	for i := MIN_PORT; i < MAX_PORT; i++ {
// 		wg.Add(1)
// 		go func(j int) {
// 			defer wg.Done()
// 			address := fmt.Sprintf("127.0.0.1:%d", j)
// 			conn, err := net.Dial("tcp", address)
// 			if err != nil {
// 				fmt.Printf("%s 关闭了\n", address)
// 				return
// 			}

// 			conn.Close()
// 			fmt.Printf("%s 打开了\n", address)
// 		}(i)
// 	}
// 	wg.Wait()

// 	elapsed := time.Since(start)
// 	fmt.Printf("\n\nSince: %d\n", elapsed)
// }

// goroutine 池并发版本 tcp 端口扫描器
func main() {
	start := time.Now()

	ports := make(chan int, 100) // 缓冲 100
	results := make(chan []int)  // 使用 results chan 达到了 WaitGroup 的效果

	var oponports []int
	var closedports []int

	// 开启 100 个 goroutine
	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	// 发送结果和底下的收集结果，应该要并行执行，所以分别放在不同的 goroutine 里面
	go func() {
		for i := MIN_PORT; i < MAX_PORT; i++ {
			ports <- i // 将数据传递到 worker 的 channel 中 (随机分配)
		}
	}()

	// 收集结果: 存储端口状态
	// 等上面分配的 MAX_PORT 次任务全部完成之后才会继续往下执行，否则会堵塞
	for i := MIN_PORT; i < MAX_PORT; i++ {
		port := <-results
		if port[0] != 0 {
			// 端口打开了
			oponports = append(oponports, port[1])
		} else {
			// 端口关闭了
			closedports = append(closedports, port[1])
		}
	}

	close(ports)
	close(results)

	// 排序并输出
	sort.Ints(oponports)
	sort.Ints(closedports)
	for _, port := range oponports {
		fmt.Printf("%d 打开了\n", port)
	}
	for _, port := range closedports {
		fmt.Printf("%d 关闭了\n", port)
	}

	elapsed := time.Since(start)
	fmt.Printf("\n\nSince: %d\n", elapsed)
}

func worker(ports <-chan int, results chan<- []int) {
	for p := range ports {
		address := fmt.Sprintf("127.0.0.1:%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			// fmt.Printf("%s 关闭了\n", address)
			results <- []int{0, p}
			continue
		}
		conn.Close()
		// fmt.Printf("%s 打开了\n", address)
		results <- []int{1, p}
	}
}
