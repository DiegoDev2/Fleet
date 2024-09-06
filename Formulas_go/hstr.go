package main

// Hstr - Bash and zsh history suggest box
// Homepage: https://github.com/dvorka/hstr

import (
	"fmt"
	
	"os/exec"
)

func installHstr() {
	// Método 1: Descargar y extraer .tar.gz
	hstr_tar_url := "https://github.com/dvorka/hstr/archive/refs/tags/3.1.tar.gz"
	hstr_cmd_tar := exec.Command("curl", "-L", hstr_tar_url, "-o", "package.tar.gz")
	err := hstr_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hstr_zip_url := "https://github.com/dvorka/hstr/archive/refs/tags/3.1.zip"
	hstr_cmd_zip := exec.Command("curl", "-L", hstr_zip_url, "-o", "package.zip")
	err = hstr_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hstr_bin_url := "https://github.com/dvorka/hstr/archive/refs/tags/3.1.bin"
	hstr_cmd_bin := exec.Command("curl", "-L", hstr_bin_url, "-o", "binary.bin")
	err = hstr_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hstr_src_url := "https://github.com/dvorka/hstr/archive/refs/tags/3.1.src.tar.gz"
	hstr_cmd_src := exec.Command("curl", "-L", hstr_src_url, "-o", "source.tar.gz")
	err = hstr_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hstr_cmd_direct := exec.Command("./binary")
	err = hstr_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
}
