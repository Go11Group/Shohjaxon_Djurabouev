package postgres

import (
    "database/sql"

    "github.com/Go11Group/at_lesson/lesson35/model"
)

type SolvedProblemStorage struct {
    Db *sql.DB
}

func NewSolvedProblemStorage(db *sql.DB) *SolvedProblemStorage {
    return &SolvedProblemStorage{Db: db}
}

func (s *SolvedProblemStorage) Create(solvedProblem *model.SolvedProblem) error {
    _, err := s.Db.Exec("INSERT INTO solved_problems (user_id, problem_id, solved_at) VALUES ($1, $2, $3)", solvedProblem.UserID, solvedProblem.ProblemID, solvedProblem.SolvedAt)
    return err
}

func (s *SolvedProblemStorage) Get(userID, problemID int) (*model.SolvedProblem, error) {
    var solvedProblem model.SolvedProblem
    err := s.Db.QueryRow("SELECT user_id, problem_id, solved_at FROM solved_problems WHERE user_id = $1 AND problem_id = $2", userID, problemID).Scan(&solvedProblem.UserID, &solvedProblem.ProblemID, &solvedProblem.SolvedAt)
    return &solvedProblem, err
}

func (s *SolvedProblemStorage) Delete(userID, problemID int) error {
    _, err := s.Db.Exec("DELETE FROM solved_problems WHERE user_id = $1 AND problem_id = $2", userID, problemID)
    return err
}
