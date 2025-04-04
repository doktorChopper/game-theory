package brownrobinson

import (
	"math"
	"math/rand"

	"github.com/doktorChopper/go-matrix/matrix"
	"github.com/jedib0t/go-pretty/v6/table"
)

func vkMax(y []int, mat *matrix.Matrix, k int) (float64, []int) {

    r := make([]float64, mat.Rows())

    for i := 0; i < mat.Rows(); i++ {
        s := 0.0
        for j := 0; j < mat.Cols(); j++ {
            s += mat.GetAt(i, j) * float64(y[j])
        }
        r[i] = s
    }

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

func vkMin(x []int, mat *matrix.Matrix, k int) (float64, []int) {

    r := make([]float64, mat.Rows())

    for i := 0; i < mat.Rows(); i++ {
        s := 0.0
        for j := 0; j < mat.Cols(); j++ {
            s += mat.GetAt(j, i) * float64(x[j])
        }
        r[i] = s
    }

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

func randStrateghy(a []int) int {
    return a[rand.Intn(len(a))]
}

func BrownRobinsonMethod(e float64, mat *matrix.Matrix, t table.Writer) ([]float64, []float64, float64) {

    yk := make([]int, mat.Rows())
    xk := make([]int, mat.Rows())

    yk[0] = 1
    xk[0] = 1

    q, eq1 := vkMax(yk, mat, 1)
    w, eq2 := vkMin(xk, mat, 1)

    xk[randStrateghy(eq1)] += 1
    yk[randStrateghy(eq2)] += 1

    epsilon := q - w

    i := 1

    t.AppendRow(table.Row{i, 
        math.Round(q * 100) / 100, 
        math.Round(w * 100) / 100, 
        math.Round(epsilon * 100) / 100})

    for epsilon > e {
        
        k, eq1 := vkMax(yk, mat, i + 1)
        l, eq2 := vkMin(xk, mat, i + 1)

        if k < q {
            q = k
        }

        if l > w {
            w = l
        }

        xk[randStrateghy(eq1)] += 1
        yk[randStrateghy(eq2)] += 1

        epsilon = q - w

        t.AppendRow(table.Row{i + 1, 
            math.Round(k * 100) / 100, 
            math.Round(l * 100) / 100, 
            math.Round(epsilon * 100) / 100})

        i++
    }

    retxk := []float64{}
    retyk := []float64{}

    for _, v := range xk {
        retxk = append(retxk, float64(v) / float64(i + 1))
    }

    for _, v := range yk {
        retyk = append(retyk, float64(v) / float64(i + 1))
    }


    return retxk, retyk, (q + w) / 2
}
