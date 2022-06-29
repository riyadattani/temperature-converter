package temphttp_test

import (
	"net/http/httptest"
	"testing"

	"temperature-converter/pkg/converter"
	"temperature-converter/pkg/temphttp"
)

func TestConvert(t *testing.T) {
	tempConverter := converter.New()
	router := temphttp.NewRouter(tempConverter)
	server := httptest.NewServer(router)
	defer server.Close()

	t.Run("converts c to f", func(t *testing.T) {

	})
}
