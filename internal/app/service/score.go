package service

import (
	"fmt"
	"rank/internal/app/model"

	"github.com/jinzhu/gorm"
	"gitlab.com/makeblock-go/mysql"
	"gitlab.com/makeblock-go/redis"
)

const (
	initScore = 0
)

func AddScore(userId int64, matchId int64, score int) error {
	if err := mysql.GetDB().First(&model.MatchUser{}).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			mysql.GetDB().Create(&model.MatchUser{
				UserID:  userId,
				MatchID: matchId,
				Score:   initScore,
			})

			return nil
		}

		return err
	}

	err := mysql.GetDB().Transaction(func(tx *gorm.DB) error {
		if err := mysql.GetDB().Exec("update match_user set score = score + ? where user_id = ? and match_id = ?", score, userId, matchId).Error; err != nil {
			return err
		}

		key := fmt.Sprintf("match:%d",matchId)
		increment := float64(score)
		member := fmt.Sprintf("user:%d", userId)
		if err := redis.GetClient().ZIncrBy(key, increment, member).Err(); err != nil {
			return err
		}
		return nil
	})

	return err
}
