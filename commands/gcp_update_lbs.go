package commands

import "github.com/cloudfoundry/bosh-bootloader/storage"

type GCPUpdateLBs struct {
	gcpCreateLBs gcpCreateLBs
}

func NewGCPUpdateLBs(gcpCreateLBs gcpCreateLBs) GCPUpdateLBs {
	return GCPUpdateLBs{
		gcpCreateLBs: gcpCreateLBs,
	}
}

func (g GCPUpdateLBs) Execute(config GCPCreateLBsConfig, state storage.State) error {
	if config.SystemDomain == "" {
		config.SystemDomain = state.LB.SystemDomain
	}

	return g.gcpCreateLBs.Execute(config, state)
}
