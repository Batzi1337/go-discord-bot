package model

import "time"

type Nightwave struct {
	ID                 string               `json:"id"`
	Activation         time.Time            `json:"activation"`
	Expiry             time.Time            `json:"expiry"`
	StartTime          string               `json:"startString"`
	Active             bool                 `json:"active"`
	Params             struct{}             `json:"params"`
	RewardTypes        []string             `json:"rewardTypes"`
	Season             int                  `json:"season"`
	Tag                string               `json:"tag"`
	Phase              int                  `json:"phase"`
	PossibleChallenges []NightwaveChallenge `json:"possibleChallenges"`
	ActiveChallenges   []NightwaveChallenge `json:"activeChallenges"`
}

type NightwaveChallenge struct {
	ID          string    `json:"id"`
	Activation  time.Time `json:"activation"`
	Expiry      time.Time `json:"expiry"`
	StartTime   string    `json:"startString"`
	Active      bool      `json:"active"`
	IsDaily     bool      `json:"isDaily"`
	IsElite     bool      `json:"isElite"`
	Title       string    `json:"title"`
	Description string    `json:"desc"`
	Reputation  int       `json:"reputation"`
}

type Mission struct {
	Reward               Reward   `json:"reward"`
	Node                 string   `json:"node"`
	NodeKey              string   `json:"nodeKey"`
	Faction              string   `json:"faction"`
	FactionKey           string   `json:"factionKey"`
	MaxEnemyLevel        int      `json:"maxEnemyLevel"`
	MinEnemyLevel        int      `json:"minEnemyLevel"`
	MaxWaveNum           int      `json:"maxWaveNum"`
	Type                 string   `json:"type"`
	TypeKey              string   `json:"typeKey"`
	Nightmare            bool     `json:"nightmare"`
	ArchwingRequired     bool     `json:"archwingRequired"`
	IsSharkwing          bool     `json:"isSharkwing"`
	EnemySpec            string   `json:"enemySpec"`
	LevelOverride        string   `json:"levelOverride"`
	AdvancedSpawners     []string `json:"advancedSpawners"`
	RequiredItems        []string `json:"requiredItems"`
	ConsumeRequiredItems bool     `json:"consumeRequiredItems"`
	LeadersAlwaysAllowed bool     `json:"leadersAlwaysAllowed"`
	LevelAuras           []string `json:"levelAuras"`
	Description          string   `json:"description"`
}

type Reward struct {
	CountedItems []CountedItem `json:"countedItems"`
	Thumbnail    string        `json:"thumbnail"`
	Color        int           `json:"color"`
	Credits      int           `json:"credits"`
	AsString     string        `json:"asString"`
	Items        []string      `json:"items"`
	ItemString   string        `json:"itemString"`
}

type CountedItem struct {
	Count int    `json:"count"`
	Type  string `json:"type"`
}

type Alert struct {
	ID          string    `json:"id"`
	Activation  time.Time `json:"activation"`
	Expiry      time.Time `json:"expiry"`
	StartString string    `json:"startString"`
	Active      bool      `json:"active"`
	Mission     Mission   `json:"mission"`
	Expired     bool      `json:"expired"`
	Eta         string    `json:"eta"`
	RewardTypes []string  `json:"rewardTypes"`
}

type Event struct {
	ID                      string                 `json:"id"`
	Activation              time.Time              `json:"activation"`
	Expiry                  time.Time              `json:"expiry"`
	StartString             string                 `json:"startString"`
	Active                  bool                   `json:"active"`
	MaximumScore            int                    `json:"maximumScore"`
	CurrentScore            int                    `json:"currentScore"`
	SmallInterval           int                    `json:"smallInterval"`
	LargeInterval           int                    `json:"largeInterval"`
	Faction                 string                 `json:"faction"`
	Description             string                 `json:"description"`
	Tooltip                 string                 `json:"tooltip"`
	Node                    string                 `json:"node"`
	ConcurrentNodes         []string               `json:"concurrentNodes"`
	VictimNode              string                 `json:"victimNode"`
	ScoreLocTag             string                 `json:"scoreLocTag"`
	Rewards                 []Reward               `json:"rewards"`
	Health                  int                    `json:"health"`
	AffiliatedWith          string                 `json:"affiliatedWith"`
	Jobs                    []Job                  `json:"jobs"`
	InterimSteps            []InterimStep          `json:"interimSteps"`
	ProgressSteps           []ProgressStep         `json:"progressSteps"`
	ProgressTotal           int                    `json:"progressTotal"`
	ShowTotalAtEndOfMission bool                   `json:"showTotalAtEndOfMission"`
	IsPersonal              bool                   `json:"isPersonal"`
	IsCommunity             bool                   `json:"isCommunity"`
	RegionDrops             []string               `json:"regionDrops"`
	ArchwingDrops           []string               `json:"archwingDrops"`
	AsString                string                 `json:"asString"`
	Metadata                map[string]interface{} `json:"metadata"`
	CompletionBonuses       []int                  `json:"completionBonuses"`
	ScoreVar                string                 `json:"scoreVar"`
	AltExpiry               time.Time              `json:"altExpiry"`
	AltActivation           time.Time              `json:"altActivation"`
	NextAlt                 Alt                    `json:"nextAlt"`
}

type Job struct {
	Activation     time.Time `json:"activation"`
	Expiry         time.Time `json:"expiry"`
	RewardPool     []string  `json:"rewardPool"`
	Type           string    `json:"type"`
	EnemyLevels    []int     `json:"enemyLevels"`
	StandingStages []int     `json:"standingStages"`
	MinMR          int       `json:"minMR"`
}

type InterimStep struct {
	Goal        int     `json:"goal"`
	Reward      Reward  `json:"reward"`
	Message     Message `json:"message"`
	WinnerCount int     `json:"winnerCount"`
}

type ProgressStep struct {
	Type        string `json:"type"`
	ProgressAmt int    `json:"progressAmt"`
}

type Alt struct {
	Expiry     time.Time `json:"expiry"`
	Activation time.Time `json:"activation"`
}

type Message struct {
	Sender      string   `json:"sender"`
	Subject     string   `json:"subject"`
	Message     string   `json:"message"`
	SenderIcon  string   `json:"senderIcon"`
	Attachments []string `json:"attachments"`
}

type SyndicateMission struct {
	Nodes      []string  `json:"nodes"`
	Eta        string    `json:"eta"`
	Jobs       []Job     `json:"jobs"`
	Syndicate  string    `json:"syndicate"`
	ID         string    `json:"id"`
	Expiry     time.Time `json:"expiry"`
	Activation time.Time `json:"activation"`
}

type Fissure struct {
	ID          string    `json:"id"`
	Activation  time.Time `json:"activation"`
	Expiry      time.Time `json:"expiry"`
	StartString string    `json:"startString"`
	Active      bool      `json:"active"`
	Node        string    `json:"node"`
	Expired     bool      `json:"expired"`
	Eta         string    `json:"eta"`
	MissionType string    `json:"missionType"`
	MissionKey  string    `json:"missionKey"`
	Tier        string    `json:"tier"`
	TierNum     int       `json:"tierNum"`
	Enemy       string    `json:"enemy"`
	EnemyKey    string    `json:"enemyKey"`
	IsStorm     bool      `json:"isStorm"`
	IsHard      bool      `json:"isHard"`
}

type Item struct {
	URLName  string `json:"url_name"`
	Thumb    string `json:"thumb"`
	ItemName string `json:"item_name"`
	ID       string `json:"id"`
}
