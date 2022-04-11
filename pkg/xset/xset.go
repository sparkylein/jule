package xset

import "encoding/json"

const (
	ModeTranspile = "transpile"
	ModeCompile   = "compile"
)

type XSet struct {
	CxxOutDir    string   `json:"cxx_out_dir"`
	CxxOutName   string   `json:"cxx_out_name"`
	OutName      string   `json:"out_name"`
	Language     string   `json:"language"`
	Mode         string   `json:"mode"`
	PostCommands []string `json:"post_commands"`
}

// Load loads XSet from json string.
func Load(bytes []byte) (*XSet, error) {
	set := XSet{
		CxxOutDir:    "./dist",
		CxxOutName:   "x.cxx",
		OutName:      "main",
		Language:     "",
		Mode:         "transpile",
		PostCommands: nil,
	}
	err := json.Unmarshal(bytes, &set)
	if err != nil {
		return nil, err
	}
	return &set, nil
}
