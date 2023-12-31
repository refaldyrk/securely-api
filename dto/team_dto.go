package dto

type TeamResponse struct {
	TeamID      string `json:"team_id" bson:"team_id"`
	Name        string `json:"name" bson:"name"`
	TotalMember int    `json:"total_member" bson:"total_member"`
	CreatedAt   int64  `json:"created_at" bson:"created_at"`
	UpdatedAt   int64  `json:"updated_at" bson:"updated_at"`
}

type TeamReq struct {
	Name string `json:"name" bson:"name" binding:"required"`
}

type TeamInviteReq struct {
	Email string `json:"email" bson:"email" binding:"required"`
}

type TeamKickReq struct {
	Email string `json:"email" bson:"email" binding:"required"`
}

type TeamPromoteReq struct {
	Role  string `json:"role" bson:"role" binding:"required"`
	Email string `json:"email" bson:"email" binding:"required"`
}
