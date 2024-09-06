package main

// Luarocks - Package manager for the Lua programming language
// Homepage: https://luarocks.org/

import (
	"fmt"
	
	"os/exec"
)

func installLuarocks() {
	// Método 1: Descargar y extraer .tar.gz
	luarocks_tar_url := "https://luarocks.org/releases/luarocks-3.11.1.tar.gz"
	luarocks_cmd_tar := exec.Command("curl", "-L", luarocks_tar_url, "-o", "package.tar.gz")
	err := luarocks_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	luarocks_zip_url := "https://luarocks.org/releases/luarocks-3.11.1.zip"
	luarocks_cmd_zip := exec.Command("curl", "-L", luarocks_zip_url, "-o", "package.zip")
	err = luarocks_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	luarocks_bin_url := "https://luarocks.org/releases/luarocks-3.11.1.bin"
	luarocks_cmd_bin := exec.Command("curl", "-L", luarocks_bin_url, "-o", "binary.bin")
	err = luarocks_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	luarocks_src_url := "https://luarocks.org/releases/luarocks-3.11.1.src.tar.gz"
	luarocks_cmd_src := exec.Command("curl", "-L", luarocks_src_url, "-o", "source.tar.gz")
	err = luarocks_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	luarocks_cmd_direct := exec.Command("./binary")
	err = luarocks_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: luajit")
exec.Command("latte", "install", "luajit")
	fmt.Println("Instalando dependencia: lua")
exec.Command("latte", "install", "lua")
}
