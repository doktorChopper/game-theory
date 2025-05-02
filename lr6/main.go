package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)


var N = 10


func PrintMatrix(m [][]float64) {
    for i := range m {
        for j := range m[i] {
            fmt.Printf("%.3f\t", m[i][j])
        }
        fmt.Println()
    }
    fmt.Println()
}

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

func main() {

    rand.New(rand.NewSource(time.Now().UnixNano()))

    /*


    Test := [][]float64 {
        {0.111, 0.178, 0.003, 0.074, 0.152, 0.131, 0.145, 0.101, 0.104},
        {0.045, 0.125, 0.076, 0.086, 0.112, 0.11, 0.207, 0.125, 0.112},
        {0.006, 0.214, 0.095, 0.038, 0.038, 0.142, 0.162, 0.275, 0.032},
        {0.168, 0.164, 0.094, 0.094, 0.007, 0.105, 0.15, 0.075, 0.143},
        {0.046, 0.258, 0.093, 0.113, 0.07, 0.006, 0.18, 0.113, 0.122},
        {0.198, 0.175, 0.07, 0.039, 0.187, 0.022, 0.064, 0.12, 0.125},
        {0.196, 0.077, 0.134, 0.151, 0.047, 0.223, 0.139, 0.017, 0.017},
        {0.049, 0.201, 0.009, 0.099, 0.077, 0.296, 0.056, 0.006, 0.207},
        {0.068, 0.036, 0.155, 0.138, 0.148, 0.058, 0.226, 0.019, 0.15},
    }
    PrintMatrix(Test)

    TestXt := [9]float64{15, 14, 19, 8, 12, 1, 16, 11, 17}
    for t := 0; t < 11; t++ {
        for i := 0; i < len(TestXt); i++ {
            x := 0.0
            for j := 0; j < len(TestXt); j++ {
                x += Test[i][j] * TestXt[j] 
            }
            TestXt[i] = x
        }
    }

    fmt.Println(TestXt)

    */

    // TestX0 := []int {15, 14, 19, 8, 12, 1, 16, 11, 17}


    A := GenerationStochasticMatrix(N)
    PrintMatrix(A)

    // Выбор рандомного подмножества агентов для игрока A

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

    // Герератор вектора начальных мнений агентов

    x0 := make([]float64, N)
    for i := range x0 {
        x0[i] = float64(rand.Intn(50))
    }

    fmt.Printf("x(0) = [  ")
    for _, v := range x0 {
        fmt.Printf("%.3f  ", v)
    }
    fmt.Println("]")

    epsilon := 1e-6

    xt := make([]float64, N)
    copy(xt, x0)

    // Вычисление результирующего мнения агентов

    t := 1
    e := MaxSlice(x0) - MinSlice(x0)
    for e > epsilon {
        for i := 0; i < len(A); i++ {
            x := 0.0
            for j := 0; j < len(A); j++ {
                x += A[i][j] * xt[j]
            }
            xt[i] = x
        }
        e = MaxSlice(xt) - MinSlice(xt)
        t += 1
    }

    fmt.Printf("x(%d) = [  ", t)
    for _, v := range xt {
        fmt.Printf("%.3f  ", v)
    }
    fmt.Println("]")
    fmt.Println()

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

    fmt.Printf("x(0) = [  ")
    for _, v := range x0 {
        fmt.Printf("%.3f  ", v)
    }
    fmt.Println("]")

    t = 1
    e = MaxSlice(x0) - MinSlice(x0)
    copy(xt, x0)
    for e > epsilon {
        for i := 0; i < len(A); i++ {
            x := 0.0
            for j := 0; j < len(A); j++ {
                x += A[i][j] * xt[j]
            }
            xt[i] = x
        }
        e = MaxSlice(xt) - MinSlice(xt)
        t += 1
    }

    fmt.Printf("x(%d) = [  ", t)
    for _, v := range xt {
        fmt.Printf("%.3f  ", v)
    }
    fmt.Println("]")

    if math.Abs(float64(u) - xt[0]) > math.Abs(float64(v) - xt[0]) {
        fmt.Println("Победил второй игрок")
    } else {
        fmt.Println("Победил первый игрок")
    }
}
