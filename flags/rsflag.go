package rsflag

import (
	"flag"
	"fmt"
)

// FlagTest test Flag func
func FlagTest() {
	config := flag.String("c", "./config/wallet-sign.yaml", "full path config file")
	_ = flag.Bool("gen", false, "gen currency address")
	_ = flag.String("hash", "", "hash for given address")
	_ = flag.Bool("pgen", false, "gen protected currency address")
	_ = flag.Bool("update", false, "update address hash of given chain in config")
	_ = flag.String("genPrivKey", "", "get the private key of the given address")
	flag.Parse()

	cmdParam := make(map[string]string, 0)
	// Visit 和 VisitAll 的区别是否是都未设置的值进行遍历
	flag.VisitAll(func(i *flag.Flag) {
		cmdParam[i.Name] = i.Value.String()
		fmt.Printf("%s: %s \n", i.Name, i.Value.String())
	})

	fmt.Println(cmdParam)
	fmt.Println(*config)
}
