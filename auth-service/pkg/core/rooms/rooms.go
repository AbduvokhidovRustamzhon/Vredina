package rooms

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Service struct {
	pool *pgxpool.Pool
}

func NewService(pool *pgxpool.Pool) *Service {
	return &Service{pool: pool}
}
type Book struct {
	Id            int64   `json:"id"`
	Name          string	`json:"name"`
	Pages         int	`json:"pages"`
	Removed       bool	`json:"removed"`
	FileName      string	`json:"file_name"`
	Category      string	`json:"category"`
	AmountOfLikes int	`json:"amount_of_likes"`
}

//
//type Rooms struct {
//	Id     int64 `json:"id"`
//	Status bool `json:"status"`
//	TimeInFour int `json:"time_in_four"`
//	TimeInMinutes int `json:"time_in_minutes"`
//	TimeOutFour int `json:"time_out_four"`
//	TimeOutMinutes int `json:"time_out_minutes"`
//	FileName string `json:"file_name"`
//}

func (service *Service) Books(ctx context.Context) (list []Book, err error) {
	list = make([]Book, 0)
	conn, err := service.pool.Acquire(ctx)
	if err != nil {
		return nil, err // TODO: wrap to specific error
	}
	defer conn.Release()
	rows, err := conn.Query(ctx, "SELECT id, name, pages, fileName FROM books WHERE removed = FALSE")
	if err != nil {
		return nil, err // TODO: wrap to specific error
	}
	defer rows.Close()

	for rows.Next() {
		item := Book{}
		err := rows.Scan(&item.Id, &item.Name, &item.Pages, &item.FileName)
		if err != nil {
			return nil, err // TODO: wrap to specific error
		}
		list = append(list, item)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return list, nil
}
