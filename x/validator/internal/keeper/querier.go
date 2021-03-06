// Copyright 2020 DSR Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/internal/types"
)

// query endpoints supported by the validator Querier.
const (
	QueryValidators = "validators"
	QueryValidator  = "validator"
)

// creates a querier for validator module.
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryValidators:
			return queryValidators(ctx, req, k)
		case QueryValidator:
			return queryValidator(ctx, path[1:], k)
		default:
			return nil, sdk.ErrUnknownRequest("unknown pki query endpoint")
		}
	}
}

func queryValidators(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {
	var params types.ListValidatorsParams
	if err := keeper.cdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, sdk.ErrUnknownRequest(fmt.Sprintf("Failed to parse request params: %s", err))
	}

	result := types.NewListValidatorItems()

	skipped := 0

	keeper.IterateValidators(ctx, func(validator types.Validator) (stop bool) {
		// filter by validator state
		// nolint:exhaustive
		switch params.State {
		case types.Active:
			if validator.IsJailed() {
				return false
			}
		case types.Jailed:
			if !validator.IsJailed() {
				return false
			}
		}

		result.Total++

		if skipped < params.Skip {
			skipped++

			return false
		}

		if len(result.Items) < params.Take || params.Take == 0 {
			result.Items = append(result.Items, validator)

			return false
		}

		return false
	})

	res = codec.MustMarshalJSONIndent(keeper.cdc, result)

	return res, nil
}

func queryValidator(ctx sdk.Context, path []string, k Keeper) ([]byte, sdk.Error) {
	validatorAddr, err := sdk.ConsAddressFromBech32(path[0])
	if err != nil {
		return nil, sdk.ErrUnknownRequest(err.Error())
	}

	if !k.IsValidatorPresent(ctx, validatorAddr) {
		return nil, types.ErrValidatorDoesNotExist(validatorAddr)
	}

	validator := k.GetValidator(ctx, validatorAddr)

	res := codec.MustMarshalJSONIndent(types.ModuleCdc, validator)

	return res, nil
}
