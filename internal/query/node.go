package query

import (
	"context"
	"time"

	promModel "github.com/prometheus/common/model"

	"github.com/noisyboy-9/data_extractor/internal/config"
	"github.com/noisyboy-9/data_extractor/internal/log"
	"github.com/noisyboy-9/data_extractor/internal/model"
	"github.com/noisyboy-9/data_extractor/internal/service"
)

func GetNodeList() []model.Node {
	ctx, cancel := context.WithTimeout(context.Background(), config.Prometheus.Timeout)
	defer cancel()

	results, warnings, err := service.Promtheus.Api.Query(ctx, "kube_node_created", time.Now())

	if err != nil {
		log.App.WithError(err).Panic("error in getting node list form prometheus")
	}

	if len(warnings) > 0 {
		log.App.Warnf("get node list warnings: %v", warnings)
	}

	nodes := make([]model.Node, 0)
	nodesVector := results.(promModel.Vector)
	for _, node := range nodesVector {
		nodes = append(nodes, model.Node{
			Name: string(node.Metric["node"]),
		})
	}

	return nodes
}
