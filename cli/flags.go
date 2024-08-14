package main

import cli "github.com/urfave/cli/v2"

// Required for commands that need an EigenPod's address
var POD_ADDRESS_FLAG = &cli.StringFlag{
	Name:        "podAddress",
	Aliases:     []string{"p", "pod"},
	Value:       "",
	Usage:       "[required] The onchain `address` of your eigenpod contract (0x123123123123)",
	Required:    true,
	Destination: &eigenpodAddress,
}

// Required for commands that need a beacon chain RPC
var BEACON_NODE_FLAG = &cli.StringFlag{
	Name:        "beaconNode",
	Aliases:     []string{"b"},
	Value:       "",
	Usage:       "[required] `URL` to a functioning beacon node RPC (https://)",
	Required:    true,
	Destination: &beacon,
}

// Required for commands that need an execution layer RPC
var EXEC_NODE_FLAG = &cli.StringFlag{
	Name:        "execNode",
	Aliases:     []string{"e"},
	Value:       "",
	Usage:       "[required] `URL` to a functioning execution-layer RPC (https://)",
	Required:    true,
	Destination: &node,
}

// Optional commands:

// Optional use for commands that want direct tx submission from a specific private key
var SENDER_PK_FLAG = &cli.StringFlag{
	Name:        "sender",
	Aliases:     []string{"s"},
	Value:       "",
	Usage:       "`Private key` of the account that will send any transactions. If set, this will automatically submit the proofs to their corresponding onchain functions after generation. If using checkpoint mode, it will also begin a checkpoint if one hasn't been started already.",
	Destination: &sender,
}

// Optional use for commands that support JSON output
var PRINT_JSON_FLAG = &cli.BoolFlag{
	Name:        "json",
	Value:       false,
	Usage:       "print only plain JSON",
	Required:    false,
	Destination: &useJson,
}

// shared flag --batch
func BatchBySize(destination *uint64, defaultValue uint64) *cli.Uint64Flag {
	return &cli.Uint64Flag{
		Name:        "batch",
		Value:       defaultValue,
		Usage:       "Submit proofs in groups of size `batchSize`, to avoid gas limit.",
		Required:    false,
		Destination: destination,
	}
}

// Hack to make a copy of a flag that sets `Required` to true
func Require(flag *cli.StringFlag) *cli.StringFlag {
	return &cli.StringFlag{
		Name:        flag.Name,
		Aliases:     flag.Aliases,
		Value:       flag.Value,
		Usage:       flag.Usage,
		Destination: flag.Destination,
		Required:    true,
	}
}