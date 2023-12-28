package instance

import (
	"testing"

	"github.com/AiRISTAFlowInc/flow-studio-core/support/log"
	"github.com/AiRISTAFlowInc/flow-studio-flow/model"
	"github.com/stretchr/testify/assert"
)


func TestSimpleChange(t *testing.T) {
	def := getDef()
	//fmt.Println(def.ModelID())

	ind , err := NewIndependentInstance("test","", def, nil, log.RootLogger())
	assert.Nil(t, err)
	assert.NotNil(t, ind)

	chgTrackingEnabled = true
	chgTrack := NewInstanceChangeTracker("test", 0)
	assert.NotNil(t, chgTrack)

	assert.NotNil(t,chgTrack.ExtractStep(false))
	chgTrack.SetStatus(1, model.FlowStatusActive)
	chgTrack.AttrChange(1, "key", "value")
	chgTrack.FlowCreated(ind)
	chgTrack.ExtractStep(true)
}