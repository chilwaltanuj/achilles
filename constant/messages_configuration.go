package constant

// Configuration loading messages adhering to Go's error string conventions
const (
	ConfigDefaultReadError   = "achilles stumbled while reading the oracle's scroll (config.default.json): %w"
	ConfigDefaultLoadSuccess = "achilles nods wisely, having read the oracle's scroll (config.default.json)"
	ConfigEnvLoadFailure     = "achilles raises an eyebrow: 'Are we sure about environment = %s?' Couldn't find the scroll %s: %s. Sticking to the oracle's advice (config.default.json)."
	ConfigEnvLoadSuccess     = "achilles finds and unrolls the scroll %s, blending its secrets with the oracle's wisdom (config.default.json)."
	ConfigUnmarshalError     = "achilles grunts in frustration: Something's amiss with marshalling the scrolls into formation: %w"
	ConfigReady              = "achilles stands ready, configurations in hand and a strategy in mind."
)
