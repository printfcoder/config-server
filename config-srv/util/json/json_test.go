package json

import "testing"

func TestPropertiesToJSON(t *testing.T) {
	m := map[string]string{
		"spring.application.name":              "fin-qingniao-api",
		"config.center.namespaces":             "application,fin.trace,fin.registry,fin.log4j2,fin.schedule_config,fin.public_redis,es,redis,fin.finz_local_gateway",
		"logging.path":                         "logs",
		"logging.config":                       "classpath:com/oppo/log/log4j2/log-conf.xml",
		"logging.file":                         "${logging.path}/${spring.application.name}-console.log",
		"com.alipay.sofa.rpc.registry.address": "${common.zookeeper.address}?file=${spring.application.name}/registry",
		"server.port":                          "10018",
		"com.alipay.sofa.rpc.restPort":         "11018",
		"com.alipay.sofa.rpc.restContextPath":  "/qingniao",
		"com.alipay.sofa.lookout.prometheus-exporter-server-port": "12018",
	}

	out, err := PropertiesToJSON(m)
	if err != nil {
		t.Error(err)
	}

	t.Log(out)
}
