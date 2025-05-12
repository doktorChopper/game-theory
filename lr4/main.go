package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"strconv"

	"github.com/awalterschulze/gographviz"
)

// "os"

// "github.com/awalterschulze/gographviz"

type node struct {
    key             string
    susessors       []*node
    // strategyCount   int 
    player          byte
    parent          *node
    costs           [][]int
    color           byte
}

type Tree struct {
    root *node
}

var leaves []*node = []*node{}
const depth int = 5

var R int = 0

func add(p byte, d int, n *node, par *node, g *gographviz.Graph) *node {
    if d == depth {
        // c := make([]int, 2)
        // for i := range c {
        //     c[i] = int(math.Pow(-1, float64(rand.Intn(2)))) * rand.Intn(21)
        // }
        c := [][]int{}
        c = append(c, []int{int(math.Pow(-1, float64(rand.Intn(2)))) * rand.Intn(21), int(math.Pow(-1, float64(rand.Intn(2)))) * rand.Intn(21)})


        r := strconv.Itoa(R)
        R++
        label := fmt.Sprintf("\"\\n%s\\n(%d, %d)\"",r, c[0][0], c[0][1])
        attrs := map[string]string {
            "label": label,
        }
        g.AddNode("G", r, attrs)
        // g.AddNode("G", r, nil)
        leaf := &node {
            key:            r,
            susessors:      nil,
            // strategyCount:  0,
            player:         p,
            parent:         par,
            costs:          c,
            color:          'w',
        }
        leaves = append(leaves, leaf)
        return leaf
    }


    if n == nil {
        size := 0
        var pl byte
        if p == 'A' {
            pl = 'B'
            size = 2
        } else {
            pl = 'A'
            size = 3
        }

        r := strconv.Itoa(R)
        R++

        label := fmt.Sprintf("\"%s\\n%s\"", string(p), r)
        attrs := map[string]string {
            "label": label,
        }
        g.AddNode("G", r, attrs)
        // g.AddNode("G", r, nil)
        
        n = &node{
            key:            r,
            susessors:      make([]*node, size),
            // strategyCount:  3,
            player:         p,
            parent:         par,
            color:          'w',
            costs:          nil,
        }

        d++

        for i := range n.susessors {
            n.susessors[i] = add(pl, d, n.susessors[i], n, g)
            g.AddEdge(r, n.susessors[i].key, true, nil)
        }
    }

    return n
}

func GenerateTree(g *gographviz.Graph) *Tree {

    players := []byte{'A', 'B'}

    // keys := make([]int, 129)
    // for i := range keys {
    //     keys[i] = i
    // }

    // fmt.Println(keys)

    t := &Tree {
        root: nil,
    }

    d := 0

    t.root = add(players[rand.Intn(len(players))], d, t.root, nil, g)

    return t
}

func ReverseInductionMethod(g *gographviz.Graph) {

    for len(leaves) > 0 {
        u := leaves[0]
        leaves = leaves[1:]
        
        if u != nil && u.parent != nil {
            if u.parent.color == 'w' {
                // maxc := []int{}
                maxc := u.costs[0]
                var idx int
                if u.player == 'A' {
                    idx = 1
                } else {
                    idx = 0
                }
                // r := u
                tmp := u
                for _, v := range u.parent.susessors {
                    for _, vv := range v.costs {
                        if vv[idx] > maxc[idx] {
                            maxc = vv
                            tmp = v
                        }
                    }

                    // if v.costs[idx] > maxc[idx] {
                    //     maxc = v.costs
                    //     // r = v
                    // }
                }
                u.parent.costs = append(u.parent.costs, maxc)
                tmp.color = 'v'
                for _, v := range u.parent.susessors {
                    for _, vv := range v.costs {
                        if vv[idx] == maxc[idx] && v.color != 'v'{
                            u.parent.costs = append(u.parent.costs, vv)
                            v.color = 'v'
                        }
                    }
                }

                n := g.Nodes.Lookup[u.parent.key]
                
                var strc string
                for _, v := range u.parent.costs {
                    strc = strc + fmt.Sprintf("(%d, %d)\\n", v[0], v[1])
                }

                // n.Attrs["label"] = fmt.Sprintf("\"%s\\n%s\\n(%d, %d)\"", string(u.parent.player), u.parent.key, u.parent.costs[0][0], u.parent.costs[0][1])
                n.Attrs["label"] = fmt.Sprintf("\"%s\\n%s\\n%s\"", string(u.parent.player), u.parent.key, strc)

                // e := g.Edges.SrcToDsts[u.parent.key][r.key][0]
                // e.Attrs["color"] = "red"
                // e.Attrs["penwidth"] = "2"
                
                leaves = append(leaves, u.parent)
                u.parent.color = 'b'
            }
        }
        // u.parent.costs 

    }

}


func PreOrderG(t *node, g *gographviz.Graph) {

    // c := t.root.costs
    for _, c := range t.costs {
        var traverse func(*node)

        check := func(a []int) bool {
            for i := range a {
                if a[i] != c[i] {
                    return false
                }
            }
            return true
        }

        traverse = func(n *node) {
            if n == nil {
                return
            }
            for _, v := range n.susessors {
                for _, vv := range v.costs {
                    if check(vv) {
                        PreOrderG(v, g)
                        e := g.Edges.SrcToDsts[n.key][v.key][0]
                        e.Attrs["color"] = "red"
                        e.Attrs["penwidth"] = "2"
                    }
                }
                traverse(v)
            }
        }

        traverse(t)

    }

}

func main() {
    // 129

    graphAst, _ := gographviz.ParseString(`digraph G {}`)
    graph := gographviz.NewGraph()
    if err := gographviz.Analyse(graphAst, graph); err != nil {
        panic(err)
    }

    t := GenerateTree(graph)
    ReverseInductionMethod(graph)
    PreOrderG(t.root, graph)
    fmt.Println(t)

    // TEST

    // graph.AddNode("G", "a", nil)
    // graph.AddNode("G", "b", nil)
    // graph.AddEdge("a", "b", true, map[string]string {
    //     "color":    "red",
    //     "penwidth": "5",
    // })
    // e := graph.Edges.SrcToDsts["a"]["b"][0]
    // e.Attrs["color"] = "green"
    // e.Attrs["penwidth"] = "3"

    output := graph.String()
    fmt.Println(output)

    os.WriteFile("graph.dot", []byte(output), 0644)

    cmd := exec.Command("dot", "-Tpng", "graph.dot", "-o", "graph.png")
    err := cmd.Run()
    if err != nil {
        log.Fatal("Error converting DOT to PNG: ", err)
    }

    fmt.Println(leaves)
    fmt.Println(len(leaves))

    // img, err := imaging.Open("graph.png") 
    // if err != nil {
    //     log.Fatal(err)
    // }
    //
    // resize := imaging.Resize(img, 1080, 1080, imaging.Lanczos)
    // err = imaging.Save(resize, "graph_resized.png")
    // if err != nil {
    //     log.Fatal(err)
    }
