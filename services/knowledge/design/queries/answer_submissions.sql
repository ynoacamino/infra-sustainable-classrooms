-- Answer submissions queries - simplified

-- name: CreateAnswerSubmission :exec
INSERT INTO answer_submissions (
    submission_id,
    question_id,
    selected_answer,
    is_correct
) VALUES ($1, $2, $3, $4);

-- name: GetAnswersBySubmission :many
SELECT 
    a.*,
    q.question_text,
    q.option_a,
    q.option_b,
    q.option_c,
    q.option_d,
    q.correct_answer
FROM answer_submissions a
JOIN questions q ON a.question_id = q.id
WHERE a.submission_id = $1
ORDER BY q.question_order ASC;
