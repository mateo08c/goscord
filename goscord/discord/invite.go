package discord

import "time"

type Invite struct {
	Uses                     int                  `json:"uses"`
	Type                     int                  `json:"type"`
	Temporary                bool                 `json:"temporary"`
	MaxUses                  int                  `json:"max_uses"`
	MaxAge                   int                  `json:"max_age"`
	GuildId                  string               `json:"guild_id,omitempty"`
	Code                     string               `json:"code"`
	Guild                    *Guild               `json:"guild,omitempty"`
	Channel                  *Channel             `json:"channel"`
	ChanelId                 string               `json:"channel_id"`
	Inviter                  *User                `json:"inviter,omitempty"`
	TargetType               int                  `json:"target_type,omitempty"`
	TargetUser               *User                `json:"target_user,omitempty"`
	TargetApplication        *Application         `json:"target_application,omitempty"`
	ApproximatePresenceCount int                  `json:"approximate_presence_count,omitempty"`
	ApproximateMemberCount   int                  `json:"approximate_member_count,omitempty" `
	CreatedAt                *time.Time           `json:"created_at"`
	ExpiresAt                *time.Time           `json:"expires_at,omitempty"`
	StageInstance            *StageInstance       `json:"stage_instance,omitempty"`
	GuildScheduledEvent      *GuildScheduledEvent `json:"guild_scheduled_event,omitempty"`
}
