package main

// RobinMap - C++ implementation of a fast hash map and hash set
// Homepage: https://github.com/Tessil/robin-map

import (
	"fmt"
	
	"os/exec"
)

func installRobinMap() {
	// Método 1: Descargar y extraer .tar.gz
	robinmap_tar_url := "https://github.com/Tessil/robin-map/archive/refs/tags/v1.3.0.tar.gz"
	robinmap_cmd_tar := exec.Command("curl", "-L", robinmap_tar_url, "-o", "package.tar.gz")
	err := robinmap_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	robinmap_zip_url := "https://github.com/Tessil/robin-map/archive/refs/tags/v1.3.0.zip"
	robinmap_cmd_zip := exec.Command("curl", "-L", robinmap_zip_url, "-o", "package.zip")
	err = robinmap_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	robinmap_bin_url := "https://github.com/Tessil/robin-map/archive/refs/tags/v1.3.0.bin"
	robinmap_cmd_bin := exec.Command("curl", "-L", robinmap_bin_url, "-o", "binary.bin")
	err = robinmap_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	robinmap_src_url := "https://github.com/Tessil/robin-map/archive/refs/tags/v1.3.0.src.tar.gz"
	robinmap_cmd_src := exec.Command("curl", "-L", robinmap_src_url, "-o", "source.tar.gz")
	err = robinmap_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	robinmap_cmd_direct := exec.Command("./binary")
	err = robinmap_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
