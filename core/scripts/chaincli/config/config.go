package config

import (
	"log"

	"github.com/spf13/viper"
)

// Config represents configuration fields
type Config struct {
	NodeURL        string   `mapstructure:"NODE_URL"`
	ChainID        int64    `mapstructure:"CHAIN_ID"`
	PrivateKey     string   `mapstructure:"PRIVATE_KEY"`
	LinkTokenAddr  string   `mapstructure:"LINK_TOKEN_ADDR"`
	Keepers        []string `mapstructure:"KEEPERS"`
	ApproveAmount  string   `mapstructure:"APPROVE_AMOUNT"`
	AddFundsAmount string   `mapstructure:"ADD_FUNDS_AMOUNT"`
	GasLimit       uint64   `mapstructure:"GAS_LIMIT"`

	// Keeper config
	LinkETHFeedAddr      string `mapstructure:"LINK_ETH_FEED"`
	FastGasFeedAddr      string `mapstructure:"FAST_GAS_FEED"`
	PaymentPremiumPBB    uint32 `mapstructure:"PAYMENT_PREMIUM_PBB"`
	FlatFeeMicroLink     uint32 `mapstructure:"FLAT_FEE_MICRO_LINK"`
	BlockCountPerTurn    int64  `mapstructure:"BLOCK_COUNT_PER_TURN"`
	CheckGasLimit        uint32 `mapstructure:"CHECK_GAS_LIMIT"`
	StalenessSeconds     int64  `mapstructure:"STALENESS_SECONDS"`
	GasCeilingMultiplier uint16 `mapstructure:"GAS_CEILING_MULTIPLIER"`
	FallbackGasPrice     int64  `mapstructure:"FALLBACK_GAS_PRICE"`
	FallbackLinkPrice    int64  `mapstructure:"FALLBACK_LINK_PRICE"`

	// Upkeep Config
	UpkeepTestRange                 int64  `mapstructure:"UPKEEP_TEST_RANGE"`
	UpkeepAverageEligibilityCadence int64  `mapstructure:"UPKEEP_AVERAGE_ELIGIBILITY_CADENCE"`
	UpkeepCheckData                 string `mapstructure:"UPKEEP_CHECK_DATA"`
	UpkeepGasLimit                  uint32 `mapstructure:"UPKEEP_GAS_LIMIT"`
	UpkeepCount                     int64  `mapstructure:"UPKEEP_COUNT"`
}

// New is the constructor of Config
func New() *Config {
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatal("failed to unmarshal config: ", err)
	}

	return &cfg
}

func init() {
	// Represented in WEI, which is 1000 Ether
	viper.SetDefault("APPROVE_AMOUNT", "1000000000000000000000")
	// Represented in WEI, which is 100 Ether
	viper.SetDefault("ADD_FUNDS_AMOUNT", "100000000000000000000")
	viper.SetDefault("GAS_LIMIT", 8000000)
	viper.SetDefault("PAYMENT_PREMIUM_PBB", 200000000)
	viper.SetDefault("FLAT_FEE_MICRO_LINK", 0)
	viper.SetDefault("BLOCK_COUNT_PER_TURN", 1)
	viper.SetDefault("CHECK_GAS_LIMIT", 650000000)
	viper.SetDefault("STALENESS_SECONDS", 90000)
	viper.SetDefault("GAS_CEILING_MULTIPLIER", 3)
	viper.SetDefault("FALLBACK_GAS_PRICE", 10000000000)
	viper.SetDefault("FALLBACK_LINK_PRICE", 200000000000000000)
	viper.SetDefault("UPKEEP_TEST_RANGE", 1)
	viper.SetDefault("UPKEEP_AVERAGE_ELIGIBILITY_CADENCE", 1)
	viper.SetDefault("UPKEEP_CHECK_DATA", "0x00")
	viper.SetDefault("UPKEEP_GAS_LIMIT", 500000)
	viper.SetDefault("UPKEEP_COUNT", 5)

	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("failed to read config: ", err)
	}
}
