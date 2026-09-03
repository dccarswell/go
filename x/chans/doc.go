// This file's package doc comment was written by Claude (Anthropic's AI assistant).

// Package chans provides small helpers for working with Go channels: bridging
// them to and from the stdlib's [iter.Seq] iterators, and splitting a
// bidirectional channel (or minting a fresh one) into separate send-only and
// receive-only views.
//
// [Iter2Chan] and [Chan2Iter] convert between iter.Seq[T] and <-chan T /
// chan T in each direction. [Pair] and [Split] both return a (chan<- T,
// <-chan T) pair backed by the same underlying channel — Pair makes a new
// channel of the given buffer length, Split takes an existing one — so an
// action on one view (a send, or closing the send side) is observable
// through the other.
package chans
