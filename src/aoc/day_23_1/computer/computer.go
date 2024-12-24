package computer

import (
	"aoc/src/aoc/readfile"
	"log"
	"slices"
	"sort"
	"strings"
)

type Computer = string

type NetworkConnections = map[Computer][]Computer

func addNetworkConnection(networkConnections NetworkConnections, computer_one, computer_two Computer) {
	connections := networkConnections[computer_one]
	if connections == nil {
		networkConnections[computer_one] = []Computer{computer_two}
	} else {
		networkConnections[computer_one] = append(networkConnections[computer_one], computer_two)
	}
}

func LoadNetworkConnections(path string) NetworkConnections {
	file := readfile.Open(path)
	defer file.Close()
	var networkConnections NetworkConnections = make(NetworkConnections)
	err := readfile.ReadLine(file, func(line string) error {
		computers := strings.Split(line, "-")
		if len(computers) != 2 {
			log.Fatal("invalid input")
		}

		addNetworkConnection(networkConnections, computers[0], computers[1])
		addNetworkConnection(networkConnections, computers[1], computers[0])

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return networkConnections
}

func getTrioString(one, two, three Computer) string {
	network := []Computer{one, two, three}
	sort.Strings(network)
	return strings.Join(network, ",")
}

func FindNetworks(identifier string, network NetworkConnections) map[string][]Computer {
	var networks_of_length map[string][]Computer = make(map[string][]Computer)
	for one, connections := range network {
		for _, two := range connections {
			sibling := network[two]
			for _, three := range sibling {
				thirdConnections := network[three]
				if slices.Contains(thirdConnections, one) {
					// this is a trio of compupters
					trio := getTrioString(one, two, three)
					if strings.HasPrefix(one, identifier) || strings.HasPrefix(two, identifier) || strings.HasPrefix(three, identifier) {

						networks_of_length[trio] = []Computer{one, two, three}
					}
				}
			}
		}
	}

	return networks_of_length
}
