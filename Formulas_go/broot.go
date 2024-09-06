package main

// Broot - New way to see and navigate directory trees
// Homepage: https://dystroy.org/broot/

import (
	"fmt"
	
	"os/exec"
)

func installBroot() {
	// Método 1: Descargar y extraer .tar.gz
	broot_tar_url := "https://github.com/Canop/broot/archive/refs/tags/v1.43.0.tar.gz"
	broot_cmd_tar := exec.Command("curl", "-L", broot_tar_url, "-o", "package.tar.gz")
	err := broot_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	broot_zip_url := "https://github.com/Canop/broot/archive/refs/tags/v1.43.0.zip"
	broot_cmd_zip := exec.Command("curl", "-L", broot_zip_url, "-o", "package.zip")
	err = broot_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	broot_bin_url := "https://github.com/Canop/broot/archive/refs/tags/v1.43.0.bin"
	broot_cmd_bin := exec.Command("curl", "-L", broot_bin_url, "-o", "binary.bin")
	err = broot_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	broot_src_url := "https://github.com/Canop/broot/archive/refs/tags/v1.43.0.src.tar.gz"
	broot_cmd_src := exec.Command("curl", "-L", broot_src_url, "-o", "source.tar.gz")
	err = broot_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	broot_cmd_direct := exec.Command("./binary")
	err = broot_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: libxcb")
exec.Command("latte", "install", "libxcb")
}
