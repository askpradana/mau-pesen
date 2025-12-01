package repository

import (
	"nfldyprdn/maupesen/internal/entity"

	"gorm.io/gorm"
)

type RatingRepository interface {
	Create(rating *entity.Rating) error
	ExistsByBookingID(bookingID string) (bool, error)
	UpdateConsultantAvg(bookingID string) error
}

type ratingRepo struct {
	db *gorm.DB
}

func NewRatingRepository(db *gorm.DB) RatingRepository {
	return &ratingRepo{db}
}

func (r *ratingRepo) Create(rating *entity.Rating) error {
	return r.db.Create(rating).Error
}

func (r *ratingRepo) ExistsByBookingID(bookingID string) (bool, error) {
	var count int64
	err := r.db.Model(&entity.Rating{}).
		Where("booking_id = ?", bookingID).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *ratingRepo) UpdateConsultantAvg(bookingID string) error {
	return r.db.Exec(`
		WITH stats AS (
			SELECT 
				b.consultant_id,
				AVG(r.rating)::numeric(3,2) as avg_rating,
				COUNT(r.id) as total
			FROM ratings r
			JOIN bookings b ON r.booking_id = b.id
			WHERE b.id = ?
			GROUP BY b.consultant_id
		)
		UPDATE consultants c
		SET 
			rating_avg = COALESCE(s.avg_rating, c.rating_avg),
			total_rating = c.total_rating + COALESCE(s.total, 0)
		FROM stats s
		WHERE c.user_id = s.consultant_id
	`, bookingID).Error
}
