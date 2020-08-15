// create by d1y<chenhonzhou@gmail.com>
// all play sound software
// link: https://stackoverflow.com/questions/34727278/how-can-i-play-an-mp3-file-in-a-cross-platform-bash-script
// github: https://github.com/oxguy3/dotfiles/blob/master/bin/playmp3
// afplay (only osx)
// ffplay -nodisp (ffmpeg only)
// mpg123 --quiet
// maplay | mplayer -really-quiet -noconsolecontrols

package audio

import (
	"fmt"
	"log"

	"github.com/d1y/yoxi/execute"
	"github.com/pkg/errors"
)

// Play play music
// cross platforms only support `mp3` music format
func Play(xPath string) execute.ExecResult {
	return afplay(xPath, .4)
}

// bind osx `afplay` audio player
func afplay(path string, volume float64) execute.ExecResult {
	const play = `afplay`
	var v = fmt.Sprintf("%v", volume)
	var x = execute.ExecTask{
		Command: play,
		Args: []string{
			"--volume",
			v,
			path,
		},
	}
	run, e := x.Execute()
	if e != nil {
		err := errors.Wrap(e, "执行失败")
		log.Printf("%+v", err)
	}
	return run
}
