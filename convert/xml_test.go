package convert

import "testing"

func TestMap2Xml(t *testing.T) {
	m := map[string]interface{}{
		"out_trade_no": "abc123",
		"total_fee":    32,
		"is_subscribe": true,
		"userinfo": map[string]interface{}{
			"nickname":   "snoopy",
			"avatar_url": "https://a.png",
		},
	}

	_, err := Map2Xml(m)
	if err != nil {
		t.Fatalf("map2xml failed: %v", err)
	}
}
