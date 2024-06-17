package dao

func DeletePost(tweetId int) error {

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM tweets WHERE id = ? ", tweetId)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func DeleteReply(replyId int) error {

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM replies WHERE id = ? ", replyId)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

//uuseEffectでuidを取得してuidが一致するデータにのみデリートボタンの表示を実装する
