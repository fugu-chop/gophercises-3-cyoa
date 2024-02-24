package story

const StoryStart = "intro"

type Chapter map[string]Page
type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}
type Page struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}
