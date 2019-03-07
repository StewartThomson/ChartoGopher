package ChartoGopher

import (
	"fmt"
	"sync"
)

type Chart interface {
	AddTimeSignatureChange(numerator int, denominator int, position int) error

	AddTempoChange(bpm int, position int)

	AddTrack(track *Track)

	Write(writer Writer) (int, error)

	GetSongInfoMap() (info map[string]interface{})

	GetTracks() []*Track

	GetSyncProperties() (properties []SyncProperty)

	GetEvents() []event
}

type SongInfo struct {
	//Song info
	SongName     string
	Artist       string
	Charter      string
	Album        string
	Year         string
	Offset       int
	Resolution   int
	Player2      string
	Difficulty   int
	PreviewStart int
	PreviewEnd   int
	Length       int
	Genre        string
	MediaType    string

	//Music file associated with this chart.
	MusicStream string
}

type chart struct {
	SongInfo  SongInfo
	SyncTrack syncTrack
	Tracks    []*Track
	Events    []event
}

func NewChart(songInfo SongInfo, bpm int, timeSigNumerator int, timeSigDenominator int) (*chart, error) {
	err := validateDenominator(timeSigDenominator, 0)
	timeSig := []timeSignature{
		{
			Numerator:   timeSigNumerator,
			Denominator: timeSigDenominator,
			Position:    0,
		},
	}
	tempo := []tempo{
		{
			Bpm:      bpm,
			Position: 0,
		},
	}
	return &chart{
		Tracks:   make([]*Track, 0),
		Events:   make([]event, 0),
		SongInfo: songInfo,
		SyncTrack: syncTrack{
			TimeSignatures: timeSig,
			Tempos:         tempo,
		},
	}, err
}

func (c *chart) AddTimeSignatureChange(numerator int, denominator int, position int) (err error) {
	err = validateDenominator(denominator, position)
	if err != nil {
		return
	}
	c.SyncTrack.TimeSignatures = append(c.SyncTrack.TimeSignatures, timeSignature{
		Numerator:   numerator,
		Denominator: denominator,
		Position:    position,
	})
	return
}

func (c *chart) AddTempoChange(bpm int, position int) {
	c.SyncTrack.Tempos = append(c.SyncTrack.Tempos, tempo{
		Bpm:      bpm,
		Position: position,
	})
}

func (c *chart) AddTrack(track *Track) {
	c.Tracks = append(c.Tracks, track)
}

func (c *chart) Write(writer Writer) (int, error) {
	c.setDefaults()

	return writer.Write(c)
}

func (c *chart) setDefaults() {
	info := &c.SongInfo
	if info.Resolution == 0 {
		info.Resolution = 192
	}
}

func (c chart) GetSongInfoMap() (info map[string]interface{}) {
	songInfo := c.SongInfo
	info = map[string]interface{}{
		"Offset":       songInfo.Offset,
		"Resolution":   songInfo.Resolution,
		"Difficulty":   songInfo.Difficulty,
		"PreviewStart": songInfo.PreviewStart,
		"PreviewEnd":   songInfo.PreviewEnd,
	}

	if songInfo.SongName != "" {
		info["Name"] = songInfo.SongName
	}
	if songInfo.Artist != "" {
		info["Artist"] = songInfo.Artist
	}
	if songInfo.Charter != "" {
		info["Charter"] = songInfo.Charter
	}
	if songInfo.Album != "" {
		info["Album"] = songInfo.Album
	}
	if songInfo.Year != "" {
		info["Year"] = songInfo.Year
	}
	if songInfo.Player2 != "" {
		info["Player2"] = songInfo.Player2
	}
	if songInfo.Length != 0 {
		info["Length"] = songInfo.Length
	}
	if songInfo.Genre != "" {
		info["Genre"] = songInfo.Genre
	}
	if songInfo.MediaType != "" {
		info["MediaType"] = songInfo.MediaType
	}
	if songInfo.MusicStream != "" {
		info["MusicStream"] = songInfo.MusicStream
	}
	return
}

func (c chart) GetSyncProperties() (properties []SyncProperty) {
	properties = make([]SyncProperty, 0)
	ch := make(chan SyncProperty)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		wg.Wait()
		close(ch)
	}()

	go func() {
		defer wg.Done()
		for _, v := range c.SyncTrack.Tempos {
			ch <- v
		}
	}()
	go func() {
		defer wg.Done()
		for _, v := range c.SyncTrack.TimeSignatures {
			ch <- v
		}
	}()

	for prop := range ch {
		properties = append(properties, prop)
	}
	return
}

func (c chart) GetTracks() []*Track {
	return c.Tracks
}

func (c *chart) AddEvent(time int, eventType EventType, comments ...string) {
	c.Events = append(c.Events, event{
		Time:    time,
		Event:   eventType,
		Comment: fmt.Sprint(comments),
	})
}

func (c chart) GetEvents() []event {
	return c.Events
}
