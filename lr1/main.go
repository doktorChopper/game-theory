package main

import (
	"fmt"

	brownrobinson "github.com/doktorChopper/game-theory/brown-robinson"
	"github.com/doktorChopper/go-matrix/matrix"
	"github.com/jedib0t/go-pretty/v6/table"
)

func initTable(t table.Writer) {
    t.SetTitle("Brown-Robinson method")
    t.AppendHeader(table.Row{"#", "_V", "V_", "EPSILON"})
}

func main() {

    C := [][]float64 {
        {2, 1, 3},
        {3, 0, 1},
        {1, 2, 1},
    }

    t := table.NewWriter()

    initTable(t)

    m := matrix.NewMatrixFromSlice(C)
    x, y, v := brownrobinson.BrownRobinsonMethod(0.1, m, t)

    fmt.Println(t.Render())

    fmt.Println(x)
    fmt.Println(y)
    fmt.Println(v)
}
