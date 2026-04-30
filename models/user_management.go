package models

type AccessItem struct {
	Code     string `json:"code"`
	Disabled bool   `json:"disabled"`
}

type InviteInfo struct {
	PhoneNumber string `json:"phoneNumber"`
	Position    string `json:"position,omitempty"`
}

type CreateInviteRequest struct {
	Access []AccessItem `json:"access,omitempty"`
	Invite InviteInfo   `json:"invite"`
}

type CreateInviteResponse struct {
	InviteID  string `json:"inviteID"`
	ExpiredAt string `json:"expiredAt"`
	IsSuccess bool   `json:"isSuccess"`
	InviteURL string `json:"inviteUrl"`
}

type GetUsersQuery struct {
	Limit        int64 `json:"-"`
	Offset       int64 `json:"-"`
	IsInviteOnly *bool `json:"-"`
}

type GetUsersResponse struct {
	Total           int64      `json:"total"`
	CountInResponse int64      `json:"countInResponse"`
	Users           []UserInfo `json:"users"`
}

type UserInfo struct {
	ID          int64        `json:"id"`
	Role        string       `json:"role"`
	Position    string       `json:"position"`
	Phone       string       `json:"phone"`
	Email       string       `json:"email"`
	IsOwner     bool         `json:"isOwner"`
	FirstName   string       `json:"firstName"`
	SecondName  string       `json:"secondName"`
	Patronymic  string       `json:"patronymic"`
	GoodsReturn bool         `json:"goodsReturn"`
	IsInvitee   bool         `json:"isInvitee"`
	InviteeInfo *InviteeInfo `json:"inviteeInfo"`
	Access      []AccessItem `json:"access"`
}

type InviteeInfo struct {
	PhoneNumber string `json:"phoneNumber"`
	Position    string `json:"position"`
	InviteUUID  string `json:"inviteUuid"`
	ExpiredAt   string `json:"expiredAt"`
	IsActive    bool   `json:"isActive"`
}

type UpdateUserAccessRequest struct {
	UsersAccesses []UserAccess `json:"usersAccesses"`
}

type UserAccess struct {
	UserID int64        `json:"userId"`
	Access []AccessItem `json:"access"`
}
