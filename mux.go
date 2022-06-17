package dgimux

import "github.com/bwmarrin/discordgo"

type InteractionHandlerFunc func(s *discordgo.Session, i *discordgo.InteractionCreate)

type Mux struct {
	handlers map[string]InteractionHandlerFunc
}

func NewRouter() *Mux {
	return &Mux{
		handlers: make(map[string]InteractionHandlerFunc),
	}
}

func (m *Mux) ApplicationCommand(name string, handler InteractionHandlerFunc) {
	m.handlers[name] = handler
}

func (m *Mux) HandleInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	name := i.ApplicationCommandData().Name

	if h, ok := m.handlers[name]; ok {
		h(s, i)
	}
}
