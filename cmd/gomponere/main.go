package main

import (
	"fmt"

	"../../internal/input"
	"github.com/emicklei/dot"
	"github.com/spf13/afero"
)

func main() {
	var err error

	fs := afero.NewOsFs()
	reader := input.NewReader(fs)

	data, err := reader.ReadAll("../test/input")
	if err != nil {
		panic(err.Error())
	}

	_, err = input.Unmarshal(data)
	if err != nil {
		panic(err.Error())
	}

	// create graph
	//	dir=LR
	g := dot.NewGraph(dot.Directed)
	g.Attr("rankdir", "LR")

	// create a subgraph for each area (recursively)
	agstudio := g.Subgraph("AgStudio", dot.ClusterOption{})

	corteva := g.Subgraph("Corteva", dot.ClusterOption{})

	platform := g.Subgraph("Platform", dot.ClusterOption{})

	// create a subgraph for each level
	//	rank=same
	//	edge[style="invisible",dir="none"]
	agstudio_apps := agstudio.Subgraph("Apps", dot.ClusterOption{})
	agstudio_apps.Attr("rank", "same")

	agstudio_apis := agstudio.Subgraph("APIs", dot.ClusterOption{})
	agstudio_apis.Attr("rank", "same")

	agstudio_databases := agstudio.Subgraph("Databases", dot.ClusterOption{})
	agstudio_databases.Attr("rank", "same")

	corteva_apis := corteva.Subgraph("APIs", dot.ClusterOption{})
	corteva_apis.Attr("rank", "same")

	platform_apis := platform.Subgraph("APIs", dot.ClusterOption{})
	platform_apis.Attr("rank", "same")

	// add components w/ edges
	agstudio_desktop := agstudio_apps.Node("AgStudio Desktop")
	agstudio_map := agstudio_apps.Node("MAP")
	agstudio_desktop.
		Edge(agstudio_map).
		Attr("style", "invisible").
		Attr("dir", "none")

	agstudio_api1 := agstudio_apis.Node("API1")
	agstudio_api2 := agstudio_apis.Node("API2")
	agstudio_api3 := agstudio_apis.Node("API3")
	agstudio_api1.
		Edge(agstudio_api2).
		Attr("style", "invisible").
		Attr("dir", "none").
		Edge(agstudio_api3).
		Attr("style", "invisible").
		Attr("dir", "none")

	agstudio_database := agstudio_databases.Node("Database")

	corteva_api1 := corteva_apis.Node("API1")
	corteva_api2 := corteva_apis.Node("API1")
	corteva_api3 := corteva_apis.Node("API1")
	corteva_api1.
		Edge(corteva_api2).
		Attr("style", "invisible").
		Attr("dir", "none").
		Edge(corteva_api3).
		Attr("style", "invisible").
		Attr("dir", "none")

	platform_api1 := corteva_apis.Node("API1")
	platform_api2 := corteva_apis.Node("API1")
	platform_api3 := corteva_apis.Node("API1")
	platform_api1.
		Edge(platform_api2).
		Attr("style", "invisible").
		Attr("dir", "none").
		Edge(platform_api3).
		Attr("style", "invisible").
		Attr("dir", "none")

	// add real edges
	//	[constraint=false]
	agstudio_desktop.Edge(agstudio_api1).Attr("constraint", "false")
	agstudio_desktop.Edge(agstudio_api3).Attr("constraint", "false")
	agstudio_map.Edge(agstudio_api1).Attr("constraint", "false")
	agstudio_map.Edge(agstudio_api2).Attr("constraint", "false")
	agstudio_map.Edge(corteva_api3).Attr("constraint", "false")
	agstudio_api1.Edge(agstudio_database).Attr("constraint", "false")
	agstudio_api1.Edge(platform_api1).Attr("constraint", "false")
	agstudio_api2.Edge(agstudio_database).Attr("constraint", "false")
	agstudio_api2.Edge(corteva_api2).Attr("constraint", "false")
	agstudio_api3.Edge(agstudio_database).Attr("constraint", "false")

	// output dot to the console
	fmt.Print(g.String())
}
