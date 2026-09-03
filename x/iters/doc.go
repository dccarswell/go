// This file's package doc comment was written by Claude (Anthropic's AI assistant).

// Package iters provides lazy, composable operations over Go 1.23+ range-over-func
// iterators.
//
// [Seq] and [Seq2] are locally-defined types with underlying types identical to
// [iter.Seq] and [iter.Seq2] respectively. They exist so this package can attach
// methods to them (Go does not allow adding methods to types declared in another
// package, including generic instantiations of stdlib types). Because Seq and Seq2
// are distinct named types from iter.Seq and iter.Seq2, passing a value between the
// two — e.g. into or out of a stdlib function like slices.Collect or maps.Keys —
// requires an explicit conversion such as iter.Seq[T](mySeq) or Seq[T](stdSeq); Go
// does not convert between named types with identical underlying types implicitly.
// Range-over-func itself does not require any conversion: "for v := range mySeq"
// works directly, since the range-over-func spec matches on the core type of the
// range expression, not its named type.
//
// Almost every operation in this package is lazy and composes by wrapping the
// source sequence in a new function: nothing runs until the resulting Seq or Seq2
// is actually ranged over (directly, or via Pull/Pull2, ForEach, Reduce, Collect,
// Sink, Count, and so on). When a consumer stops iterating early — by returning
// false from a yield function, or by "break"-ing out of a "for range" loop — that
// early stop propagates back through every operation in the chain to the original
// source, so no unnecessary work happens beyond the last element actually consumed.
// The exceptions are [Reverse] and [Sort], which must fully drain their source up
// front — see their doc comments for why laziness isn't possible there.
//
// The operations provided are:
//
//   - [Filter] / [Filter2]: keep only elements matching a predicate.
//   - [Limit] / [Limit2]: truncate a sequence to at most a given number of elements.
//   - [Monitor] / [Monitor2]: stop a sequence early once a context.Context is done.
//   - [Map] / [Map2]: transform each element (and, for Map2, its type) via a function.
//   - [Expand] / [Expand2]: map each element to a sub-sequence and flatten the results
//     (flat-map).
//   - [Reverse]: reverse element order (not lazy — see its doc comment).
//   - [Sort] (with [DefaultSortFunc] for the common ascending case): sort element
//     order (not lazy — see its doc comment).
//   - [Coalesce]: fold each Seq2 pair into a single value, producing a Seq.
//   - [Reduce] / [Reduce2] (with [Sum] as a ready-made summing [ReduceFunc]): fold a
//     sequence into a single accumulated value.
//   - [ForEach] / [ForEach2]: invoke a callback for every element, for side effects.
//   - [Sink] / [Sink2]: fully drain a sequence without collecting anything.
//   - [Count] / [Count2] (via [Seq.Count] / [Seq2.Count]): count the elements.
//   - [Seq.Collect]: collect a Seq into a []T, via the stdlib's slices.Collect.
//   - [Values]: build a Seq from a []T. [Unslice]: build a Seq2 of (index, value)
//     pairs from a []T. [Unmap]: build a Seq2 from a map[K]V.
//   - [Create]: build a Seq directly from a variadic argument list.
//   - [Stringify] / [Join]: convert a Seq[T] to strings and concatenate them.
//   - [ToChan] (via [Seq.ToChan]): drain a Seq onto a channel from a goroutine.
//   - [Pull] / [Pull2]: convert a Seq/Seq2 into an explicit next/stop pull-style
//     iterator, via the stdlib's iter.Pull/iter.Pull2.
package iters
