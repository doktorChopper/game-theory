package main

import (
	"fmt"
)

type Pair struct {
    x int
    y int
}

func (p *Pair) First() int {
    return p.x
}

func (p *Pair) Second() int {
    return p.y
} 

func Pareto(m [][]Pair) []Pair {

    return []Pair{}
}

func Nash(m [][]Pair) []Pair {

    s := []Pair{}

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
            }
        }
    } 

    return s
}

func main() {

    // m := [][]Pair {
    //     {Pair{-5, -5}, Pair{0, -10}},
    //     {Pair{-10, 0}, Pair{-1, -1}},
    // }

    // m := [][]Pair {
    //     {Pair{4, 1}, Pair{0, 0}},
    //     {Pair{0, 0}, Pair{1, 4}},
    // }

    m := [][]Pair {
        {Pair{24, 32}, Pair{2, 2}, Pair{40, 17}, Pair{-46, -29}, Pair{27, -42}, Pair{-46, -8}, Pair{35, -20}, Pair{-23, -16}, Pair{-33, 12}, Pair{-50, -1}},
        {Pair{-40, -11}, Pair{28, -22}, Pair{-15, -47}, Pair{38, 49}, Pair{26, -25}, Pair{-38, 26}, Pair{-35, 35}, Pair{-26, 30}, Pair{-37, 23}, Pair{-19, -30}},
        {Pair{-1, 46}, Pair{-36, -14}, Pair{30, -17}, Pair{-35, 5}, Pair{4, -22}, Pair{7, -33}, Pair{17, -2}, Pair{-6, -15}, Pair{-8, -48}, Pair{-33, -14}},
        {Pair{46, 28}, Pair{1, -6}, Pair{2, 18}, Pair{-48, -38}, Pair{45, 49}, Pair{-3, -28}, Pair{-29, 37}, Pair{-24, 7}, Pair{36, 17}, Pair{24, -19}},
        {Pair{-9, 25}, Pair{41, -48}, Pair{-9, 6}, Pair{-17, -22}, Pair{-30, 12}, Pair{-20, -5}, Pair{-44, 32}, Pair{41, -35}, Pair{-43, 42}, Pair{9, -33}},
        {Pair{19, 13}, Pair{49, -37}, Pair{0, -10}, Pair{-30, 39}, Pair{-48, -16}, Pair{38, 42}, Pair{-18, -31}, Pair{-27, -6}, Pair{35, -16}, Pair{11, -43}},
        {Pair{37, 10}, Pair{4, -38}, Pair{-9, 36}, Pair{-7, -30}, Pair{-21, 5}, Pair{-9, -9}, Pair{15, -30}, Pair{-49, 34}, Pair{-23, 43}, Pair{-20, -15}},
        {Pair{-33, 37}, Pair{29, 23}, Pair{-29, 5}, Pair{24, 17}, Pair{-21, 49}, Pair{17, -30}, Pair{49, 47}, Pair{41, 8}, Pair{-36, 6}, Pair{40, 34}},
        {Pair{-40, 9}, Pair{24, -32}, Pair{-45, 41}, Pair{49, 34}, Pair{1, -12}, Pair{47, 43}, Pair{49, -11}, Pair{-17, -39}, Pair{26, 24}, Pair{-15, -3}},
        {Pair{46, -49}, Pair{9, -5}, Pair{28, 36}, Pair{-38, 24}, Pair{-11, -39}, Pair{32, -41}, Pair{9, -13}, Pair{42, 10}, Pair{19, -18}, Pair{37, 4}},
    }

    for i := range m {
        for j := range m[i] {
            fmt.Printf("(%v, %v)    ", m[i][j].First(), m[i][j].Second())
        }

        fmt.Println()
    }

    r := Nash(m)

    fmt.Println(r)
}
