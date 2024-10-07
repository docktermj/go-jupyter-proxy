package httpproxyservice

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	httpProxyServiceSingleton HTTPProxyService
)

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

func TestBasicHTTPProxyService(test *testing.T) {
	ctx := context.TODO()
	httpProxyService := getTestObject(ctx, test)
	_ = httpProxyService.Handler(ctx)
	testError(ctx, test, nil)

}

// ----------------------------------------------------------------------------
// Internal functions
// ----------------------------------------------------------------------------

func getTestObject(ctx context.Context, test *testing.T) HTTPProxyService {
	_ = ctx
	_ = test
	if httpProxyServiceSingleton == nil {
		httpProxyServiceSingleton = &BasicHTTPProxyService{}
	}
	return httpProxyServiceSingleton
}

func testError(ctx context.Context, test *testing.T, err error) {
	_ = ctx
	if err != nil {
		test.Log("Error:", err.Error())
		assert.FailNow(test, err.Error())
	}
}
