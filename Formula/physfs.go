package main

// Physfs - Library to provide abstract access to various archives
// Homepage: https://icculus.org/physfs/

import (
	"fmt"
	
	"os/exec"
)

func installPhysfs() {
	// Método 1: Descargar y extraer .tar.gz
	physfs_tar_url := "https://github.com/icculus/physfs/archive/refs/tags/release-3.2.0.tar.gz"
	physfs_cmd_tar := exec.Command("curl", "-L", physfs_tar_url, "-o", "package.tar.gz")
	err := physfs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	physfs_zip_url := "https://github.com/icculus/physfs/archive/refs/tags/release-3.2.0.zip"
	physfs_cmd_zip := exec.Command("curl", "-L", physfs_zip_url, "-o", "package.zip")
	err = physfs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	physfs_bin_url := "https://github.com/icculus/physfs/archive/refs/tags/release-3.2.0.bin"
	physfs_cmd_bin := exec.Command("curl", "-L", physfs_bin_url, "-o", "binary.bin")
	err = physfs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	physfs_src_url := "https://github.com/icculus/physfs/archive/refs/tags/release-3.2.0.src.tar.gz"
	physfs_cmd_src := exec.Command("curl", "-L", physfs_src_url, "-o", "source.tar.gz")
	err = physfs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	physfs_cmd_direct := exec.Command("./binary")
	err = physfs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
}
