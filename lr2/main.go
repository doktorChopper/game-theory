package main

import (
	"errors"
	"fmt"
	"math"

	brownrobinson "github.com/doktorChopper/game-theory/brown-robinson"
	"github.com/doktorChopper/go-matrix/matrix"
	"github.com/jedib0t/go-pretty/v6/table"
)

func Hxy(x, y float64) float64 {

    // H := [][]float64 {
    //     {0.0, -(72.0 / 25.0), 3.0 / 2.0},
    //     {-(18.0 / 50.0), 18.0 / 5.0, 0.0},
    //     {-3.0, 0.0, 0.0},
    // }

    H := [][]float64 {
        {0.0, -(81.0 / 5.0), 9.0},
        {-(9.0 / 5.0), 18.0, 0.0},
        {-3.0, 0.0, 0.0},
    }

    ret := 0.0

    for i := range len(H) {
        for j := range len(H[0]) {
            ret += H[i][j] * math.Pow(x, float64(i)) * math.Pow(y, float64(j))
        }
    }

    return ret
}

func HN(n int) *matrix.Matrix {

    ret := matrix.NewMatrixNM(n + 1, n + 1)

    var v float64

    for i := 0; i <= n; i++ {
        for j := 0; j <= n; j++ {
            v = Hxy(float64(i) / float64(n), float64(j) / float64(n))
            ret.SetAt(i, j, v)
        }
    }

    return ret
}

func maxMin(m *matrix.Matrix) (float64, int) {
    a := []float64{}

    var min float64

    for i := 0; i < m.Rows(); i++ {

        min = m.GetAt(i, 0)
        for j := 0; j < m.Cols(); j++ {
            if m.GetAt(i, j) < min {
                min = m.GetAt(i, j)
            }
        }

        a = append(a, min)
    }

    maxMin := a[0]
    idx := 0
    for i, v := range a {
        if maxMin < v {
            maxMin = v
            idx = i
        }
    }

    return maxMin, idx
}

func minMax(m *matrix.Matrix) (float64, int) {

    a := []float64{}
    var max float64

    for j := 0; j < m.Cols(); j++ {

        max = m.GetAt(0, j)
        for i := 0; i < m.Rows(); i++ {
            if m.GetAt(i, j) > max {
                max = m.GetAt(i, j)
            }
        }

        a = append(a, max)
    }

    minMax := a[0]
    idx := 0
    for i, v := range a {
        if minMax > v {
            minMax = v
            idx = i
        }
    }

    return minMax, idx
}

func maxInd(a []float64) (int, error) {
    
    if len(a) == 0 {
        return 0, errors.New("empty array")
    }

    var ret int
    max := a[0]

    for i, v := range a {
        if max < v {
            max = v
            ret = i
        }
    }

    return ret, nil
}

func checkSaddlePoint(m *matrix.Matrix) (int, int, bool) {

    a, ai := maxMin(m)
    b, bi := minMax(m)

    if a == b {
        return ai, bi, true
    }

    return -1, -1, false
}

func initTable(t table.Writer) {
    t.SetTitle("Brown-Robinson method")
    t.AppendHeader(table.Row{"#", "_V", "V_", "EPSILON"})
}

func fxy(x, y float64, c []float64) float64 {
    return c[0] * x * x + c[1] * y * y + c[2] * x * y + c[3] * x + c[4] * y
}

type F func(float64, float64, []float64) float64

// func derfx(f F, x, y, h float64, c []float64) float64 {
//
//     return 
// }
//
// func derfy(f F, x, y, h float64, c []float64) float64 {
//
// }

func AnalyticalSolution() {

    vec := []float64 {-3.0, 9.0, 18.0, -(9.0 / 5.0), -(81.0 / 5.0)}
    // vec := []float64 {-(3.0), 3.0 / 2.0, 18.0 / 5.0, -(18.0 / 50.0), -(72.0 / 25.0)}

    fmt.Printf("H = %v\n", vec)
    fmt.Println()
    fmt.Printf("H(x, y) = %.2fx\u00B2 + %.2fy\u00B2 + %.2fxy + %.2fx + %.2fy\n", vec[0], vec[1], vec[2], vec[3], vec[4])
    fmt.Println()
    fmt.Printf("Hx = 2 * %.2fx + %.2fy + %.2f\n", vec[0], vec[2], vec[3])
    fmt.Printf("Hy = 2 * %.2fy + %.2fx + %.2f\n", vec[1], vec[2], vec[4])
    fmt.Println()
    fmt.Printf("Hxx = 2a = %.2f < 0\n", 2.0 * vec[0])
    fmt.Printf("Hyy = 2b = %.2f > 0\n", 2.0 * vec[1])
    fmt.Println()

    tmp := (vec[2] * vec[3] - 2 * vec[0] * vec[4]) / (4 * vec[0] * vec[1] - vec[2] * vec[2])

    if 2 * vec[0] < 0 && 2 * vec[1] > 0 {
        fmt.Println("Игра является выпукло-вогнутой")
    } else {
        fmt.Println("Игра НЕ является выпукло-вогнутой")
    }
    fmt.Println()

    x := func(y float64) float64 {
        if y < -(vec[3] / vec[2]) {
            return 0
        }
        return -(vec[2] * y + vec[3]) / (2 * vec[0])
    }

    y := func(x float64) float64 {
        if x > -(vec[4] / vec[2]) {
            return 0
        }
        return -(vec[2] * x + vec[4]) / (2 * vec[1])
    }

    xa := x(tmp)
    ya := y(xa)
    fmt.Printf("x = %.3f\n", x(ya))
    fmt.Printf("y = %.3f\n", y(xa))
    fmt.Printf("H = %.3f\n", Hxy(x(ya), y(xa)))
}

func PrintMatrix(m *matrix.Matrix) {
    for i := range m.Rows() {
        for j := range m.Cols() {
            fmt.Printf("%-5.3f\t", m.GetAt(i, j))
        }
        fmt.Println()
    }
    fmt.Println()
}

func average(s []float64) float64 {
    if len(s) == 0 {
        return 0
    }

    var sum float64

    for _, v := range s {
        sum += v
    }

    return sum / float64(len(s))
}

func standartDeviation(s []float64) float64 {

    if len(s) == 0 {
        return 0
    }

    avr := average(s)
    var sum float64

    for _, v := range s {
        deviation := v - avr
        sum += deviation * deviation
    }

    variance := sum / float64(len(s))
    return math.Sqrt(variance)
}

func main() {

    t := table.NewWriter()
    initTable(t)

    fmt.Println()
    fmt.Println("********** Numerical solution **********")
    fmt.Println()

    sH := []float64{}

    n := 2
    for ; n <= 1000; n++ {
        if n < 12 {
            fmt.Printf("N = %d\n", n)
        }
        m := HN(n)

        ai, bi, flag := checkSaddlePoint(m)

        if flag {

            x := float64(ai) / float64(n)
            y := float64(bi) / float64(n)

            if n < 12 {
                PrintMatrix(m)
                fmt.Println("Есть седловая точка")
                fmt.Printf("x = %.3f\n", x)
                fmt.Printf("y = %.3f\n", y)
                fmt.Printf("H = %.3f\n", Hxy(x, y))
            }

            sH = append(sH, Hxy(x, y))
            if len(sH) > 2 && standartDeviation(sH[len(sH)-5:]) < 0.001 {
                break
            }
        } else {
            x, y, _ := brownrobinson.BrownRobinsonMethod(0.001, m, t)
            xmi, _ := maxInd(x)
            ymi, _ := maxInd(y)

            xf := float64(xmi) / float64(n)
            yf := float64(ymi) / float64(n)

            if n < 12 {
                PrintMatrix(m)
                fmt.Println("Нет седловой точки")

                fmt.Printf("x = %.3f\n", xf)
                fmt.Printf("y = %.3f\n", yf)
                fmt.Printf("H = %.3f\n", Hxy(xf, yf))
            }

            sH = append(sH, Hxy(xf, yf))
            if len(sH) > 5 && standartDeviation(sH[len(sH)-5:]) < 0.001 {
                break
            }
        }

        if n < 12 {
            fmt.Println()
            fmt.Println()
        }
    }

    fmt.Printf("Всего итераций: N = %d\n\n", n)


    fmt.Println("********** Analytical solution **********")
    fmt.Println()
    AnalyticalSolution()
}
