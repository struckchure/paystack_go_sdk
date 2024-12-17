/*
Paystack

Testing DedicatedVirtualAccountAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package paystack

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	openapiclient "github.com/struckchure/paystack_go_sdk"
)

func Test_paystack_DedicatedVirtualAccountAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DedicatedVirtualAccountAPIService DedicatedAccountAddSplit", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DedicatedVirtualAccountAPI.DedicatedAccountAddSplit(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DedicatedVirtualAccountAPIService DedicatedAccountAvailableProviders", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DedicatedVirtualAccountAPI.DedicatedAccountAvailableProviders(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DedicatedVirtualAccountAPIService DedicatedAccountCreate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DedicatedVirtualAccountAPI.DedicatedAccountCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DedicatedVirtualAccountAPIService DedicatedAccountDeactivate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var accountId string

		resp, httpRes, err := apiClient.DedicatedVirtualAccountAPI.DedicatedAccountDeactivate(context.Background(), accountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DedicatedVirtualAccountAPIService DedicatedAccountFetch", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var accountId string

		resp, httpRes, err := apiClient.DedicatedVirtualAccountAPI.DedicatedAccountFetch(context.Background(), accountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DedicatedVirtualAccountAPIService DedicatedAccountList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DedicatedVirtualAccountAPI.DedicatedAccountList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DedicatedVirtualAccountAPIService DedicatedAccountRemoveSplit", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DedicatedVirtualAccountAPI.DedicatedAccountRemoveSplit(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
