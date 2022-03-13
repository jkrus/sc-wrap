package blockchain

type (
	Service interface {
		// Start tries start service.
		Start() error

		// Stop tries stop service.
		Stop() error

		// Reconnect tries reconnect service.
		Reconnect() error
	}
)
