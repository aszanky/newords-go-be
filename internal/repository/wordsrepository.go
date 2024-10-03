package repository

func (r *repository) AddNewWords(word, indonesia, notes string) (err error) {
	queryAddNewWords := `INSERT INTO translation(word, indonesian, notes) VALUES ($1, $2, $3)`

	_, err = r.db.DB.Exec(queryAddNewWords, word, indonesia, notes)
	if err != nil {
		return
	}

	return nil
}
