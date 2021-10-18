package user_agent

import (
	"testing"
)

var ua_list = []struct {
	ua     string
	except string
}{
	{
		ua:     "Mozilla/5.0 (Linux; Android 7.1.1; OPPO R9sk) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.111 Mobile Safari/537.36",
		except: "OPPO R9sk",
	},
	{
		ua:     "Mozilla/5.0 (Linux; arm_64; Android 9; ANE-LX2J) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.136 YaBrowser/20.2.6.114.00 Mobile Safari/537.36",
		except: "ANE-LX2J",
	},
	{
		ua:     "Dalvik/1.6.0 (Linux; U; Android 4.4.2; ASUS_T00Q Build/KVT49L)/CLDC-1.1",
		except: "ASUS_T00Q",
	},
	{
		ua:     "Mozilla/5.0 (Linux; U; Android 2.3.7; en-us; Nexus One Build/FRF91) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1",
		except: "Nexus One",
	},
	{
		ua:     "Mozilla/5.0 (Linux; Android 9; ELE-AL00 Build/HUAWEIELE-AL0001; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/63.0.3239.83 Mobile Safari/537.36 T7/11.15 baiduboxapp/11.15.5.10 (Baidu; P1 9)",
		except: "ELE-AL00",
	},
	{
		ua:     "Mozilla/5.0 (Linux; U; Android 5.0.2; zh-cn; Redmi Note 3 Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/53.0.2785.146 Mobile Safari/537.36 XiaoMi/MiuiBrowser/8.8.7",
		except: "Redmi Note 3",
	},{
		ua:     "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.97 Safari/537.11",
		except: "",
	},
	{
		ua:     "Mozilla/5.0 (Linux; Android 8.0.0; MI 6 Build/OPR1.170623.027; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/76.0.3809.89 Mobile Safari/537.36 T7/11.12 swan/2.11.0 baiduboxapp/11.15.0.0 (Baidu; P1 8.0.0)",
		except: "MI 6",
	},
	{
		ua:     "Mozilla/5.0 (Linux; U; Android 2.2.1; zh-cn; HTC_Wildfire_A3333 Build/FRG83D) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1",
		except: "HTC_Wildfire_A3333",
	},
	{
		ua:     "Mozilla/5.0 (iPad; U; CPU OS 4_3_3 like Mac OS X; en-us) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8J2 Safari/6533.18.5",
		except: "iPad",
	},
	{
		ua:     "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1",
		except: "iPhone",
	},
}

func TestUserAgentModel(t *testing.T) {
	for _, item := range ua_list {
		ua := New(item.ua)
		if ua.model != item.except {
			t.Errorf("ua.model = %s \n except = %s \n", ua.model, item.except)
		}
	}
}
