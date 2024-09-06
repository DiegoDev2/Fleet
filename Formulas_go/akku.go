package main

// Akku - Package manager for Scheme
// Homepage: https://akkuscm.org/

import (
	"fmt"
	
	"os/exec"
)

func installAkku() {
	// Método 1: Descargar y extraer .tar.gz
	akku_tar_url := "https://gitlab.com/akkuscm/akku/uploads/819fd1f988c6af5e7df0dfa70aa3d3fe/akku-1.1.0.tar.gz"
	akku_cmd_tar := exec.Command("curl", "-L", akku_tar_url, "-o", "package.tar.gz")
	err := akku_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	akku_zip_url := "https://gitlab.com/akkuscm/akku/uploads/819fd1f988c6af5e7df0dfa70aa3d3fe/akku-1.1.0.zip"
	akku_cmd_zip := exec.Command("curl", "-L", akku_zip_url, "-o", "package.zip")
	err = akku_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	akku_bin_url := "https://gitlab.com/akkuscm/akku/uploads/819fd1f988c6af5e7df0dfa70aa3d3fe/akku-1.1.0.bin"
	akku_cmd_bin := exec.Command("curl", "-L", akku_bin_url, "-o", "binary.bin")
	err = akku_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	akku_src_url := "https://gitlab.com/akkuscm/akku/uploads/819fd1f988c6af5e7df0dfa70aa3d3fe/akku-1.1.0.src.tar.gz"
	akku_cmd_src := exec.Command("curl", "-L", akku_src_url, "-o", "source.tar.gz")
	err = akku_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	akku_cmd_direct := exec.Command("./binary")
	err = akku_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: guile")
exec.Command("latte", "install", "guile")
}
