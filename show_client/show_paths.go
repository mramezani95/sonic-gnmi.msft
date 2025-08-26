package show_client

import (
	sdc "github.com/sonic-net/sonic-gnmi/sonic_data_client"
)

// All SHOW path and getters are defined here
func init() {
	sdc.RegisterCliPath(
		[]string{"SHOW", "reboot-cause"},
		getPreviousRebootCause,
		map[string]string{
			"history": "show/reboot-cause/history: Show history of reboot-cause",
		},
	)
	sdc.RegisterCliPath(
		[]string{"SHOW", "reboot-cause", "history"},
		getRebootCauseHistory,
		nil,
	)
	sdc.RegisterCliPath(
		[]string{"SHOW", "clock"},
		getDate,
		map[string]string{
			"timezones": "show/clock/timezones: List of available timezones",
		},
		showCmdOptionVerbose,
	)
	sdc.RegisterCliPath(
		[]string{"SHOW", "clock", "timezones"},
		getDateTimezone,
		nil,
		showCmdOptionVerbose,
	)
	sdc.RegisterCliPath(
		[]string{"SHOW", "ipv6", "bgp", "summary"},
		getIPv6BGPSummary,
		nil,
		sdc.UnimplementedOption(showCmdOptionNamespace),
		showCmdOptionDisplay,
	)
	sdc.RegisterCliPath(
		[]string{"SHOW", "interface", "counters"},
		getInterfaceCounters,
		nil,
		sdc.UnimplementedOption(showCmdOptionNamespace),
		showCmdOptionDisplay,
		showCmdOptionInterfaces,
		showCmdOptionPeriod,
		showCmdOptionJson,
		showCmdOptionVerbose,
	)
	sdc.RegisterCliPath(
		[]string{"SHOW", "interface", "errors"},
		getInterfaceErrors,
		nil,
		sdc.RequiredOption(showCmdOptionInterface),
	)
	sdc.RegisterCliPath(
		[]string{"SHOW", "interface", "fec", "status"},
		getInterfaceFecStatus,
		nil,
		showCmdOptionInterface,
	)
	sdc.RegisterCliPath(
		[]string{"SHOW", "watermark", "telemetry", "interval"},
		getWatermarkTelemetryInterval,
		nil,
	)
	sdc.RegisterCliPath(
		[]string{"SHOW", "interface", "transceiver", "error-status"},
		getTransceiverErrorStatus,
		nil,
		showCmdOptionVerbose,
		sdc.UnimplementedOption(showCmdOptionNamespace),
		sdc.UnimplementedOption(showCmdOptionFetchFromHW),
		showCmdOptionInterface,
	)
	sdc.RegisterCliPath(
		[]string{"SHOW", "interface", "transceiver", "presence"},
		getInterfaceTransceiverPresence,
		nil,
		showCmdOptionInterface,
	)
	sdc.RegisterCliPath(
		[]string{"SHOW", "headroom_pool", "watermark"},
		getHeadroomPoolWatermark,
		nil,
		sdc.UnimplementedOption(showCmdOptionNamespace),
	)
	sdc.RegisterCliPath(
		[]string{"SHOW", "headroom_pool", "persistent-watermark"},
		getHeadroomPoolPersistentWatermark,
		nil,
		sdc.UnimplementedOption(showCmdOptionNamespace),
	)
	sdc.RegisterCliPath(
		[]string{"SHOW", "buffer_pool", "watermark"},
		getBufferPoolWatermark,
		nil,
		sdc.UnimplementedOption(showCmdOptionNamespace),
	)
	sdc.RegisterCliPath(
		[]string{"SHOW", "buffer_pool", "persistent-watermark"},
		getBufferPoolPersistentWatermark,
		nil,
		sdc.UnimplementedOption(showCmdOptionNamespace),
	)
	sdc.RegisterCliPath(
		[]string{"SHOW", "mac", "aging-time"},
		getMacAgingTime,
		nil,
	)
	sdc.RegisterCliPath(
		[]string{"SHOW", "ipv6", "interfaces"},
		getIPv6Interfaces,
		nil,
		sdc.UnimplementedOption(showCmdOptionNamespace),
		showCmdOptionDisplay,
	)
	sdc.RegisterCliPath(
		[]string{"SHOW", "lldp", "table"},
		getLLDPTable,
		nil,
	)
	sdc.RegisterCliPath(
		[]string{"SHOW", "interface", "alias"},
		getInterfaceAlias,
		nil,
		showCmdOptionInterface,
	)
	sdc.RegisterCliPath(
		[]string{"SHOW", "queue", "counters"},
		getQueueCounters,
		nil,
		showCmdOptionInterfaces,
		showCmdOptionNonzero,
		showCmdOptionTrim,
	)
}
