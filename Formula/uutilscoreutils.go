package main

// UutilsCoreutils - Cross-platform Rust rewrite of the GNU coreutils
// Homepage: https://github.com/uutils/coreutils

import (
	"fmt"
	
	"os/exec"
)

func installUutilsCoreutils() {
	// Método 1: Descargar y extraer .tar.gz
	uutilscoreutils_tar_url := "https://github.com/uutils/coreutils/archive/refs/tags/0.0.27.tar.gz"
	uutilscoreutils_cmd_tar := exec.Command("curl", "-L", uutilscoreutils_tar_url, "-o", "package.tar.gz")
	err := uutilscoreutils_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	uutilscoreutils_zip_url := "https://github.com/uutils/coreutils/archive/refs/tags/0.0.27.zip"
	uutilscoreutils_cmd_zip := exec.Command("curl", "-L", uutilscoreutils_zip_url, "-o", "package.zip")
	err = uutilscoreutils_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	uutilscoreutils_bin_url := "https://github.com/uutils/coreutils/archive/refs/tags/0.0.27.bin"
	uutilscoreutils_cmd_bin := exec.Command("curl", "-L", uutilscoreutils_bin_url, "-o", "binary.bin")
	err = uutilscoreutils_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	uutilscoreutils_src_url := "https://github.com/uutils/coreutils/archive/refs/tags/0.0.27.src.tar.gz"
	uutilscoreutils_cmd_src := exec.Command("curl", "-L", uutilscoreutils_src_url, "-o", "source.tar.gz")
	err = uutilscoreutils_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	uutilscoreutils_cmd_direct := exec.Command("./binary")
	err = uutilscoreutils_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: make")
	exec.Command("latte", "install", "make").Run()
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: sphinx-doc")
	exec.Command("latte", "install", "sphinx-doc").Run()
}
