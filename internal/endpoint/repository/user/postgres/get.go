package postgres

import (
	"context"

	"github.com/nickylogan/guestbook/internal/endpoint/models"
)

func (b *postgresRepository) Fetch(ctx context.Context, count, offset int) (users []models.User, err error) {
	// Build query
	q := `
		SELECT
			user_id, full_name, user_email, msisdn, birth_date, create_time, update_time, 
			date_part('year', age(birth_date)) as age
		FROM ` + b.tableName + `
		WHERE
			full_name IS NOT NULL AND
			msisdn IS NOT NULL AND
			birth_date IS NOT NULL
		ORDER BY
			full_name ASC
		LIMIT $1
		OFFSET $2;
	`

	// Execute query
	rows, err := b.db.Query(q, count, offset)
	if err != nil {
		return
	}
	defer func() {
		err = rows.Close()
	}()

	// Initialize users to empty array
	users = make([]models.User, 0)
	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.MSISDN, &user.BirthDate, &user.CreatedAt, &user.UpdatedAt, &user.Age)

		users = append(users, user)
	}
	return
}

func (b *postgresRepository) CountFetch(ctx context.Context, count, offset int) (res int, err error) {
	// Build query
	q := `
		SELECT
			COUNT(user_id)
		FROM ` + b.tableName + `
		WHERE
			full_name IS NOT NULL AND
			msisdn IS NOT NULL AND
			birth_date IS NOT NULL;
	`

	// Execute query
	rows, err := b.db.Query(q)
	if err != nil {
		return
	}
	defer func() {
		err = rows.Close()
	}()

	// Scan results
	for rows.Next() {
		err = rows.Scan(&res)
	}
	return
}

func (b *postgresRepository) FetchFilterName(ctx context.Context, name string, count, offset int) (users []models.User, err error) {
	// Build query
	q := `
		SELECT
			user_id, full_name, user_email, msisdn, birth_date, create_time, update_time, 
			date_part('year', age(birth_date)) as age
		FROM ` + b.tableName + `
		WHERE
			msisdn IS NOT NULL AND
			birth_date IS NOT NULL AND
			full_name ILIKE $1
		ORDER BY
			full_name ASC
		LIMIT $2
		OFFSET $3;
	`

	// Execute query
	filter := "%" + name + "%"
	rows, err := b.db.Query(q, filter, count, offset)
	if err != nil {
		return
	}
	defer func() {
		err = rows.Close()
	}()

	// Initialize users to empty array
	users = make([]models.User, 0)
	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.MSISDN, &user.BirthDate, &user.CreatedAt, &user.UpdatedAt, &user.Age)

		users = append(users, user)
	}
	return
}

func (b *postgresRepository) CountFetchFilterName(ctx context.Context, name string, count, offset int) (res int, err error) {
	// Build query
	q := `
		SELECT
			COUNT(user_id)
		FROM ` + b.tableName + `
		WHERE
			msisdn IS NOT NULL AND
			birth_date IS NOT NULL AND
			full_name ILIKE $1;
	`

	// Execute query
	filter := "%" + name + "%"
	rows, err := b.db.Query(q, filter)
	if err != nil {
		return
	}
	defer func() {
		err = rows.Close()
	}()

	// Scan results
	for rows.Next() {
		err = rows.Scan(&res)
	}
	return
}
