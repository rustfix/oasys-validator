// Copyright 2017 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

// These are network parameters that need to be constant between clients, but
// aren't necessarily consensus related.

const (
	// BloomBitsBlocks is the number of blocks a single bloom bit section vector
	// contains on the server side.
	BloomBitsBlocks uint64 = 4096

	// BloomConfirms is the number of confirmation blocks before a bloom section is
	// considered probably final and its rotated bits are calculated.
	BloomConfirms = 256

	// StableStateThreshold is the reserve number of block state save to disk before delete ancientdb
	StableStateThreshold uint64 = 128
)

var (
	// FullImmutabilityThreshold is the number of blocks after which a chain segment is
	// considered immutable (i.e. soft finality). It is used by the downloader as a
	// hard limit against deep ancestors, by the blockchain against deep reorgs, by
	// the freezer as the cutoff threshold and by clique as the snapshot trust limit.
	FullImmutabilityThreshold uint64 = 360_000 * MaxwellBlockTimeReductionFactorForBSC // =90000
)
