package main

// Fennel - Lua Lisp Language
// Homepage: https://fennel-lang.org

import (
	"fmt"
	
	"os/exec"
)

func installFennel() {
	// Método 1: Descargar y extraer .tar.gz
	fennel_tar_url := "https://github.com/bakpakin/Fennel/archive/refs/tags/1.5.1.tar.gz"
	fennel_cmd_tar := exec.Command("curl", "-L", fennel_tar_url, "-o", "package.tar.gz")
	err := fennel_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fennel_zip_url := "https://github.com/bakpakin/Fennel/archive/refs/tags/1.5.1.zip"
	fennel_cmd_zip := exec.Command("curl", "-L", fennel_zip_url, "-o", "package.zip")
	err = fennel_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fennel_bin_url := "https://github.com/bakpakin/Fennel/archive/refs/tags/1.5.1.bin"
	fennel_cmd_bin := exec.Command("curl", "-L", fennel_bin_url, "-o", "binary.bin")
	err = fennel_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fennel_src_url := "https://github.com/bakpakin/Fennel/archive/refs/tags/1.5.1.src.tar.gz"
	fennel_cmd_src := exec.Command("curl", "-L", fennel_src_url, "-o", "source.tar.gz")
	err = fennel_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fennel_cmd_direct := exec.Command("./binary")
	err = fennel_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: lua")
exec.Command("latte", "install", "lua")
}
