/*
 * Copyright (C) 2018 The NaturalChain Authors
 * This file is part of The NaturalChain library.
 *
 * The NaturalChain is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The NaturalChain is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with The NaturalChain.  If not, see <http://www.gnu.org/licenses/>.
 */

package config

import "github.com/micro/cli"

var Version = "0.0.1"

var Authors = []cli.Author{
	{
		Name: "Sean.Xiao",
		Email: "jxzsxsp@qq.com",
	},
}

const (
	Name = "NaturalChain"
	Usage = "A new public chain project"
	UsageText = "natural [global options] command [command options] [arguments...]"
	Copyright = "Copyright in 2018 The NaturalChain Authors"
)
