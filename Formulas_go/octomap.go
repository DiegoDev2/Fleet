package main

// Octomap - Efficient probabilistic 3D mapping framework based on octrees
// Homepage: https://octomap.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installOctomap() {
	// Método 1: Descargar y extraer .tar.gz
	octomap_tar_url := "https://github.com/OctoMap/octomap/archive/refs/tags/v1.10.0.tar.gz"
	octomap_cmd_tar := exec.Command("curl", "-L", octomap_tar_url, "-o", "package.tar.gz")
	err := octomap_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	octomap_zip_url := "https://github.com/OctoMap/octomap/archive/refs/tags/v1.10.0.zip"
	octomap_cmd_zip := exec.Command("curl", "-L", octomap_zip_url, "-o", "package.zip")
	err = octomap_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	octomap_bin_url := "https://github.com/OctoMap/octomap/archive/refs/tags/v1.10.0.bin"
	octomap_cmd_bin := exec.Command("curl", "-L", octomap_bin_url, "-o", "binary.bin")
	err = octomap_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	octomap_src_url := "https://github.com/OctoMap/octomap/archive/refs/tags/v1.10.0.src.tar.gz"
	octomap_cmd_src := exec.Command("curl", "-L", octomap_src_url, "-o", "source.tar.gz")
	err = octomap_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	octomap_cmd_direct := exec.Command("./binary")
	err = octomap_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
