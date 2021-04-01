# Progress - Section 22

- Channels
- Channels are the pipes that connect concurrent goroutines. You can send values into channels from one goroutine and receive those values into another goroutine.
- Send a value into a channel using the channel <- syntax
- The <-channel syntax receives a value from the channel.
- Channels Block - By default sends and receives block until both the sender and receiver are ready.
- A channel may be constrained only to send or only to receive [general to specific] by conversion or assignment
- By default channels are unbuffered, meaning that they will only accept sends (chan <-) if there is a corresponding receive (<- chan) ready to receive the sent value.
- Buffered channels accept a limited number of values without a corresponding receiver for those values.
- Because this channel is buffered, we can send these values into the channel without a corresponding concurrent receive.
- When using channels as function parameters, you can specify if a channel is meant to only send or receive values. This specificity increases the type-safety of the program.

- Go’s select lets you wait on multiple channel operations. Combining goroutines and channels with select is a powerful feature of Go.

- Closing a channel indicates that no more values will be sent on it. This can be useful to communicate completion to the channel’s receivers.
- This range iterates over each element as it’s received from channel. Because we closed the channel above, the iteration terminates after receiving the 2 elements.
- It’s possible to close a non-empty channel but still have the remaining values be received.
