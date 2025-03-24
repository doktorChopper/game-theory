package main

import (
	"fmt"
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

}

func main() {
    BrownRobinson(0.1)
}

