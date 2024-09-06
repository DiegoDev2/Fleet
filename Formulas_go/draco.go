package main

// Draco - 3D geometric mesh and point cloud compression library
// Homepage: https://google.github.io/draco/

import (
	"fmt"
	
	"os/exec"
)

func installDraco() {
	// Método 1: Descargar y extraer .tar.gz
	draco_tar_url := "https://github.com/google/draco/archive/refs/tags/1.5.7.tar.gz"
	draco_cmd_tar := exec.Command("curl", "-L", draco_tar_url, "-o", "package.tar.gz")
	err := draco_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	draco_zip_url := "https://github.com/google/draco/archive/refs/tags/1.5.7.zip"
	draco_cmd_zip := exec.Command("curl", "-L", draco_zip_url, "-o", "package.zip")
	err = draco_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	draco_bin_url := "https://github.com/google/draco/archive/refs/tags/1.5.7.bin"
	draco_cmd_bin := exec.Command("curl", "-L", draco_bin_url, "-o", "binary.bin")
	err = draco_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	draco_src_url := "https://github.com/google/draco/archive/refs/tags/1.5.7.src.tar.gz"
	draco_cmd_src := exec.Command("curl", "-L", draco_src_url, "-o", "source.tar.gz")
	err = draco_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	draco_cmd_direct := exec.Command("./binary")
	err = draco_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
