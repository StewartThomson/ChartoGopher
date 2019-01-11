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

func (w *chartWriter) write(chart cg.Chart) (int, error) {
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
			songBlock += TAB + fmt.Sprintf("%s = %v", k, v) + LINE_ENDING
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
		go func() {
			defer wg.Done()
			trackBlock := "[" + track.Header() + "]" + LINE_ENDING + START_BRACKET

			for _, note := range track.Notes {
				trackBlock += TAB + fmt.Sprintf("%d = N %d %d", note.Time, note.Colour, note.Duration) + LINE_ENDING
			}

			trackBlock += END_BRACKET
		}()
	}

	go func() {
		defer wg.Done()
		close(ch)
	}()

	for block := range ch {
		output += block
	}

	return w.output.Write([]byte(output))
}
