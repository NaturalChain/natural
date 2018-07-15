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

package main

import (
	"fmt"
	"github.com/NaturalChain/natural/common/config"
	"github.com/micro/cli"
	"os"
	"runtime"
)

func setupApp() *cli.App {
	app := cli.NewApp()
	app.Name = config.Name
	app.Usage = config.Usage
	app.Copyright = config.Copyright
	app.Version = config.Version
	app.UsageText = config.UsageText
	app.Authors = config.Authors
	//app.Action
	//app.After
	//app.ArgsUsage
	//app.BashComplete
	//app.CommandNotFound
	//app.Commands

	app.Before = func(context *cli.Context) error {
		runtime.GOMAXPROCS(runtime.NumCPU())
		return nil
	}

	return app
}

func main() {
	err := setupApp().Run(os.Args)

	if err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}
}
