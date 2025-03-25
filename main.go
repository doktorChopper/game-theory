package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
)

func vkMax(y [3]int, m [3][3]int, k int) (float64, []int) {
    
    r := [3]int{}

    for i := 0; i < len(m); i++ {
        s := 0
        for j := 0; j < len(m[0]); j++ {       
            s += m[i][j] * y[j]
        }
        r[i] = s
    }
    fmt.Print("x: ")
    fmt.Print(r)
    fmt.Print("         ")

    max := r[0]

    for i := 0; i < len(r); i++ {
        if r[i] > max {
            max = r[i] 
        }
    }

    eq := []int{}

    for i, v := range r {
        if v == max {
            eq = append(eq, i)
        }
    }

    return float64(max) / float64(k), eq
}

func vkMin(x [3]int, m [3][3]int, k int) (float64, []int) {

    r := [3]int{}

    for i := 0; i < len(m); i++ {
        s := 0
        for j := 0; j < len(m[0]); j++ {       
            s += m[j][i] * x[j]
        }
        r[i] = s
    }

    fmt.Print("y: ")
    fmt.Print(r)

    min := r[0]

    for i := 0; i < len(r); i++ {
        if r[i] < min {
            min = r[i] 
        }
    }

    eq := []int{}

    for i, v := range r {
        if v == min {
            eq = append(eq, i)
        }
    }

    return float64(min) / float64(k), eq
}

func randStrategy(a []int) int {
    return a[rand.Intn(len(a))]
}

func BrownRobinson(e float64) {

    C := [3][3]int {
        {0, 16, 19},
        {5, 19, 12},
        {16, 12, 7},
    }

    // C := [3][3]int {
    //     {2, 1, 3},
    //     {3, 0, 1},
    //     {1, 2, 1},
    // }

    yk := [3]int{1, 0, 0}
    xk := [3]int{1, 0, 0}

    fmt.Printf("%d:  ", 1)

    q, eq1 := vkMax(yk, C, 1)
    w, eq2 := vkMin(xk, C, 1)

    xk[randStrategy(eq1)] += 1
    yk[randStrategy(eq2)] += 1

    fmt.Printf("    %.2f  %.2f\n", q, w)

    epsilon := q - w

    i := 1
    for epsilon > e {
        
        fmt.Printf("%d:  ", i + 1)

        k, eq1 := vkMax(yk, C, i + 1)
        l, eq2 := vkMin(xk, C, i + 1)

        if k < q {
            q = k
        }

        if l > w {
            w = l
        }

        xk[randStrategy(eq1)] += 1
        yk[randStrategy(eq2)] += 1

        epsilon = q - w

        fmt.Printf("    %.2f  %.2f  epsilon: %.2f\n", k, l, epsilon)

        i++
    }

    fmt.Printf("~x = ( ")
    for _,v := range xk {
        fmt.Printf("%.3f ", float64(v) / float64(i + 1))
    }
    fmt.Println(")")

    fmt.Printf("~y = ( ")
    for _,v := range yk {
        fmt.Printf("%.3f ", float64(v) / float64(i + 1))
    }
    fmt.Println(")")

}

type stack []float64

func (s *stack) isEmpty() bool {
    return len(*s) == 0
}

func (s *stack) push(v float64) {
    *s = append(*s, v)
}

func (s *stack) pop() (float64, bool) {
    if s.isEmpty() {
        return 0, false
    }

    ret := (*s)[len(*s) - 1]
    *s = (*s)[:len(*s) - 1]

    return ret, true
}

func (s *stack) toSlice() []float64 {
    return *s
}

func subMat(m [][]float64, p int) [][]float64 {
    stacks := make([]stack, len(m))

    for n := range m {
        stacks[n] = stack{}
        for j := range m[n] {
            if j != p {
                stacks[n].push(m[n][j])
            }
        }
    }

    out := make([][]float64, len(m))
    for k := range stacks {
        out[k] = stacks[k].toSlice()
    }
    return out
}

func det(m [][]float64) (float64, error) {

    if len(m) != len(m[0]) {
        return 0.0, errors.New("not square matrix")
    }

    if len(m) == 1 {
        return (m[0][0]), nil
    }

    if len(m) == 2 {
        return (m[0][0] * m[1][1] - m[0][1] * m[1][0]), nil
    }

    s := 0.0

    for i := 0; i < len(m[0]); i++ {
        sm := subMat(m[1:][:], i)
        z, err := det(sm)
        if err == nil {
            if i % 2 != 0 {
                s -= m[0][i] * z
            } else {
                s += m[0][i] * z
            }
        }
    }

    return s, nil
}

func minor(m [][]float64, row, col int) [][]float64 {

    ret := make([][]float64, len(m) - 1)
    for i := range ret {
        ret[i] = make([]float64, len(m) - 1)
    }

    r := 0
    for i := 0; i < len(m); i++ {
        if i == row {
            continue
        }

        c := 0
        for j := 0; j < len(m); j++ {
            if j == col {
                continue
            }

            ret[r][c] = m[i][j]
            c++
        }
        r++
    }

    return ret
}

func algAddition(m [][]float64) [][]float64 {

    a := make([][]float64, len(m))
    for i := range a {
        a[i] = make([]float64, len(m))
    }

    for i := 0; i < len(m); i++ {
        for j := 0; j < len(m[0]); j++ {
            d, err := det(minor(m, i, j))
            if err != nil {
                fmt.Println(err.Error())
                return nil
            }
            am := math.Pow(-1, float64(i) + float64(j)) * d
            a[i][j] = am
        }
    }

    return a
}

func transposition(m [][]float64) [][]float64 {

    t := make([][]float64, len(m))
    for i := range m {
        t[i] = make([]float64, len(m))
    }

    for i := 0; i < len(m); i++ {
        for j := 0; j < len(m); j++ {
            t[j][i] = m[i][j]
        }
    }

    return t
}

func buildReverseMatrix(m [][]float64) [][]float64 {

    d, _ := det(m)
    a := transposition(algAddition(m))

    for i := range a {
        for j := range a {
            a[i][j] /= d
        }
    }

    return a
}

func main() {

    fmt.Printf("Brown-Robinson method\n\n")
    BrownRobinson(0.1)

    C := [][]float64 {
        {2, 1, 3},
        {3, 0, 1},
        {1, 2, 1},
    }

    // fmt.Println(det(C))
    // fmt.Println(transposition(algAddition(C)))
    fmt.Println(buildReverseMatrix(C))
}

