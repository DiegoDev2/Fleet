package main

// Luacheck - Tool for linting and static analysis of Lua code
// Homepage: https://luacheck.readthedocs.io/

import (
	"fmt"
	
	"os/exec"
)

func installLuacheck() {
	// Método 1: Descargar y extraer .tar.gz
	luacheck_tar_url := "https://github.com/lunarmodules/luacheck/archive/refs/tags/v1.2.0.tar.gz"
	luacheck_cmd_tar := exec.Command("curl", "-L", luacheck_tar_url, "-o", "package.tar.gz")
	err := luacheck_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	luacheck_zip_url := "https://github.com/lunarmodules/luacheck/archive/refs/tags/v1.2.0.zip"
	luacheck_cmd_zip := exec.Command("curl", "-L", luacheck_zip_url, "-o", "package.zip")
	err = luacheck_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	luacheck_bin_url := "https://github.com/lunarmodules/luacheck/archive/refs/tags/v1.2.0.bin"
	luacheck_cmd_bin := exec.Command("curl", "-L", luacheck_bin_url, "-o", "binary.bin")
	err = luacheck_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	luacheck_src_url := "https://github.com/lunarmodules/luacheck/archive/refs/tags/v1.2.0.src.tar.gz"
	luacheck_cmd_src := exec.Command("curl", "-L", luacheck_src_url, "-o", "source.tar.gz")
	err = luacheck_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	luacheck_cmd_direct := exec.Command("./binary")
	err = luacheck_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: luarocks")
exec.Command("latte", "install", "luarocks")
	fmt.Println("Instalando dependencia: lua")
exec.Command("latte", "install", "lua")
}
