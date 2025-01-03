/*
Val Town API

Testing ValsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package valgo

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/404wolf/valgo"
)

func Test_valgo_ValsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ValsAPIService RunGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var valname string

		httpRes, err := apiClient.ValsAPI.RunGet(context.Background(), valname).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ValsAPIService RunPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var valname string

		httpRes, err := apiClient.ValsAPI.RunPost(context.Background(), valname).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ValsAPIService ValsCancel", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var valId string
		var evaluationId string

		resp, httpRes, err := apiClient.ValsAPI.ValsCancel(context.Background(), valId, evaluationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ValsAPIService ValsCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ValsAPI.ValsCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ValsAPIService ValsCreateVersion", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var valId string

		resp, httpRes, err := apiClient.ValsAPI.ValsCreateVersion(context.Background(), valId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ValsAPIService ValsDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var valId string

		httpRes, err := apiClient.ValsAPI.ValsDelete(context.Background(), valId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ValsAPIService ValsDeleteVersion", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var valId string
		var version int32

		httpRes, err := apiClient.ValsAPI.ValsDeleteVersion(context.Background(), valId, version).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ValsAPIService ValsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var valId string

		resp, httpRes, err := apiClient.ValsAPI.ValsGet(context.Background(), valId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ValsAPIService ValsGetVersion", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var valId string
		var version int32

		resp, httpRes, err := apiClient.ValsAPI.ValsGetVersion(context.Background(), valId, version).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ValsAPIService ValsList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var valId string

		resp, httpRes, err := apiClient.ValsAPI.ValsList(context.Background(), valId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ValsAPIService ValsPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ValsAPI.ValsPut(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ValsAPIService ValsUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var valId string

		httpRes, err := apiClient.ValsAPI.ValsUpdate(context.Background(), valId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
