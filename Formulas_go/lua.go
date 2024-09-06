package main

// Lua - Powerful, lightweight programming language
// Homepage: https://www.lua.org/

import (
	"fmt"
	
	"os/exec"
)

func installLua() {
	// Método 1: Descargar y extraer .tar.gz
	lua_tar_url := "https://www.lua.org/ftp/lua-5.4.7.tar.gz"
	lua_cmd_tar := exec.Command("curl", "-L", lua_tar_url, "-o", "package.tar.gz")
	err := lua_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lua_zip_url := "https://www.lua.org/ftp/lua-5.4.7.zip"
	lua_cmd_zip := exec.Command("curl", "-L", lua_zip_url, "-o", "package.zip")
	err = lua_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lua_bin_url := "https://www.lua.org/ftp/lua-5.4.7.bin"
	lua_cmd_bin := exec.Command("curl", "-L", lua_bin_url, "-o", "binary.bin")
	err = lua_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lua_src_url := "https://www.lua.org/ftp/lua-5.4.7.src.tar.gz"
	lua_cmd_src := exec.Command("curl", "-L", lua_src_url, "-o", "source.tar.gz")
	err = lua_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lua_cmd_direct := exec.Command("./binary")
	err = lua_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
}
