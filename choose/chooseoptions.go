package choose

type Direction int

const (
	DirectionAll Direction = iota
	DirectionH
	DirectionV
)

type Option func(*Model)

func WithTheme(theme Theme) Option {
	return func(m *Model) {
		m.theme = theme
	}
}