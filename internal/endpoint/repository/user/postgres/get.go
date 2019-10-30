package postgres

import (
	"context"
	"fmt"

	"github.com/nickylogan/guestbook/internal/endpoint/models"
)

func (b *postgresRepository) Fetch(ctx context.Context, name string, count, offset int) (users []models.User, err error) {
	// Build query
	q := fmt.Sprintf(`
		SELECT
			user_id, full_name, user_email, birth_date, create_time, update_time, 
			date_part('year', age(birth_date)) as age
		FROM
			%s
		WHERE
			full_name ILIKE %%$1%%
		LIMIT $2
		OFFSET $3
	`, b.tableName)
	rows, err := b.db.Query(q, name, count, offset)
	if err != nil {
		return
	}
	defer rows.Close()

	// Initialize users to empty array
	users = make([]models.User, 0)
	for rows.Next() {
		var user models.User
		rows.Scan(&user.ID, &user.Name, &user.Email, &user.BirthDate, &user.CreatedAt, &user.UpdatedAt, &user.Age)

		users = append(users, user)
	}
	return
}
