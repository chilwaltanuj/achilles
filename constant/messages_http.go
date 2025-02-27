package constant

const (
	DependenciesLoaded = "The stage is set! Configuration loaded, logger ready, and dependencies summoned to action."
	DependencyFailed   = "One of the dependency responded unexpectedly."

	ServerStartFailure    = "Oops, Achilles stumbled. Failed to unleash the server."
	ServerStartSuccess    = "Achilles, the server, is unleashed, armed with code mightier than the sword."
	ServerShutdownRequest = "\n Signal received: %v. Achilles bids farewell,but the server's legend lives on."

	GracefulShutdownError   = "Achilles's shield has a crack! Graceful retreat is not an option."
	GracefulShutdownSuccess = "With valor, Achilles lays down his sword. Server rests in peace."

	HttpOk                 = "Victory! Your request has triumphed over the challenges."
	HttpRequestFailure     = "The Oracle speaks of doom! Achilles cannot shield it from the fates. (Non-2xx Response)"
	HttpUnauthorized       = "Alert! Achilles detects a Trojan horse in your midst. Security breach identified (401 - Unauthorized)."
	HttpRouteNotFound      = "Achilles scours the horizon, but alas, Troy remains hidden. 404 - Troy Not Found."
	HttpMethodNotSupported = "Achilles arches an eyebrow in confusion. This method isn't part of the epic script. 405 - Method Not Supported."
	HttpServerErrorPanic   = "Oh no, the gods of Olympus have intervened (500). Internal Server Error"
)
