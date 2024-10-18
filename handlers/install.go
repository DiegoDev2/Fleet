// Copyright (C) 2024 Fleet Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package handlers

import (
	lib "LattePkg/lib"

	"github.com/fatih/color"
)

var (
	boldGreen = color.New(color.FgGreen, color.Bold)
	redBold   = color.New(color.FgRed, color.Bold)
	green     = color.New(color.FgGreen)
	yellow    = color.New(color.FgYellow)
)

func Install(pkg string) {
	boldGreen.Println("Instalando " + pkg + " 📥")

	installFunc, exists := lib.GetTool(pkg)
	if !exists {
		redBold.Printf("Error: El paquete %s no es reconocido.\n", pkg)
		return
	}

	installFunc()

	boldGreen.Println(pkg + " instalado correctamente 🎉")
}
