package main

// Busted - Elegant Lua unit testing
// Homepage: https://lunarmodules.github.io/busted/

import (
	"fmt"
	
	"os/exec"
)

func installBusted() {
	// Método 1: Descargar y extraer .tar.gz
	busted_tar_url := "https://github.com/lunarmodules/busted/archive/refs/tags/v2.2.0.tar.gz"
	busted_cmd_tar := exec.Command("curl", "-L", busted_tar_url, "-o", "package.tar.gz")
	err := busted_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	busted_zip_url := "https://github.com/lunarmodules/busted/archive/refs/tags/v2.2.0.zip"
	busted_cmd_zip := exec.Command("curl", "-L", busted_zip_url, "-o", "package.zip")
	err = busted_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	busted_bin_url := "https://github.com/lunarmodules/busted/archive/refs/tags/v2.2.0.bin"
	busted_cmd_bin := exec.Command("curl", "-L", busted_bin_url, "-o", "binary.bin")
	err = busted_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	busted_src_url := "https://github.com/lunarmodules/busted/archive/refs/tags/v2.2.0.src.tar.gz"
	busted_cmd_src := exec.Command("curl", "-L", busted_src_url, "-o", "source.tar.gz")
	err = busted_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	busted_cmd_direct := exec.Command("./binary")
	err = busted_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: luarocks")
	exec.Command("latte", "install", "luarocks").Run()
	fmt.Println("Instalando dependencia: lua")
	exec.Command("latte", "install", "lua").Run()
}
