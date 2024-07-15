package events

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TestEvent struct {
	Name    string
	Payload interface{}
}

func (e *TestEvent) GetName() string {
	return e.Name
}

func (e *TestEvent) GetPayload() interface{} {
	return e.Payload
}

func (e *TestEvent) GetDateTime() time.Time {
	return time.Now()
}

type TestEventHandler struct {
	ID int
}

func (h *TestEventHandler) Handle(event EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	println("Handling event:", event.GetName())
}

type EventDispatcherTestSuite struct {
	suite.Suite
	event      TestEvent
	event2     TestEvent
	handler    TestEventHandler
	handler2   TestEventHandler
	handler3   TestEventHandler
	dispatcher *EventDispatcher
}

func (s *EventDispatcherTestSuite) SetupTest() {
	s.event = TestEvent{Name: "test", Payload: "test"}
	s.event2 = TestEvent{Name: "test2", Payload: "test2"}
	s.handler = TestEventHandler{ID: 1}
	s.handler2 = TestEventHandler{ID: 2}
	s.handler3 = TestEventHandler{ID: 3}
	s.dispatcher = NewEventDispatcher()
}

func (s *EventDispatcherTestSuite) TestEventDispatcherRegister() {
	err := s.dispatcher.Register(s.event.GetName(), &s.handler)
	s.Nil(err)
	s.Equal(1, len(s.dispatcher.handlers[s.event.GetName()]))

	err = s.dispatcher.Register(s.event.GetName(), &s.handler)
	s.Equal(ErrHandlerAlreadyRegistered, err)

	err = s.dispatcher.Register(s.event.GetName(), &s.handler2)
	s.Nil(err)
	s.Equal(2, len(s.dispatcher.handlers[s.event.GetName()]))

	s.Equal(&s.handler, s.dispatcher.handlers[s.event.GetName()][0])
	s.Equal(&s.handler2, s.dispatcher.handlers[s.event.GetName()][1])
}

func (s *EventDispatcherTestSuite) TestEventDispatcherRegisterWithSameHandler() {
	err := s.dispatcher.Register(s.event.GetName(), &s.handler)
	s.Nil(err)
	s.Equal(1, len(s.dispatcher.handlers[s.event.GetName()]))

	err = s.dispatcher.Register(s.event.GetName(), &s.handler)
	s.Equal(ErrHandlerAlreadyRegistered, err)
}

func (s *EventDispatcherTestSuite) TestEventDispatcherHas() {
	s.False(s.dispatcher.Has(s.event.GetName(), &s.handler))
	s.False(s.dispatcher.Has(s.event.GetName(), &s.handler2))

	err := s.dispatcher.Register(s.event.GetName(), &s.handler)
	s.Nil(err)
	s.True(s.dispatcher.Has(s.event.GetName(), &s.handler))
	s.False(s.dispatcher.Has(s.event.GetName(), &s.handler2))
}

func (s *EventDispatcherTestSuite) TestEventDispatcherClear() {
	s.dispatcher.Register(s.event.GetName(), &s.handler)
	s.dispatcher.Register(s.event.GetName(), &s.handler2)
	s.dispatcher.Register(s.event2.GetName(), &s.handler)
	s.dispatcher.Register(s.event2.GetName(), &s.handler2)
	s.dispatcher.Register(s.event2.GetName(), &s.handler3)

	s.Equal(2, len(s.dispatcher.handlers[s.event.GetName()]))
	s.Equal(3, len(s.dispatcher.handlers[s.event2.GetName()]))

	s.dispatcher.Clear()
	s.Equal(0, len(s.dispatcher.handlers[s.event.GetName()]))
	s.Equal(0, len(s.dispatcher.handlers[s.event2.GetName()]))
}

type MockHandler struct {
	mock.Mock
}

func (m *MockHandler) Handle(event EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	m.Called(event)
}

func (s *EventDispatcherTestSuite) TestEventDispatcherDispatch() {
	eh := MockHandler{}
	eh2 := MockHandler{}
	eh3 := MockHandler{}
	eh.On("Handle", &s.event)
	eh2.On("Handle", &s.event)

	s.dispatcher.Register(s.event.GetName(), &eh)
	s.dispatcher.Register(s.event.GetName(), &eh2)
	s.dispatcher.Register(s.event2.GetName(), &eh3)
	s.dispatcher.Dispatch(&s.event)

	eh.AssertExpectations(s.T())
	eh.AssertCalled(s.T(), "Handle", &s.event)
	eh.AssertNumberOfCalls(s.T(), "Handle", 1)

	eh2.AssertExpectations(s.T())
	eh2.AssertCalled(s.T(), "Handle", &s.event)
	eh2.AssertNumberOfCalls(s.T(), "Handle", 1)

	eh3.AssertExpectations(s.T())
	eh3.AssertNotCalled(s.T(), "Handle", &s.event)
}

func (s *EventDispatcherTestSuite) TestEventDispatcherRemove() {
	s.dispatcher.Register(s.event.GetName(), &s.handler)
	s.dispatcher.Register(s.event.GetName(), &s.handler2)
	s.dispatcher.Register(s.event2.GetName(), &s.handler)
	s.dispatcher.Register(s.event2.GetName(), &s.handler2)
	s.dispatcher.Register(s.event2.GetName(), &s.handler3)

	s.Equal(2, len(s.dispatcher.handlers[s.event.GetName()]))
	s.Equal(3, len(s.dispatcher.handlers[s.event2.GetName()]))

	s.dispatcher.Remove(s.event.GetName(), &s.handler)
	s.Equal(1, len(s.dispatcher.handlers[s.event.GetName()]))
	s.False(s.dispatcher.Has(s.event.GetName(), &s.handler))
	s.True(s.dispatcher.Has(s.event.GetName(), &s.handler2))
	s.Equal(3, len(s.dispatcher.handlers[s.event2.GetName()]))

	s.dispatcher.Remove(s.event2.GetName(), &s.handler)
	s.Equal(1, len(s.dispatcher.handlers[s.event.GetName()]))
	s.Equal(2, len(s.dispatcher.handlers[s.event2.GetName()]))
	s.False(s.dispatcher.Has(s.event2.GetName(), &s.handler))
	s.True(s.dispatcher.Has(s.event2.GetName(), &s.handler2))
	s.True(s.dispatcher.Has(s.event2.GetName(), &s.handler3))
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuite))
}
