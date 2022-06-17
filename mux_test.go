package dgimux

import (
	"testing"

	"github.com/bwmarrin/discordgo"
)

func stubHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
}

func TestRouter(t *testing.T) {
	r := NewRouter()

	t.Run("added routes are registered in mux", func(t *testing.T) {
		commands := []string{"ping", "hello", "world"}

		for _, cmd := range commands {
			r.ApplicationCommand(cmd, stubHandler)
		}

		for _, cmd := range commands {
			if _, ok := r.handlers[cmd]; !ok {
				t.Error("handler was not added to struct")
			}
		}
	})

	t.Run("handler is called when interaction matches", func(t *testing.T) {
		called := false
		stubH := InteractionHandlerFunc(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			called = true
		})

		r.ApplicationCommand("pong", stubH)

		r.HandleInteraction(&discordgo.Session{}, &discordgo.InteractionCreate{
			Interaction: &discordgo.Interaction{
				Type: discordgo.InteractionApplicationCommand,
				Data: discordgo.ApplicationCommandInteractionData{
					Name: "pong",
				},
			},
		})

		if !called {
			t.Error("interaction handler for 'pong' was not called")
		}
	})

	t.Run("handler is not called when interaction does not match", func(t *testing.T) {
		called := false
		stubH := InteractionHandlerFunc(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			called = true
		})

		r.ApplicationCommand("pong", stubH)

		r.HandleInteraction(&discordgo.Session{}, &discordgo.InteractionCreate{
			Interaction: &discordgo.Interaction{
				Type: discordgo.InteractionApplicationCommand,
				Data: discordgo.ApplicationCommandInteractionData{
					Name: "ban",
				},
			},
		})

		if called {
			t.Error("interaction handler for 'pong' was called when it shouldn't have been")
		}
	})

	t.Run("multiple interaction types are supported", func(t *testing.T) {
		r.ApplicationCommand("cmd", stubHandler)
		r.MessageComponent("btn", stubHandler)
		r.AutoComplete("auto", stubHandler)
	})
}
