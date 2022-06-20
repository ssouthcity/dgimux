package dgimux

import "github.com/bwmarrin/discordgo"

type ResponseWriter struct {
	res *discordgo.InteractionResponse
}

func NewResponseWriter() *ResponseWriter {
	return &ResponseWriter{
		res: &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{},
		},
	}
}

func (r *ResponseWriter) Type(itype discordgo.InteractionResponseType) {
	r.res.Type = itype
}

func (r *ResponseWriter) Text(content string) {
	r.res.Data.Content = content
}

func (r *ResponseWriter) Embed(embed *discordgo.MessageEmbed) {
	r.res.Data.Embeds = append(r.res.Data.Embeds, embed)
}

func (r *ResponseWriter) Ephemral() {
	r.res.Data.Flags = 1 << 6
}

func (r *ResponseWriter) ClearComponentRows() {
	r.res.Data.Components = []discordgo.MessageComponent{}
}

func (r *ResponseWriter) ComponentRow(comps ...discordgo.MessageComponent) {
	r.res.Data.Components = append(r.res.Data.Components, &discordgo.ActionsRow{
		Components: comps,
	})
}

func (r *ResponseWriter) Response() *discordgo.InteractionResponse {
	return r.res
}
