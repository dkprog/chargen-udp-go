package main

import (
	"net"
	"regexp"
	"testing"
	"time"
)

func TestAsciiTableLength(t *testing.T) {
	ascii := getASCIITable()
	desiredLength := 95

	if len(ascii) != desiredLength {
		t.Fatalf("ASCII has %v elements, want %v", len(ascii), desiredLength)
	}
}

func TestAsciiTableContent(t *testing.T) {
	ascii := getASCIITable()
	pattern := "[ -~]"

	matched, err := regexp.Match(pattern, ascii)

	if err != nil {
		t.Fatalf("Could not match pattern %s: %s", pattern, err)
	}

	if !matched {
		t.Fatalf("ASCII table does not match to pattern %s", pattern)
	}
}

func TestAsciiStreamLength(t *testing.T) {
	stream := getASCIIStream(0)
	desiredLength := 72

	if len(stream) != desiredLength {
		t.Fatalf("ASCII stream has %v elements, want %v", len(stream), desiredLength)
	}
}

func TestAsciiStreamShift(t *testing.T) {

	a := getASCIIStream(0)
	b := getASCIIStream(1)

	for i := 1; i < len(b)-1; i++ {
		if a[i] != b[i-1] {
			t.Fatalf("ASCII stream doesn't shift. Want a[%d] = b[%d]. Got %v and %v", i, i-1, a[i], b[i-1])
		}
	}
}

func TestServe(t *testing.T) {
	go serve(3019)

	conn, err := net.ListenPacket("udp", ":0")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	dst, err := net.ResolveUDPAddr("udp", "127.0.0.1:3019")
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(1 * time.Second)
	_, err = conn.WriteTo([]byte{'a'}, dst)
	if err != nil {
		t.Fatal(err)
	}
	buf := make([]byte, 1024)
	conn.SetReadDeadline(time.Now().Add(1 * time.Second))
	n, _, err := conn.ReadFrom(buf)
	if err != nil {
		t.Fatal(err)
	}
	desiredLength := 72
	if n != desiredLength {
		t.Fatalf("ASCII stream has %v elements, want %v", n, desiredLength)
	}
}
