package chart_writer

import (
	"fmt"
	cg "github.com/StewartThomson/ChartoGopher"
	"io"
	"sync"
)

const (
	LINE_ENDING   = "\r\n"
	END_BRACKET   = "}" + LINE_ENDING
	START_BRACKET = "{" + LINE_ENDING
	TAB           = "  "
)

type chartWriter struct {
	output io.Writer
}

func New(output io.Writer) *chartWriter {
	return &chartWriter{
		output: output,
	}
}

func (w *chartWriter) Write(chart cg.Chart) (int, error) {
	tracks := chart.GetTracks()

	ch := make(chan string)
	var wg sync.WaitGroup
	//Tracks + song + synctrack + events
	wg.Add(len(tracks) + 3)

	output := ""
	go func() {
		wg.Wait()
		close(ch)
	}()
	go func() {
		defer wg.Done()
		songBlock := "[Song]" + LINE_ENDING + START_BRACKET

		songInfo := chart.GetSongInfoMap()
		for k, v := range songInfo {
			if _, ok := v.(string); ok {
				songBlock += TAB + fmt.Sprintf("%s = \"%v\"", k, v) + LINE_ENDING
			} else {
				songBlock += TAB + fmt.Sprintf("%s = %v", k, v) + LINE_ENDING
			}
		}
		songBlock += END_BRACKET
		ch <- songBlock
	}()
	go func() {
		defer wg.Done()
		syncBlock := "[SyncTrack]" + LINE_ENDING + START_BRACKET

		props := chart.GetSyncProperties()
		for _, v := range props {
			syncBlock += TAB + v.Location() + " = " + v.Key() + " " + v.Value() + LINE_ENDING
		}
		syncBlock += END_BRACKET
		ch <- syncBlock
	}()

	for _, track := range tracks {
		go func(track cg.Track) {
			defer wg.Done()
			trackBlock := "[" + track.Header() + "]" + LINE_ENDING + START_BRACKET

			for _, note := range track.Notes {
				trackBlock += TAB + fmt.Sprintf("%d = %s %d %d", note.Time, note.NoteChar, note.Colour, note.Duration) + LINE_ENDING
			}

			trackBlock += END_BRACKET
			ch <- trackBlock
		}(*track)
	}

	events := chart.GetEvents()

	go func() {
		defer wg.Done()
		eventBlock := "[Events]" + LINE_ENDING + START_BRACKET
		for _, event := range events {
			eventString := string(event.Event)
			if event.Comment != "" {
				eventString += " " + event.Comment
			}
			eventBlock += TAB + fmt.Sprintf("%d = E \"%s\"", event.Time, eventString) + LINE_ENDING
		}
		eventBlock += END_BRACKET
		ch <- eventBlock
	}()

	for block := range ch {
		output += block
	}

	return w.output.Write([]byte(output))
}
