package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"log"
	"path"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

type Config struct {
	vp   *viper.Viper
	data map[string]interface{}
}

var (
	configPath  = ""
	configType  = ""
	defaultName = ""
	container   sync.Map
)

func SetDefaultName(name string) {
	if defaultName == "" {
		defaultName = name
	}
}

func SetConfigPath(p string) {
	if configPath == "" {
		configPath = p
	}
}

func SetConfigType(t string) {
	if configType == "" {
		configType = t
	}
}

func GetDefaultName() string {
	if defaultName == "" {
		SetDefaultName("config")
	}

	return defaultName
}

func GetConfigType() string {
	if configType == "" {
		SetConfigType("yaml")
	}

	return configType
}

func GetConfigPath() string {
	if configPath == "" {
		SetConfigPath("etc")
	}

	return configPath
}

func Read(filename string) (c *Config) {
	name := strings.TrimSuffix(filepath.Base(filename), path.Ext(filename))
	v := viper.New()
	c = &Config{
		vp:   v,
		data: make(map[string]interface{}),
	}
	v.SetConfigFile(filename)
	err := v.ReadInConfig()
	container.Store(name, c)
	if err != nil {
		log.Println(fmt.Errorf("config file %s read failed, %s", filename, err))
		return c
	}
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		if err = v.Unmarshal(&c.data); err != nil {
			log.Println(fmt.Errorf("config file %s on change, read failed", filename))
		}
	})
	if err = v.Unmarshal(&c.data); err != nil {
		log.Println(fmt.Errorf("config file %s unmarshal failed", filename))
	}

	return c
}

func Default() *Config {
	return Use(GetDefaultName())
}

func Use(name string) *Config {
	if config, ok := container.Load(name); ok {
		return config.(*Config)
	}

	return Read(path.Join(GetConfigPath(), name+"."+GetConfigType()))
}

func IsSet(key string) bool {
	return Default().IsSet(key)
}

func (c *Config) IsSet(key string) bool {
	return c.vp.IsSet(key)
}

func (c *Config) All() map[string]interface{} {
	return c.vp.AllSettings()
}

func (c *Config) Keys() []string {
	return c.vp.AllKeys()
}

// Get can retrieve any value given the key to use.
func Get(key string, def ...interface{}) interface{} {
	return Default().Get(key, def...)
}

func (c *Config) Get(key string, def ...interface{}) interface{} {
	val := c.vp.Get(key)
	if val == nil && len(def) > 0 {
		val = def[0]
	}
	return val
}

func GetString(key string, def ...string) string {
	return Default().GetString(key, def...)
}

func (c *Config) GetString(key string, def ...string) string {
	val := c.vp.Get(key)
	if val == nil && len(def) > 0 {
		val = def[0]
	}
	return cast.ToString(val)
}

func GetBool(key string, def ...bool) bool {
	return Default().GetBool(key, def...)
}

func (c *Config) GetBool(key string, def ...bool) bool {
	val := c.vp.Get(key)
	if val == nil && len(def) > 0 {
		val = def[0]
	}
	return cast.ToBool(val)
}

func GetInt(key string, def ...int) int {
	return Default().GetInt(key, def...)
}

func (c *Config) GetInt(key string, def ...int) int {
	val := c.vp.Get(key)
	if val == nil && len(def) > 0 {
		val = def[0]
	}
	return cast.ToInt(val)
}

func GetInt8(key string, def ...int8) int8 {
	return Default().GetInt8(key, def...)
}

func (c *Config) GetInt8(key string, def ...int8) int8 {
	val := c.vp.Get(key)
	if val == nil && len(def) > 0 {
		val = def[0]
	}
	return cast.ToInt8(val)
}

func GetInt16(key string, def ...int16) int16 {
	return Default().GetInt16(key, def...)
}

func (c *Config) GetInt16(key string, def ...int16) int16 {
	val := c.vp.Get(key)
	if val == nil && len(def) > 0 {
		val = def[0]
	}
	return cast.ToInt16(val)
}

func GetInt32(key string, def ...int32) int32 {
	return Default().GetInt32(key, def...)
}

func (c *Config) GetInt32(key string, def ...int32) int32 {
	val := c.vp.Get(key)
	if val == nil && len(def) > 0 {
		val = def[0]
	}
	return cast.ToInt32(val)
}

func GetInt64(key string, def ...int64) int64 {
	return Default().GetInt64(key, def...)
}

func (c *Config) GetInt64(key string, def ...int64) int64 {
	val := c.vp.Get(key)
	if val == nil && len(def) > 0 {
		val = def[0]
	}
	return cast.ToInt64(val)
}

func GetUint(key string, def ...uint) uint {
	return Default().GetUint(key, def...)
}

func (c *Config) GetUint(key string, def ...uint) uint {
	val := c.vp.Get(key)
	if val == nil && len(def) > 0 {
		val = def[0]
	}
	return cast.ToUint(val)
}

func GetUint8(key string, def ...uint8) uint8 {
	return Default().GetUint8(key, def...)
}

func (c *Config) GetUint8(key string, def ...uint8) uint8 {
	val := c.vp.Get(key)
	if val == nil && len(def) > 0 {
		val = def[0]
	}
	return cast.ToUint8(val)
}

func GetUint16(key string, def ...uint16) uint16 {
	return Default().GetUint16(key, def...)
}

func (c *Config) GetUint16(key string, def ...uint16) uint16 {
	val := c.vp.Get(key)
	if val == nil && len(def) > 0 {
		val = def[0]
	}
	return cast.ToUint16(val)
}

func GetUint32(key string, def ...uint32) uint32 {
	return Default().GetUint32(key, def...)
}

func (c *Config) GetUint32(key string, def ...uint32) uint32 {
	val := c.vp.Get(key)
	if val == nil && len(def) > 0 {
		val = def[0]
	}
	return cast.ToUint32(val)
}

func GetUint64(key string, def ...uint64) uint64 {
	return Default().GetUint64(key, def...)
}

func (c *Config) GetUint64(key string, def ...uint64) uint64 {
	val := c.vp.Get(key)
	if val == nil && len(def) > 0 {
		val = def[0]
	}
	return cast.ToUint64(val)
}

func GetTime(key string, def ...interface{}) time.Time {
	return cast.ToTime(Get(key, def...))
}

func (c *Config) GetTime(key string, def ...interface{}) time.Time {
	val := c.vp.Get(key)
	if val == nil && len(def) > 0 {
		val = def[0]
	}
	return cast.ToTime(val)
}

func GetDuration(key string, def ...interface{}) time.Duration {
	return cast.ToDuration(Get(key, def...))
}

func (c *Config) GetDuration(key string, def ...interface{}) time.Duration {
	val := c.vp.Get(key)
	if val == nil && len(def) > 0 {
		val = def[0]
	}
	return cast.ToDuration(val)
}

func GetIntSlice(key string, def ...interface{}) []int {
	return Default().GetIntSlice(key, def)
}

func (c *Config) GetIntSlice(key string, def ...interface{}) []int {
	val := c.vp.Get(key)
	if val == nil && len(def) > 0 {
		val = def[0]
	}
	return cast.ToIntSlice(val)
}

func GetStringSlice(key string, def ...string) []string {
	return Default().GetStringSlice(key, def...)
}

func (c *Config) GetStringSlice(key string, def ...string) []string {
	val := c.vp.Get(key)
	if val == nil && len(def) > 0 {
		val = def[0]
	}
	return cast.ToStringSlice(val)
}

func GetStringMap(key string, def ...map[string]interface{}) map[string]interface{} {
	return Default().GetStringMap(key, def...)
}

func (c *Config) GetStringMap(key string, def ...map[string]interface{}) map[string]interface{} {
	val := c.vp.Get(key)
	if val == nil && len(def) > 0 {
		val = def[0]
	}
	return cast.ToStringMap(val)
}

func GetStringMapString(key string, def ...map[string]string) map[string]string {
	return Default().GetStringMapString(key, def...)
}

func (c *Config) GetStringMapString(key string, def ...map[string]string) map[string]string {
	val := c.vp.Get(key)
	if val == nil && len(def) > 0 {
		val = def[0]
	}
	return cast.ToStringMapString(val)
}

func GetStringMapStringSlice(key string, def ...map[string][]string) map[string][]string {
	return Default().GetStringMapStringSlice(key, def...)
}

func (c *Config) GetStringMapStringSlice(key string, def ...map[string][]string) map[string][]string {
	val := c.vp.Get(key)
	if val == nil && len(def) > 0 {
		val = def[0]
	}
	return cast.ToStringMapStringSlice(val)
}

func GetSizeInBytes(key string, def ...uint) uint {
	return Default().GetSizeInBytes(key, def...)
}

func (c *Config) GetSizeInBytes(key string, def ...uint) uint {
	if c.vp.IsSet(key) {
		return c.vp.GetSizeInBytes(key)
	}
	if len(def) == 0 {
		return 0
	}

	return def[0]
}
