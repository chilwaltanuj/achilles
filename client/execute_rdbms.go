package client

import (
	"achilles/model"
	"context"
	"fmt"

	"github.com/afex/hystrix-go/hystrix"
)

// ExecuteQuery performs a Hystrix-wrapped RDBMS query and returns the results.
// func ExecuteQuery[T any](ctx context.Context, clientRDBMS *model.ClientRDBMS, query string, args ...interface{}) ([]T, error) {
func ExecuteQuery[T any](ctx context.Context, clientRDBMS *model.ClientRDBMS, request model.RequestClientRDBMS) ([]T, error) {
	logQueryExecutionStart(clientRDBMS.Logger, request.Query, request.Arguments)

	var results []T
	output := make(chan []T, 1)
	errors := hystrix.GoC(ctx, "rdbms_query", func(ctx context.Context) error {
		rows, err := clientRDBMS.DB.QueryContext(ctx, request.Query, request.Arguments)
		if err != nil {
			return fmt.Errorf("query execution failed: %w", err)
		}
		defer rows.Close()
		for rows.Next() {
			var result T
			if err := rows.Scan(&result); err != nil {
				logQueryEnding(clientRDBMS.Logger, request.Query, err)
				return fmt.Errorf("query reading failed: %w", err)
			}
			results = append(results, result)
		}
		output <- results
		return nil
	}, func(ctx context.Context, err error) error {
		return fmt.Errorf("hystrix fallback : %w", err)
	})

	var result []T
	var finalError error
	select {
	case result = <-output:
		return result, nil
	case err := <-errors:
		finalError = err
	case <-ctx.Done():
		finalError = ctx.Err()
	}

	return result, finalError
}
