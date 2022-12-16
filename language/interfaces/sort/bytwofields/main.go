package main

import (
	"fmt"
	"sort"
)

func RemainderSorting(strArr []string) []string {
	res := make([]string, len(strArr))
	copy(res, strArr)

	sf := func(i, j int) bool {
		r1 := len(res[i]) % 3
		r2 := len(res[j]) % 3

		if r1 < r2 {
			return true
		}

		if r1 == r2 && res[i] < res[j] {
			return true
		}

		return false
	}

	sort.SliceStable(res, sf)

	return res
}

func main() {
	src := []string{
		"ABrezayfnXdwvqguWadkdHhqxZPDAsuJiNCuXAipwnxogviZurTrNGnBv",
		"ApItmTroSNPkcbNktqJjQMzDuDHZcEmoevMsDoWvYMKcPsTwYsTuTYWmZaIj",
		"AExnZWzqWTOtQTus",
		"BodDUORuWLMLnogCmYCBsRGagaDBTcnIlAOZczIONrMFkXauofdFmWccsPQaePocvpkFVSSZTIyehViOkwrVzfgsKhdUhAnhZTP",
		"AqBKVgejHBWbaUlSzdHTUUqBizczJYKHEQLklEGdgCMmQoeHSHqQPTMHMNYeVAHMLdYbLzCsRSIutQoeqktdRc",
		"BKiHiShozewtnQCtenLgFmufkfgLYigFUhQEHUgjgmjyCyvwLduULN",
		"BNDxtABWGagZiGmPuCKCsufzAvLkuLfnBladKXSfxliTlMOTHnNGrKZEv",
		"BafkYJpqwFgogrLkDWIlLXssGpJDGPMkvFSdLsIwZbGmFTAroKVNIApInmXuFsMXlqWtPQpnUYvojFdPoKNGkrrQYo",
		"BpvyqewelGJMwiadWfxokIFTijcMnDzyrQoENOKMKJuGJGBMwsTGAOunxhUKQKrAqs",
		"CEMAMVlzAZKE",
		"CwNRcpcAXALTkEDUiYEzutDVXHFvfNoJebLvRcOKDtIIAmNDLSMFGBbXtGtKojW",
		"EMgheZ",
		"DvPjvDDFDE",
		"EjFwmACPWlmJHgxuRnGefUMklbPXnyYSIkBzSll",
		"AEgPJYdtVYFiVsAxtKKbzriPnOxhepwtluwPGtjWxpveQcpXdhYotBogeC",
		"DBrcvUwzzIMmMeqDYrxcpJcrMRF",
		"FCLIdpQBxssSRhqFgSBOH",
		"CUKpUwAEMZaGSsdFllM",
		"DRQEdpXBSYdfvkiGMazTERMJxPqHJnaeZHztE",
		"EoprllSRJunhBILpyz",
		"CcoaCePfrcIlhclaIOeAgwMwbObzwNsrVyCpshsfEFvViZpKSActrkTprY",
		"CtjlTHzIjDITdhYeuWXUZiWgFXIvqfqlWWxahDEctmCTJBJAQkl",
	}

	src = []string{
		"a",
		"ab",
		"bc",
		"abc",
	}

	fmt.Println("init: ")
	for _, s := range src {
		fmt.Printf("%s \n", s)
	}

	res := RemainderSorting(src)

	fmt.Println("sorted: len%3 # string ")
	for _, s := range res {
		fmt.Printf("%d %s \n", len(s)%3, s)
	}
}
