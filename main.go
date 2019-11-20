// Copyright © 2016 Abcum Ltd
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"os"

	"runtime"

	"github.com/pkg/profile"

	"github.com/abcum/surreal/cli"

	"github.com/stackimpact/stackimpact-go"
)

func main() {

	switch os.Getenv("DEBUG") {
	case "cpu":
		defer profile.Start(profile.CPUProfile, profile.ProfilePath("."), profile.NoShutdownHook).Stop()
	case "mem":
		defer profile.Start(profile.MemProfile, profile.ProfilePath("."), profile.NoShutdownHook).Stop()
	case "block":
		defer profile.Start(profile.BlockProfile, profile.ProfilePath("."), profile.NoShutdownHook).Stop()
	}

	stackimpact.Start(stackimpact.Options{
		AgentKey: "3326cc7b7c7de70cf2a29f8320c42c31149f39da",
		AppName:  "Surreal",
	})

	runtime.GOMAXPROCS(runtime.NumCPU())

	cli.Init()

}
