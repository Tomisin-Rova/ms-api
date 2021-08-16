// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package connections

import (
	"fmt"

	"ms.api/graph/models"
	"ms.api/types"
)

// TypesPaymentEdgeMaker define a function that take a *types.Payment and an offset and
// create an Edge.
type PaymentLookupEdgeMaker func(value *types.Payment, offset int) Edge

// PaymentLookupConMaker define a function that create a types.PaymentConnection
type PaymentLookupConMaker func(
	edges []*types.PaymentEdge,
	nodes []*types.Payment,
	info *types.PageInfo,
	totalCount int) (*types.PaymentConnection, error)

// PaymentLookupCon will paginate a source according to the input of a relay connection
func PaymentLookupCon(source []*types.Payment, edgeMaker PaymentLookupEdgeMaker, conMaker PaymentLookupConMaker, input models.ConnectionInput) (*types.PaymentConnection, error) {
	var nodes []*types.Payment
	var edges []*types.PaymentEdge
	var cursors []string
	var pageInfo = &types.PageInfo{}
	var totalCount = len(source)

	emptyCon, _ := conMaker(edges, nodes, pageInfo, 0)

	offset := 0

	if input.After != nil {
		for i, value := range source {
			edge := edgeMaker(value, i)
			if edge.GetCursor() == *input.After {
				// remove all previous element including the "after" one
				source = source[i+1:]
				offset = i + 1
				pageInfo.HasPreviousPage = true
				break
			}
		}
	}

	if input.Before != nil {
		for i, value := range source {
			edge := edgeMaker(value, i+offset)

			if edge.GetCursor() == *input.Before {
				// remove all after element including the "before" one
				pageInfo.HasNextPage = true
				break
			}

			e := edge.(types.PaymentEdge)
			edges = append(edges, &e)
			cursors = append(cursors, edge.GetCursor())
			nodes = append(nodes, value)
		}
	} else {
		edges = make([]*types.PaymentEdge, len(source))
		cursors = make([]string, len(source))
		nodes = source

		for i, value := range source {
			edge := edgeMaker(value, i+offset)
			e := edge.(types.PaymentEdge)
			edges[i] = &e
			cursors[i] = edge.GetCursor()
		}
	}

	if input.First != nil {
		if *input.First < 0 {
			return emptyCon, fmt.Errorf("first less than zero")
		}

		if int64(len(edges)) > *input.First {
			// Slice result to be of length first by removing edges from the end
			edges = edges[:*input.First]
			cursors = cursors[:*input.First]
			nodes = nodes[:*input.First]
			pageInfo.HasNextPage = true
		}
	}

	if input.Last != nil {
		if *input.Last < 0 {
			return emptyCon, fmt.Errorf("last less than zero")
		}

		if int64(len(edges)) > *input.Last {
			// Slice result to be of length last by removing edges from the start
			edges = edges[int64(len(edges))-*input.Last:]
			cursors = cursors[int64(len(cursors))-*input.Last:]
			nodes = nodes[int64(len(nodes))-*input.Last:]
			pageInfo.HasPreviousPage = true
		}
	}

	// Fill up pageInfo cursors
	if len(cursors) > 0 {
		pageInfo.StartCursor = &cursors[0]
		pageInfo.EndCursor = &cursors[len(cursors)-1]
	}

	return conMaker(edges, nodes, pageInfo, totalCount)
}