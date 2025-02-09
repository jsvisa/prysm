package v2

import (
	"testing"

	ethpb "github.com/prysmaticlabs/prysm/proto/prysm/v1alpha1"
	"github.com/prysmaticlabs/prysm/testing/require"
)

func TestBeaconState_AppendCurrentEpochAttestations(t *testing.T) {
	s, err := InitializeFromProtoUnsafe(&ethpb.BeaconStateAltair{})
	require.NoError(t, err)
	require.ErrorContains(t, "AppendCurrentEpochAttestations is not supported", s.AppendCurrentEpochAttestations(nil))
}

func TestBeaconState_AppendPreviousEpochAttestations(t *testing.T) {
	s, err := InitializeFromProtoUnsafe(&ethpb.BeaconStateAltair{})
	require.NoError(t, err)
	require.ErrorContains(t, "AppendPreviousEpochAttestations is not supported", s.AppendPreviousEpochAttestations(nil))
}
