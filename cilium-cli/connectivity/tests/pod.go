// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package tests

import (
	"context"
	"fmt"

	"github.com/cilium/cilium-cli/connectivity/check"
)

// PodToPod generates one HTTP request from each client pod
// to each echo (server) pod in the test context. The remote Pod is contacted
// directly, no DNS is involved.
func PodToPod(opts ...Option) check.Scenario {
	options := &labelsOption{}
	for _, opt := range opts {
		opt(options)
	}
	return &podToPod{
		sourceLabels:      options.sourceLabels,
		destinationLabels: options.destinationLabels,
		method:            options.method,
	}
}

// podToPod implements a Scenario.
type podToPod struct {
	sourceLabels      map[string]string
	destinationLabels map[string]string
	method            string
}

func (s *podToPod) Name() string {
	return "pod-to-pod"
}

func (s *podToPod) Run(ctx context.Context, t *check.Test) {
	var i int

	for _, client := range t.Context().ClientPods() {
		client := client // copy to avoid memory aliasing when using reference
		if !hasAllLabels(client, s.sourceLabels) {
			continue
		}
		for _, echo := range t.Context().EchoPods() {
			if !hasAllLabels(echo, s.destinationLabels) {
				continue
			}
			t.NewAction(s, fmt.Sprintf("curl-%d", i), &client, echo).Run(func(a *check.Action) {
				if s.method == "" {
					a.ExecInPod(ctx, curl(echo))
				} else {
					a.ExecInPod(ctx, curl(echo, "-X", s.method))
				}

				a.ValidateFlows(ctx, client, a.GetEgressRequirements(check.FlowParameters{}))
				a.ValidateFlows(ctx, echo, a.GetIngressRequirements(check.FlowParameters{}))
			})

			i++
		}
	}
}

func PodToPodWithEndpoints(opts ...Option) check.Scenario {
	options := &labelsOption{}
	for _, opt := range opts {
		opt(options)
	}
	return &podToPodWithEndpoints{
		sourceLabels:      options.sourceLabels,
		destinationLabels: options.destinationLabels,
		method:            options.method,
	}
}

// podToPodWithEndpoints implements a Scenario.
type podToPodWithEndpoints struct {
	sourceLabels      map[string]string
	destinationLabels map[string]string
	method            string
}

func (s *podToPodWithEndpoints) Name() string {
	return "pod-to-pod-with-endpoints"
}

func (s *podToPodWithEndpoints) Run(ctx context.Context, t *check.Test) {
	var i int

	for _, client := range t.Context().ClientPods() {
		client := client // copy to avoid memory aliasing when using reference
		if !hasAllLabels(client, s.sourceLabels) {
			continue
		}
		for _, echo := range t.Context().EchoPods() {
			if !hasAllLabels(echo, s.destinationLabels) {
				continue
			}

			if s.method == "" {
				curlEndpoints(ctx, s, t, fmt.Sprintf("curl-%d", i), &client, echo)
			} else {
				curlEndpoints(ctx, s, t, fmt.Sprintf("curl-%d", i), &client, echo, "-X", s.method)
			}

			i++
		}
	}
}

func curlEndpoints(ctx context.Context, s check.Scenario, t *check.Test,
	name string, client *check.Pod, echo check.TestPeer, curlOpts ...string) {

	baseURL := fmt.Sprintf("%s://%s:%d", echo.Scheme(), echo.Address(), echo.Port())

	// Manually construct an HTTP endpoint for each API endpoint.
	for _, path := range []string{"public", "private"} {
		epName := fmt.Sprintf("%s-%s", name, path)
		url := fmt.Sprintf("%s/%s", baseURL, path)
		ep := check.HTTPEndpoint(epName, url)

		t.NewAction(s, epName, client, ep).Run(func(a *check.Action) {
			a.ExecInPod(ctx, curl(ep, curlOpts...))

			a.ValidateFlows(ctx, client, a.GetEgressRequirements(check.FlowParameters{}))
			a.ValidateFlows(ctx, ep, a.GetIngressRequirements(check.FlowParameters{}))
		})

		// Additionally test private endpoint access with HTTP header expected by policy.
		if path == "private" {
			epName += "with-header"
			ep = check.HTTPEndpointWithLabels(epName, url, map[string]string{
				"X-Very-Secret-Token": "42",
			})
			t.NewAction(s, epName, client, ep).Run(func(a *check.Action) {
				opts := make([]string, 0, len(curlOpts)+2)
				opts = append(opts, curlOpts...)
				opts = append(opts, "-H", "X-Very-Secret-Token: 42")

				a.ExecInPod(ctx, curl(ep, opts...))

				a.ValidateFlows(ctx, client, a.GetEgressRequirements(check.FlowParameters{}))
				a.ValidateFlows(ctx, ep, a.GetIngressRequirements(check.FlowParameters{}))
			})
		}
	}
}
