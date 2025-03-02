//go:build integration
// +build integration

package shield

import (
	"context"
	"testing"
	"time"

	"github.com/mniehe/aws-sdk-go-v2/service/internal/integrationtest"
	"github.com/mniehe/aws-sdk-go-v2/service/shield"
)

func TestInteg_00_ListAttacks(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	cfg, err := integrationtest.LoadConfigWithDefaultRegion("us-east-1")
	if err != nil {
		t.Fatalf("failed to load config, %v", err)
	}

	client := shield.NewFromConfig(cfg)
	params := &shield.ListAttacksInput{}
	_, err = client.ListAttacks(ctx, params)
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
}
