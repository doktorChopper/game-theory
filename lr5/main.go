package main

import (
	"fmt"
	"strconv"
	"strings"
)

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
    0b00000001: 3,
    0b00000010: 1,
    0b00000011: 2,
    0b00000100: 4,
    0b00000101: 4,
    0b00000110: 5,
    0b00000111: 8,
    0b00001000: 3,
    0b00001001: 5,
    0b00001010: 6,
    0b00001011: 7,
    0b00001100: 10,
    0b00001101: 11,
    0b00001110: 8,
    0b00001111: 13,
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
                    w := IntSliceToString(BinToArr(Union(a[i], a[j])))
                    fmt.Printf("%2d = %-17s%-5s%d + %d\n", vFunc[Union(a[i], a[j])], "v(" + w + ")", ">=", vFunc[a[i]], vFunc[a[j]])
                    fmt.Println(a[i])
                    fmt.Println(a[j])
                    return false
                }

                // w := IntSliceToString(BinToArr(Union(a[i], a[j])))
                // fmt.Printf("%2d = %-17s%-5s%d + %d\n", vFunc[Union(a[i], a[j])], "v(" + w + ")", ">=", vFunc[a[i]], vFunc[a[j]])
            }
        }
    }
    fmt.Println()

    return true
}

func CheckConvexity(a [len(I)]uint8) bool {

    for i := 0; i < len(a); i++ {
        for j := 0; j < len(a); j++ {
            if vFunc[Union(a[i], a[j])] + vFunc[Intersection(a[i], a[j])] < vFunc[a[i]] + vFunc[a[j]] {
                p := IntSliceToString(BinToArr(Union(a[i], a[j])))
                q := IntSliceToString(BinToArr(Intersection(a[i], a[j])))

                fmt.Printf("%-27s = %d + %d %s %d + %d\n","v(" + p + ") + v(" + q + ")", vFunc[Union(a[i], a[j])], vFunc[Intersection(a[i], a[j])], "<", vFunc[a[i]], vFunc[a[j]])
                return false
            }
            // p := IntSliceToString(BinToArr(Union(a[i], a[j])))
            // q := IntSliceToString(BinToArr(Intersection(a[i], a[j])))

            // fmt.Printf("%-27s = %d + %d %s %d + %d\n","v(" + p + ") + v(" + q + ")", vFunc[Union(a[i], a[j])], vFunc[Intersection(a[i], a[j])], ">=", vFunc[a[i]], vFunc[a[j]])
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

func CheckGroupRationalization(a []float64) bool {

    sum := 0.0

    for i, v := range a {
        if i == len(a) - 1 {
            fmt.Printf("%.2f ", v)
        } else {
            fmt.Printf("%.2f + ", v)
        }
        sum += v
    }

    if sum == float64(vFunc[uint8(len(I) - 1)]) {
        fmt.Printf("== %d\n", vFunc[uint8(len(I) - 1)])
        return true
    }

    fmt.Printf("!= %d\n", vFunc[uint8(len(I) - 1)])
    return false
}

func CheckIndividualRationalization(a []float64) bool {

    for i := range a {
        if a[i] < float64(vFunc[1 << i]) {
            fmt.Printf("x%d = %.2f < v({%d}) = %d\n\n", i + 1, a[i], i + 1, vFunc[1 << i])
            return false
        }
        fmt.Printf("x%d = %.2f >= v({%d}) = %d\n", i + 1, a[i], i + 1, vFunc[1 << i])
    }
    fmt.Println()

    return true
}

func BinToArr(b uint8) []int {

    r := []int{}
    
    for i := 0; i < N; i++ {
        if 1 << i & b != 0 {
            r = append(r, i + 1)
        }
    }

    return r
}

func IntSliceToString(a []int) string {

    str := make([]string, len(a))

    for i := range a {
        str[i] = strconv.Itoa(a[i])
    }

    return "{" + strings.Join(str, ", ") + "}"
}

func main() {

    fmt.Println()
    fmt.Println("===== Характеристическая функция =====")
    fmt.Println()

    for _, v := range I {
        fmt.Printf("%-15s%-5s%-10v\n", IntSliceToString(BinToArr(v)), "-->", vFunc[v])
    } 
    fmt.Println()
    
    fmt.Printf("Является супераддитивной: %v\n", CheckSuperAdditive(I))
    fmt.Println()
    fmt.Printf("Является выпуклой: %v\n", CheckConvexity(I))
    fmt.Println()

    vec := VectorShapley(I)

    fmt.Print("Вектор Шепли: [ ")
    for _, v := range vec {
        fmt.Printf("%.2f ", v)
    }
    fmt.Println("]")
    fmt.Println()

    fmt.Printf("Условие групповой рационализации: %v\n", CheckGroupRationalization(vec))
    fmt.Println()
    fmt.Printf("Условие индивидуальной рационализации: %v\n", CheckIndividualRationalization(vec))
    fmt.Println()
}
