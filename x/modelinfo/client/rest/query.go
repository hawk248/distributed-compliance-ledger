package rest

import (
	"encoding/json"
	"fmt"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/modelinfo/internal/keeper"
	"net/http"
	"strconv"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/modelinfo/internal/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
)

func getModelsHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := parsePaginationParams(cliCtx, r)

		res, height, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/all_models", storeName), params)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		respond(w, cliCtx, res, height)
	}
}

func getModelHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		vid := vars[vid]
		pid := vars[pid]

		res, height, err := cliCtx.QueryStore([]byte(keeper.ModelInfoId(vid, pid)), storeName)
		if err != nil || res == nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, "Model Not Found")
			return
		}

		var modelInfo types.ModelInfo
		cliCtx.Codec.MustUnmarshalBinaryBare(res, &modelInfo)

		out, err := json.Marshal(modelInfo)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		}

		respond(w, cliCtx, out, height)
	}
}

func getVendorsHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := parsePaginationParams(cliCtx, r)

		res, height, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/vendors", storeName), params)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		respond(w, cliCtx, res, height)
	}
}

func getVendorModelsHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		vid := vars[vid]

		res, height, err := cliCtx.QueryStore([]byte(keeper.VendorProductsId(vid)), storeName)
		if err != nil || res == nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, "VendorProducts Not Found")
			return
		}

		var vendorProducts types.VendorProducts
		cliCtx.Codec.MustUnmarshalBinaryBare(res, &vendorProducts)

		out, err := json.Marshal(vendorProducts)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		}

		respond(w, cliCtx, out, height)
	}
}

func parsePaginationParams(cliCtx context.CLIContext, r *http.Request) []byte {
	skip, _ := strconv.Atoi(r.FormValue("skip"))
	take, _ := strconv.Atoi(r.FormValue("take"))
	params := types.NewPaginationParams(skip, take)
	return cliCtx.Codec.MustMarshalJSON(params)
}

func respond(w http.ResponseWriter, cliCtx context.CLIContext, data []byte, height int64) {
	cliCtx.Height = height
	rest.PostProcessResponse(w, cliCtx, data)

}