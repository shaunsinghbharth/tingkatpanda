package goutils

import (
	"container/heap"
	"math"
	"sort"

	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
)

type WeightedBuilder interface {
	AddNode(graph.Node)
	SetWeightedEdge(graph.WeightedEdge)
}

func Prim(dst WeightedBuilder, g graph.WeightedUndirected) float64 {
	nodes := graph.NodesOf(g.Nodes())
	if len(nodes) == 0 {
		return 0
	}

	q := &primQueue{
		indexOf: make(map[int64]int, len(nodes)-1),
		nodes:   make([]simple.WeightedEdge, 0, len(nodes)-1),
	}
	dst.AddNode(nodes[0])
	for _, u := range nodes[1:] {
		dst.AddNode(u)
		heap.Push(q, simple.WeightedEdge{F: u, W: math.Inf(1)})
	}

	u := nodes[0]
	uid := u.ID()
	for _, v := range graph.NodesOf(g.From(uid)) {
		w, ok := g.Weight(uid, v.ID())
		if !ok {
			panic("prim: unexpected invalid weight")
		}
		q.update(v, u, w)
	}

	var w float64
	for q.Len() > 0 {
		e := heap.Pop(q).(simple.WeightedEdge)
		if e.To() != nil && g.HasEdgeBetween(e.From().ID(), e.To().ID()) {
			dst.SetWeightedEdge(g.WeightedEdge(e.From().ID(), e.To().ID()))
			w += e.Weight()
		}

		u = e.From()
		uid := u.ID()
		for _, n := range graph.NodesOf(g.From(uid)) {
			if key, ok := q.key(n); ok {
				w, ok := g.Weight(uid, n.ID())
				if !ok {
					panic("prim: unexpected invalid weight")
				}
				if w < key {
					q.update(n, u, w)
				}
			}
		}
	}
	return w
}

type primQueue struct {
	indexOf map[int64]int
	nodes   []simple.WeightedEdge
}

func (q *primQueue) Less(i, j int) bool {
	return q.nodes[i].Weight() < q.nodes[j].Weight()
}

func (q *primQueue) Swap(i, j int) {
	q.indexOf[q.nodes[i].From().ID()] = j
	q.indexOf[q.nodes[j].From().ID()] = i
	q.nodes[i], q.nodes[j] = q.nodes[j], q.nodes[i]
}

func (q *primQueue) Len() int {
	return len(q.nodes)
}

func (q *primQueue) Push(x interface{}) {
	n := x.(simple.WeightedEdge)
	q.indexOf[n.From().ID()] = len(q.nodes)
	q.nodes = append(q.nodes, n)
}

func (q *primQueue) Pop() interface{} {
	n := q.nodes[len(q.nodes)-1]
	q.nodes = q.nodes[:len(q.nodes)-1]
	delete(q.indexOf, n.From().ID())
	return n
}

func (q *primQueue) key(u graph.Node) (key float64, ok bool) {
	i, ok := q.indexOf[u.ID()]
	if !ok {
		return math.Inf(1), false
	}
	return q.nodes[i].Weight(), ok
}

func (q *primQueue) update(u, v graph.Node, key float64) {
	id := u.ID()
	i, ok := q.indexOf[id]
	if !ok {
		return
	}
	q.nodes[i].T = v
	q.nodes[i].W = key
	heap.Fix(q, i)
}

type UndirectedWeightLister interface {
	graph.WeightedUndirected
	WeightedEdges() graph.WeightedEdges
}

func Kruskal(dst WeightedBuilder, g UndirectedWeightLister) float64 {
	edges := graph.WeightedEdgesOf(g.WeightedEdges())
	sort.Sort(byWeight(edges))

	ds := make(djSet)
	it := g.Nodes()
	for it.Next() {
		n := it.Node()
		dst.AddNode(n)
		ds.add(n.ID())
	}

	var w float64
	for _, e := range edges {
		if s1, s2 := ds.find(e.From().ID()), ds.find(e.To().ID()); s1 != s2 {
			ds.union(s1, s2)
			dst.SetWeightedEdge(g.WeightedEdge(e.From().ID(), e.To().ID()))
			w += e.Weight()
		}
	}
	return w
}

type byWeight []graph.WeightedEdge

func (e byWeight) Len() int           { return len(e) }
func (e byWeight) Less(i, j int) bool { return e[i].Weight() < e[j].Weight() }
func (e byWeight) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }

type djSet map[int64]*dsNode

func (s djSet) add(e int64) {
	if _, ok := s[e]; ok {
		return
	}
	s[e] = &dsNode{}
}

func (djSet) union(a, b *dsNode) {
	ra := find(a)
	rb := find(b)
	if ra == rb {
		return
	}
	if ra.rank < rb.rank {
		ra.parent = rb
		return
	}
	rb.parent = ra
	if ra.rank == rb.rank {
		ra.rank++
	}
}

func (s djSet) find(e int64) *dsNode {
	n, ok := s[e]
	if !ok {
		return nil
	}
	return find(n)
}

func find(n *dsNode) *dsNode {
	for ; n.parent != nil; n = n.parent {
	}
	return n
}

type dsNode struct {
	parent *dsNode
	rank   int
}
