package models

type Task struct {
    ID          int       `json:"id"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
    OwnerID     int       `json:"owner_id"`
}

type User struct {
    ID        int       `json:"id"`
    Name      string    `json:"name"`
    Email     string    `json:"email"`
    CreatedAt time.Time `json:"created_at"`
}

type TaskComment struct {
    ID      int       `json:"id"`
    TaskID  int       `json:"task_id"`
    UserID  int       `json:"user_id"`
    Content  string    `json:"content"`
    CreatedAt time.Time `json:"created_at"`
}

type TaskAttachment struct {
    ID      int       `json:"id"`
    TaskID  int       `json:"task_id"`
    FileURL string    `json:"file_url"`
    CreatedAt time.Time `json:"created_at"`
}

type TaskStatistics struct {
    TaskID          int `json:"task_id"`
    CommentsCount    int `json:"comments_count"`
    AttachmentsCount  int `json:"attachments_count"`
}