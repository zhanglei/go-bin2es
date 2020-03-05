package bin2es

import (
	"io/ioutil"
	"encoding/json"

	"github.com/BurntSushi/toml"
	"github.com/juju/errors"
)

type Config struct {
	MasterInfo MasterInfo `toml:"master_info"`
	Es       Es       `toml:"es"`
	Mysql    Mysql    `toml:"mysql"`
	Sources  []Source `toml:"source"`
}
type MasterInfo struct {
	Addr    string `toml:"addr"`
	Port    int    `toml:"port"`
	User    string `toml:"user"`
	Pwd     string `toml:"pwd"`
	Charset string `toml:"charset"`
	Schema  string `toml:"schema"`
	Table   string `toml:"table"`
}
type Es struct {
	Nodes         []string `toml:"nodes"`
	BulkSize      int      `toml:"bulk_size"`
	FlushDuration int      `toml:"flush_duration"`
}
type Mysql struct {
	Addr     string `toml:"addr"`
	Port     uint64 `toml:"port"`
	User     string `toml:"user"`
	Pwd      string `toml:"pwd"`
	Charset  string `toml:"charset"`
	ServerID uint32 `toml:"server_id"`
}
type Source struct {
	Schema  string   `toml:"schema"`
	Tables  []string `toml:"tables"`
}

func NewConfigWithFile(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return NewConfig(string(data))
}

// NewConfig creates a Config from data.
func NewConfig(data string) (*Config, error) {
	var c Config

	_, err := toml.Decode(data, &c)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return &c, nil
}

type Pos2Field  map[string]int
type Field      map[string]string
type Replace    map[string]string
type Pipeline   map[string]interface{}
type ROWS       []map[string]interface{}

type Bin2esConfig []struct {
	Schema   	string   	`json:"schema"`
	Tables   	[]string 	`json:"tables"`
	Actions  	[]string 	`json:"actions"`
	Pipelines 	[]Pipeline 	`json:"pipelines"`
	Dest     	Dest     	`json:"dest"`
}
type Dest struct {
	Index  string `json:"index"`
}

func NewBin2esConfig(path string, config *Bin2esConfig) error {
	data, err := ioutil.ReadFile(path)
    if err != nil {
        return err
    }

    //读取的数据为json格式，需要进行解码
    err = json.Unmarshal(data, config)
    if err != nil {
        return err
    }

    return nil
}