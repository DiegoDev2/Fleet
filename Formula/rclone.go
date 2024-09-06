package main

// Rclone - Rsync for cloud storage
// Homepage: https://rclone.org/

import (
	"fmt"
	
	"os/exec"
)

func installRclone() {
	// Método 1: Descargar y extraer .tar.gz
	rclone_tar_url := "https://github.com/rclone/rclone/archive/refs/tags/v1.67.0.tar.gz"
	rclone_cmd_tar := exec.Command("curl", "-L", rclone_tar_url, "-o", "package.tar.gz")
	err := rclone_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rclone_zip_url := "https://github.com/rclone/rclone/archive/refs/tags/v1.67.0.zip"
	rclone_cmd_zip := exec.Command("curl", "-L", rclone_zip_url, "-o", "package.zip")
	err = rclone_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rclone_bin_url := "https://github.com/rclone/rclone/archive/refs/tags/v1.67.0.bin"
	rclone_cmd_bin := exec.Command("curl", "-L", rclone_bin_url, "-o", "binary.bin")
	err = rclone_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rclone_src_url := "https://github.com/rclone/rclone/archive/refs/tags/v1.67.0.src.tar.gz"
	rclone_cmd_src := exec.Command("curl", "-L", rclone_src_url, "-o", "source.tar.gz")
	err = rclone_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rclone_cmd_direct := exec.Command("./binary")
	err = rclone_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
