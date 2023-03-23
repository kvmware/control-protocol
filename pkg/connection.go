/*
Copyright 2021 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package control

// Connection handles the low level stuff, reading and writing to the wire
type Connection interface {
	// WriteMessage writes a message to the wire.
	// This method may not perform the actual writing, but it could just add the message to a local queue, which then dispatch the actual message.
	WriteMessage(msg *Message)

	// ReadMessage reads a message from the wire.
	// This blocks untild there's a message available to read.
	// Returns nil if the Connection is closed.
	ReadMessage() *Message

	// Errors returns a channel that signals very bad, usually fatal, errors
	// (like cannot re-establish the connection after several attempts)
	Errors() <-chan error
}
