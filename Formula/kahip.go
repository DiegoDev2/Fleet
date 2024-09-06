package main

// Kahip - Karlsruhe High Quality Partitioning
// Homepage: https://algo2.iti.kit.edu/documents/kahip/index.html

import (
	"fmt"
	
	"os/exec"
)

func installKahip() {
	// Método 1: Descargar y extraer .tar.gz
	kahip_tar_url := "https://github.com/KaHIP/KaHIP/archive/refs/tags/v3.16.tar.gz"
	kahip_cmd_tar := exec.Command("curl", "-L", kahip_tar_url, "-o", "package.tar.gz")
	err := kahip_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kahip_zip_url := "https://github.com/KaHIP/KaHIP/archive/refs/tags/v3.16.zip"
	kahip_cmd_zip := exec.Command("curl", "-L", kahip_zip_url, "-o", "package.zip")
	err = kahip_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kahip_bin_url := "https://github.com/KaHIP/KaHIP/archive/refs/tags/v3.16.bin"
	kahip_cmd_bin := exec.Command("curl", "-L", kahip_bin_url, "-o", "binary.bin")
	err = kahip_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kahip_src_url := "https://github.com/KaHIP/KaHIP/archive/refs/tags/v3.16.src.tar.gz"
	kahip_cmd_src := exec.Command("curl", "-L", kahip_src_url, "-o", "source.tar.gz")
	err = kahip_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kahip_cmd_direct := exec.Command("./binary")
	err = kahip_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: open-mpi")
	exec.Command("latte", "install", "open-mpi").Run()
	fmt.Println("Instalando dependencia: gcc")
	exec.Command("latte", "install", "gcc").Run()
}
