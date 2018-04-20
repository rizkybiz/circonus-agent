// Copyright © 2017 Circonus, Inc. <support@circonus.com>
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package config

import (
	"path/filepath"

	"github.com/circonus-labs/circonus-agent/internal/config/defaults"
)

// Log defines the running config.log structure
type Log struct {
	Level  string `json:"level" yaml:"level" toml:"level"`
	Pretty bool   `json:"pretty" yaml:"pretty" toml:"pretty"`
}

// API defines the running config.api structure
type API struct {
	App    string `json:"app" yaml:"app" toml:"app"`
	CAFile string `mapstructure:"ca_file" json:"ca_file" yaml:"ca_file" toml:"ca_file"`
	Key    string `json:"key" yaml:"key" toml:"key"`
	URL    string `json:"url" yaml:"url" toml:"url"`
}

// ReverseCreateCheckOptions defines the running config.reverse.check structure
type ReverseCreateCheckOptions struct {
	Broker string `json:"broker" yaml:"broker" toml:"broker"`
	Tags   string `json:"tags" yaml:"tags" toml:"tags"`
	Title  string `json:"title" yaml:"title" toml:"title"`
}

// Check defines the check parameters
type Check struct {
	Broker           string `json:"broker" yaml:"broker" toml:"broker"`
	BundleID         string `mapstructure:"bundle_id" json:"bundle_id" yaml:"bundle_id" toml:"bundle_id"`
	Create           bool   `mapstructure:"create" json:"create" yaml:"create" toml:"create"`
	EnableNewMetrics bool   `mapstructure:"enable_new_metrics" json:"enable_new_metrics" yaml:"enable_new_metrics" toml:"enable_new_metrics"`
	MetricStateDir   string `mapstructure:"metric_state_dir" json:"metric_state_dir" yaml:"metric_state_dir" toml:"metric_state_dir"`
	MetricRefreshTTL string `mapstructure:"metric_refresh_ttl" json:"metric_refresh_ttl" yaml:"metric_refresh_ttl" toml:"metric_refresh_ttl"`
	Tags             string `json:"tags" yaml:"tags" toml:"tags"`
	Target           string `mapstructure:"target" json:"target" yaml:"target" toml:"target"`
	Title            string `json:"title" yaml:"title" toml:"title"`
}

// Reverse defines the running config.reverse structure
type Reverse struct {
	BrokerCAFile string `mapstructure:"broker_ca_file" json:"broker_ca_file" yaml:"broker_ca_file" toml:"broker_ca_file"`
	Enabled      bool   `json:"enabled" yaml:"enabled" toml:"enabled"`
}

// SSL defines the running config.ssl structure
type SSL struct {
	CertFile string `mapstructure:"cert_file" json:"cert_file" yaml:"cert_file" toml:"cert_file"`
	KeyFile  string `mapstructure:"key_file" json:"key_file" yaml:"key_file" toml:"key_file"`
	Listen   string `json:"listen" yaml:"listen" toml:"listen"`
	Verify   bool   `json:"verify" yaml:"verify" toml:"verify"`
}

// StatsDHost defines the running config.statsd.host structure
type StatsDHost struct {
	Category     string `json:"category" yaml:"category" toml:"category"`
	MetricPrefix string `mapstructure:"metric_prefix" json:"metric_prefix" yaml:"metric_prefix" toml:"metric_prefix"`
}

// StatsDGroup defines the running config.statsd.group structure
type StatsDGroup struct {
	CheckBundleID string `mapstructure:"check_bundle_id" json:"check_bundle_id" yaml:"check_bundle_id" toml:"check_bundle_id"`
	Counters      string `json:"counters" yaml:"counters" toml:"counters"`
	Gauges        string `json:"gauges" yaml:"gauges" toml:"gauges"`
	MetricPrefix  string `mapstructure:"metric_prefix" json:"metric_prefix" yaml:"metric_prefix" toml:"metric_prefix"`
	Sets          string `json:"sets" yaml:"sets" toml:"sets"`
}

// StatsD defines the running config.statsd structure
type StatsD struct {
	Disabled bool        `json:"disabled" yaml:"disabled" toml:"disabled"`
	Group    StatsDGroup `json:"group" yaml:"group" toml:"group"`
	Host     StatsDHost  `json:"host" yaml:"host" toml:"host"`
	Port     string      `json:"port" yaml:"port" toml:"port"`
}

// Config defines the running config structure
type Config struct {
	API              API      `json:"api" yaml:"api" toml:"api"`
	Check            Check    `json:"check" yaml:"check" toml:"check"`
	Collectors       []string `json:"collectors" yaml:"collectors" toml:"collectors"`
	Debug            bool     `json:"debug" yaml:"debug" toml:"debug"`
	DebugCGM         bool     `mapstructure:"debug_cgm" json:"debug_cgm" yaml:"debug_cgm" toml:"debug_cgm"`
	DebugDumpMetrics string   `mapstructure:"debug_dump_metrics" json:"debug_dump_metrics" yaml:"debug_dump_metrics" toml:"debug_dump_metrics"`
	Listen           []string `json:"listen" yaml:"listen" toml:"listen"`
	ListenSocket     []string `mapstructure:"listen_socket" json:"listen_socket" yaml:"listen_socket" toml:"listen_socket"`
	Log              Log      `json:"log" yaml:"log" toml:"log"`
	PluginDir        string   `mapstructure:"plugin_dir" json:"plugin_dir" yaml:"plugin_dir" toml:"plugin_dir"`
	PluginTTLUnits   string   `mapstructure:"plugin_ttl_units" json:"plugin_ttl_units" yaml:"plugin_ttl_units" toml:"plugin_ttl_units"`
	Reverse          Reverse  `json:"reverse" yaml:"reverse" toml:"reverse"`
	SSL              SSL      `json:"ssl" yaml:"ssl" toml:"ssl"`
	StatsD           StatsD   `json:"statsd" yaml:"statsd" toml:"statsd"`
}

type cosiCheckConfig struct {
	CID string `json:"_cid"`
}

//
// NOTE: adding a Key* MUST be reflected in the Config structures above
//
const (
	// KeyAPICAFile custom ca for circonus api (e.g. inside)
	KeyAPICAFile = "api.ca_file"

	// KeyAPITokenApp circonus api token key application name
	KeyAPITokenApp = "api.app"

	// KeyAPITokenKey circonus api token key
	KeyAPITokenKey = "api.key"

	// KeyAPIURL custom circonus api url (e.g. inside)
	KeyAPIURL = "api.url"

	// KeyDebug enables debug messages
	KeyDebug = "debug"

	// KeyDebugCGM enables debug messages for circonus-gometrics
	KeyDebugCGM = "debug_cgm"

	// KeyDebugDumpMetrics enables dumping metrics to a file as they are submitted to circonus
	// it should contain a directory name where the user running circonus-agentd has write
	// permissions. metrics will be dumped for each _successful_ request.
	KeyDebugDumpMetrics = "debug_dump_metrics"

	// KeyListen primary address and port to listen on
	KeyListen = "listen"

	// KeyListenSocket identifies one or more unix socket files to create
	KeyListenSocket = "listen_socket"

	// KeyLogLevel logging level (panic, fatal, error, warn, info, debug, disabled)
	KeyLogLevel = "log.level"

	// KeyLogPretty output formatted log lines (for running in foreground)
	KeyLogPretty = "log.pretty"

	// KeyPluginDir plugin directory
	KeyPluginDir = "plugin_dir"

	// KeyPluginTTLUnits plugin run ttl units
	KeyPluginTTLUnits = "plugin_ttl_units"

	// KeyReverse indicates whether to use reverse connections
	KeyReverse = "reverse.enabled"

	// KeyReverseBrokerCAFile custom broker ca file
	KeyReverseBrokerCAFile = "reverse.broker_ca_file"

	// KeyShowConfig - show configuration and exit
	KeyShowConfig = "show-config"

	// KeyShowVersion - show version information and exit
	KeyShowVersion = "version"

	// KeySSLCertFile pem certificate file for SSL
	KeySSLCertFile = "ssl.cert_file"

	// KeySSLKeyFile key for ssl.cert_file
	KeySSLKeyFile = "ssl.key_file"

	// KeySSLListen ssl address and prot to listen on
	KeySSLListen = "ssl.listen"

	// KeySSLVerify controls verification for ssl connections
	KeySSLVerify = "ssl.verify"

	// KeyStatsdDisabled disables the default statsd listener
	KeyStatsdDisabled = "statsd.disabled"

	// KeyStatsdGroupCID circonus check bundle id for "group" metrics sent to statsd
	KeyStatsdGroupCID = "statsd.group.check_bundle_id"

	// KeyStatsdGroupCounters operator for group counters (sum|average)
	KeyStatsdGroupCounters = "statsd.group.counters"

	// KeyStatsdGroupGauges operator for group gauges (sum|average)
	KeyStatsdGroupGauges = "statsd.group.gauges"

	// KeyStatsdGroupPrefix metrics prefixed with this string are considered "group" metrics
	KeyStatsdGroupPrefix = "statsd.group.metric_prefix"

	// KeyStatsdGroupSets operator for group sets (sum|average)
	KeyStatsdGroupSets = "statsd.group.sets"

	// KeyStatsdHostCategory "plugin" name to put metrics sent to host
	KeyStatsdHostCategory = "statsd.host.category"

	// KeyStatsdHostPrefix metrics prefixed with this string are considered "host" metrics
	KeyStatsdHostPrefix = "statsd.host.metric_prefix"

	// KeyStatsdPort port for statsd listener (note, address will always be 'localhost')
	KeyStatsdPort = "statsd.port"

	// KeyCollectors defines the builtin collectors to enable
	KeyCollectors = "collectors"

	// KeyDisableGzip disables gzip on http responses
	KeyDisableGzip = "server.disable_gzip"

	// KeyCheckBundleID the check bundle id to use
	KeyCheckBundleID = "check.bundle_id"

	// KeyCheckTarget the check bundle target to use to search for or create a check bundle
	// note: if not using reverse, this must be an IP address reachable by the broker
	KeyCheckTarget = "check.target"

	// KeyCheckEnableNewMetrics toggles automatically enabling new metrics
	KeyCheckEnableNewMetrics = "check.enable_new_metrics"
	// KeyCheckMetricStateDir defines the path where check metric state will be maintained when --check-enable-new-metrics is turned on
	KeyCheckMetricStateDir = "check.metric_state_dir"
	// KeyCheckMetricRefreshTTL determines how often to refresh check bundle metrics from API when enable new metrics is turned on
	KeyCheckMetricRefreshTTL = "check.metric_refresh_ttl"

	// KeyCheckCreate toggles creating a new check bundle when a check bundle id is not supplied
	KeyCheckCreate = "check.create"

	// KeyCheckBroker a specific broker ID to use when creating a new check bundle
	KeyCheckBroker = "check.broker"

	// KeyCheckTitle a specific title to use when creating a new check bundle
	KeyCheckTitle = "check.title"

	// KeyCheckTags a specific set of tags to use when creating a new check bundle
	KeyCheckTags = "check.tags"

	cosiName = "cosi"
)

var (
	cosiCfgFile = filepath.Join(defaults.BasePath, "..", cosiName, "etc", "cosi.json")

	// MetricNameSeparator defines character used to delimit metric name parts
	MetricNameSeparator = defaults.MetricNameSeparator // var, TBD whether it will become configurable
)
