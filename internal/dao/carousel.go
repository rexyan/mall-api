// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"mall-api/internal/dao/internal"
)

// internalCarouselDao is internal type for wrapping internal DAO implements.
type internalCarouselDao = *internal.CarouselDao

// carouselDao is the data access object for table tb_newbee_mall_carousel.
// You can define custom methods on it to extend its functionality as you wish.
type carouselDao struct {
	internalCarouselDao
}

var (
	// Carousel is globally public accessible object for table tb_newbee_mall_carousel operations.
	Carousel = carouselDao{
		internal.NewCarouselDao(),
	}
)

// Fill with you ideas below.
