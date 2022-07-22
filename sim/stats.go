package sim

import (
	"strconv"
	"strings"
)

type Stats struct {
	// Common Stats
	GP int

	// Skater Stats
	G                int
	A                int
	PTS              int
	PlusMinus        int
	PIM              int
	EV_G             int
	PP_G             int
	SH_G             int
	GW               int
	EV_A             int
	PP_A             int
	SH_A             int
	S                int
	SPercent         float64
	TOI              int
	ATOI             int
	OPS              float64
	DPS              float64
	PS               float64
	BLK              int
	HIT              int
	FOW              int
	FOL              int
	FOPercent        float64
	CF               int
	CA               int
	CFPercent        float64
	CFPercent_rel    float64
	FF               int
	FA               int
	FFPercent        float64
	FFPercent_rel    float64
	OISHPercent      float64
	OISVPercent      float64
	PDO              float64
	OZSPercent       float64
	DZSPercent       float64
	TOI60            int
	TOIEV            int
	TK               int
	GV               int
	EPlusMinus       float64
	SAtt             int
	ThruPercent      float64
	AvgShift         int
	EV_TOI           int
	EV_CFPercent_rel float64
	EV_GF60          float64
	EV_GA60          float64
	PP_TOI           int
	PP_CFPercent_rel float64
	PP_GF60          float64
	PP_GA60          float64
	SH_TOI           int
	SH_CFPercent_rel float64
	SH_GF60          float64
	SH_GA60          float64

	// Goalie Stats
	GS             int
	W              int
	L              int
	TO             int
	GA             int
	SA             int
	SV             int
	SVPercent      float64
	GAA            float64
	SO             int
	MIN            int
	QS             int
	QSPercent      float64
	RBS            int
	GAPercentMinus int
	GSAA           float64
	AdjGAA         float64
	GPS            float64
}

func NewPlayerFromSlice(record []string) (p *Player, err error) {
	p = new(Player)

	name := strings.Fields(record[0])
	p.Name.First = name[0]
	p.Name.Last = name[1]

	if p.Age, err = ParseIntFromCSV(record[1]); err != nil {
		return nil, err
	}

	if p.Position, err = GetPositionEnum(record[2]); err != nil {
		return nil, err
	}

	if p.GP, err = ParseIntFromCSV(record[3]); err != nil {
		return nil, err
	}

	if p.G, err = ParseIntFromCSV(record[4]); err != nil {
		return nil, err
	}

	if p.A, err = ParseIntFromCSV(record[5]); err != nil {
		return nil, err
	}

	if p.PTS, err = ParseIntFromCSV(record[6]); err != nil {
		return nil, err
	}

	if p.PlusMinus, err = ParseIntFromCSV(record[7]); err != nil {
		return nil, err
	}

	if p.PIM, err = ParseIntFromCSV(record[8]); err != nil {
		return nil, err
	}

	if p.EV_G, err = ParseIntFromCSV(record[9]); err != nil {
		return nil, err
	}

	if p.PP_G, err = ParseIntFromCSV(record[10]); err != nil {
		return nil, err
	}

	if p.SH_G, err = ParseIntFromCSV(record[11]); err != nil {
		return nil, err
	}

	if p.GW, err = ParseIntFromCSV(record[12]); err != nil {
		return nil, err
	}

	if p.EV_A, err = ParseIntFromCSV(record[13]); err != nil {
		return nil, err
	}

	if p.PP_A, err = ParseIntFromCSV(record[14]); err != nil {
		return nil, err
	}

	if p.SH_A, err = ParseIntFromCSV(record[15]); err != nil {
		return nil, err
	}

	if p.S, err = ParseIntFromCSV(record[16]); err != nil {
		return nil, err
	}

	// Needs Nil Check
	if p.SPercent, err = ParseFloatFromCSV(record[17]); err != nil {
		return nil, err
	}

	if p.TOI, err = ParseIntFromCSV(record[18]); err != nil {
		return nil, err
	}

	if p.ATOI, err = TimeToSeconds(record[19]); err != nil {
		return nil, err
	}

	if p.OPS, err = ParseFloatFromCSV(record[20]); err != nil {
		return nil, err
	}

	if p.DPS, err = ParseFloatFromCSV(record[21]); err != nil {
		return nil, err
	}

	if p.PS, err = ParseFloatFromCSV(record[22]); err != nil {
		return nil, err
	}

	if p.BLK, err = ParseIntFromCSV(record[23]); err != nil {
		return nil, err
	}

	if p.HIT, err = ParseIntFromCSV(record[24]); err != nil {
		return nil, err
	}

	if p.FOW, err = ParseIntFromCSV(record[25]); err != nil {
		return nil, err
	}

	if p.FOL, err = ParseIntFromCSV(record[26]); err != nil {
		return nil, err
	}

	// Possible Nil
	if p.FOPercent, err = ParseFloatFromCSV(record[27]); err != nil {
		return nil, err
	}

	p.url = record[28]

	// All can be possible Nil
	//
	//
	if p.CF, err = ParseIntFromCSV(record[29]); err != nil {
		return nil, err
	}

	if p.CA, err = ParseIntFromCSV(record[30]); err != nil {
		return nil, err
	}

	if p.CFPercent, err = ParseFloatFromCSV(record[31]); err != nil {
		return nil, err
	}

	if p.CFPercent_rel, err = ParseFloatFromCSV(record[32]); err != nil {
		return nil, err
	}

	if p.FF, err = ParseIntFromCSV(record[33]); err != nil {
		return nil, err
	}

	if p.FA, err = ParseIntFromCSV(record[34]); err != nil {
		return nil, err
	}

	if p.FFPercent, err = ParseFloatFromCSV(record[35]); err != nil {
		return nil, err
	}

	if p.FFPercent_rel, err = ParseFloatFromCSV(record[36]); err != nil {
		return nil, err
	}

	if p.OISHPercent, err = ParseFloatFromCSV(record[37]); err != nil {
		return nil, err
	}

	if p.OISVPercent, err = ParseFloatFromCSV(record[38]); err != nil {
		return nil, err
	}

	if p.PDO, err = ParseFloatFromCSV(record[39]); err != nil {
		return nil, err
	}

	if p.OZSPercent, err = ParseFloatFromCSV(record[40]); err != nil {
		return nil, err
	}

	if p.DZSPercent, err = ParseFloatFromCSV(record[41]); err != nil {
		return nil, err
	}

	if p.TOI60, err = TimeToSeconds(record[42]); err != nil {
		return nil, err
	}

	if p.TOIEV, err = TimeToSeconds(record[43]); err != nil {
		return nil, err
	}

	if p.TK, err = ParseIntFromCSV(record[44]); err != nil {
		return nil, err
	}

	if p.GV, err = ParseIntFromCSV(record[45]); err != nil {
		return nil, err
	}

	if p.EPlusMinus, err = ParseFloatFromCSV(record[46]); err != nil {
		return nil, err
	}

	if p.SAtt, err = ParseIntFromCSV(record[47]); err != nil {
		return nil, err
	}

	if p.ThruPercent, err = ParseFloatFromCSV(record[48]); err != nil {
		return nil, err
	}

	if p.AvgShift, err = TimeToSeconds(record[49]); err != nil {
		return nil, err
	}

	if p.EV_TOI, err = TimeToSeconds(record[50]); err != nil {
		return nil, err
	}

	if p.EV_CFPercent_rel, err = ParseFloatFromCSV(record[51]); err != nil {
		return nil, err
	}

	if p.EV_GF60, err = ParseFloatFromCSV(record[52]); err != nil {
		return nil, err
	}

	if p.EV_GA60, err = ParseFloatFromCSV(record[53]); err != nil {
		return nil, err
	}

	if p.PP_TOI, err = TimeToSeconds(record[54]); err != nil {
		return nil, err
	}

	if p.PP_CFPercent_rel, err = ParseFloatFromCSV(record[55]); err != nil {
		return nil, err
	}

	if p.PP_GF60, err = ParseFloatFromCSV(record[56]); err != nil {
		return nil, err
	}

	if p.PP_GA60, err = ParseFloatFromCSV(record[57]); err != nil {
		return nil, err
	}

	if p.SH_TOI, err = TimeToSeconds(record[58]); err != nil {
		return nil, err
	}

	if p.SH_CFPercent_rel, err = ParseFloatFromCSV(record[59]); err != nil {
		return nil, err
	}

	if p.SH_GF60, err = ParseFloatFromCSV(record[60]); err != nil {
		return nil, err
	}

	if p.SH_GA60, err = ParseFloatFromCSV(record[61]); err != nil {
		return nil, err
	}

	if p.GS, err = ParseIntFromCSV(record[62]); err != nil {
		return nil, err
	}

	if p.W, err = ParseIntFromCSV(record[63]); err != nil {
		return nil, err
	}

	if p.L, err = ParseIntFromCSV(record[64]); err != nil {
		return nil, err
	}

	if p.TO, err = ParseIntFromCSV(record[65]); err != nil {
		return nil, err
	}

	if p.GA, err = ParseIntFromCSV(record[66]); err != nil {
		return nil, err
	}

	if p.SA, err = ParseIntFromCSV(record[67]); err != nil {
		return nil, err
	}

	if p.SV, err = ParseIntFromCSV(record[68]); err != nil {
		return nil, err
	}

	if p.SVPercent, err = ParseFloatFromCSV(record[69]); err != nil {
		return nil, err
	}

	if p.GAA, err = ParseFloatFromCSV(record[70]); err != nil {
		return nil, err
	}

	if p.SO, err = ParseIntFromCSV(record[71]); err != nil {
		return nil, err
	}

	if p.MIN, err = ParseIntFromCSV(record[72]); err != nil {
		return nil, err
	}

	if p.QS, err = ParseIntFromCSV(record[73]); err != nil {
		return nil, err
	}

	if p.QSPercent, err = ParseFloatFromCSV(record[74]); err != nil {
		return nil, err
	}

	if p.RBS, err = ParseIntFromCSV(record[75]); err != nil {
		return nil, err
	}

	if p.GAPercentMinus, err = ParseIntFromCSV(record[76]); err != nil {
		return nil, err
	}

	if p.GSAA, err = ParseFloatFromCSV(record[77]); err != nil {
		return nil, err
	}

	if p.AdjGAA, err = ParseFloatFromCSV(record[78]); err != nil {
		return nil, err
	}

	if p.GPS, err = ParseFloatFromCSV(record[79]); err != nil {
		return nil, err
	}

	return p, nil
}

func TimeToSeconds(time string) (int, error) {
	if time == "" {
		return 0, nil
	}

	s := strings.Split(time, ":")
	t := 0
	mins, err := strconv.Atoi(s[0])
	if err != nil {
		return 0, err
	}
	t += 60 * mins

	secs, err := strconv.Atoi(s[1])
	if err != nil {
		return 0, err
	}
	t += secs
	return t, nil
}

func ParseFloatFromCSV(float string) (f float64, err error) {
	err = nil
	if f, err = strconv.ParseFloat(float, 64); err != nil {
		f, err = 0.0, nil
	}
	return
}

func ParseIntFromCSV(s string) (i int, err error) {
	err = nil
	if i, err = strconv.Atoi(s); err != nil {
		i, err = 0, nil
	}
	return
}
