package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"sync"

	"github.com/k0kubun/go-ansi"
	"github.com/schollz/progressbar/v3"
)

type Downloader struct {
	concurrency int                      // 协程数量
	resume      bool                     // 断点续传
	bar         *progressbar.ProgressBar // 下载进度条
}

func NewDownloader(concurrency int, resume bool) *Downloader {
	return &Downloader{concurrency: concurrency, resume: resume}
}

func (d *Downloader) Download(strURL, filename string) error {
	if filename == "" {
		filename = path.Base(strURL)
	}

	// 并发下载之前先发送 Head 请求, 判断服务器是否支持并发下载 (Accept-Ranges)
	resp, err := http.Head(strURL)
	if err != nil {
		return err
	}

	// 支持并发下载 (支持部分请求下载)
	// 服务器通过该头 Accept-Ranges 来标识自身支持部分请求
	// - bytes：部分请求的单位是 bytes （字节）
	// - none：不支持任何部分请求单位，由于其等同于没有返回此头部，因此很少使用。不过一些浏览器，比如 IE9，会依据该头部去禁用或者移除下载管理器的暂停按钮。
	if resp.StatusCode == http.StatusOK && resp.Header.Get("Accept-Ranges") == "bytes" {
		return d.multiDownload(strURL, filename, int(resp.ContentLength))
	}

	// 不支持并发下载, 直接下载整个文件
	return d.singleDownload(strURL, filename)
}

// 并发下载
func (d *Downloader) multiDownload(strURL, filename string, contentLen int) error {
	d.setBar(contentLen)

	// 计算出每个部分的大小
	partSize := contentLen / d.concurrency

	// 创建部分文件的存放目录
	partDir := d.getPartDir(filename)
	os.Mkdir(partDir, 0777)
	defer os.RemoveAll(partDir)

	var wg sync.WaitGroup
	wg.Add(d.concurrency)

	rangeStart := 0

	for i := 0; i < d.concurrency; i++ {
		// 并发请求
		go func(i, rangeStart int) {
			defer wg.Done()

			rangeEnd := rangeStart + partSize
			// 最后一部分，总长度不能超过 ContentLength
			if i == d.concurrency-1 {
				rangeEnd = contentLen
			}

			downloaded := 0
			// 如果支持断点续传, 先看看是否已有下载好了的部分文件
			if d.resume {
				partFileName := d.getPartFilename(filename, i)
				content, err := os.ReadFile(partFileName)
				if err == nil {
					downloaded = len(content)
				}
				d.bar.Add(downloaded)
			}

			d.downloadPartial(strURL, filename, rangeStart+downloaded, rangeEnd, i)

		}(i, rangeStart)

		rangeStart += partSize + 1
	}

	wg.Wait()

	// 合并文件
	d.merge(filename)

	return nil
}

// 部分请求 (请求头携带 Range)
//   - https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Headers/Range
// Range 告知服务器返回文件的哪一部分
//   - 在一个  Range 头部中，可以一次性请求多个部分，服务器会以 multipart 文件的形式将其返回。
//   - 如果服务器返回的是范围响应，需要使用 206 Partial Content 状态码。
//   - 假如所请求的范围不合法，那么服务器会返回  416 Range Not Satisfiable 状态码，表示客户端错误
//   - 服务器允许忽略  Range  首部，从而返回整个文件，状态码用 200
// 语法如下
//   - Range: <unit>=<range-start>-
//   - Range: <unit>=<range-start>-<range-end>
//   - Range: <unit>=<range-start>-<range-end>, <range-start>-<range-end>
//   - Range: <unit>=<range-start>-<range-end>, <range-start>-<range-end>, <range-start>-<range-end>
func (d *Downloader) downloadPartial(strURL, filename string, rangeStart, rangeEnd, i int) {
	if rangeStart >= rangeEnd {
		return
	}

	req, err := http.NewRequest("GET", strURL, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", rangeStart, rangeEnd))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	flags := os.O_CREATE | os.O_WRONLY
	if d.resume {
		flags |= os.O_APPEND
	}

	partFile, err := os.OpenFile(d.getPartFilename(filename, i), flags, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer partFile.Close()

	// 写入文件和更新进度条
	buf := make([]byte, 32*1024)
	_, err = io.CopyBuffer(io.MultiWriter(partFile, d.bar), resp.Body, buf)
	if err != nil {
		if err == io.EOF {
			return
		}
		log.Fatal(err)
	}
}

// 合并文件
func (d *Downloader) merge(filename string) error {
	destFile, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer destFile.Close()

	for i := 0; i < d.concurrency; i++ {
		partFileName := d.getPartFilename(filename, i)
		partFile, err := os.Open(partFileName)
		if err != nil {
			return err
		}
		io.Copy(destFile, partFile)
		partFile.Close()
		os.Remove(partFileName)
	}

	return nil
}

// getPartDir 部分文件存放的目录
func (d *Downloader) getPartDir(filename string) string {
	return strings.SplitN(filename, ".", 2)[0]
}

// getPartFilename 构造部分文件的名字 (文件名上加上序号, 方便后续合并)
func (d *Downloader) getPartFilename(filename string, partNum int) string {
	partDir := d.getPartDir(filename)
	return fmt.Sprintf("%s/%s-%d", partDir, filename, partNum)
}

// 整个文件下载
func (d *Downloader) singleDownload(strURL, filename string) error {
	resp, err := http.Get(strURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	d.setBar(int(resp.ContentLength))

	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	buf := make([]byte, 32*1024)
	_, err = io.CopyBuffer(io.MultiWriter(f, d.bar), resp.Body, buf)
	return err
}

// 设置进度条
func (d *Downloader) setBar(length int) {
	d.bar = progressbar.NewOptions(
		length,
		progressbar.OptionSetWriter(ansi.NewAnsiStdout()),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowCount(),
		progressbar.OptionShowBytes(true),
		progressbar.OptionSetWidth(50),
		progressbar.OptionSetDescription("downloading..."),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]=[reset]",
			SaucerHead:    "[green]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}),
	)
}
