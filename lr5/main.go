package main

import "fmt"

// Test

// var vFunc = map[uint8]int{
//     0b00000000: 0,
//     0b00000001: 1,
//     0b00000010: 1,
//     0b00000011: 3,
//     0b00000100: 1,
//     0b00000101: 3,
//     0b00000110: 3,
//     0b00000111: 4,
// }

// var I = [...]uint8 {
//     0b00000000,
//     0b00000001,
//     0b00000010,
//     0b00000011,
//     0b00000100,
//     0b00000101,
//     0b00000110,
//     0b00000111,
// }

// Metr

var vFunc = map[uint8]int {
    0b00000000: 0,
    0b00000001: 2,
    0b00000010: 1,
    0b00000011: 4,
    0b00000100: 1,
    0b00000101: 4,
    0b00000110: 2,
    0b00000111: 7,
    0b00001000: 2,
    0b00001001: 4,
    0b00001010: 4,
    0b00001011: 8,
    0b00001100: 4,
    0b00001101: 8,
    0b00001110: 6,
    0b00001111: 10,
}

var I = [...]uint8 {
    0b00000000,
    0b00000001,
    0b00000010,
    0b00000011,
    0b00000100,
    0b00000101,
    0b00000110,
    0b00000111,
    0b00001000,
    0b00001001,
    0b00001010,
    0b00001011,
    0b00001100,
    0b00001101,
    0b00001110,
    0b00001111,
}

var N = 4

func HasIntersection(a, b uint8) bool {
    return a & b != 0b00000000 
}

func Union(a, b uint8) uint8 {
    return a | b
}

func Intersection(a, b uint8) uint8 {
    return a & b
}

func CheckSuperAdditive(a [len(I)]uint8) bool {

    for i := 0; i < len(a); i++ {
        for j := i + 1; j < len(a); j++ {
            if !HasIntersection(a[i], a[j]) {
                if vFunc[Union(a[i], a[j])] < vFunc[a[i]] + vFunc[a[j]] {
                    return false
                }
            }
        }
    }

    return true
}

func CheckConvexity(a [len(I)]uint8) bool {

    for i := 0; i < len(a); i++ {
        for j := 0; j < len(a); j++ {
            if vFunc[Union(a[i], a[j])] + vFunc[Intersection(a[i], a[j])] < vFunc[a[i]] + vFunc[a[j]] {
                return false
            }
        }
    }

    return true
}

func Factorial(n int) int {

    r := 1
    for n > 0 {
        r *= n
        n--
    }

    return r
}

func CountOne(b uint8) int {

    count := 0

    for i := 0; i < N; i++ {
        if 1 << i & b != 0 {
            count++
        }
    }

    return count
}

func VectorShapley(a [len(I)]uint8) []float64 {

    r := []float64{}

    n := float64(Factorial(N))

    for i := 0; i < N; i++ {
        xiv := 0.0
        for _, s := range a {
            if 1 << i & s != 0 {
                xiv += float64(Factorial(CountOne(s) - 1) * Factorial(N - CountOne(s)) * (vFunc[s] - vFunc[1 << i ^ s]))
            }
        }

        r = append(r, xiv / n)

    } 

    return r
}

func main() {
    
    fmt.Println(CheckSuperAdditive(I))
    fmt.Println(CheckConvexity(I))

    fmt.Println(VectorShapley(I))

    fmt.Println("Hello, World!")
}
