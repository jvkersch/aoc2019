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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func totalEnergy(p []int, v []int) int {
	P := abs(p[0]) + abs(p[1]) + abs(p[2])
	V := abs(v[0]) + abs(v[1]) + abs(v[2])
	return P * V
}

func main() {
	xs := []int{-19, -9, -4, 1}
	ys := []int{-4, 8, 5, 9}
	zs := []int{2, -16, -11, -13}

	vx := []int{0, 0, 0, 0}
	vy := []int{0, 0, 0, 0}
	vz := []int{0, 0, 0, 0}

	for i := 0; i < 1000; i++ {
		vx = applyGravity(xs, vx)
		vy = applyGravity(ys, vy)
		vz = applyGravity(zs, vz)

		xs = applyVelocity(xs, vx)
		ys = applyVelocity(ys, vy)
		zs = applyVelocity(zs, vz)
	}

	E := 0
	for i := 0; i < len(xs); i++ {
		p := []int{xs[i], ys[i], zs[i]}
		v := []int{vx[i], vy[i], vz[i]}
		E += totalEnergy(p, v)
	}
	fmt.Println(E)
}
