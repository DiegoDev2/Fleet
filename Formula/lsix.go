package main

// Lsix - Shows thumbnails in terminal using sixel graphics
// Homepage: https://github.com/hackerb9/lsix

import (
	"fmt"
	
	"os/exec"
)

func installLsix() {
	// Método 1: Descargar y extraer .tar.gz
	lsix_tar_url := "https://github.com/hackerb9/lsix/archive/refs/tags/1.9.1.tar.gz"
	lsix_cmd_tar := exec.Command("curl", "-L", lsix_tar_url, "-o", "package.tar.gz")
	err := lsix_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lsix_zip_url := "https://github.com/hackerb9/lsix/archive/refs/tags/1.9.1.zip"
	lsix_cmd_zip := exec.Command("curl", "-L", lsix_zip_url, "-o", "package.zip")
	err = lsix_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lsix_bin_url := "https://github.com/hackerb9/lsix/archive/refs/tags/1.9.1.bin"
	lsix_cmd_bin := exec.Command("curl", "-L", lsix_bin_url, "-o", "binary.bin")
	err = lsix_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lsix_src_url := "https://github.com/hackerb9/lsix/archive/refs/tags/1.9.1.src.tar.gz"
	lsix_cmd_src := exec.Command("curl", "-L", lsix_src_url, "-o", "source.tar.gz")
	err = lsix_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lsix_cmd_direct := exec.Command("./binary")
	err = lsix_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bash")
	exec.Command("latte", "install", "bash").Run()
	fmt.Println("Instalando dependencia: imagemagick")
	exec.Command("latte", "install", "imagemagick").Run()
}
