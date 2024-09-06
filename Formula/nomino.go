package main

// Nomino - Batch rename utility
// Homepage: https://github.com/yaa110/nomino

import (
	"fmt"
	
	"os/exec"
)

func installNomino() {
	// Método 1: Descargar y extraer .tar.gz
	nomino_tar_url := "https://github.com/yaa110/nomino/archive/refs/tags/1.3.5.tar.gz"
	nomino_cmd_tar := exec.Command("curl", "-L", nomino_tar_url, "-o", "package.tar.gz")
	err := nomino_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nomino_zip_url := "https://github.com/yaa110/nomino/archive/refs/tags/1.3.5.zip"
	nomino_cmd_zip := exec.Command("curl", "-L", nomino_zip_url, "-o", "package.zip")
	err = nomino_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nomino_bin_url := "https://github.com/yaa110/nomino/archive/refs/tags/1.3.5.bin"
	nomino_cmd_bin := exec.Command("curl", "-L", nomino_bin_url, "-o", "binary.bin")
	err = nomino_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nomino_src_url := "https://github.com/yaa110/nomino/archive/refs/tags/1.3.5.src.tar.gz"
	nomino_cmd_src := exec.Command("curl", "-L", nomino_src_url, "-o", "source.tar.gz")
	err = nomino_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nomino_cmd_direct := exec.Command("./binary")
	err = nomino_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
