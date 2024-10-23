package belajargolangviper

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestViper(t *testing.T) {
	var config *viper.Viper = viper.New()
	assert.NotNil(t, config)
}

func TestJSON(t *testing.T) {
	var config *viper.Viper = viper.New()
	config.SetConfigName("config")
	config.SetConfigType("json")
	config.AddConfigPath(".") // ini mengikuti kaidah UNIX atau mungkin windows juga sama

	err := config.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "belajar-golang-viper", config.GetString("app.name"))
	assert.Equal(t, "Fadli Izurohman", config.GetString("app.author"))
	assert.Equal(t, "localhost", config.GetString("database.host"))
	assert.Equal(t, 3306, config.GetInt("database.port"))
	assert.Equal(t, true, config.GetBool("database.show_sql"))
}
func TestYAML(t *testing.T) {
	var config *viper.Viper = viper.New()
	// config.SetConfigName("config")
	// config.SetConfigType("json")
	config.SetConfigFile("config.yaml")
	config.AddConfigPath(".") // ini mengikuti kaidah UNIX atau mungkin windows juga sama

	err := config.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "belajar-golang-viper", config.GetString("app.name"))
	assert.Equal(t, "Fadli Izurohman", config.GetString("app.author"))
	assert.Equal(t, "localhost", config.GetString("database.host"))
	assert.Equal(t, 3306, config.GetInt("database.port"))
	assert.Equal(t, true, config.GetBool("database.show_sql"))
}
func TestENV(t *testing.T) {
	var config *viper.Viper = viper.New()
	// config.SetConfigName("config")
	// config.SetConfigType("json")
	config.SetConfigFile("config.env")
	config.AddConfigPath(".") // ini mengikuti kaidah UNIX atau mungkin windows juga sama

	err := config.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "belajar-golang-viper", config.GetString("APP_NAME"))
	assert.Equal(t, "Fadli Izurohman", config.GetString("AUTHOR"))
	assert.Equal(t, "localhost", config.GetString("DATABASE_HOST"))
	assert.Equal(t, 3306, config.GetInt("DATABASE_PORT"))
	assert.Equal(t, true, config.GetBool("DATABASE_SHOW_SQL"))
}
func TestENVFILE(t *testing.T) {
	var config *viper.Viper = viper.New()
	// config.SetConfigName("config")
	// config.SetConfigType("json")
	config.SetConfigFile("config.env")
	config.AddConfigPath(".") // ini mengikuti kaidah UNIX atau mungkin windows juga sama
	config.AutomaticEnv()

	err := config.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "belajar-golang-viper", config.GetString("APP_NAME"))
	assert.Equal(t, "Fadli Izurohman", config.GetString("AUTHOR"))
	assert.Equal(t, "localhost", config.GetString("DATABASE_HOST"))
	assert.Equal(t, 3306, config.GetInt("DATABASE_PORT"))
	assert.Equal(t, true, config.GetBool("DATABASE_SHOW_SQL"))

	assert.Equal(t, "Hello", config.GetString("FROM_ENV")) // harus pake terminal dlu pake " export FROM_ENV=Hello "
}