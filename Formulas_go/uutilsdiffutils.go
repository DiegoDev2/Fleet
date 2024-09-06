package main

// UutilsDiffutils - Cross-platform Rust rewrite of the GNU diffutils
// Homepage: https://github.com/uutils/diffutils

import (
	"fmt"
	
	"os/exec"
)

func installUutilsDiffutils() {
	// Método 1: Descargar y extraer .tar.gz
	uutilsdiffutils_tar_url := "https://github.com/uutils/diffutils/archive/refs/tags/v0.4.1.tar.gz"
	uutilsdiffutils_cmd_tar := exec.Command("curl", "-L", uutilsdiffutils_tar_url, "-o", "package.tar.gz")
	err := uutilsdiffutils_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	uutilsdiffutils_zip_url := "https://github.com/uutils/diffutils/archive/refs/tags/v0.4.1.zip"
	uutilsdiffutils_cmd_zip := exec.Command("curl", "-L", uutilsdiffutils_zip_url, "-o", "package.zip")
	err = uutilsdiffutils_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	uutilsdiffutils_bin_url := "https://github.com/uutils/diffutils/archive/refs/tags/v0.4.1.bin"
	uutilsdiffutils_cmd_bin := exec.Command("curl", "-L", uutilsdiffutils_bin_url, "-o", "binary.bin")
	err = uutilsdiffutils_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	uutilsdiffutils_src_url := "https://github.com/uutils/diffutils/archive/refs/tags/v0.4.1.src.tar.gz"
	uutilsdiffutils_cmd_src := exec.Command("curl", "-L", uutilsdiffutils_src_url, "-o", "source.tar.gz")
	err = uutilsdiffutils_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	uutilsdiffutils_cmd_direct := exec.Command("./binary")
	err = uutilsdiffutils_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
