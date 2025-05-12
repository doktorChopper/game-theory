package main

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/doktorChopper/go-matrix/matrix"
)

type Pair struct {
    x float64
    y float64
}

func (p *Pair) First() float64 {
    return p.x
}

func (p *Pair) Second() float64 {
    return p.y
} 

func Pareto(m [][]Pair) ([]Pair, []Pair) {

    s := []Pair{}
    idxs := []Pair{}
    
    flag1 := true
    flag2 := false
    flag3 := false

    for i := 0; i < len(m); i++ {

        for j := 0; j < len(m[0]); j++ {

            flag1 = true

            for k := 0; k < len(m); k++ {
                for l := 0; l < len(m[0]); l++ {
                    if k == i && l == j {
                        continue
                    }

                    if m[k][l].First() >= m[i][j].First() && m[k][l].Second() >= m[i][j].Second() {
                        flag1 = false
                    }
                    
                    if m[k][l].First() <= m[i][j].First() {
                        flag2 = true
                    }

                    if m[k][l].Second() <= m[i][j].Second() {
                        flag3 = true
                    }
                }
            }

            if flag1 && flag2 && flag3 {
                s = append(s, m[i][j])
                idxs = append(idxs, Pair{float64(i), float64(j)})
            }

        }

    }

    return s, idxs
}



func Nash(m [][]Pair) ([]Pair, []Pair) {

    s := []Pair{}
    idxs := []Pair{}

    flag := true

    for i := 0; i < len(m); i++ {

        for j := 0; j < len(m[0]); j++ {

            flag = true

            for k := 0; k < len(m); k++ {
                if !(m[i][j].First() >= m[k][j].First()) {
                    flag = false
                }
            }

            for k := 0; k < len(m[0]); k++ {
                if !(m[i][j].Second() >= m[i][k].Second()) {
                    flag = false
                }
            }

            if flag {
                s = append(s, m[i][j])
                idxs = append(idxs, Pair{float64(i), float64(j)})
            }
        }
    } 

    return s, idxs
}

func Search(p []Pair, i, j int) bool {

    for _, v := range p {
        if i == int(v.First()) && j == int(v.Second()) {
            return true
        }
    }

    return false
}

func Solution2() {
    // mat := [][]Pair {
    //     {Pair{6, 8}, Pair{7, 4}},
    //     {Pair{0, 1}, Pair{9, 3}},
    // }
}

func main() {

    const (
        green  = "\033[32m"
        reset  = "\033[0m"
    )

    printMatrix := func(m [][]Pair, idxs []Pair) {
        for i := range m {
            for j := range m[i] {

                if Search(idxs, i, j) {
                    fmt.Printf("%s(%-3.0f, %-3.0f)%s     ", green, m[i][j].First(), m[i][j].Second(), reset)
                } else {
                    fmt.Printf("(%-3.0f, %-3.0f)     ", m[i][j].First(), m[i][j].Second())
                }

            }
            fmt.Println()
        }
        fmt.Println()
        fmt.Println()
    }


    // mat := [][]Pair {
    //     {Pair{24, 32}, Pair{2, 2}, Pair{40, 17}, Pair{-46, -29}, Pair{27, -42}, Pair{-46, -8}, Pair{35, -20}, Pair{-23, -16}, Pair{-33, 12}, Pair{-50, -1}},
    //     {Pair{-40, -11}, Pair{28, -22}, Pair{-15, -47}, Pair{38, 49}, Pair{26, -25}, Pair{-38, 26}, Pair{-35, 35}, Pair{-26, 30}, Pair{-37, 23}, Pair{-19, -30}},
    //     {Pair{-1, 46}, Pair{-36, -14}, Pair{30, -17}, Pair{-35, 5}, Pair{4, -22}, Pair{7, -33}, Pair{17, -2}, Pair{-6, -15}, Pair{-8, -48}, Pair{-33, -14}},
    //     {Pair{46, 28}, Pair{1, -6}, Pair{2, 18}, Pair{-48, -38}, Pair{45, 49}, Pair{-3, -28}, Pair{-29, 37}, Pair{-24, 7}, Pair{36, 17}, Pair{24, -19}},
    //     {Pair{-9, 25}, Pair{41, -48}, Pair{-9, 6}, Pair{-17, -22}, Pair{-30, 12}, Pair{-20, -5}, Pair{-44, 32}, Pair{41, -35}, Pair{-43, 42}, Pair{9, -33}},
    //     {Pair{19, 13}, Pair{49, -37}, Pair{0, -10}, Pair{-30, 39}, Pair{-48, -16}, Pair{38, 42}, Pair{-18, -31}, Pair{-27, -6}, Pair{35, -16}, Pair{11, -43}},
    //     {Pair{37, 10}, Pair{4, -38}, Pair{-9, 36}, Pair{-7, -30}, Pair{-21, 5}, Pair{-9, -9}, Pair{15, -30}, Pair{-49, 34}, Pair{-23, 43}, Pair{-20, -15}},
    //     {Pair{-33, 37}, Pair{29, 23}, Pair{-29, 5}, Pair{24, 17}, Pair{-21, 49}, Pair{17, -30}, Pair{49, 47}, Pair{41, 8}, Pair{-36, 6}, Pair{40, 34}},
    //     {Pair{-40, 9}, Pair{24, -32}, Pair{-45, 41}, Pair{49, 34}, Pair{1, -12}, Pair{47, 43}, Pair{49, -11}, Pair{-17, -39}, Pair{26, 24}, Pair{-15, -3}},
    //     {Pair{46, -49}, Pair{9, -5}, Pair{28, 36}, Pair{-38, 24}, Pair{-11, -39}, Pair{32, -41}, Pair{9, -13}, Pair{42, 10}, Pair{19, -18}, Pair{37, 4}},
    // }

    prisoner := [][]Pair {
        {Pair{-5, -5}, Pair{0, -10}},
        {Pair{-10, 0}, Pair{-1, -1}},
    }

    fmt.Println()

    pn, idxsPN := Nash(prisoner)
    fmt.Printf("********** Prisoner **********\n\n")
    fmt.Println("Nash: ", pn)
    fmt.Println()
    printMatrix(prisoner, idxsPN)

    pp, idxsPP := Pareto(prisoner)
    fmt.Println("Pareto: ", pp)
    fmt.Println() 
    printMatrix(prisoner, idxsPP)

    family := [][]Pair {
        {Pair{4, 1}, Pair{0, 0}},
        {Pair{0, 0}, Pair{1, 4}},
    }

    fn, idxsFN := Nash(family)
    fmt.Printf("********** Family **********\n\n")
    fmt.Println("Nash: ", fn)
    fmt.Println()
    printMatrix(family, idxsFN)

    fp, idxsFP := Pareto(family)
    fmt.Println("Pareto: ", fp)
    fmt.Println()
    printMatrix(family, idxsFP)

    cross := [][]Pair {
        {Pair{1, 1}, Pair{1 - 1e-10, 2}},
        {Pair{2, 1 - 1e-10}, Pair{0, 0}},
    }

    cn, idxsCN := Nash(cross)
    fmt.Printf("********** Cross **********\n\n")
    // fmt.Println("Nash: ", cn)
    fmt.Print("Nash: [")
    for _, v := range cn {
        fmt.Printf("%.2f ", v)
    }
    fmt.Printf("]\n")
    fmt.Println()
    printMatrix(cross, idxsCN)

    cp, idxsCP := Pareto(cross)
    fmt.Print("Pareto: [")
    for _, v := range cp {
        fmt.Printf("%.2f ", v)
    }
    fmt.Printf("]\n")
    fmt.Println()

    printMatrix(cross, idxsCP)

    mat := make([][]Pair, 10)
    for i := range mat {
        mat[i] = make([]Pair, 10)
    }

    for i := range mat {
        for j := range mat[i] {
            mat[i][j] = Pair{float64(int(math.Pow(-1, float64(rand.Intn(2)))) * rand.Intn(51)), float64(int(math.Pow(-1, float64(rand.Intn(2)))) * rand.Intn(51))}
        }
    }

    rn, idxsRN := Nash(mat)
    fmt.Printf("********** Random Matrix **********\n\n")
    fmt.Println("Nash: ", rn)
    fmt.Println()
    printMatrix(mat, idxsRN)


    rp, idxsRP := Pareto(mat)
    fmt.Println("Pareto: ", rp)
    fmt.Println()
    printMatrix(mat, idxsRP)


    mat2 := [][]Pair {
        {Pair{6, 8}, Pair{7, 4}},
        {Pair{0, 1}, Pair{9, 3}},
    }

    // mat2 := [][]Pair {
    //     {Pair{0, 1}, Pair{11, 4}},
    //     {Pair{7, 8}, Pair{6, 3}},
    // }

    Amat2 := make([][]float64, 2)
    Amat2[0] = make([]float64, 2)
    Amat2[1] = make([]float64, 2)

    Bmat2 := make([][]float64, 2)
    Bmat2[0] = make([]float64, 2)
    Bmat2[1] = make([]float64, 2)

    fmt.Printf("********** B-Matrix **********\n\n")
    for i := range mat2 {
        for j := range mat2[i] {
            Amat2[i][j] = mat2[i][j].First()
            Bmat2[i][j] = mat2[i][j].Second()
        }
        fmt.Println()
    }

    bn, idxsBM := Nash(mat2)
    fmt.Println("Nash: ", bn)
    fmt.Println()
    printMatrix(mat2, idxsBM)
    
    printMatrix2 := func(m *matrix.Matrix) {
        for i := range m.Rows() {
            for j := range m.Cols() {
                fmt.Printf("%.3f ", m.GetAt(i, j))
            }
            fmt.Println()
        }
        fmt.Println()
    }

    // printMatrix2(mat2)

    Amat2M := matrix.NewMatrixFromSlice(Amat2)
    Bmat2M := matrix.NewMatrixFromSlice(Bmat2)
    
    fmt.Println("A matrix:")
    Amat2M.Display()
    fmt.Println("B matrix:")
    Bmat2M.Display()

    aI := Amat2M.InverseMatrix()
    bI := Bmat2M.InverseMatrix()

    fmt.Println("A inverse matrix:")
    aI.Display()
    fmt.Println("B inverse matrix:")
    bI.Display()

    u := matrix.NewMatrixFromSlice([][]float64{{1, 1}})
    uT := u.Transposition()

    v1 := 1 / u.Mult(aI).Mult(uT).GetAt(0, 0)
    v2 := 1 / u.Mult(bI).Mult(uT).GetAt(0, 0)

    fmt.Printf("v1 = %.3f\n", v1)
    fmt.Printf("v2 = %.3f\n", v2)
    fmt.Println()

    x := u.Mult(bI)
    x.ScalarMult(v2)

    y := aI.Mult(uT)
    y.ScalarMult(v1)

    // x.Display()
    fmt.Print("x = ")
    printMatrix2(x)

    fmt.Print("y = ")
    printMatrix2(y.Transposition())

}
