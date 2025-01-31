package gstadapter_test

import (
	"os"
	"testing"

	"github.com/DineshKuppan/quicktime_video_hack/screencapture/gstadapter"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func skipCI(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip("Skipping testing in CI environment")
	}
}

func TestCustomPipelineParsing(t *testing.T) {
	linuxCI := os.Getenv("LINUX_CI")
	log.Infof("linuxCI: %s", linuxCI)
	if linuxCI == "true" {
		log.Info("Skipping gstreamer test on headless containerized CI")
		t.SkipNow()
	}

	_, err := gstadapter.NewWithCustomPipeline("daniel")
	assert.Error(t, err)

	_, err = gstadapter.NewWithCustomPipeline("queue name=my_filesrc ! fakesink")
	assert.Error(t, err)

	_, err = gstadapter.NewWithCustomPipeline("queue name=audio_target ! fakesink")
	assert.Error(t, err)

	gsta, err := gstadapter.NewWithCustomPipeline("rtpmux name=mux ! fakesink \n queue name=audio_target ! mux.sink_0 \n queue name=video_target ! mux.sink_1")
	assert.NoError(t, err)
	assert.NotNil(t, gsta)
}
