// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pullsync

import (
	m "github.com/ethersphere/bee/pkg/metrics"
	"github.com/prometheus/client_golang/prometheus"
)

type metrics struct {
	Offered       prometheus.Counter  // number of chunks offered
	Wanted        prometheus.Counter  // number of chunks wanted
	Delivered     prometheus.Counter  // number of chunk deliveries
	DbOps         prometheus.Counter  // number of db ops
	DuplicateRuid prometheus.Counter  // number of duplicate RUID requests we got
	LastReceived  prometheus.GaugeVec // last timestamp of the received chunks per bin
}

func newMetrics() metrics {
	subsystem := "pullsync"

	return metrics{
		Offered: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: m.Namespace,
			Subsystem: subsystem,
			Name:      "chunks_offered",
			Help:      "Total chunks offered.",
		}),
		Wanted: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: m.Namespace,
			Subsystem: subsystem,
			Name:      "chunks_wanted",
			Help:      "Total chunks wanted.",
		}),
		Delivered: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: m.Namespace,
			Subsystem: subsystem,
			Name:      "chunks_delivered",
			Help:      "Total chunks delivered.",
		}),
		DbOps: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: m.Namespace,
			Subsystem: subsystem,
			Name:      "db_ops",
			Help:      "Total Db Ops.",
		}),
		DuplicateRuid: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: m.Namespace,
			Subsystem: subsystem,
			Name:      "duplicate_ruids",
			Help:      "Total duplicate RUIDs.",
		}),
		LastReceived: *prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: m.Namespace,
				Subsystem: subsystem,
				Name:      "last_received",
				Help:      `The last timestamp of the received chunks per bin.`,
			}, []string{"bin"}),
	}
}

func (s *Syncer) Metrics() []prometheus.Collector {
	return m.PrometheusCollectorsFromFields(s.metrics)
}
