package segtree

import (
  "strconv"
  "strings"
)

func IP2Uint(ipv4 string) uint32 {
  bytes := strings.Split(ipv4, ".")

  b0, _ := strconv.Atoi(bytes[0])
  b1, _ := strconv.Atoi(bytes[1])
  b2, _ := strconv.Atoi(bytes[2])
  b3, _ := strconv.Atoi(bytes[3])

  var sum uint32
  sum += uint32(b0) << 24
  sum += uint32(b1) << 16
  sum += uint32(b2) << 8
  sum += uint32(b3)

  return sum
}

func Uint2IP(ipv4 uint32) string {
  b0 := strconv.FormatUint(uint64(ipv4>>24)&0xff, 10)
  b1 := strconv.FormatUint(uint64(ipv4>>16)&0xff, 10)
  b2 := strconv.FormatUint(uint64(ipv4>>8)&0xff, 10)
  b3 := strconv.FormatUint(uint64(ipv4>>0)&0xff, 10)

  return b0 + "." + b1 + "." + b2 + "." + b3
}
