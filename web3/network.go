package web3

import (
	"log"

	"org.donghyusn.com/chain/collector/database"
)

func CreateNewNetwork(networkName string, rpcUrl string) error {
	dbCon, dbErr := database.GetConnection()

	if dbErr != nil {
		return dbErr
	}

	_, insertErr := dbCon.InsertQuery(InsertNetworkQuery, networkName, rpcUrl)

	if insertErr != nil {
		return insertErr
	}

	return nil
}

// Get Network List
func GetNetworkList() ([]NetworkListResult, error) {
	dbCon, dbErr := database.GetConnection()

	if dbErr != nil {
		return []NetworkListResult{}, dbErr
	}

	result, queryErr := dbCon.SelectMultipleRows(SelectNetworkList)

	if queryErr != nil {
		return []NetworkListResult{}, queryErr
	}

	var networkList []NetworkListResult

	for result.Next() {
		var network NetworkListResult

		if scanErr := result.Scan(&network.NetworkSeq, &network.NetworkName, &network.NetworkUrl); scanErr != nil {
			log.Printf("[NETWORK] Scan Network List Error: %v", scanErr)
			return []NetworkListResult{}, scanErr
		}

		networkList = append(networkList, network)
	}

	return networkList, nil
}

// Get Network By Network Name
func GetRpcUrlByNetworkName(networkName string) (NetworkListResult, error) {
	dbCon, dbErr := database.GetConnection()

	if dbErr != nil {
		return NetworkListResult{}, dbErr
	}

	result, queryErr := dbCon.SelectOneRow(SelectNetworkByNetworkNameQuery, networkName)

	if queryErr != nil {
		return NetworkListResult{}, queryErr
	}

	var network NetworkListResult

	if scanErr := result.Scan(&network.NetworkSeq, &network.NetworkName, &network.NetworkUrl); scanErr != nil {
		log.Printf("[NETWORK] Scan Network By Name Error: %v", scanErr)
		return NetworkListResult{}, scanErr
	}

	return network, nil
}
