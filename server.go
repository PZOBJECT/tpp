package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"
)

const (
	listenAddr = ":28080"
)

var (
	cliNum int
	cliMu  sync.Mutex
)

func nextID() int {
	cliMu.Lock()
	defer cliMu.Unlock()
	cliNum++
	return cliNum
}

// handleConn 处理单个客户端连接的完整生命周期
func handleConn(conn net.Conn, id int) {
	defer func() {
		conn.Close()
		log.Printf("[%d] 客户端断开连接", id)
	}()

	log.Printf("[%d] 新客户端接入: %s", id, conn.RemoteAddr())

	// 发送欢迎消息
	welcome := fmt.Sprintf("欢迎连接到 Hermes TCP 服务 (ID: %d)\n", id)
	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	if _, err := io.WriteString(conn, welcome); err != nil {
		log.Printf("[%d] 发送欢迎消息失败: %v", id, err)
		return
	}

	r := bufio.NewReader(conn)
	for {
		// 设置读取超时（300 秒无数据断开）
		conn.SetReadDeadline(time.Now().Add(300 * time.Second))

		msg, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				log.Printf("[%d] 客户端正常关闭连接", id)
			} else if neterr, ok := err.(net.Error); ok && neterr.Timeout() {
				log.Printf("[%d] 读取超时，断开连接", id)
				io.WriteString(conn, "连接超时，服务端断开\n")
			} else {
				log.Printf("[%d] 读取错误: %v", id, err)
			}
			return
		}

		msg = strings.TrimSpace(msg)
		if msg == "" {
			continue
		}

		log.Printf("[%d] 收到消息: %s", id, msg)

		// 处理特殊命令
		switch strings.ToLower(msg) {
		case "exit", "quit":
			io.WriteString(conn, "再见！\n")
			return
		case "ping":
			io.WriteString(conn, "pong\n")
			continue
		case "time":
			io.WriteString(conn, time.Now().Format(time.RFC3339)+"\n")
			continue
		case "echo off":
			continue
		}

		// 默认：回声
		resp := fmt.Sprintf("收到: %s\n", msg)
		conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
		if _, err := io.WriteString(conn, resp); err != nil {
			log.Printf("[%d] 发送响应失败: %v", id, err)
			return
		}
	}
}

func main() {
	l, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatalf("监听失败 %s: %v", listenAddr, err)
	}
	defer l.Close()

	log.Printf("✅ Hermes TCP 服务已启动，监听 %s", listenAddr)

	// 优雅退出：监听 SIGINT / SIGTERM
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigCh
		log.Println("🛑 收到退出信号，停止接受新连接...")
		l.Close()
	}()

	for {
		conn, err := l.Accept()
		if err != nil {
			// 关闭后 Accept 返回错误，正常退出
			log.Println("TCP 服务已关闭")
			return
		}

		id := nextID()
		go handleConn(conn, id)
	}
}
