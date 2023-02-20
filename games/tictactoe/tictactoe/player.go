// Copyright (c) 2023 Zaba505
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package tictactoe

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/google/uuid"
)

// Action
type Action interface {
	apply(*Game) (*Game, error)
}

type actionFunc func(*Game) (*Game, error)

func (f actionFunc) apply(g *Game) (*Game, error) {
	return f(g)
}

var (
	// ErrPositionOutOfBounds
	ErrPositionOutOfBounds = errors.New("tictactoe: piece position out of bounds")
)

// PlacePiece
func PlacePiece(pos uint8, piece rune) Action {
	return actionFunc(func(g *Game) (*Game, error) {
		if pos > 8 {
			return nil, ErrPositionOutOfBounds
		}

		id, err := uuid.NewRandom()
		if err != nil {
			return nil, err
		}

		ev := event.New()
		ev.SetID(id.String())
		ev.SetSource("game")
		ev.SetType("tictactoe.piece.place")
		ev.SetTime(time.Now())

		err = g.eventlog.Append(context.Background(), ev)
		if err != nil {
			return nil, err
		}
		g.board[pos] = piece
		g.publishNotification(context.Background(), ev)
		return g, err
	})
}

// Player
type Player interface {
	Play(*Game) (Action, error)
}

// PlayerFunc
type PlayerFunc func(*Game) (Action, error)

// Play
func (f PlayerFunc) Play(g *Game) (Action, error) {
	return f(g)
}

// HumanPlayer
func HumanPlayer(piece rune) Player {
	return PlayerFunc(func(g *Game) (Action, error) {
		var pos uint8
		_, err := fmt.Print("Enter a position (0-8): ")
		if err != nil {
			return nil, err
		}
		_, err = fmt.Scanln(&pos)
		if err != nil {
			return nil, err
		}

		return PlacePiece(pos, piece), nil
	})
}

type randomAIOption struct {
	seed int64
}

// RandomAIOption
type RandomAIOption func(*randomAIOption)

// WithRandomSeed
func WithRandomSeed(seed int64) RandomAIOption {
	return func(ra *randomAIOption) {
		ra.seed = seed
	}
}

// RandomAIPlayer
func RandomAIPlayer(piece rune, opts ...RandomAIOption) Player {
	ra := randomAIOption{
		seed: time.Now().Unix(),
	}
	for _, opt := range opts {
		opt(&ra)
	}
	r := rand.New(rand.NewSource(ra.seed))

	return PlayerFunc(func(g *Game) (Action, error) {
		n := r.Intn(9)
		return PlacePiece(uint8(n), piece), nil
	})
}
