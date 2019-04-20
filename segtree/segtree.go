package segtree

import (
"fmt"
"math"
)

type SegmentTree struct {
  root  *node
  start uint64
  end   uint64
}

func NewSegTree(start uint64, end uint64) (*SegmentTree, error) {
  if end < start {
    return nil, fmt.Errorf("invalid range: start (%d) greater than end (%d)", start, end)
  }

  if start < 0 {
    return nil, fmt.Errorf("invalid range: start (%d) must be greather than 0", start)
  }

  if end > math.MaxUint32 {
    return nil, fmt.Errorf("invalid range: end (%d) must be lesser than %d", end, math.MaxUint32)
  }

  return &SegmentTree{
    root: &node{
      left:  nil,
      right: nil,
      k:     0,
      lazy:  0,
      start: start,
      end:   end,
    },
    start: start,
    end:   end,
  }, nil
}

func (t *SegmentTree) AddRange(start uint64, end uint64) error {
  if end < start {
    return fmt.Errorf("invalid range: start (%d) greater than end (%d)", start, end)
  }

  if start < t.start {
    return fmt.Errorf("invalid range: start must be greather than %d", t.start)
  }

  if end > t.end {
    return fmt.Errorf("invalid range: end must be lesser than %d", t.end)
  }

  update(t.root, start, end)
  return nil
}

func (t *SegmentTree) Contains(value uint64) bool {
  if value < t.start || value > t.end {
    return false
  }

  return query(t.root, value)
}

type node struct {
  left  *node
  right *node

  k     uint64
  lazy  uint64
  start uint64
  end   uint64
}

func query(n *node, value uint64) bool {
  switch {
  case n == nil:
    return false
  case value < n.start || n.end < value:
    return false
  case n.k > 0 && n.start <= value && value <= n.end:
    return true
  default:
    if n.left != nil && n.left.start <= value && value <= n.left.end {
      return query(n.left, value)
    }

    // Commented code are hints, they are unnecessary
    // if n.right != nil && n.right.start <= value && value <= n.right.end {
    return query(n.right, value)
    // }

    // return false
  }
}

func update(n *node, start uint64, end uint64) {
  // No overlap
  if end < n.start || n.end < start {
    return
  }

  propagate(n)
  switch {
  // Full overlap
  case start <= n.start && n.end <= end:
    n.k += 1

    // If there is a range, it means there are child nodes from `propagate`
    if n.start < n.end && n.left != nil && n.right != nil {
      n.left.lazy += 1
      n.right.lazy += 1
    }
  // Partial overlap
  default:
    update(n.left, start, end)
    update(n.right, start, end)
  }
}

func propagate(n *node) {
  n.k += n.lazy
  mid := (n.start + n.end) >> 1

  // Skip unless there is a range
  if n.start < n.end {
    if n.left == nil || n.right == nil {
      n.left = &node{k: n.k, start: n.start, end: mid}
      n.right = &node{k: n.k, start: mid + 1, end: n.end}
    } else {
      n.left.lazy += n.lazy
      n.right.lazy += n.lazy
    }
  }
  n.lazy = 0
}
