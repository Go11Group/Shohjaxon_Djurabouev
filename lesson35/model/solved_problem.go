package model

type SolvedProblem struct {
	UserID    int    `json:"user_id"`
	ProblemID int    `json:"problem_id"`
	SolvedAt  string `json:"solved_at"`
}
