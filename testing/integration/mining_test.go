package integration

import (
	"math/rand"
	"testing"
	"time"

	"github.com/brics18/bricsd/app/appmessage"
	"github.com/brics18/bricsd/domain/consensus/model/externalapi"
	"github.com/brics18/bricsd/domain/consensus/utils/mining"
)

func mineNextBlock(t *testing.T, harness *appHarness) *externalapi.DomainBlock {
	blockTemplate, err := harness.rpcClient.GetBlockTemplate(harness.miningAddress, "integration")
	if err != nil {
		t.Fatalf("Error getting block template: %+v", err)
	}

	block, err := appmessage.RPCBlockToDomainBlock(blockTemplate.Block)
	if err != nil {
		t.Fatalf("Error converting block: %s", err)
	}

	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	mining.SolveBlock(block, rd)

	_, err = harness.rpcClient.SubmitBlockAlsoIfNonDAA(block)
	if err != nil {
		t.Fatalf("Error submitting block: %s", err)
	}

	return block
}
