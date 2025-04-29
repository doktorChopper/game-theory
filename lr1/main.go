package main

import (
	"fmt"

	brownrobinson "github.com/doktorChopper/game-theory/brown-robinson"
	"github.com/doktorChopper/go-matrix/matrix"
	"github.com/jedib0t/go-pretty/v6/table"
)

func initTable(t table.Writer) {
    t.SetTitle("Brown-Robinson method")
    t.AppendHeader(table.Row{"#", "A", "B", "_V", "V_", "EPSILON"})
}

func main() {

    // C := [][]float64 {
    //     {2, 1, 3},
    //     {3, 0, 1},
    //     {1, 2, 1},
    // }

    C := [][]float64 {
        {19, 7, 3},
        {6, 9, 9},
        {8, 2, 11},
    }

    // C := [][]float64 {
    //     {13, 3, 14},
    //     {15, 5, 0},
    //     {7, 19, 13},
    // }

    m := matrix.NewMatrixFromSlice(C)

    a := matrix.NewMatrixFromSlice(C)

    fmt.Println()
    fmt.Println("M")
    m.Display()

    fmt.Printf("\n*************** Analytics method ***************\n\n")

    u := matrix.NewMatrixFromSlice([][]float64{{1, 1, 1}})
    uT := u.Transposition()

    ia := a.InverseMatrix()

    xc := u.Mult(ia)
    yc := ia.Mult(uT)

    z := xc.Mult(uT)

    f := 1 / z.GetAt(0, 0)

    fmt.Print("x = ")
    yc.ScalarDiv(z.GetAt(0, 0))
    ycc := yc.Transposition()
    ycc.Display()

    fmt.Print("y = ")
    xc.ScalarDiv(z.GetAt(0, 0))
    xc.Display()

    fmt.Printf("v = %.2f\n\n", f)

    t := table.NewWriter()

    initTable(t)

    x, y, v := brownrobinson.BrownRobinsonMethod(0.1, m, t)

    fmt.Println(t.Render())

    fmt.Println()
    fmt.Printf("~x = ")
    for _, v := range x {
        fmt.Printf("%.3f ", v)
    }
    fmt.Println()

    fmt.Printf("~y = ")
    for _, v := range y {
        fmt.Printf("%.3f ", v)
    }
    fmt.Println()

    fmt.Printf("v = %.2f\n", v)
}
