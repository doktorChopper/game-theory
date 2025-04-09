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

func MaxMin(m *matrix.Matrix) (float64, int) {
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

func MinMax(m *matrix.Matrix) (float64, int) {

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

func MaxInd(a []float64) (int, error) {
    
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

func initTable(t table.Writer) {
    t.SetTitle("Brown-Robinson method")
    t.AppendHeader(table.Row{"#", "_V", "V_", "EPSILON"})
}

func main() {

    m := HN(2)
    m.Display()

    a, ia := MaxMin(m)
    b, ib := MinMax(m)

    fmt.Println("MaxMin: ", a, ia)
    fmt.Println("MinMax: ", b, ib)

    t := table.NewWriter()
    initTable(t)

    x, y, v := brownrobinson.BrownRobinsonMethod(0.001, m, t)

    fmt.Println(t.Render())

    prt := func(a []float64) {
        for _, v := range a {
            fmt.Printf("%.3f  ", v)
        }

        fmt.Println()
    }

    prt(x)
    prt(y)

    xmi, _ := MaxInd(x)
    ymi, _ := MaxInd(y)

    fmt.Println(float64(xmi) / 2)
    fmt.Println(float64(ymi) / 2)

    fmt.Println(v)

}
