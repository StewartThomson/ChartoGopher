package ChartoGopher

type Chart struct {
	Tracks []Track
}

func (c *Chart) AddTrack(track Track) {
	c.Tracks = append(c.Tracks, track)
}