package discord

type GuildAuditType int

const (
	AuditTypeGuildUpdate                             GuildAuditType = 1
	AuditTypeChannelCreate                           GuildAuditType = 10
	AuditTypeChannelUpdate                           GuildAuditType = 11
	AuditTypeChannelDelete                           GuildAuditType = 12
	AuditTypeChannelOverwriteCreate                  GuildAuditType = 13
	AuditTypeChannelOverwriteUpdate                  GuildAuditType = 14
	AuditTypeChannelOverwriteDelete                  GuildAuditType = 15
	AuditTypeMemberKick                              GuildAuditType = 20
	AuditTypeMemberPrune                             GuildAuditType = 21
	AuditTypeMemberBanAdd                            GuildAuditType = 22
	AuditTypeMemberBanRemove                         GuildAuditType = 23
	AuditTypeMemberUpdate                            GuildAuditType = 24
	AuditTypeMemberRoleUpdate                        GuildAuditType = 25
	AuditTypeMemberMove                              GuildAuditType = 26
	AuditTypeMemberDisconnect                        GuildAuditType = 27
	AuditTypeBotAdd                                  GuildAuditType = 28
	AuditTypeRoleCreate                              GuildAuditType = 30
	AuditTypeRoleUpdate                              GuildAuditType = 31
	AuditTypeRoleDelete                              GuildAuditType = 32
	AuditTypeInviteCreate                            GuildAuditType = 40
	AuditTypeInviteUpdate                            GuildAuditType = 41
	AuditTypeInviteDelete                            GuildAuditType = 42
	AuditTypeWebhookCreate                           GuildAuditType = 50
	AuditTypeWebhookUpdate                           GuildAuditType = 51
	AuditTypeWebhookDelete                           GuildAuditType = 52
	AuditTypeEmojiCreate                             GuildAuditType = 60
	AuditTypeEmojiUpdate                             GuildAuditType = 61
	AuditTypeEmojiDelete                             GuildAuditType = 62
	AuditTypeMessageDelete                           GuildAuditType = 72
	AuditTypeMessageBulkDelete                       GuildAuditType = 73
	AuditTypeMessagePin                              GuildAuditType = 74
	AuditTypeMessageUnpin                            GuildAuditType = 75
	AuditTypeIntegrationCreate                       GuildAuditType = 80
	AuditTypeIntegrationUpdate                       GuildAuditType = 81
	AuditTypeIntegrationDelete                       GuildAuditType = 82
	AuditTypeStageInstanceCreate                     GuildAuditType = 83
	AuditTypeStageInstanceUpdate                     GuildAuditType = 84
	AuditTypeStageInstanceDelete                     GuildAuditType = 85
	AuditTypeStickerCreate                           GuildAuditType = 90
	AuditTypeStickerUpdate                           GuildAuditType = 91
	AuditTypeStickerDelete                           GuildAuditType = 92
	AuditTypeGuildScheduledEventCreate               GuildAuditType = 100
	AuditTypeGuildScheduledEventUpdate               GuildAuditType = 101
	AuditTypeGuildScheduledEventDelete               GuildAuditType = 102
	AuditTypeThreadCreate                            GuildAuditType = 110
	AuditTypeThreadUpdate                            GuildAuditType = 111
	AuditTypeThreadDelete                            GuildAuditType = 112
	AuditTypeApplicationCommandPermissionUpdate      GuildAuditType = 121
	AuditTypeAutoModerationRuleCreate                GuildAuditType = 140
	AuditTypeAutoModerationRuleUpdate                GuildAuditType = 141
	AuditTypeAutoModerationRuleDelete                GuildAuditType = 142
	AuditTypeAutoModerationBlockMessage              GuildAuditType = 143
	AuditTypeAutoModerationFlagToChannel             GuildAuditType = 144
	AuditTypeAutoModerationUserCommunicationDisabled GuildAuditType = 145
)

type GuildAuditLog struct {
	TargetId   string                `json:"target_id"`
	Changes    []GuildAuditLogChange `json:"changes"`
	UserId     string                `json:"user_id"`
	Id         string                `json:"id"`
	ActionType GuildAuditType        `json:"action_type"`
	GuildId    string                `json:"guild_id"`
	Options    *AuditLogOptions      `json:"options"`
	Reason     string                `json:"reason"`
}

type GuildAuditLogChange struct {
	OldValue string `json:"old_value"`
	NewValue string `json:"new_value"`
	Key      string `json:"key"`
}

type AuditLogOptions struct {
	//TODO: add all options
}
