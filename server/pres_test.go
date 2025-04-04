package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tinode/chat/server/store/types"
)

// Test generated using Keploy
func TestPresOfflineFilter_VariousScenarios_777(t *testing.T) {
	tests := []struct {
		name     string
		mode     types.AccessMode
		what     string
		pf       *presFilters
		expected bool
	}{
		// Always true cases
		{name: "what=acs", mode: types.ModeRead, what: "acs", pf: nil, expected: true},
		{name: "what=gone", mode: types.ModeRead, what: "gone", pf: nil, expected: true},
		{name: "what=upd, mode=Joiner", mode: types.ModeJoin, what: "upd", pf: nil, expected: true},
		{name: "what=upd, mode!=Joiner", mode: types.ModeRead, what: "upd", pf: nil, expected: false}, // Fails IsJoiner, then needs IsPresencer which is also false

		// Mode check (IsPresencer)
		{name: "mode=Presencer, pf=nil", mode: types.ModePres, what: "on", pf: nil, expected: true},
		{name: "mode!=Presencer, pf=nil", mode: types.ModeRead, what: "on", pf: nil, expected: false},

		// FilterIn checks
		{name: "mode=Presencer, filterIn=None", mode: types.ModePres, what: "on", pf: &presFilters{filterIn: types.ModeNone}, expected: true},
		{name: "mode=Presencer+Read, filterIn=Read (match)", mode: types.ModePres | types.ModeRead, what: "on", pf: &presFilters{filterIn: types.ModeRead}, expected: true},
		{name: "mode=Presencer+Read, filterIn=Write (no match)", mode: types.ModePres | types.ModeRead, what: "on", pf: &presFilters{filterIn: types.ModeWrite}, expected: false},

		// FilterOut checks
		{name: "mode=Presencer, filterOut=None", mode: types.ModePres, what: "on", pf: &presFilters{filterOut: types.ModeNone}, expected: true},
		{name: "mode=Presencer+Read, filterOut=Read (match)", mode: types.ModePres | types.ModeRead, what: "on", pf: &presFilters{filterOut: types.ModeRead}, expected: false},
		{name: "mode=Presencer+Read, filterOut=Write (no match)", mode: types.ModePres | types.ModeRead, what: "on", pf: &presFilters{filterOut: types.ModeWrite}, expected: true},

		// Combined filters
		{name: "mode=Presencer+Read, filterIn=Read, filterOut=Write (pass)", mode: types.ModePres | types.ModeRead, what: "on", pf: &presFilters{filterIn: types.ModeRead, filterOut: types.ModeWrite}, expected: true},
		{name: "mode=Presencer+Read, filterIn=Read, filterOut=Read (fail out)", mode: types.ModePres | types.ModeRead, what: "on", pf: &presFilters{filterIn: types.ModeRead, filterOut: types.ModeRead}, expected: false},
		{name: "mode=Presencer+Read, filterIn=Write, filterOut=Write (fail in)", mode: types.ModePres | types.ModeRead, what: "on", pf: &presFilters{filterIn: types.ModeWrite, filterOut: types.ModeWrite}, expected: false},
		{name: "mode=Presencer+Read+Write, filterIn=Read, filterOut=Write (fail out)", mode: types.ModePres | types.ModeRead | types.ModeWrite, what: "on", pf: &presFilters{filterIn: types.ModeRead, filterOut: types.ModeWrite}, expected: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := presOfflineFilter(tt.mode, tt.what, tt.pf)
			assert.Equal(t, tt.expected, result)
		})
	}
}
