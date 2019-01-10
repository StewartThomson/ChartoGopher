package chart_writer

import (
	"fmt"
	cg "github.com/StewartThomson/ChartoGopher"
	"io"
)

const (
	LINE_ENDING   = "\r\n"
	END_BRACKET   = "}" + LINE_ENDING
	START_BRACKET = "{" + LINE_ENDING
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

	output := "[Song]" + LINE_ENDING + START_BRACKET
	songInfo := chart.GetSongInfoMap()

	for k, v := range songInfo {
		output += fmt.Sprintf("%s = %v", k, v) + LINE_ENDING
	}

	output += END_BRACKET

	return w.output.Write([]byte(output))
}
