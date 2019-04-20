package segtree

import (
  "fmt"
  "math"
  "math/rand"
  "strconv"
  "strings"
  "testing"

  "github.com/stretchr/testify/assert"
)

func TestTree(t *testing.T) {
  segt, _ := NewSegTree(0, math.MaxUint32)

  ips := []struct {
    start string
    end   string
  }{
    {"127.0.0.1", "127.0.0.2"},
    {"127.0.0.1", "127.0.0.3"},
    {"127.1.0.0", "127.2.0.0"},
    {"192.169.169.1", "192.169.169.255"},
    {"162.221.207.0", "162.221.207.255"},
    {"223.27.168.0", "223.27.175.255"},
    {"1.0.0.0", "2.0.0.255"},
  }

  for _, ipr := range ips {
    t.Logf("Adding %s to %s", ipr.start, ipr.end)
    _ = segt.AddRange(uint64(IP2Uint(ipr.start)), uint64(IP2Uint(ipr.end)))
  }

  assert.Equal(t, true, segt.Contains(uint64(IP2Uint("127.0.0.1"))), "127.0.0.1")
  assert.Equal(t, true, segt.Contains(uint64(IP2Uint("127.0.0.2"))), "127.0.0.2")
  assert.Equal(t, true, segt.Contains(uint64(IP2Uint("127.0.0.3"))), "127.0.0.3")
  assert.Equal(t, false, segt.Contains(uint64(IP2Uint("127.0.0.4"))), "127.0.0.4")
  assert.Equal(t, true, segt.Contains(uint64(IP2Uint("127.1.0.0"))), "127.1.0.0")
  assert.Equal(t, true, segt.Contains(uint64(IP2Uint("127.1.255.255"))), "127.1.255.255")
  assert.Equal(t, true, segt.Contains(uint64(IP2Uint("127.2.0.0"))), "127.2.0.0")
  assert.Equal(t, false, segt.Contains(uint64(IP2Uint("127.2.0.1"))), "127.2.0.1")
  assert.Equal(t, false, segt.Contains(uint64(IP2Uint("192.169.169.0"))), "192.169.169.0")
  assert.Equal(t, true, segt.Contains(uint64(IP2Uint("192.169.169.1"))), "192.169.169.1")
  assert.Equal(t, true, segt.Contains(uint64(IP2Uint("192.169.169.255"))), "192.169.169.255")
  assert.Equal(t, true, segt.Contains(uint64(IP2Uint("192.169.169.169"))), "192.169.169.169")
  assert.Equal(t, true, segt.Contains(uint64(IP2Uint("162.221.207.0"))), "162.221.207.0")
  assert.Equal(t, false, segt.Contains(uint64(IP2Uint("162.221.206.255"))), "162.221.206.255")
  assert.Equal(t, true, segt.Contains(uint64(IP2Uint("162.221.207.255"))), "162.221.207.255")
  assert.Equal(t, false, segt.Contains(uint64(IP2Uint("162.221.208.0"))), "162.221.208.0")
  assert.Equal(t, true, segt.Contains(uint64(IP2Uint("223.27.168.0"))), "223.27.168.0")
  assert.Equal(t, false, segt.Contains(uint64(IP2Uint("223.27.167.255"))), "223.27.167.255")
  assert.Equal(t, true, segt.Contains(uint64(IP2Uint("223.27.169.255"))), "223.27.169.255")
  assert.Equal(t, true, segt.Contains(uint64(IP2Uint("223.27.174.2"))), "223.27.174.2")
  assert.Equal(t, true, segt.Contains(uint64(IP2Uint("223.27.175.254"))), "223.27.175.254")
  assert.Equal(t, true, segt.Contains(uint64(IP2Uint("223.27.175.255"))), "223.27.175.255")
  assert.Equal(t, false, segt.Contains(uint64(IP2Uint("223.27.176.0"))), "223.27.176.0")
  assert.Equal(t, true, segt.Contains(uint64(IP2Uint("1.0.1.0"))), "1.0.1.0")
  assert.Equal(t, true, segt.Contains(uint64(IP2Uint("1.2.1.0"))), "1.2.1.0")
}

func TestTree2(t *testing.T) {
  end := 1000000
  segt, _ := NewSegTree(0, uint64(end)-1)

  for i := 0; i < end; i++ {
    _ = segt.AddRange(uint64(i), uint64(i))
  }

  for i := 0; i < end; i++ {
    assert.Equal(t, true, segt.Contains(uint64(i)))
  }
}

func TestTree3(t *testing.T) {
  end := 1000000
  segt, _ := NewSegTree(0, uint64(end)-1)

  for i := 0; i < end; i++ {
    if rand.Intn(end) < 500 {
      _ = segt.AddRange(uint64(i), uint64(i))
      assert.Equal(t, true, segt.Contains(uint64(i)))
    } else {
      assert.Equal(t, false, segt.Contains(uint64(i)))
    }
  }
}

func BenchmarkTreeAddRange(b *testing.B) {
  segt, _ := NewSegTree(0, math.MaxUint32)
  b.ResetTimer()

  for i := 0; i < b.N; i++ {
    // Hmm...
    i, j := rand.Uint32(), rand.Uint32()
    if i < j {
      _ = segt.AddRange(uint64(i), uint64(j))
    } else {
      _ = segt.AddRange(uint64(j), uint64(i))
    }
  }
}

func BenchmarkTreeQuery(b *testing.B) {
  segt, _ := NewSegTree(0, math.MaxUint32)
  for i := 0; i < 1000000; i++ {
    i, j := rand.Uint32(), rand.Uint32()
    if i < j {
      _ = segt.AddRange(uint64(i), uint64(j))
    } else {
      _ = segt.AddRange(uint64(j), uint64(i))
    }
  }

  b.ResetTimer()
  for i := 0; i < b.N; i++ {
    segt.Contains(uint64(0))
  }
}

func BenchmarkTreeQueryParallel(b *testing.B) {
  segt, _ := NewSegTree(0, math.MaxUint32)
  for i := 0; i < 1000000; i++ {
    i, j := rand.Uint32(), rand.Uint32()
    if i < j {
      _ = segt.AddRange(uint64(i), uint64(j))
    } else {
      _ = segt.AddRange(uint64(j), uint64(i))
    }
  }

  b.ResetTimer()
  b.RunParallel(func(pb *testing.PB) {
    for pb.Next() {
      segt.Contains(uint64(0))
    }
  })
}

var ipBlacklistMap = make(map[string]bool)

func OldIsBannedIp(ip string) bool {
  ipParts := strings.Split(ip, ".")
  ipLength := len(ip)

  for i := len(ipParts) - 1; i >= 0; i-- {
    if _, contains := ipBlacklistMap[ip[0:ipLength]]; contains {
      return true
    }
    ipLength -= len(ipParts[i]) + 1
  }
  return false
}

func OldAddIpRange(ip1 string, ip2 string) {
  ip1Parts := strings.Split(ip1, ".")
  ip2Parts := strings.Split(ip2, ".")
  for i := len(ip1Parts) - 1; i >= 0; {
    if ip1Parts[i] == "0" && ip2Parts[i] == "255" {
      i--
    } else if ip1Parts[i] == ip2Parts[i] {
      var ipMask string
      for j := 0; j < i; j++ {
        ipMask += ip1Parts[j]
        ipMask += "."
      }
      ipMask += ip1Parts[i]
      ipBlacklistMap[ipMask] = true
      break
    } else {
      ip1Val, err := strconv.Atoi(ip1Parts[i])
      ip2Val, err := strconv.Atoi(ip2Parts[i])

      if err != nil {
        fmt.Println("error")
      }

      var ipMask string
      for j := 0; j <= i-1; j++ {
        ipMask += ip1Parts[j]
        ipMask += "."
      }
      for ip1Val <= ip2Val {
        ipBlacklistMap[ipMask+strconv.Itoa(ip1Val)] = true
        ip1Val++
      }
      break
    }
  }
}

func TestMap(t *testing.T) {
  ips := []struct {
    start string
    end   string
  }{
    {"127.0.0.1", "127.0.0.2"},
    {"127.0.0.1", "127.0.0.3"},
    {"127.1.0.0", "127.2.0.0"},
    {"192.169.169.1", "192.169.169.255"},
    {"162.221.207.0", "162.221.207.255"},
    {"223.27.168.0", "223.27.175.255"},
    {"1.0.0.0", "2.0.0.255"},
  }

  for _, ipr := range ips {
    t.Logf("Adding %s to %s", ipr.start, ipr.end)
    OldAddIpRange(ipr.start, ipr.end)
  }

  assert.Equal(t, true, OldIsBannedIp("127.0.0.1"), "127.0.0.1")
  assert.Equal(t, true, OldIsBannedIp("127.0.0.2"), "127.0.0.2")
  assert.Equal(t, true, OldIsBannedIp("127.0.0.3"), "127.0.0.3")
  assert.Equal(t, false, OldIsBannedIp("127.0.0.4"), "127.0.0.4")
  assert.Equal(t, true, OldIsBannedIp("127.1.0.0"), "127.1.0.0")
  assert.Equal(t, true, OldIsBannedIp("127.1.255.255"), "127.1.255.255")
  assert.Equal(t, true, OldIsBannedIp("127.2.0.0"), "127.2.0.0")
  assert.Equal(t, false, OldIsBannedIp("127.2.0.1"), "127.2.0.1")
  assert.Equal(t, false, OldIsBannedIp("192.169.169.0"), "192.169.169.0")
  assert.Equal(t, true, OldIsBannedIp("192.169.169.1"), "192.169.169.1")
  assert.Equal(t, true, OldIsBannedIp("192.169.169.255"), "192.169.169.255")
  assert.Equal(t, true, OldIsBannedIp("192.169.169.169"), "192.169.169.169")
  assert.Equal(t, true, OldIsBannedIp("162.221.207.0"), "162.221.207.0")
  assert.Equal(t, false, OldIsBannedIp("162.221.206.255"), "162.221.206.255")
  assert.Equal(t, true, OldIsBannedIp("162.221.207.255"), "162.221.207.255")
  assert.Equal(t, false, OldIsBannedIp("162.221.208.0"), "162.221.208.0")
  assert.Equal(t, true, OldIsBannedIp("223.27.168.0"), "223.27.168.0")
  assert.Equal(t, false, OldIsBannedIp("223.27.167.255"), "223.27.167.255")
  assert.Equal(t, true, OldIsBannedIp("223.27.169.255"), "223.27.169.255")
  assert.Equal(t, true, OldIsBannedIp("223.27.174.2"), "223.27.174.2")
  assert.Equal(t, true, OldIsBannedIp("223.27.175.254"), "223.27.175.254")
  assert.Equal(t, true, OldIsBannedIp("223.27.175.255"), "223.27.175.255")
  assert.Equal(t, false, OldIsBannedIp("223.27.176.0"), "223.27.176.0")
  assert.Equal(t, true, OldIsBannedIp("1.0.1.0"), "1.0.1.0")
  assert.Equal(t, true, OldIsBannedIp("1.2.1.0"), "1.2.1.0")
  // assert.Equal(t, false, OldIsBannedIp("rex"), "rex")
  // assert.Equal(t, false, OldIsBannedIp("rex"), "rex")
}

func BenchmarkMapInsert(b *testing.B) {
  // Prep array
  starts := make([]string, 0, b.N)
  ends := make([]string, 0, b.N)
  for i := 0; i < b.N; i++ {
    i, j := rand.Uint32(), rand.Uint32()
    if i < j {
      starts = append(starts, Uint2IP(i))
      ends = append(ends, Uint2IP(j))
    } else {
      starts = append(starts, Uint2IP(j))
      ends = append(ends, Uint2IP(i))
    }
  }

  b.Logf("Generated an array of length %d", b.N)
  b.ResetTimer()

  for i := 0; i < b.N; i++ {
    OldAddIpRange(starts[i], ends[i])
  }
}

func BenchmarkMapQuery(b *testing.B) {
  // Generate ip
  for i := 0; i < 1000000; i++ {
    i, j := rand.Uint32(), rand.Uint32()
    if i < j {
      OldAddIpRange(Uint2IP(i), Uint2IP(j))
    } else {
      OldAddIpRange(Uint2IP(j), Uint2IP(i))
    }
  }

  b.ResetTimer()
  for i := 0; i < b.N; i++ {
    OldIsBannedIp("0.0.0.0")
  }
}

func BenchmarkMapQueryParallel(b *testing.B) {
  // Generate ip
  for i := 0; i < 1000000; i++ {
    i, j := rand.Uint32(), rand.Uint32()
    if i < j {
      OldAddIpRange(Uint2IP(i), Uint2IP(j))
    } else {
      OldAddIpRange(Uint2IP(j), Uint2IP(i))
    }
  }

  b.ResetTimer()
  b.RunParallel(func(pb *testing.PB) {
    for pb.Next() {
      OldIsBannedIp("0.0.0.0")
    }
  })
}
