package main

// RobloxTs - TypeScript-to-Luau Compiler for Roblox
// Homepage: https://roblox-ts.com/

import (
	"fmt"
	
	"os/exec"
)

func installRobloxTs() {
	// Método 1: Descargar y extraer .tar.gz
	robloxts_tar_url := "https://registry.npmjs.org/roblox-ts/-/roblox-ts-2.3.0.tgz"
	robloxts_cmd_tar := exec.Command("curl", "-L", robloxts_tar_url, "-o", "package.tar.gz")
	err := robloxts_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	robloxts_zip_url := "https://registry.npmjs.org/roblox-ts/-/roblox-ts-2.3.0.tgz"
	robloxts_cmd_zip := exec.Command("curl", "-L", robloxts_zip_url, "-o", "package.zip")
	err = robloxts_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	robloxts_bin_url := "https://registry.npmjs.org/roblox-ts/-/roblox-ts-2.3.0.tgz"
	robloxts_cmd_bin := exec.Command("curl", "-L", robloxts_bin_url, "-o", "binary.bin")
	err = robloxts_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	robloxts_src_url := "https://registry.npmjs.org/roblox-ts/-/roblox-ts-2.3.0.tgz"
	robloxts_cmd_src := exec.Command("curl", "-L", robloxts_src_url, "-o", "source.tar.gz")
	err = robloxts_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	robloxts_cmd_direct := exec.Command("./binary")
	err = robloxts_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
}
