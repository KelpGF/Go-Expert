# GO - Contexts

A context control the lifecycle of a goroutine. It allows to cancel a goroutine or to pass values to a goroutine.

Contexts are commonly used in the following situations:

- **Deadlines**: to set a deadline for a goroutine.
- **Cancelation**: to cancel a goroutine.
- **Values**: to pass data to a goroutine.
