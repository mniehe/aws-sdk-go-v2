// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisvideowebrtcstorage

import (
	"context"
	"github.com/mniehe/aws-sdk-go-v2/aws"
	"os"
	"reflect"
	"testing"
)

type mockConfigSource struct {
	global  string
	service string
	ignore  bool
}

// GetIgnoreConfiguredEndpoints is used in knowing when to disable configured
// endpoints feature.
func (m mockConfigSource) GetIgnoreConfiguredEndpoints(context.Context) (bool, bool, error) {
	return m.ignore, m.ignore, nil
}

// GetServiceBaseEndpoint is used to retrieve a normalized SDK ID for use
// with configured endpoints.
func (m mockConfigSource) GetServiceBaseEndpoint(ctx context.Context, sdkID string) (string, bool, error) {
	if m.service != "" {
		return m.service, true, nil
	}
	return "", false, nil
}

func TestResolveBaseEndpoint(t *testing.T) {
	cases := map[string]struct {
		envGlobal      string
		envService     string
		envIgnore      bool
		configGlobal   string
		configService  string
		configIgnore   bool
		clientEndpoint *string
		expectURL      *string
	}{
		"env ignore": {
			envGlobal:     "https://env-global.dev",
			envService:    "https://env-kinesis-video-webrtc-storage.dev",
			envIgnore:     true,
			configGlobal:  "http://config-global.dev",
			configService: "http://config-kinesis-video-webrtc-storage.dev",
			expectURL:     nil,
		},
		"env global": {
			envGlobal:     "https://env-global.dev",
			configGlobal:  "http://config-global.dev",
			configService: "http://config-kinesis-video-webrtc-storage.dev",
			expectURL:     aws.String("https://env-global.dev"),
		},
		"env service": {
			envGlobal:     "https://env-global.dev",
			envService:    "https://env-kinesis-video-webrtc-storage.dev",
			configGlobal:  "http://config-global.dev",
			configService: "http://config-kinesis-video-webrtc-storage.dev",
			expectURL:     aws.String("https://env-kinesis-video-webrtc-storage.dev"),
		},
		"config ignore": {
			envGlobal:     "https://env-global.dev",
			envService:    "https://env-kinesis-video-webrtc-storage.dev",
			configGlobal:  "http://config-global.dev",
			configService: "http://config-kinesis-video-webrtc-storage.dev",
			configIgnore:  true,
			expectURL:     nil,
		},
		"config global": {
			configGlobal: "http://config-global.dev",
			expectURL:    aws.String("http://config-global.dev"),
		},
		"config service": {
			configGlobal:  "http://config-global.dev",
			configService: "http://config-kinesis-video-webrtc-storage.dev",
			expectURL:     aws.String("http://config-kinesis-video-webrtc-storage.dev"),
		},
		"client": {
			envGlobal:      "https://env-global.dev",
			envService:     "https://env-kinesis-video-webrtc-storage.dev",
			configGlobal:   "http://config-global.dev",
			configService:  "http://config-kinesis-video-webrtc-storage.dev",
			clientEndpoint: aws.String("https://client-kinesis-video-webrtc-storage.dev"),
			expectURL:      aws.String("https://client-kinesis-video-webrtc-storage.dev"),
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			os.Clearenv()

			awsConfig := aws.Config{}
			ignore := c.envIgnore || c.configIgnore

			if c.configGlobal != "" && !ignore {
				awsConfig.BaseEndpoint = aws.String(c.configGlobal)
			}

			if c.envGlobal != "" {
				t.Setenv("AWS_ENDPOINT_URL", c.envGlobal)
				if !ignore {
					awsConfig.BaseEndpoint = aws.String(c.envGlobal)
				}
			}

			if c.envService != "" {
				t.Setenv("AWS_ENDPOINT_URL_KINESIS_VIDEO_WEBRTC_STORAGE", c.envService)
			}

			awsConfig.ConfigSources = []interface{}{
				mockConfigSource{
					global:  c.envGlobal,
					service: c.envService,
					ignore:  c.envIgnore,
				},
				mockConfigSource{
					global:  c.configGlobal,
					service: c.configService,
					ignore:  c.configIgnore,
				},
			}

			client := NewFromConfig(awsConfig, func(o *Options) {
				if c.clientEndpoint != nil {
					o.BaseEndpoint = c.clientEndpoint
				}
			})

			if e, a := c.expectURL, client.options.BaseEndpoint; !reflect.DeepEqual(e, a) {
				t.Errorf("expect endpoint %v , got %v", e, a)
			}
		})
	}
}
