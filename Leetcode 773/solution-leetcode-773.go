func helper(board [][]int) (enc int) {
	for _, r := range board {
		for _, c := range r {
			enc = enc<<4 | c
		}
	}
	return
}

func slidingPuzzle(board [][]int) int {

	type move struct {
		n, b int
	}

	m := make(map[int]bool, 720)
	q := make([]move, 0, 720)

	push := func(n, b int) {
		if !m[b] {
			q, m[b] = append(q, move{n, b}), true
		}
	}

	push(0, helper(board))

	hSwap := func(b int, l, r int) int {
		return b&^(l|r) | b<<4&l | b>>4&r
	}

	vSwap := func(b int, u, d int) int {
		return b&^(u|d) | b<<12&u | b>>12&d
	}

	for len(q) != 0 {
		t := q[0]
		q = q[1:]
		if t.b == 0x123450 {
			return t.n
		}
		switch n := t.n + 1; {
		case t.b&0x000_00f == 0:
			push(n, hSwap(t.b, 0x000_0f0, 0x000_00f))
			push(n, vSwap(t.b, 0x00f_000, 0x000_00f))
		case t.b&0x000_0f0 == 0:
			push(n, hSwap(t.b, 0x000_0f0, 0x000_00f))
			push(n, hSwap(t.b, 0x000_f00, 0x000_0f0))
			push(n, vSwap(t.b, 0x0f0_000, 0x000_0f0))
		case t.b&0x000_f00 == 0:
			push(n, hSwap(t.b, 0x000_f00, 0x000_0f0))
			push(n, vSwap(t.b, 0xf00_000, 0x000_f00))
		case t.b&0x00f_000 == 0:
			push(n, hSwap(t.b, 0x0f0_000, 0x00f_000))
			push(n, vSwap(t.b, 0x00f_000, 0x000_00f))
		case t.b&0x0f0_000 == 0:
			push(n, hSwap(t.b, 0x0f0_000, 0x00f_000))
			push(n, hSwap(t.b, 0xf00_000, 0x0f0_000))
			push(n, vSwap(t.b, 0x0f0_000, 0x000_0f0))
		case t.b&0xf00_000 == 0:
			push(n, hSwap(t.b, 0xf00_000, 0x0f0_000))
			push(n, vSwap(t.b, 0xf00_000, 0x000_f00))
		}
	}

	return -1
}