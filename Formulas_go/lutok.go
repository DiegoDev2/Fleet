package main

// Lutok - Lightweight C++ API for Lua
// Homepage: https://github.com/freebsd/lutok

import (
	"fmt"
	
	"os/exec"
)

func installLutok() {
	// Método 1: Descargar y extraer .tar.gz
	lutok_tar_url := "https://github.com/freebsd/lutok/releases/download/lutok-0.4/lutok-0.4.tar.gz"
	lutok_cmd_tar := exec.Command("curl", "-L", lutok_tar_url, "-o", "package.tar.gz")
	err := lutok_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lutok_zip_url := "https://github.com/freebsd/lutok/releases/download/lutok-0.4/lutok-0.4.zip"
	lutok_cmd_zip := exec.Command("curl", "-L", lutok_zip_url, "-o", "package.zip")
	err = lutok_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lutok_bin_url := "https://github.com/freebsd/lutok/releases/download/lutok-0.4/lutok-0.4.bin"
	lutok_cmd_bin := exec.Command("curl", "-L", lutok_bin_url, "-o", "binary.bin")
	err = lutok_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lutok_src_url := "https://github.com/freebsd/lutok/releases/download/lutok-0.4/lutok-0.4.src.tar.gz"
	lutok_cmd_src := exec.Command("curl", "-L", lutok_src_url, "-o", "source.tar.gz")
	err = lutok_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lutok_cmd_direct := exec.Command("./binary")
	err = lutok_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: lua")
exec.Command("latte", "install", "lua")
}
