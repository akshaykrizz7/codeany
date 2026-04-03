package theme

import "runtime"

// UI figures matching Claude Code's visual style
var (
	// Status indicators
	BlackCircle = func() string {
		if runtime.GOOS == "darwin" {
			return "⏺"
		}
		return "●"
	}()
	BulletOp      = "∙"
	UpArrow       = "↑"
	DownArrow     = "↓"
	LightningBolt = "↯"

	// Effort indicators
	EffortLow    = "○"
	EffortMedium = "◐"
	EffortHigh   = "●"
	EffortMax    = "◉"

	// Media
	PlayIcon  = "▶"
	PauseIcon = "⏸"

	// Status
	RefreshArrow = "↻"
	DiamondOpen  = "◇"
	DiamondFull  = "◆"
	FlagIcon     = "⚑"
	BlockquoteBar = "▎"

	// Checkmarks
	CheckMark = "✓"
	CrossMark = "✗"
	TreeConn  = "⎿"
)
