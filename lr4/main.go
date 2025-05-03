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
    key             int
    susessors       []*node
    // strategyCount   int 
    player          byte
    costs           []int

}

type Tree struct {
    root *node
}

const depth int = 5

func add(p byte, d int, n *node, g *gographviz.Graph) *node {
    if d == depth {
        c := make([]int, 2)
        for i := range c {
            c[i] = int(math.Pow(-1, float64(rand.Intn(2)))) * rand.Intn(21)
        }


        r := rand.Intn(1000000)
        label := fmt.Sprintf("\"%s\\n%d\\n(%d, %d)\"", string(p), r, c[0], c[1])
        attrs := map[string]string {
            "label": label,
        }
        g.AddNode("G", strconv.Itoa(r), attrs)
        return &node {
            key:            r,
            susessors:      nil,
            // strategyCount:  0,
            player:         p,
            costs:          c,
        }
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

        r := rand.Intn(1000000)
        label := fmt.Sprintf("\"%s\\n%d\"", string(p), r)
        attrs := map[string]string {
            "label": label,
        }
        g.AddNode("G", strconv.Itoa(r), attrs)
        
        n = &node{
            key:            r,
            susessors:      make([]*node, size),
            // strategyCount:  3,
            player:         p,
            costs:          nil,
        }

        d++

        for i := range n.susessors {
            n.susessors[i] = add(pl, d, n.susessors[i], g)
            g.AddEdge(strconv.Itoa(r), strconv.Itoa(n.susessors[i].key), true, nil)
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

    t.root = add(players[rand.Intn(len(players))], d, t.root, g)

    return t
}

func main() {
    // 129

    graphAst, _ := gographviz.ParseString(`digraph G {}`)
    graph := gographviz.NewGraph()
    if err := gographviz.Analyse(graphAst, graph); err != nil {
        panic(err)
    }

    t := GenerateTree(graph)
    fmt.Println(t)

    // TEST


    // graph.AddNode("G", "a", nil)
    // graph.AddNode("G", "b", nil)
    // graph.AddEdge("a", "b", true, nil)
    output := graph.String()
    fmt.Println(output)

    os.WriteFile("graph.dot", []byte(output), 0644)

    cmd := exec.Command("dot", "-Tpng", "graph.dot", "-o", "graph.png")
    err := cmd.Run()
    if err != nil {
        log.Fatal("Error converting DOT to PNG: ", err)
    }

    // img, err := imaging.Open("graph.png") 
    // if err != nil {
    //     log.Fatal(err)
    // }
    //
    // resize := imaging.Resize(img, 1080, 1080, imaging.Lanczos)
    // err = imaging.Save(resize, "graph_resized.png")
    // if err != nil {
    //     log.Fatal(err)
    // }
}
