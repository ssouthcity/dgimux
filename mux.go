package dgimux

import "github.com/bwmarrin/discordgo"

type InteractionHandler interface {
	HandleInteraction(s *discordgo.Session, i *discordgo.InteractionCreate)
}

type InteractionHandlerFunc func(s *discordgo.Session, i *discordgo.InteractionCreate)

func (f InteractionHandlerFunc) HandleInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	f(s, i)
}

type key struct {
	discordgo.InteractionType
	string
}

type Mux struct {
	handlers map[key]InteractionHandler
}

func NewRouter() *Mux {
	return &Mux{
		handlers: make(map[key]InteractionHandler),
	}
}

func (m *Mux) AddInteractionHandler(kind discordgo.InteractionType, id string, handler InteractionHandler) {
	m.handlers[key{kind, id}] = handler
}

func (m *Mux) ApplicationCommand(id string, handler InteractionHandler) {
	m.AddInteractionHandler(discordgo.InteractionApplicationCommand, id, handler)
}

func (m *Mux) ApplicationCommandAutoComplete(id string, handler InteractionHandler) {
	m.AddInteractionHandler(discordgo.InteractionApplicationCommandAutocomplete, id, handler)
}

func (m *Mux) MessageComponent(id string, handler InteractionHandler) {
	m.AddInteractionHandler(discordgo.InteractionMessageComponent, id, handler)
}

func (m *Mux) ModalSubmit(id string, handler InteractionHandler) {
	m.AddInteractionHandler(discordgo.InteractionModalSubmit, id, handler)
}

func (m *Mux) HandleInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	kind := i.Type
	id := m.resolveID(i)

	if h, ok := m.handlers[key{kind, id}]; ok {
		h.HandleInteraction(s, i)
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
