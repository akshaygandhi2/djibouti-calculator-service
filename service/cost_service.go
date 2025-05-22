package service

import (
	"buildingcost/model"
	"encoding/json"
	"errors"
)

func ProcessApplications(request map[string]interface{}) error {
	apps, ok := request["Application"].([]interface{})
	if !ok {
		return errors.New("invalid Application array")
	}

	for _, raw := range apps {
		app, ok := raw.(map[string]interface{})
		if !ok {
			continue
		}

		addDetails, ok := app["additionalDetails"].(map[string]interface{})
		if !ok {
			continue
		}

		// Extract and marshal costEstimation to our model
		costData, ok := addDetails["costEstimation"].(map[string]interface{})
		if !ok {
			continue
		}
		costJson, _ := json.Marshal(costData)

		var ce model.CostEstimation
		if err := json.Unmarshal(costJson, &ce); err != nil {
			continue
		}

		var total float64
		for i := range ce.Floors {
			f := &ce.Floors[i]
			f.TotalAreaPerLevel = (f.BuiltUpAreaLiving + f.BuiltupAreaCommercial)
			f.FloorCost = (f.BuiltUpAreaLiving * ce.CostPerSqmLivingSpace) + (f.BuiltupAreaCommercial * ce.CostPerSqmCommercialSpace)
			total += f.FloorCost
		}

		ce.TotalBuildingCost = total
		ce.RoyaltyFee = total * ce.RoyaltyPer / 100
		ce.EqResistanceCost = total * ce.EqResistancePer / 100

		ce.TotalTax = ce.RoyaltyFee + ce.EqResistanceCost
		ce.TotalTaxWithServiceCharge = ce.TotalTax + ce.RegistryServiceFee

		// Put back the updated model
		updatedJson, _ := json.Marshal(ce)
		var updatedMap map[string]interface{}
		json.Unmarshal(updatedJson, &updatedMap)

		addDetails["costEstimation"] = updatedMap
		app["additionalDetails"] = addDetails
	}
	return nil
}
