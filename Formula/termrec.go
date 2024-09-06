package main

// Termrec - Record videos of terminal output
// Homepage: https://angband.pl/termrec.html

import (
	"fmt"
	
	"os/exec"
)

func installTermrec() {
	// Método 1: Descargar y extraer .tar.gz
	termrec_tar_url := "https://github.com/kilobyte/termrec/archive/refs/tags/v0.19.tar.gz"
	termrec_cmd_tar := exec.Command("curl", "-L", termrec_tar_url, "-o", "package.tar.gz")
	err := termrec_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	termrec_zip_url := "https://github.com/kilobyte/termrec/archive/refs/tags/v0.19.zip"
	termrec_cmd_zip := exec.Command("curl", "-L", termrec_zip_url, "-o", "package.zip")
	err = termrec_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	termrec_bin_url := "https://github.com/kilobyte/termrec/archive/refs/tags/v0.19.bin"
	termrec_cmd_bin := exec.Command("curl", "-L", termrec_bin_url, "-o", "binary.bin")
	err = termrec_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	termrec_src_url := "https://github.com/kilobyte/termrec/archive/refs/tags/v0.19.src.tar.gz"
	termrec_cmd_src := exec.Command("curl", "-L", termrec_src_url, "-o", "source.tar.gz")
	err = termrec_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	termrec_cmd_direct := exec.Command("./binary")
	err = termrec_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: xz")
	exec.Command("latte", "install", "xz").Run()
}
