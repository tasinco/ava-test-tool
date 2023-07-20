package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/ava-labs/avalanchego/app"
	"github.com/ava-labs/avalanchego/config"
	"github.com/ava-labs/avalanchego/node"
	"golang.org/x/sync/errgroup"
)

const (
	baseDir   = "/tmp/t"
	baseDB    = baseDir + "/db"
	baseLogs  = baseDir + "/logs"
	pluginDir = "./"
	certs     = "./certs/"
)

func main() {
	// clean up old db's
	os.RemoveAll(baseDir)

	var baseHTTPPort uint = 9650
	baseStakingPort := baseHTTPPort + 1
	baseStakingPortS := fmt.Sprintf("%d", baseStakingPort)
	bootstrapIPs := "127.0.0.1:" + baseStakingPortS
	bootstrapIds := "NodeID-7Xhw2mDxuDS44j42TCB6U5579esbSt3Lg"
	indexEnabled := true
	loglevel := "info"

	df1 := defaultFlags()
	df1.LogLevel = loglevel
	df1.StakingEnabled = true
	df1.HTTPPort = baseHTTPPort
	df1.StakingPort = baseStakingPort
	df1.BootstrapIPs = ""
	df1.StakingTLSCertFile = certs + "keys1/staker.crt"
	df1.StakingTLSKeyFile = certs + "keys1/staker.key"
	df1.IndexEnabled = indexEnabled
	nodeConfig1, err := createNodeConfig(flagsToArgs(df1))
	if err != nil {
		log.Fatal(err)
	}

	df2 := defaultFlags()
	df2.LogLevel = loglevel
	df2.StakingEnabled = true
	df2.HTTPPort = 9652
	df2.StakingPort = 9653
	df2.BootstrapIPs = bootstrapIPs
	df2.BootstrapIDs = bootstrapIds
	df2.StakingTLSCertFile = certs + "keys2/staker.crt"
	df2.StakingTLSKeyFile = certs + "keys2/staker.key"
	df2.IndexEnabled = indexEnabled
	nodeConfig2, err := createNodeConfig(flagsToArgs(df2))
	if err != nil {
		log.Fatal(err)
	}

	df3 := defaultFlags()
	df3.LogLevel = loglevel
	df3.StakingEnabled = true
	df3.HTTPPort = 9654
	df3.StakingPort = 9655
	df3.BootstrapIPs = bootstrapIPs
	df3.BootstrapIDs = bootstrapIds
	df3.StakingTLSCertFile = certs + "keys3/staker.crt"
	df3.StakingTLSKeyFile = certs + "keys3/staker.key"
	df3.IndexEnabled = indexEnabled
	nodeConfig3, err := createNodeConfig(flagsToArgs(df3))
	if err != nil {
		log.Fatal(err)
	}

	df4 := defaultFlags()
	df4.LogLevel = loglevel
	df4.StakingEnabled = true
	df4.HTTPPort = 9656
	df4.StakingPort = 9657
	df4.BootstrapIPs = bootstrapIPs
	df4.BootstrapIDs = bootstrapIds
	df4.StakingTLSCertFile = certs + "keys4/staker.crt"
	df4.StakingTLSKeyFile = certs + "keys4/staker.key"
	df4.IndexEnabled = indexEnabled
	nodeConfig4, err := createNodeConfig(flagsToArgs(df4))
	if err != nil {
		log.Fatal(err)
	}
	df5 := defaultFlags()
	df5.LogLevel = loglevel
	df5.StakingEnabled = true
	df5.HTTPPort = 9658
	df5.StakingPort = 9659
	df5.BootstrapIPs = bootstrapIPs
	df5.BootstrapIDs = bootstrapIds
	df5.StakingTLSCertFile = certs + "keys5/staker.crt"
	df5.StakingTLSKeyFile = certs + "keys5/staker.key"
	df5.IndexEnabled = indexEnabled
	nodeConfig5, err := createNodeConfig(flagsToArgs(df5))
	if err != nil {
		log.Fatal(err)
	}

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go startApp(wg, nodeConfig2)
	time.Sleep(2 * time.Second)

	wg.Add(1)
	go startApp(wg, nodeConfig3)
	time.Sleep(2 * time.Second)

	wg.Add(1)
	go startApp(wg, nodeConfig4)
	time.Sleep(2 * time.Second)

	wg.Add(1)
	go startApp(wg, nodeConfig5)
	time.Sleep(2 * time.Second)

	wg.Add(1)
	go startApp(wg, nodeConfig1)
	time.Sleep(2 * time.Second)

	wg.Wait()
}

func startApp(wg *sync.WaitGroup, config node.Config) {
	err := runApp(wg, config)
	if err != nil {
		log.Println("err", err)
	}
}

func runApp(wg *sync.WaitGroup, config node.Config) error {
	defer wg.Done()

	app := app.New(config)

	// start running the application
	if err := app.Start(); err != nil {
		return err
	}

	// register signals to kill the application
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT)
	signal.Notify(signals, syscall.SIGTERM)

	// start up a new go routine to handle attempts to kill the application
	var eg errgroup.Group
	eg.Go(func() error {
		for range signals {
			return app.Stop()
		}
		return nil
	})

	// wait for the app to exit and get the exit code response
	exitCode, err := app.ExitCode()

	// shut down the signal go routine
	signal.Stop(signals)
	close(signals)

	// if there was an error closing or running the application, report that error
	if eg.Wait() != nil || err != nil {
		return err
	}

	if exitCode != 0 {
		log.Println("exit code", exitCode)
	}
	return nil
}

func createNodeConfig(args []string) (node.Config, error) {
	fs := config.BuildFlagSet()
	v, err := config.BuildViper(fs, args)
	if err != nil {
		return node.Config{}, err
	}

	return config.GetNodeConfig(v)
}

// Flags represents available CLI flags when starting a node
type Flags struct {

	// Version
	Version bool

	// TX fees
	TxFee uint

	// IP
	PublicIP        string
	DynamicPublicIP string

	// Network ID
	NetworkID string

	// APIs
	APIAdminEnabled    bool
	APIIPCsEnabled     bool
	APIKeystoreEnabled bool
	APIMetricsEnabled  bool
	APIHealthEnabled   bool
	APIInfoEnabled     bool

	// HTTP
	HTTPHost        string
	HTTPPort        uint
	HTTPTLSEnabled  bool
	HTTPTLSCertFile string
	HTTPTLSKeyFile  string

	// Bootstrapping
	BootstrapIPs                     string
	BootstrapIDs                     string
	BootstrapBeaconConnectionTimeout string

	// Build
	BuildDir string

	// Plugins
	PluginDir string

	// Logging
	LogLevel        string
	LogDir          string
	LogDisplayLevel string

	// Consensus
	SnowSampleSize              int
	SnowQuorumSize              int
	SnowVirtuousCommitThreshold int
	SnowRogueCommitThreshold    int
	SnowConcurrentRepolls       int
	MinDelegatorStake           int
	ConsensusShutdownTimeout    string
	ConsensusGossipFrequency    string
	MinDelegationFee            int
	MinValidatorStake           int
	MaxStakeDuration            string
	MaxValidatorStake           int

	// Staking
	StakingEnabled        bool
	StakeMintingPeriod    string
	StakingPort           uint
	StakingDisabledWeight int
	StakingTLSKeyFile     string
	StakingTLSCertFile    string

	// Auth
	APIAuthRequired        bool
	APIAuthPasswordFileKey string
	MinStakeDuration       string

	// Whitelisted Subnets
	WhitelistedSubnets string

	// Config
	ConfigFile string

	// IPCS
	IPCSChainIDs string
	IPCSPath     string

	// File Descriptor Limit
	FDLimit int

	// Benchlist
	BenchlistFailThreshold      int
	BenchlistMinFailingDuration string
	BenchlistDuration           string
	// Network Timeout
	NetworkInitialTimeout                   string
	NetworkMinimumTimeout                   string
	NetworkMaximumTimeout                   string
	NetworkHealthMaxSendFailRateKey         float64
	NetworkHealthMaxPortionSendQueueFillKey float64
	NetworkHealthMaxTimeSinceMsgSentKey     string
	NetworkHealthMaxTimeSinceMsgReceivedKey string
	NetworkHealthMinConnPeers               int
	NetworkTimeoutCoefficient               int
	NetworkTimeoutHalflife                  string

	// Peer List Gossiping
	NetworkPeerListGossipFrequency string

	// Uptime Requirement
	UptimeRequirement float64

	// Retry
	RetryBootstrap bool

	// Health
	HealthCheckAveragerHalflifeKey string
	HealthCheckFreqKey             string

	// Router
	RouterHealthMaxOutstandingRequestsKey int
	RouterHealthMaxDropRateKey            float64

	IndexEnabled bool
}

// defaultFlags returns Avash-specific default node flags
func defaultFlags() Flags {
	return Flags{
		Version:                                 false,
		TxFee:                                   1000000,
		PublicIP:                                "127.0.0.1",
		DynamicPublicIP:                         "",
		NetworkID:                               "local",
		APIAdminEnabled:                         true,
		APIIPCsEnabled:                          true,
		APIKeystoreEnabled:                      true,
		APIMetricsEnabled:                       true,
		HTTPHost:                                "127.0.0.1",
		HTTPPort:                                9650,
		HTTPTLSEnabled:                          false,
		HTTPTLSCertFile:                         "",
		HTTPTLSKeyFile:                          "",
		BootstrapIPs:                            "",
		BootstrapIDs:                            "",
		BootstrapBeaconConnectionTimeout:        "60s",
		BuildDir:                                "",
		PluginDir:                               "",
		LogLevel:                                "info",
		LogDir:                                  baseLogs,
		LogDisplayLevel:                         "", // defaults to the value provided to --log-level
		SnowSampleSize:                          2,
		SnowQuorumSize:                          2,
		SnowVirtuousCommitThreshold:             5,
		SnowRogueCommitThreshold:                10,
		SnowConcurrentRepolls:                   4,
		MinDelegatorStake:                       5000000,
		ConsensusShutdownTimeout:                "5s",
		ConsensusGossipFrequency:                "10s",
		MinDelegationFee:                        20000,
		MinValidatorStake:                       5000000,
		MaxStakeDuration:                        "8760h",
		MaxValidatorStake:                       3000000000000000,
		StakeMintingPeriod:                      "8760h",
		NetworkInitialTimeout:                   "5s",
		NetworkMinimumTimeout:                   "5s",
		NetworkMaximumTimeout:                   "10s",
		NetworkHealthMaxSendFailRateKey:         0.9,
		NetworkHealthMaxPortionSendQueueFillKey: 0.9,
		NetworkHealthMaxTimeSinceMsgSentKey:     "1m",
		NetworkHealthMaxTimeSinceMsgReceivedKey: "1m",
		NetworkHealthMinConnPeers:               1,
		NetworkTimeoutCoefficient:               2,
		NetworkTimeoutHalflife:                  "5m",
		NetworkPeerListGossipFrequency:          "1m",
		StakingEnabled:                          false,
		StakingPort:                             9651,
		StakingDisabledWeight:                   1,
		StakingTLSKeyFile:                       "",
		StakingTLSCertFile:                      "",
		APIAuthRequired:                         false,
		APIAuthPasswordFileKey:                  "",
		MinStakeDuration:                        "336h",
		APIHealthEnabled:                        true,
		ConfigFile:                              "",
		WhitelistedSubnets:                      "",
		APIInfoEnabled:                          true,
		IPCSChainIDs:                            "",
		IPCSPath:                                "/tmp",
		FDLimit:                                 32768,
		BenchlistDuration:                       "1h",
		BenchlistFailThreshold:                  10,
		BenchlistMinFailingDuration:             "5m",
		UptimeRequirement:                       0.6,
		RetryBootstrap:                          true,
		HealthCheckAveragerHalflifeKey:          "10s",
		HealthCheckFreqKey:                      "30s",
		RouterHealthMaxOutstandingRequestsKey:   1024,
		RouterHealthMaxDropRateKey:              1,
		IndexEnabled:                            false,
	}
}

// flagsToArgs converts a `Flags` struct into a CLI command flag string
func flagsToArgs(flags Flags) []string {
	// Port targets
	httpPortString := strconv.FormatUint(uint64(flags.HTTPPort), 10)
	stakingPortString := strconv.FormatUint(uint64(flags.StakingPort), 10)

	// Paths/directories
	dbPath := baseDB + "/" + stakingPortString
	logPath := baseLogs + "/" + stakingPortString

	wd, _ := os.Getwd()
	// If the path given in the flag doesn't begin with "/", treat it as relative
	// to the directory of the avash binary
	httpCertFile := flags.HTTPTLSCertFile
	if httpCertFile != "" && string(httpCertFile[0]) != "/" {
		httpCertFile = fmt.Sprintf("%s/%s", wd, httpCertFile)
	}

	httpKeyFile := flags.HTTPTLSKeyFile
	if httpKeyFile != "" && string(httpKeyFile[0]) != "/" {
		httpKeyFile = fmt.Sprintf("%s/%s", wd, httpKeyFile)
	}

	stakerCertFile := flags.StakingTLSCertFile
	if stakerCertFile != "" && string(stakerCertFile[0]) != "/" {
		stakerCertFile = fmt.Sprintf("%s/%s", wd, stakerCertFile)
	}

	stakerKeyFile := flags.StakingTLSKeyFile
	if stakerKeyFile != "" && string(stakerKeyFile[0]) != "/" {
		stakerKeyFile = fmt.Sprintf("%s/%s", wd, stakerKeyFile)
	}

	args := []string{
		"--version=" + strconv.FormatBool(flags.Version),
		"--tx-fee=" + strconv.FormatUint(uint64(flags.TxFee), 10),
		"--public-ip=" + flags.PublicIP,
		"--dynamic-public-ip=" + flags.DynamicPublicIP,
		"--network-id=" + flags.NetworkID,
		"--api-admin-enabled=" + strconv.FormatBool(flags.APIAdminEnabled),
		"--api-ipcs-enabled=" + strconv.FormatBool(flags.APIIPCsEnabled),
		"--api-keystore-enabled=" + strconv.FormatBool(flags.APIKeystoreEnabled),
		"--api-metrics-enabled=" + strconv.FormatBool(flags.APIMetricsEnabled),
		"--http-host=" + flags.HTTPHost,
		"--http-port=" + httpPortString,
		"--http-tls-enabled=" + strconv.FormatBool(flags.HTTPTLSEnabled),
		"--http-tls-cert-file=" + httpCertFile,
		"--http-tls-key-file=" + httpKeyFile,
		"--bootstrap-ips=" + flags.BootstrapIPs,
		"--bootstrap-ids=" + flags.BootstrapIDs,
		"--bootstrap-beacon-connection-timeout=" + flags.BootstrapBeaconConnectionTimeout,
		"--db-dir=" + dbPath,
		// "--db-type=memdb",
		"--plugin-dir=" + flags.PluginDir,
		"--build-dir=" + flags.BuildDir,
		"--log-level=" + flags.LogLevel,
		"--log-dir=" + logPath,
		"--log-display-level=" + flags.LogDisplayLevel,
		"--snow-sample-size=" + strconv.Itoa(flags.SnowSampleSize),
		"--snow-quorum-size=" + strconv.Itoa(flags.SnowQuorumSize),
		"--snow-virtuous-commit-threshold=" + strconv.Itoa(flags.SnowVirtuousCommitThreshold),
		"--snow-rogue-commit-threshold=" + strconv.Itoa(flags.SnowRogueCommitThreshold),
		"--min-delegator-stake=" + strconv.Itoa(flags.MinDelegatorStake),
		"--consensus-shutdown-timeout=" + flags.ConsensusShutdownTimeout,
		"--consensus-accepted-frontier-gossip-frequency=" + flags.ConsensusGossipFrequency,
		"--min-delegation-fee=" + strconv.Itoa(flags.MinDelegationFee),
		"--min-validator-stake=" + strconv.Itoa(flags.MinValidatorStake),
		"--max-stake-duration=" + flags.MaxStakeDuration,
		"--max-validator-stake=" + strconv.Itoa(flags.MaxValidatorStake),
		"--snow-concurrent-repolls=" + strconv.Itoa(flags.SnowConcurrentRepolls),
		"--stake-minting-period=" + flags.StakeMintingPeriod,
		"--network-initial-timeout=" + flags.NetworkInitialTimeout,
		"--network-minimum-timeout=" + flags.NetworkMinimumTimeout,
		"--network-maximum-timeout=" + flags.NetworkMaximumTimeout,
		fmt.Sprintf("--network-health-max-send-fail-rate=%f", flags.NetworkHealthMaxSendFailRateKey),
		fmt.Sprintf("--network-health-max-portion-send-queue-full=%f", flags.NetworkHealthMaxPortionSendQueueFillKey),
		"--network-health-max-time-since-msg-sent=" + flags.NetworkHealthMaxTimeSinceMsgSentKey,
		"--network-health-max-time-since-msg-received=" + flags.NetworkHealthMaxTimeSinceMsgReceivedKey,
		"--network-health-min-conn-peers=" + strconv.Itoa(flags.NetworkHealthMinConnPeers),
		"--network-timeout-coefficient=" + strconv.Itoa(flags.NetworkTimeoutCoefficient),
		"--network-timeout-halflife=" + flags.NetworkTimeoutHalflife,
		"--network-peer-list-gossip-frequency=" + flags.NetworkPeerListGossipFrequency,
		"--sybil-protection-enabled=" + strconv.FormatBool(flags.StakingEnabled),
		"--staking-port=" + stakingPortString,
		"--sybil-protection-disabled-weight=" + strconv.Itoa(flags.StakingDisabledWeight),
		"--staking-tls-key-file=" + stakerKeyFile,
		"--staking-tls-cert-file=" + stakerCertFile,
		"--api-auth-required=" + strconv.FormatBool(flags.APIAuthRequired),
		"--api-auth-password-file=" + flags.APIAuthPasswordFileKey,
		"--min-stake-duration=" + flags.MinStakeDuration,
		"--whitelisted-subnets=" + flags.WhitelistedSubnets,
		"--api-health-enabled=" + strconv.FormatBool(flags.APIHealthEnabled),
		"--config-file=" + flags.ConfigFile,
		"--api-info-enabled=" + strconv.FormatBool(flags.APIInfoEnabled),
		"--ipcs-chain-ids=" + flags.IPCSChainIDs,
		"--ipcs-path=" + flags.IPCSPath,
		"--fd-limit=" + strconv.Itoa(flags.FDLimit),
		"--benchlist-duration=" + flags.BenchlistDuration,
		"--benchlist-fail-threshold=" + strconv.Itoa(flags.BenchlistFailThreshold),
		"--benchlist-min-failing-duration=" + flags.BenchlistMinFailingDuration,
		fmt.Sprintf("--uptime-requirement=%f", flags.UptimeRequirement),
		"--bootstrap-retry-enabled=" + strconv.FormatBool(flags.RetryBootstrap),
		"--health-check-averager-halflife=" + flags.HealthCheckAveragerHalflifeKey,
		"--health-check-frequency=" + flags.HealthCheckFreqKey,
		"--router-health-max-outstanding-requests=" + strconv.Itoa(flags.RouterHealthMaxOutstandingRequestsKey),
		fmt.Sprintf("--router-health-max-drop-rate=%f", flags.RouterHealthMaxDropRateKey),
		"--index-enabled=" + strconv.FormatBool(flags.IndexEnabled),
		"--chain-config-dir=" + "./cconfig",
	}
	args = removeEmptyFlags(args)

	return args
}

func removeEmptyFlags(args []string) []string {
	var res []string
	for _, f := range args {
		tmp := strings.TrimSpace(f)
		if !strings.HasSuffix(tmp, "=") {
			res = append(res, tmp)
		}
	}
	return res
}
