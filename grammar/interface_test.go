package grammar

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInterface(t *testing.T) {
	threeNet := ThreeNetResponse{code: 200, msg: "成功"}
	xfl := XflResponse{returnMsg: "成功"}

	intf := []client{&threeNet, &xfl}
	req := make(map[string]string)
	req["auth"] = "Bear token"
	req["url"] = "https://www.baidu.com"

	getResp := intf[0].get(req)
	postResp := intf[0].post(req)
	assert.Equal(t, "ThreeNet resp(get): https://www.baidu.com 成功 200", getResp)
	assert.Equal(t, "ThreeNet resp(post): Bear token 成功 200", postResp)

	xGetResp := intf[1].get(req)
	xPostResp := intf[1].post(req)
	assert.Equal(t, "Xfl resp(get): https://www.baidu.com 成功", xGetResp)
	assert.Equal(t, "Xfl resp(post): Bear token 成功", xPostResp)

	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	assert.Equal(t, "127.0.0.1\n", fmt.Sprintln(hosts["loopback"]))
	assert.Equal(t, "8.8.8.8", fmt.Sprint(hosts["googleDNS"]))

}
