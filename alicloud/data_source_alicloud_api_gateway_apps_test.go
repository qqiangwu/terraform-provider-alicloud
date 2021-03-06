package alicloud

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/acctest"
	"testing"
)

func SkipTestAccAlicloudApigatewayAppsDataSource(t *testing.T) {
	rand := acctest.RandIntRange(1000000, 9999999)
	resourceId := "data.alicloud_api_gateway_apps.default"

	testAccConfig := dataSourceTestAccConfigFunc(resourceId,
		fmt.Sprintf("tf_testAccApp_%d", rand),
		dataSourceApiGatewayAppsConfigDependence)

	allConf := dataSourceTestAccConfig{
		existConfig: testAccConfig(map[string]interface{}{
			"name_regex": "${alicloud_api_gateway_app.default.name}",
		}),
		fakeConfig: testAccConfig(map[string]interface{}{
			"name_regex": "${alicloud_api_gateway_app.default.name}_fake",
		}),
	}

	var existApiGatewayAppsMapFunc = func(rand int) map[string]string {
		return map[string]string{
			"ids.#":              "1",
			"names.#":            "1",
			"names.0":            fmt.Sprintf("tf_testAccApp_%d", rand),
			"apps.#":             "1",
			"apps.0.name":        fmt.Sprintf("tf_testAccApp_%d", rand),
			"apps.0.description": "tf_testAcc api gateway description",
		}
	}

	var fakeApiGatewayAppsMapFunc = func(rand int) map[string]string {
		return map[string]string{
			"ids.#":   "0",
			"names.#": "0",
			"apps.#":  "0",
		}
	}

	var apiGatewayAppsCheckInfo = dataSourceAttr{
		resourceId:   resourceId,
		existMapFunc: existApiGatewayAppsMapFunc,
		fakeMapFunc:  fakeApiGatewayAppsMapFunc,
	}

	apiGatewayAppsCheckInfo.dataSourceTestCheck(t, rand, allConf)
}

func dataSourceApiGatewayAppsConfigDependence(name string) string {
	return fmt.Sprintf(`

variable "name" {
  default = "%s"
}

variable "description" {
  default = "tf_testAcc api gateway description"
}

resource "alicloud_api_gateway_app" "default" {
  name = "${var.name}"
  description = "${var.description}"
}

`, name)
}
