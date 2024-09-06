package main

// Automake - Tool for generating GNU Standards-compliant Makefiles
// Homepage: https://www.gnu.org/software/automake/

import (
	"fmt"
	
	"os/exec"
)

func installAutomake() {
	// Método 1: Descargar y extraer .tar.gz
	automake_tar_url := "https://ftp.gnu.org/gnu/automake/automake-1.17.tar.xz"
	automake_cmd_tar := exec.Command("curl", "-L", automake_tar_url, "-o", "package.tar.gz")
	err := automake_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	automake_zip_url := "https://ftp.gnu.org/gnu/automake/automake-1.17.tar.xz"
	automake_cmd_zip := exec.Command("curl", "-L", automake_zip_url, "-o", "package.zip")
	err = automake_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	automake_bin_url := "https://ftp.gnu.org/gnu/automake/automake-1.17.tar.xz"
	automake_cmd_bin := exec.Command("curl", "-L", automake_bin_url, "-o", "binary.bin")
	err = automake_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	automake_src_url := "https://ftp.gnu.org/gnu/automake/automake-1.17.tar.xz"
	automake_cmd_src := exec.Command("curl", "-L", automake_src_url, "-o", "source.tar.gz")
	err = automake_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	automake_cmd_direct := exec.Command("./binary")
	err = automake_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
}
