package service

import (
	"context"
	"errors"

	"github.com/manuonda/go-projects/inventario/internal/models"
)

var validRolesToAddProduct []int64 = []int64{1, 2}

var (
	ErrInvalidPermissions = errors.New("User dont have permissions")
)

func (s *service) SaveProduct(ctx context.Context, product models.ProductDTO, email string) error {
	u, err := s.repo.GetUserByEmail(ctx, email)

	if err != nil {
		return err
	}

	roles, err := s.repo.GetUserRolesByUserId(ctx, u.ID)
	if err != nil {
		return err
	}

	userCanAdd := false

	for _, r := range roles {
		for _, vr := range validRolesToAddProduct {
			if vr == r.RoleID {
				userCanAdd = true
			}
		}
	}
	if !userCanAdd {
		return ErrInvalidPermissions
	}

	return s.repo.SaveProduct(
		ctx,
		product.Name,
		product.Description,
		product.Price,
		u.ID,
	)

}

func (s *service) GetProduct(ctx context.Context) ([]models.ProductDTO, error) {
	pp, err := s.repo.GetProducts(ctx)
	if err != nil {
		return nil, err
	}

	products := []models.ProductDTO{}

	for _, p := range pp {
		products = append(products, models.ProductDTO{
			ID:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
		})
	}
	return products, nil
}

func (s *service) GetByIdProduct(ctx context.Context, ID int64) (*models.ProductDTO, error) {
	p, err := s.repo.GetProduct(ctx, ID)
	if err != nil {
		return nil, err
	}

	return &models.ProductDTO{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
	}, err
}
