package main

// UutilsFindutils - Cross-platform Rust rewrite of the GNU findutils
// Homepage: https://github.com/uutils/findutils

import (
	"fmt"
	
	"os/exec"
)

func installUutilsFindutils() {
	// Método 1: Descargar y extraer .tar.gz
	uutilsfindutils_tar_url := "https://github.com/uutils/findutils/archive/refs/tags/0.6.0.tar.gz"
	uutilsfindutils_cmd_tar := exec.Command("curl", "-L", uutilsfindutils_tar_url, "-o", "package.tar.gz")
	err := uutilsfindutils_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	uutilsfindutils_zip_url := "https://github.com/uutils/findutils/archive/refs/tags/0.6.0.zip"
	uutilsfindutils_cmd_zip := exec.Command("curl", "-L", uutilsfindutils_zip_url, "-o", "package.zip")
	err = uutilsfindutils_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	uutilsfindutils_bin_url := "https://github.com/uutils/findutils/archive/refs/tags/0.6.0.bin"
	uutilsfindutils_cmd_bin := exec.Command("curl", "-L", uutilsfindutils_bin_url, "-o", "binary.bin")
	err = uutilsfindutils_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	uutilsfindutils_src_url := "https://github.com/uutils/findutils/archive/refs/tags/0.6.0.src.tar.gz"
	uutilsfindutils_cmd_src := exec.Command("curl", "-L", uutilsfindutils_src_url, "-o", "source.tar.gz")
	err = uutilsfindutils_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	uutilsfindutils_cmd_direct := exec.Command("./binary")
	err = uutilsfindutils_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: oniguruma")
exec.Command("latte", "install", "oniguruma")
}
