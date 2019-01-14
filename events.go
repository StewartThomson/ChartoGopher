package ChartoGopher

type event struct {
	Time    int
	Event   EventType
	Comment string
}

type EventType string

const (
	EVENT_EVENT_PHRASE_START  EventType = "phrase_start"
	EVENT_PHRASE_END          EventType = "phrase_end"
	EVENT_LYRIC               EventType = "lyric"
	EVENT_IDLE                EventType = "idle"
	EVENT_PLAY                EventType = "play"
	EVENT_HALF_TEMPO          EventType = "half_tempo"
	EVENT_NORMAL_TEMPO        EventType = "normal_tempo"
	EVENT_VERSE               EventType = "verse"
	EVENT_CHORUS              EventType = "chorus"
	EVENT_END                 EventType = "end"
	EVENT_MUSIC_START         EventType = "music_start"
	EVENT_LIGHTING            EventType = "lighting ()"
	EVENT_LIGHTING_FLARE      EventType = "lighting (flare)"
	EVENT_LIGHTING_BLACKOUT   EventType = "lighting (blackout)"
	EVENT_LIGHTING_CHASE      EventType = "lighting (chase)"
	EVENT_LIGHTING_STROBE     EventType = "lighting (strobe)"
	EVENT_LIGHTING_COLOR1     EventType = "lighting (color1)"
	EVENT_LIGHTING_COLOR2     EventType = "lighting (color2)"
	EVENT_LIGHTING_SWEEP      EventType = "lighting (sweep)"
	EVENT_CROWD_LIGHTERS_FAST EventType = "crowd_lighters_fast"
	EVENT_CROWD_LIGHTERS_OFF  EventType = "crowd_lighters_off"
	EVENT_CROWD_LIGHTERS_SLOW EventType = "crowd_lighters_slow"
	EVENT_CROWD_TEMPO_HALF    EventType = "crowd_half_tempo"
	EVENT_CROWD_TEMPO_NORMAL  EventType = "crowd_normal_tempo"
	EVENT_CROWD_TEMPO_DOUBLE  EventType = "crowd_double_tempo"
	EVENT_BAND_JUMP           EventType = "band_jump"
	EVENT_SYNC_HEAD_BANG      EventType = "sync_head_bang"
	EVENT_SYNC_WAG            EventType = "sync_wag"
	EVENT_SECTION             EventType = "section"
)
