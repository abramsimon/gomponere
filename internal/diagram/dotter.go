package diagram

import (
	"../model"
	"github.com/emicklei/dot"
)

func MakeDot(diagram model.Diagram) (string, error) {
	// start our directed graph
	g := dot.NewGraph(dot.Directed)

	// rank direction of LR is important for our layout
	g.Attr("rankdir", "LR")

	// add the teams to the graph (kinda like a legend)
	if err := MakeTeams(diagram, g); err != nil {
		return "", err
	}

	// keep a running list of nodes to make edges from
	nodes := make(map[string]dot.Node, len(diagram.Components))

	// add areas and components to the graph, starting with root areas (no parents) and traversing recursively
	for k, a := range diagram.Areas {
		if a.ParentKey == "" {
			n, err := MakeArea(diagram, g, a, k)
			if err != nil {
				return "", err
			}

			// merge the new nodes into the running list
			for nk, nv := range n {
				nodes[nk] = nv
			}
		}
	}

	// add the edges for all component dependencies
	for lk, lc := range diagram.Components {
		for _, rk := range lc.DependencyKeys {
			nodes[lk].Edge(nodes[rk]).Attr("constraint", "false")
		}
	}

	return g.String(), nil
}

func MakeArea(diagram model.Diagram, graph *dot.Graph, area model.Area, areaKey string) (map[string]dot.Node, error) {
	nodes := map[string]dot.Node{}

	// add the area to the graph
	g := graph.Subgraph(area.Name, dot.ClusterOption{})

	// add child areas to the graph
	for k, a := range diagram.Areas {
		if a.ParentKey == areaKey {
			n, err := MakeArea(diagram, g, a, k)
			if err != nil {
				return nil, err
			}

			// merge the new nodes into the running list
			for nk, nv := range n {
				nodes[nk] = nv
			}
		}
	}

	// group components in this area by levels
	componentsByLevel := make(map[string]map[string]model.Component)
	for k, c := range diagram.Components {
		if c.AreaKey == areaKey {
			if _, exists := componentsByLevel[c.LevelKey]; exists {
				componentsByLevel[c.LevelKey][k] = c
			} else {
				componentsByLevel[c.LevelKey] = map[string]model.Component{k: c}
			}
		}
	}

	// create a subgraph for each level with it's components
	for k, c := range componentsByLevel {
		n, err := MakeLevels(diagram, g, diagram.Levels[k], k, c)
		if err != nil {
			return nil, err
		}

		// merge the new nodes into the running list
		for nk, nv := range n {
			nodes[nk] = nv
		}
	}

	return nodes, nil
}

func MakeLevels(diagram model.Diagram, graph *dot.Graph, level model.Level, levelKey string, components map[string]model.Component) (map[string]dot.Node, error) {
	nodes := make(map[string]dot.Node, len(components))

	// create a subgraph so the components can be ranked the same
	g := graph.Subgraph("")
	g.Attr("rank", "same")

	// add child components to the graph
	var prev *dot.Node
	for k, c := range components {
		t := diagram.Teams[c.TeamKey]

		n := g.Node(c.Name)
		n.Attr("style", "filled").
			Attr("color", t.Display.BackgroundColor).
			Attr("fontcolor", t.Display.ForegroundColor)

		// create an invisible edge to the previous node
		// this is the trick that makes left-right ranking work
		if prev != nil {
			prev.Edge(n).
				Attr("style", "invisible").
				Attr("dir", "none")
		}

		// make the current the previous so we can create an edge to it
		prev = &n

		// capture the node for creating edges later
		nodes[k] = n
	}

	return nodes, nil
}

func MakeTeams(diagram model.Diagram, graph *dot.Graph) error {
	// add the teams subgraph to the graph
	g := graph.Subgraph("Teams", dot.ClusterOption{})
	g.Attr("rank", "same")

	var prev *dot.Node
	for _, t := range diagram.Teams {
		n := g.Node(t.Name)
		n.Attr("style", "filled").
			Attr("color", t.Display.BackgroundColor).
			Attr("fontcolor", t.Display.ForegroundColor)

		// create an invisible edge to the previous node
		// this is the trick that makes left-right ranking work
		if prev != nil {
			prev.Edge(n).
				Attr("style", "invisible").
				Attr("dir", "none")
		}

		prev = &n
	}

	return nil
}
