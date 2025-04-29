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

    H := [][]float64 {
        {0.0, -(72.0 / 25.0), 3.0 / 2.0},
        {-(18.0 / 50.0), 18.0 / 5.0, 0.0},
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


}

func main() {

    t := table.NewWriter()
    initTable(t)

    for n := 2; n <= 10; n++ {
        m := HN(n)
        m.Display()

        ai, bi, flag := checkSaddlePoint(m)

        if flag {
            fmt.Println("Есть седловая точка")

            x := float64(ai) / float64(n)
            y := float64(bi) / float64(n)

            fmt.Printf("x = %.3f\n", x)
            fmt.Printf("y = %.3f\n", y)
            fmt.Printf("H = %.3f\n", Hxy(x, y))
        } else {
            fmt.Println("Нет седловой точки")

            x, y, _ := brownrobinson.BrownRobinsonMethod(0.001, m, t)
            xmi, _ := maxInd(x)
            ymi, _ := maxInd(y)

            xf := float64(xmi) / float64(n)
            yf := float64(ymi) / float64(n)

            fmt.Printf("x = %.3f\n", xf)
            fmt.Printf("y = %.3f\n", yf)
            fmt.Printf("H = %.3f\n", Hxy(xf, yf))
        }

        fmt.Println()
        fmt.Println()
    }



    // fmt.Println(t.Render())

    // prt := func(a []float64) {
    //     for _, v := range a {
    //         fmt.Printf("%.3f  ", v)
    //     }
    //
    //     fmt.Println()
    // }
    //
    // prt(x)
    // prt(y)
    //
    //
    // fmt.Println(v)

}
