package main

// Macvim - GUI for vim, made for macOS
// Homepage: https://github.com/macvim-dev/macvim

import (
	"fmt"
	
	"os/exec"
)

func installMacvim() {
	// Método 1: Descargar y extraer .tar.gz
	macvim_tar_url := "https://github.com/macvim-dev/macvim/archive/refs/tags/release-179.tar.gz"
	macvim_cmd_tar := exec.Command("curl", "-L", macvim_tar_url, "-o", "package.tar.gz")
	err := macvim_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	macvim_zip_url := "https://github.com/macvim-dev/macvim/archive/refs/tags/release-179.zip"
	macvim_cmd_zip := exec.Command("curl", "-L", macvim_zip_url, "-o", "package.zip")
	err = macvim_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	macvim_bin_url := "https://github.com/macvim-dev/macvim/archive/refs/tags/release-179.bin"
	macvim_cmd_bin := exec.Command("curl", "-L", macvim_bin_url, "-o", "binary.bin")
	err = macvim_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	macvim_src_url := "https://github.com/macvim-dev/macvim/archive/refs/tags/release-179.src.tar.gz"
	macvim_cmd_src := exec.Command("curl", "-L", macvim_src_url, "-o", "source.tar.gz")
	err = macvim_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	macvim_cmd_direct := exec.Command("./binary")
	err = macvim_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: libsodium")
	exec.Command("latte", "install", "libsodium").Run()
	fmt.Println("Instalando dependencia: cscope")
	exec.Command("latte", "install", "cscope").Run()
	fmt.Println("Instalando dependencia: lua")
	exec.Command("latte", "install", "lua").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: ruby")
	exec.Command("latte", "install", "ruby").Run()
}
