package json

import (
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestConvertSynopsysDetect(t *testing.T) {
	tests := []struct {
		json string
		want Node
	}{
		{
			json: `
{
  "spanId": "de0464382a8230c3",
  "traceId": "bc02fa7897ebd5d4f79323768403caa2",
  "parent": "synopsysDetect",
  "all-info": "span(name: synopsys_detect, spanId: de0464382a8230c3, parentSpanId: 65d2aa46fdde017d, traceId: bc02fa7897ebd5d4f79323768403caa2, attr: ci.pipeline.run.user:SYSTEM;harness-attribute:{\n  \"detectProperties\" : \"--detect.project.name=jenkinstest --detect.project.version.name=v1.0\"\n};harness-others:-DISPLAY_NAME-staticField com.synopsys.integration.jenkins.detect.extensions.pipeline.DetectPipelineStep DISPLAY_NAME-com.synopsys.integration.jenkins.detect.extensions.pipeline.DetectPipelineStep.DISPLAY_NAME-class java.lang.String-PIPELINE_NAME-staticField com.synopsys.integration.jenkins.detect.extensions.pipeline.DetectPipelineStep PIPELINE_NAME-com.synopsys.integration.jenkins.detect.extensions.pipeline.DetectPipelineStep.PIPELINE_NAME-class java.lang.String;jenkins.pipeline.step.id:45;jenkins.pipeline.step.name:Synopsys Detect;jenkins.pipeline.step.plugin.name:blackduck-detect;jenkins.pipeline.step.plugin.version:9.0.0;jenkins.pipeline.step.type:synopsys_detect;)",
  "name": "synopsysDetect #1",
  "attributesMap": {
    "harness-others": "-DISPLAY_NAME-staticField com.synopsys.integration.jenkins.detect.extensions.pipeline.DetectPipelineStep DISPLAY_NAME-com.synopsys.integration.jenkins.detect.extensions.pipeline.DetectPipelineStep.DISPLAY_NAME-class java.lang.String-PIPELINE_NAME-staticField com.synopsys.integration.jenkins.detect.extensions.pipeline.DetectPipelineStep PIPELINE_NAME-com.synopsys.integration.jenkins.detect.extensions.pipeline.DetectPipelineStep.PIPELINE_NAME-class java.lang.String",
    "jenkins.pipeline.step.name": "Synopsys Detect",
    "ci.pipeline.run.user": "SYSTEM",
    "jenkins.pipeline.step.id": "7",
    "jenkins.pipeline.step.type": "synopsys_detect",
    "harness-attribute": "{\n  \"detectProperties\" : \"--detect.project.name=jenkinstest --detect.project.version.name=v1.0\"\n}",
    "jenkins.pipeline.step.plugin.name": "blackduck-detect",
    "jenkins.pipeline.step.plugin.version": "9.0.0"
  },
  "type": "Run Phase Span",
  "parentSpanId": "65d2aa46fdde017d",
  "parameterMap": {"detectProperties": "--detect.project.name=jenkinstest --detect.project.version.name=v1.0"},
  "spanName": "synopsys_detect"
}
			`,
			want: Node{
				AttributesMap: map[string]string{
					"ci.pipeline.run.user":                 "SYSTEM",
					"jenkins.pipeline.step.id":             "7",
					"jenkins.pipeline.step.name":           "Synopsys Detect",
					"jenkins.pipeline.step.plugin.name":    "blackduck-detect",
					"jenkins.pipeline.step.plugin.version": "9.0.0",
					"jenkins.pipeline.step.type":           "synopsys_detect",
					"harness-attribute":                    "{\n  \"detectProperties\" : \"--detect.project.name=jenkinstest --detect.project.version.name=v1.0\"\n}",
					"harness-others":                       "-DISPLAY_NAME-staticField com.synopsys.integration.jenkins.detect.extensions.pipeline.DetectPipelineStep DISPLAY_NAME-com.synopsys.integration.jenkins.detect.extensions.pipeline.DetectPipelineStep.DISPLAY_NAME-class java.lang.String-PIPELINE_NAME-staticField com.synopsys.integration.jenkins.detect.extensions.pipeline.DetectPipelineStep PIPELINE_NAME-com.synopsys.integration.jenkins.detect.extensions.pipeline.DetectPipelineStep.PIPELINE_NAME-class java.lang.String",
				},
				Name:         "synopsysDetect #1",
				Parent:       "synopsysDetect",
				ParentSpanId: "65d2aa46fdde017d",
				SpanId:       "de0464382a8230c3",
				SpanName:     "synopsys_detect",
				TraceId:      "bc02fa7897ebd5d4f79323768403caa2",
				Type:         "Run Phase Span",
				ParameterMap: map[string]any{"detectProperties": string("--detect.project.name=jenkinstest --detect.project.version.name=v1.0")},
			},
		},
	}
	for i, test := range tests {
		got := new(Node)
		if err := json.Unmarshal([]byte(test.json), got); err != nil {
			t.Error(err)
			return
		}
		if diff := cmp.Diff(got, &test.want); diff != "" {
			t.Errorf("Unexpected parsing results for test %v", i)
			t.Log(diff)
		}
	}
}
