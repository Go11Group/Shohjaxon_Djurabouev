package postgres

import (
    "database/sql"

    "github.com/Go11Group/at_lesson/lesson35/model"
)

type ProblemStorage struct {
    Db *sql.DB
}

func CreateProblemStorage(db *sql.DB) *ProblemStorage {
    return &ProblemStorage{Db: db}
}

func (p *ProblemStorage) Create(problem *model.Problem) error {
    _, err := p.Db.Exec("INSERT INTO problems (title, description) VALUES ($1, $2)", problem.Title, problem.Description)
    return err
}

func (p *ProblemStorage) Get(id int) (*model.Problem, error) {
    var problem model.Problem
    err := p.Db.QueryRow("SELECT id, title, description FROM problems WHERE id = $1", id).Scan(&problem.ID, &problem.Title, &problem.Description)
    return &problem, err
}

func (p *ProblemStorage) Update(problem *model.Problem) error {
    _, err := p.Db.Exec("UPDATE problems SET title = $1, description = $2 WHERE id = $3", problem.Title, problem.Description, problem.ID)
    return err
}

func (p *ProblemStorage) Delete(id int) error {
    _, err := p.Db.Exec("DELETE FROM problems WHERE id = $1", id)
    return err
}
