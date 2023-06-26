package db

import (
	"context"
	"database/sql"

	"github.com/elbombardi/squrl/util"
	_ "github.com/lib/pq"
)

type CustomersRepository interface {
	CheckApiKeyExists(ctx context.Context, apiKey string) (bool, error)
	CheckEmailExists(ctx context.Context, email string) (bool, error)
	CheckUsernameExists(ctx context.Context, username string) (bool, error)
	GetCustomerByApiKey(ctx context.Context, apiKey string) (Customer, error)
	GetCustomerByPrefix(ctx context.Context, prefix string) (Customer, error)
	GetCustomerByUsername(ctx context.Context, username string) (Customer, error)
	InsertNewCustomer(ctx context.Context, arg InsertNewCustomerParams) error
	UpdateCustomerStatusByUsername(ctx context.Context, arg UpdateCustomerStatusByUsernameParams) error
}

type ShortURLsRepository interface {
	CheckShortUrlKeyExists(ctx context.Context, arg CheckShortUrlKeyExistsParams) (bool, error)
	GetShortURLByCustomerIDAndShortURLKey(ctx context.Context, arg GetShortURLByCustomerIDAndShortURLKeyParams) (ShortUrl, error)
	IncrementShortURLClickCount(ctx context.Context, arg IncrementShortURLClickCountParams) error
	InsertNewShortURL(ctx context.Context, arg InsertNewShortURLParams) error
	SetShortURLFirstClickDate(ctx context.Context, arg SetShortURLFirstClickDateParams) error
	SetShortURLLastClickDate(ctx context.Context, arg SetShortURLLastClickDateParams) error
	UpdateShortURLLongURL(ctx context.Context, arg UpdateShortURLLongURLParams) error
	UpdateShortURLStatus(ctx context.Context, arg UpdateShortURLStatusParams) error
	UpdateShortURLTrackingStatus(ctx context.Context, arg UpdateShortURLTrackingStatusParams) error
}
type ClicksRepository interface {
}

type Store interface {
	CustomersRepository
	ShortURLsRepository
	ClicksRepository
	// Transactional calls
	// RegisterMatching(ctx context.Context, params RegisterMatchingParams) error
	// CancelMatching(ctx context.Context, params CancelMatchingParams) error
}

type SQLStore struct {
	*Queries
	db *sql.DB
}

// type RegisterMatchingParams struct {
// 	Order1 OrderInfo
// 	Order2 OrderInfo
// 	Amount int32
// }

// type CancelMatchingParams struct {
// 	Order1 OrderInfo
// 	Order2 OrderInfo
// }

// func (store *SQLStore) RegisterMatching(ctx context.Context, params RegisterMatchingParams) error {
// 	return store.transactional(ctx, func(queries *Queries) error {
// 		//Sort by ID to avoid deadlocks
// 		order1, order2 := params.Order1, params.Order2
// 		if params.Order1.ID > params.Order2.ID {
// 			order1, order2 = params.Order2, params.Order1
// 		}

// 		// Insert new line into order_matching_info
// 		_, err := queries.InsertOrderMatching(ctx, InsertOrderMatchingParams{
// 			Order1ID: sql.NullInt32{
// 				Int32: order1.ID,
// 				Valid: true,
// 			},
// 			Order2ID: sql.NullInt32{
// 				Int32: order2.ID,
// 				Valid: true,
// 			},
// 			Amount: params.Amount,
// 		})
// 		if err != nil {
// 			return err
// 		}

// 		// Updating remaining_amount
// 		matchedAmount1, err := queries.GetMatchedAmount(ctx, sql.NullInt32{
// 			Int32: order1.ID,
// 			Valid: true,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		matchedAmount2, err := queries.GetMatchedAmount(ctx, sql.NullInt32{
// 			Int32: order2.ID,
// 			Valid: true,
// 		})
// 		if err != nil {
// 			return err
// 		}

// 		if order1.Amount < int32(matchedAmount1) {
// 			return &errors.InvalidInput{
// 				Message: fmt.Sprintf(
// 					"Matched amount (%v) cannot be bigger than the original order amount (%v). OrderUUID : %v",
// 					matchedAmount1, order1.Amount, order1.OrderUuid.String()),
// 			}
// 		}
// 		if order2.Amount < int32(matchedAmount2) {
// 			return &errors.InvalidInput{
// 				Message: fmt.Sprintf(
// 					"Matched amount (%v) cannot be bigger than the original order amount (%v). OrderUUID : %v",
// 					matchedAmount2, order2.Amount, order2.OrderUuid),
// 			}
// 		}
// 		err = queries.UpdateOrderRemainingAmount(ctx, UpdateOrderRemainingAmountParams{
// 			OrderUuid:       order1.OrderUuid,
// 			RemainingAmount: order1.Amount - int32(matchedAmount1),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		err = queries.UpdateOrderRemainingAmount(ctx, UpdateOrderRemainingAmountParams{
// 			OrderUuid:       order2.OrderUuid,
// 			RemainingAmount: order2.Amount - int32(matchedAmount2),
// 		})
// 		if err != nil {
// 			return err
// 		}

// 		// Updating orders status
// 		status1 := OrderStatusPARTIALLYMATCHED
// 		if order1.Amount == int32(matchedAmount1) {
// 			status1 = OrderStatusFULLYMATCHED
// 		}
// 		err = queries.UpdateOrderStatus(ctx, UpdateOrderStatusParams{
// 			OrderUuid: order1.OrderUuid,
// 			Status:    status1,
// 		})
// 		if err != nil {
// 			return err
// 		}

// 		status2 := OrderStatusPARTIALLYMATCHED
// 		if order2.Amount == int32(matchedAmount2) {
// 			status2 = OrderStatusFULLYMATCHED
// 		}
// 		err = queries.UpdateOrderStatus(ctx, UpdateOrderStatusParams{
// 			OrderUuid: order2.OrderUuid,
// 			Status:    status2,
// 		})
// 		return err
// 	})
// }

// func (store *SQLStore) CancelMatching(ctx context.Context, params CancelMatchingParams) error {
// 	return store.transactional(ctx, func(queries *Queries) error {
// 		//Sort by ID to avoid deadlocks
// 		order1, order2 := params.Order1, params.Order2
// 		if params.Order1.ID > params.Order2.ID {
// 			order1, order2 = params.Order2, params.Order1
// 		}

// 		// Delete the line from order_matching_info
// 		err := queries.DeleteOrderMatching(ctx, DeleteOrderMatchingParams{
// 			Order1ID: sql.NullInt32{
// 				Int32: order1.ID,
// 				Valid: true,
// 			},
// 			Order2ID: sql.NullInt32{
// 				Int32: order2.ID,
// 				Valid: true,
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}

// 		// Updating order1, and order2 remaining_amount
// 		matchedAmount1, err := queries.GetMatchedAmount(ctx, sql.NullInt32{
// 			Int32: order1.ID,
// 			Valid: true,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		matchedAmount2, err := queries.GetMatchedAmount(ctx, sql.NullInt32{
// 			Int32: order2.ID,
// 			Valid: true,
// 		})
// 		if err != nil {
// 			return err
// 		}

// 		err = queries.UpdateOrderRemainingAmount(ctx, UpdateOrderRemainingAmountParams{
// 			OrderUuid:       order1.OrderUuid,
// 			RemainingAmount: order1.Amount - int32(matchedAmount1),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		err = queries.UpdateOrderRemainingAmount(ctx, UpdateOrderRemainingAmountParams{
// 			OrderUuid:       order2.OrderUuid,
// 			RemainingAmount: order2.Amount - int32(matchedAmount2),
// 		})
// 		if err != nil {
// 			return err
// 		}

// 		// Updating order1, and order2 status
// 		status1 := OrderStatusPARTIALLYMATCHED
// 		if matchedAmount1 == 0 {
// 			status1 = OrderStatusSUBMITTED

// 		}
// 		err = queries.UpdateOrderStatus(ctx, UpdateOrderStatusParams{
// 			OrderUuid: order1.OrderUuid,
// 			Status:    status1,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		status2 := OrderStatusPARTIALLYMATCHED
// 		if matchedAmount2 == 0 {
// 			status2 = OrderStatusSUBMITTED

// 		}
// 		err = queries.UpdateOrderStatus(ctx, UpdateOrderStatusParams{
// 			OrderUuid: order2.OrderUuid,
// 			Status:    status2,
// 		})
// 		return err
// 	})
// }

var dbInstance *sql.DB

func GetStoreInstance() (*SQLStore, error) {
	if dbInstance == nil {
		var err error
		dbInstance, err = sql.Open(*util.ConfigDBDriver(), *util.ConfigDBSource())
		if err != nil {
			return nil, err
		}
		err = dbInstance.Ping()
		if err != nil {
			return nil, err
		}
		value, _ := util.ConfigDBMaxIdleConns()
		dbInstance.SetMaxIdleConns(value)
		value, _ = util.ConfigDBMaxOpenConns()
		dbInstance.SetMaxOpenConns(value)
		duration, _ := util.ConfigDBConnMaxIdleTime()
		dbInstance.SetConnMaxIdleTime(duration)
		duration, _ = util.ConfigDBConnMaxLifeTime()
		dbInstance.SetConnMaxLifetime(duration)
	}

	return &SQLStore{
		db:      dbInstance,
		Queries: &Queries{db: dbInstance},
	}, nil
}

func Finalize() error {
	if dbInstance == nil {
		return nil
	}
	return dbInstance.Close()
}

func (store *SQLStore) transactional(ctx context.Context, fn func(queries *Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return rollbackErr
		}
		return err
	}
	return tx.Commit()
}
