package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)


var N = 10
var EPSILON = 1e-6


func PrintMatrix(m [][]float64) {
    for i := range m {
        for j := range m[i] {
            fmt.Printf("%.3f\t", m[i][j])
        }
        fmt.Println()
    }
    fmt.Println()
}

// Генерация случайной стохастической матрицы доверия

func GenerationStochasticMatrix(n int) [][]float64 {

    r := make([][]float64, n)
    for i := range r {
        r[i] = make([]float64, n)
    }

    for i := range r {
        s := 0
        for j := range r[i] {
            g := rand.Intn(1000)
            s += g
            r[i][j] = float64(g)
        }

        for j := range r[i] {
            r[i][j] = r[i][j] / float64(s)
        }
    }

    return r
} 

// Нахождение максимального и минимального значения в срезе

func MaxSlice(a []float64) float64 {
    max := a[0]
    for _, v := range a {
        if v > max {
            max = v
        }
    }

    return max
}

func MinSlice(a []float64) float64 {
    min := a[0]
    for _, v := range a {
        if v < min {
            min = v
        }
    }

    return min
} 

// Функция вычисления итогового мнения агентов, до схождения матрицы с заданной точностью epsilon

func Calc(epsilon float64, A [][]float64, x0 []float64) (int, []float64) {
    t := 1
    e := MaxSlice(x0) - MinSlice(x0)
    xt := make([]float64, N)
    copy(xt, x0)

    for e > epsilon {
        for i := 0; i < len(A); i++ {
            x := 0.0
            for j := 0; j < len(A); j++ {
                x += A[i][j] * xt[j]
                // s := 0.0
                // for k := 0; k < len(A); k++ {
                    // s += A[i][k] * B[k][j]
                // }
                // B[i][j] = s
            }
            xt[i] = x
        }
        e = MaxSlice(xt) - MinSlice(xt)
        t += 1
    }

    return t, xt
}


func main() {

    rand.New(rand.NewSource(time.Now().UnixNano()))

    A := GenerationStochasticMatrix(N)
    PrintMatrix(A)

    // Выбор случайного подмножества агентов для игрока A

    agents := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

    var agentNumA int
    for {
        agentNumA = rand.Intn(len(agents))
        if agentNumA != 0 && agentNumA != 10 {
            break
        }
    }

    // Выбор рандомного подмножества агентов для игрока B, непересекающегося с множеством агентов для игрока А

    var agentNumB int
    for {
        agentNumB = rand.Intn(len(agents) - agentNumA)
        if agentNumB == 0 {
            agentNumB = 1
            break
        } else {
            break
        }
    }


    agentsA := []int{}
    for agentNumA > 0 {
        idx := rand.Intn(len(agents))
        agentsA = append(agentsA, agents[idx])
        agents = append(agents[:idx], agents[idx + 1:]...)
        agentNumA--
    }
    sort.Slice(agentsA, func(i int, j int) bool {
        return agentsA[i] < agentsA[j]
    })
    fmt.Printf("agents A: %v\n", agentsA)

    agentsB := []int{}
    for agentNumB > 0 {
        idx := rand.Intn(len(agents))
        agentsB = append(agentsB, agents[idx])
        agents = append(agents[:idx], agents[idx + 1:]...)
        agentNumB--
    }

    sort.Slice(agentsB, func(i int, j int) bool {
        return agentsB[i] < agentsB[j]
    })
    fmt.Printf("agents B: %v\n", agentsB)
    fmt.Println()

    printXt := func(t int, xt []float64) {
        fmt.Printf("x(%d) = [ ", t)
        for _, v := range xt {
            fmt.Printf("%.3f  ", v)
        }
        fmt.Println("]")
        fmt.Println()
    }

    // Герератор вектора начальных мнений агентов

    x0 := make([]float64, N)
    for i := range x0 {
        x0[i] = float64(rand.Intn(50))
    }
    printXt(0, x0)


    // xt := make([]float64, N)
    // copy(xt, x0)

    // Вычисление результирующего мнения агентов

    // t := 1
    // e := MaxSlice(x0) - MinSlice(x0)
    // B := make([][]float64, len(A))
    // for i := range B {
    //     B[i] = make([]float64, len(A))
    // }
    //
    // for i := range A {
    //     for j := range A[i] {
    //         B[i][j] = A[i][j]
    //     }
    // }
    //
    // for e > epsilon {
    //     for i := 0; i < len(A); i++ {
    //         x := 0.0
    //         for j := 0; j < len(A); j++ {
    //             x += A[i][j] * xt[j]
    //             s := 0.0
    //             for k := 0; k < len(A); k++ {
    //                 s += A[i][k] * B[k][j]
    //             }
    //             B[i][j] = s
    //         }
    //         xt[i] = x
    //     }
    //     e = MaxSlice(xt) - MinSlice(xt)
    //     t += 1
    // }

    // fmt.Println()
    // PrintMatrix(B)


    t, xt := Calc(EPSILON, A, x0)
    printXt(t, xt)

    // генерация начального мнения u, v для агентов первого (А) и второго (В) игроков соответственно

    u := rand.Intn(100)
    v := rand.Intn(100) - 100

    fmt.Printf("u = %d\n", u)
    fmt.Printf("v = %d\n", v)
    fmt.Println()

    for _, x := range agentsA {
        x0[x - 1] = float64(u)
    }

    for _, x := range agentsB {
        x0[x - 1] = float64(v)
    }
    printXt(0, x0)

    t, xt = Calc(EPSILON, A, x0)
    printXt(t, xt)

    if math.Abs(float64(u) - xt[0]) > math.Abs(float64(v) - xt[0]) {
        fmt.Println("Победил второй игрок")
    } else {
        fmt.Println("Победил первый игрок")
    }
}
