package actions

type BaseAction interface {
	GetInfo() map[string]interface{}
	GetName() string
	GetLevel() int
}

const (
	ACTION_LEVEL_BASE     = -100
	ACTION_LEVEL_KAI_PAI  = -30
	ACTION_LEVEL_MO_PAI   = -20
	ACTION_LEVEL_CHU_PAI  = -10
	ACTION_LEVEL_GUO_PAI  = 0
	ACTION_LEVEL_CHI_PAI  = 10
	ACTION_LEVEL_PENG_PAI = 20
	ACTION_LEVEL_GANG_PAI = 30
	ACTION_LEVEL_TING_PAI = 35
	ACTION_LEVEL_HU_PAI   = 40
	ACTION_LEVEL_LIU_JU   = 50
)
