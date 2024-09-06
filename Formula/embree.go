package main

// Embree - High-performance ray tracing kernels
// Homepage: https://www.embree.org/

import (
	"fmt"
	
	"os/exec"
)

func installEmbree() {
	// Método 1: Descargar y extraer .tar.gz
	embree_tar_url := "https://github.com/RenderKit/embree/archive/refs/tags/v4.3.3.tar.gz"
	embree_cmd_tar := exec.Command("curl", "-L", embree_tar_url, "-o", "package.tar.gz")
	err := embree_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	embree_zip_url := "https://github.com/RenderKit/embree/archive/refs/tags/v4.3.3.zip"
	embree_cmd_zip := exec.Command("curl", "-L", embree_zip_url, "-o", "package.zip")
	err = embree_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	embree_bin_url := "https://github.com/RenderKit/embree/archive/refs/tags/v4.3.3.bin"
	embree_cmd_bin := exec.Command("curl", "-L", embree_bin_url, "-o", "binary.bin")
	err = embree_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	embree_src_url := "https://github.com/RenderKit/embree/archive/refs/tags/v4.3.3.src.tar.gz"
	embree_cmd_src := exec.Command("curl", "-L", embree_src_url, "-o", "source.tar.gz")
	err = embree_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	embree_cmd_direct := exec.Command("./binary")
	err = embree_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: ispc")
	exec.Command("latte", "install", "ispc").Run()
	fmt.Println("Instalando dependencia: tbb")
	exec.Command("latte", "install", "tbb").Run()
}
