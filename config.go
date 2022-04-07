/*
 * This file is part of the dupman/config project.
 *
 * (c) 2022. dupman <info@dupman.cloud>
 *
 * For the full copyright and license information, please view the LICENSE
 * file that was distributed with this source code.
 *
 * Written by Temuri Takalandze <me@abgeo.dev>
 */

package config

import (
	"github.com/mcuadros/go-defaults"
	"github.com/spf13/viper"
)

func Load(path string, conf interface{}) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AllowEmptyEnv(true)
	viper.AutomaticEnv()

	defaults.SetDefaults(conf)

	_ = viper.ReadInConfig()
	_ = viper.Unmarshal(conf)
}
