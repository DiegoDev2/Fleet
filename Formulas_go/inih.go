package main

// Inih - Simple .INI file parser in C
// Homepage: https://github.com/benhoyt/inih

import (
	"fmt"
	
	"os/exec"
)

func installInih() {
	// Método 1: Descargar y extraer .tar.gz
	inih_tar_url := "https://github.com/benhoyt/inih/archive/refs/tags/r58.tar.gz"
	inih_cmd_tar := exec.Command("curl", "-L", inih_tar_url, "-o", "package.tar.gz")
	err := inih_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	inih_zip_url := "https://github.com/benhoyt/inih/archive/refs/tags/r58.zip"
	inih_cmd_zip := exec.Command("curl", "-L", inih_zip_url, "-o", "package.zip")
	err = inih_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	inih_bin_url := "https://github.com/benhoyt/inih/archive/refs/tags/r58.bin"
	inih_cmd_bin := exec.Command("curl", "-L", inih_bin_url, "-o", "binary.bin")
	err = inih_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	inih_src_url := "https://github.com/benhoyt/inih/archive/refs/tags/r58.src.tar.gz"
	inih_cmd_src := exec.Command("curl", "-L", inih_src_url, "-o", "source.tar.gz")
	err = inih_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	inih_cmd_direct := exec.Command("./binary")
	err = inih_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
}
