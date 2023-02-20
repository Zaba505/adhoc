// Copyright (c) 2023 Zaba505
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

// Package tictactoe provides a game engine for Tic-Tac-Toe
package tictactoe

import (
	"context"
	"io"
	"math/rand"
	"os"
	"time"

	"github.com/Zaba505/ibento"
	"github.com/cloudevents/sdk-go/v2/event"
	"go.uber.org/zap"
)

type engineOption struct {
	logger *zap.Logger
}

// EngineOption
type EngineOption func(*engineOption)

// WithLogger
func WithLogger(logger *zap.Logger) EngineOption {
	return func(eo *engineOption) {
		eo.logger = logger
	}
}

// GameEngine
type GameEngine struct {
	eventlog *ibento.Log
	logger   *zap.Logger
}

// InitEngine
func InitEngine(dir string, opts ...EngineOption) (*GameEngine, error) {
	eo := engineOption{
		logger: zap.NewNop(),
	}

	for _, opt := range opts {
		opt(&eo)
	}

	eventlog, err := ibento.Open(dir, ibento.WithLogger(eo.logger))
	if err != nil {
		return nil, err
	}

	eng := &GameEngine{
		eventlog: eventlog,
		logger:   eo.logger,
	}
	return eng, nil
}

// Shutdown
func (eng *GameEngine) Shutdown() error {
	return eng.eventlog.Close()
}

type gameOption struct {
	inputSrc   io.Reader
	viewBuffer io.Writer
}

// GameOption
type GameOption func(*gameOption)

// WithInputSource
func WithInputSource(r io.Reader) GameOption {
	return func(gOpt *gameOption) {
		gOpt.inputSrc = r
	}
}

// WithViewBuffer
func WithViewBuffer(w io.Writer) GameOption {
	return func(gOpt *gameOption) {
		gOpt.viewBuffer = w
	}
}

type gameState struct {
}

// Game
type Game struct {
	id            uint64
	notifications chan Notification
	subscriptions chan *Subscription

	currentPlayer uint8
	players       [2]Player

	eventlog *ibento.Log
	board    [9]rune
}

// NewGame will initialize a new TicTacToe game.
//
// Default input source and view buffer are STDIN and STDOUT, respectively.
func (eng *GameEngine) NewGame(playerOne, playerTwo Player, opts ...GameOption) *Game {
	gameOpts := gameOption{
		inputSrc:   os.Stdin,
		viewBuffer: os.Stdout,
	}
	for _, opt := range opts {
		opt(&gameOpts)
	}

	// TODO: game ids should be monotonically increasing
	// TODO: or would uuids be better?
	id := rand.Uint64()

	return &Game{
		id:            id,
		eventlog:      eng.eventlog,
		notifications: make(chan Notification),
		subscriptions: make(chan *Subscription),
		players:       [2]Player{playerOne, playerTwo},
	}
}

// LoadGame will load a TicTacToe game from the stored game data.
func (eng *GameEngine) LoadGame(id uint64) (*Game, error) {
	g := &Game{
		id:       id,
		eventlog: eng.eventlog,
	}
	// TODO: init g.board from event log
	return g, nil
}

// Stop
func (g *Game) Stop() error {
	// TODO
	return nil
}

// Run
func (g *Game) Run() error {
	return nil
}

// Notification
type Notification struct {
	eventId     string
	eventSource string
	eventType   string
}

// Subscription
type Subscription struct {
	notifications chan Notification
}

// Subscribe
func (g *Game) Subscribe() *Subscription {
	s := &Subscription{
		notifications: make(chan Notification),
	}

	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		select {
		case <-ctx.Done():
		case g.subscriptions <- s:
		}
	}()

	return s
}

// Next will block until a notification is received.
func (s *Subscription) Next() Notification {
	return <-s.notifications
}

func (g *Game) publishNotification(ctx context.Context, ev event.Event) error {
	return nil
}

// Events will return all game events logged corresponding to this game.
// This method is safe to call during or after a game has been run.
func (g *Game) Events() ([]event.Event, error) {
	return nil, nil
}
