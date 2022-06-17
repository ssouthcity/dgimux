package dgimux

import "github.com/bwmarrin/discordgo"

type InteractionHandlerFunc func(s *discordgo.Session, i *discordgo.InteractionCreate)

type key struct {
	discordgo.InteractionType
	string
}

type Mux struct {
	handlers map[key]InteractionHandlerFunc
}

func NewRouter() *Mux {
	return &Mux{
		handlers: make(map[key]InteractionHandlerFunc),
	}
}

func (m *Mux) AddInteractionHandler(kind discordgo.InteractionType, id string, handler InteractionHandlerFunc) {
	m.handlers[key{kind, id}] = handler
}

func (m *Mux) HandleInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	kind := i.Type
	id := m.resolveID(i)

	if h, ok := m.handlers[key{kind, id}]; ok {
		h(s, i)
	}
}

func (m *Mux) resolveID(i *discordgo.InteractionCreate) string {
	switch i.Type {
	case discordgo.InteractionApplicationCommand:
		return i.ApplicationCommandData().Name
	case discordgo.InteractionApplicationCommandAutocomplete:
		return i.ApplicationCommandData().Name
	case discordgo.InteractionMessageComponent:
		return i.MessageComponentData().CustomID
	case discordgo.InteractionModalSubmit:
		return i.ModalSubmitData().CustomID
	}
	return ""
}
