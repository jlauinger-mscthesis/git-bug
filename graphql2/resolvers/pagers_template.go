package resolvers

import (
	"fmt"
	"github.com/cheekybits/genny/generic"
)

type NodeType generic.Type
type EdgeType generic.Type

type NodeTypeEdger func(value NodeType, offset int) Edge

func NodeTypePaginate(source []NodeType, edger NodeTypeEdger, input ConnectionInput) ([]EdgeType, PageInfo, error) {
	var result []EdgeType
	var pageInfo PageInfo

	offset := 0

	if input.After != nil {
		for i, value := range source {
			edge := edger(value, i)
			if edge.GetCursor() == *input.After {
				// remove all previous element including the "after" one
				source = source[i+1:]
				offset = i + 1
				break
			}
		}
	}

	if input.Before != nil {
		for i, value := range source {
			edge := edger(value, i+offset)

			if edge.GetCursor() == *input.Before {
				// remove all after element including the "before" one
				break
			}

			result = append(result, edge.(EdgeType))
		}
	} else {
		result = make([]EdgeType, len(source))

		for i, value := range source {
			result[i] = edger(value, i+offset).(EdgeType)
		}
	}

	if input.First != nil {
		if *input.First < 0 {
			return nil, PageInfo{}, fmt.Errorf("first less than zero")
		}

		if len(result) > *input.First {
			// Slice result to be of length first by removing edges from the end
			result = result[:*input.First]
			pageInfo.HasNextPage = true
		}
	}

	if input.Last != nil {
		if *input.Last < 0 {
			return nil, PageInfo{}, fmt.Errorf("last less than zero")
		}

		if len(result) > *input.Last {
			// Slice result to be of length last by removing edges from the start
			result = result[len(result)-*input.Last:]
			pageInfo.HasPreviousPage = true
		}
	}

	return result, pageInfo, nil
}

// Apply the before/after cursor params to the source and return an array of edges
//func ApplyCursorToEdges(source []interface{}, edger Edger, input ConnectionInput) []Edge {
//	var result []Edge
//
//	if input.After != nil {
//		for i, value := range source {
//			edge := edger(value)
//			if edge.Cursor() == *input.After {
//				// remove all previous element including the "after" one
//				source = source[i+1:]
//				break
//			}
//		}
//	}
//
//	if input.Before != nil {
//		for _, value := range source {
//			edge := edger(value)
//
//			if edge.Cursor() == *input.Before {
//				// remove all after element including the "before" one
//				break
//			}
//
//			result = append(result, edge)
//		}
//	} else {
//		result = make([]Edge, len(source))
//
//		for i, value := range source {
//			result[i] = edger(value)
//		}
//	}
//
//	return result
//}

// Apply the first/last cursor params to the edges
//func EdgesToReturn(edges []Edge, input ConnectionInput) ([]Edge, PageInfo, error) {
//	hasPreviousPage := false
//	hasNextPage := false
//
//	if input.First != nil {
//		if *input.First < 0 {
//			return nil, nil, fmt.Errorf("first less than zero")
//		}
//
//		if len(edges) > *input.First {
//			// Slice result to be of length first by removing edges from the end
//			edges = edges[:*input.First]
//			hasNextPage = true
//		}
//	}
//
//	if input.Last != nil {
//		if *input.Last < 0 {
//			return nil, nil, fmt.Errorf("last less than zero")
//		}
//
//		if len(edges) > *input.Last {
//			// Slice result to be of length last by removing edges from the start
//			edges = edges[len(edges)-*input.Last:]
//			hasPreviousPage = true
//		}
//	}
//
//	pageInfo := PageInfo{
//		HasNextPage:     hasNextPage,
//		HasPreviousPage: hasPreviousPage,
//	}
//
//	return edges, pageInfo, nil
//}

//func EdgesToReturn(allEdges []Edge, before *cursor, after *cursor, first *int, last *int) ([]Edge, error) {
//	result := ApplyCursorToEdges(allEdges, before, after)
//
//	if first != nil {
//		if *first < 0 {
//			return nil, fmt.Errorf("first less than zero")
//		}
//
//		if len(result) > *first {
//			// Slice result to be of length first by removing edges from the end
//			result = result[:*first]
//		}
//	}
//
//	if last != nil {
//		if *last < 0 {
//			return nil, fmt.Errorf("last less than zero")
//		}
//
//		if len(result) > *last {
//			// Slice result to be of length last by removing edges from the start
//			result = result[len(result)-*last:]
//		}
//	}
//
//	return result, nil
//}

//func ApplyCursorToEdges(allEdges []Edge, before *cursor, after *cursor) []Edge {
//	result := allEdges
//
//	if after != nil {
//		for i, edge := range result {
//			if edge.Cursor() == *after {
//				// remove all previous element including the "after" one
//				result = result[i+1:]
//				break
//			}
//		}
//	}
//
//	if before != nil {
//		for i, edge := range result {
//			if edge.Cursor() == *before {
//				// remove all after element including the "before" one
//				result = result[:i]
//			}
//		}
//	}
//
//	return result
//}

//func HasPreviousPage(allEdges []Edge, before *cursor, after *cursor, last *int) bool {
//	if last != nil {
//		edges := ApplyCursorToEdges(allEdges, before, after)
//		return len(edges) > *last
//	}
//
//	// TODO: handle "after", but according to the spec it's ok to return false
//
//	return false
//}
//
//func HasNextPage(allEdges []Edge, before *cursor, after *cursor, first *int) bool {
//	if first != nil {
//		edges := ApplyCursorToEdges(allEdges, before, after)
//		return len(edges) > *first
//	}
//
//	// TODO: handle "before", but according to the spec it's ok to return false
//
//	return false
//}
