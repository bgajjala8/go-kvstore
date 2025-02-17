package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Simple store acting as kv
type Store struct {
	data map[string]string
}

func NewStore() *Store {
	return &Store{data: make(map[string]string)}
}

func (s *Store) Set(key, value string) {
	s.data[key] = value
}

func (s *Store) Get(key string) string {
	val, exists := s.data[key]
	if !exists {
		fmt.Println("KEY NOT EXIST B")
	}
	return val
}

// TODO: Move to cli class to act as driver
func main() {
	store := NewStore()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		parts := strings.Fields(scanner.Text())
		if len(parts) == 0 {
			continue
		}

		switch parts[0] {
		case "set":
			if len(parts) < 3 {
				continue
			}
			store.Set(parts[1], strings.Join(parts[2:], " "))
			fmt.Println("OK")
		case "get":
			if len(parts) < 2 {
				continue
			}
			fmt.Println(store.Get(parts[1]))
		case "exit":
			return
		default:
			fmt.Println("Unknown command")
		}
	}
}
