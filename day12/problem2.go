package main

import "fmt"

func applyGravity(pos []int, velocity []int) []int {
	// naieve n^2 approach -- we only have 4 masses
	updates := make([]int, len(velocity))
	for i := 0; i < len(velocity); i++ {
		for j := 0; j < len(velocity); j++ {
			if pos[j] < pos[i] {
				updates[i]--
			} else if pos[j] > pos[i] {
				updates[i]++
			}
		}
	}

	for i, update := range updates {
		velocity[i] += update
	}
	return velocity
}

func applyVelocity(pos []int, velocity []int) []int {
	for i, v := range velocity {
		pos[i] += v
	}
	return pos
}

func find1DPeriod(p []int, v []int) int {
	x0 := p[0]
	x1 := p[1]
	x2 := p[2]
	x3 := p[3]

	period := 0
	for {
		period++

		v = applyGravity(p, v)
		p = applyVelocity(p, v)

		if p[0] == x0 && p[1] == x1 && p[2] == x2 && p[3] == x3 &&
			v[0] == 0 && v[1] == 0 && v[2] == 0 && v[3] == 0 {
			break
		}

	}
	return period
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func main() {
	xs := []int{-19, -9, -4, 1}
	ys := []int{-4, 8, 5, 9}
	zs := []int{2, -16, -11, -13}

	vx := []int{0, 0, 0, 0}
	vy := []int{0, 0, 0, 0}
	vz := []int{0, 0, 0, 0}

	xp := find1DPeriod(xs, vx)
	yp := find1DPeriod(ys, vy)
	zp := find1DPeriod(zs, vz)

	period := lcm(xp, lcm(yp, zp))
	fmt.Println(period)
}
