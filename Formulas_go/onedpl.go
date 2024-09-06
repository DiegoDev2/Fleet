package main

// Onedpl - C++ standard library algorithms with support for execution policies
// Homepage: https://software.intel.com/content/www/us/en/develop/tools/oneapi/components/dpc-library.html

import (
	"fmt"
	
	"os/exec"
)

func installOnedpl() {
	// Método 1: Descargar y extraer .tar.gz
	onedpl_tar_url := "https://github.com/oneapi-src/oneDPL/archive/refs/tags/oneDPL-2022.0.0-release.tar.gz"
	onedpl_cmd_tar := exec.Command("curl", "-L", onedpl_tar_url, "-o", "package.tar.gz")
	err := onedpl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	onedpl_zip_url := "https://github.com/oneapi-src/oneDPL/archive/refs/tags/oneDPL-2022.0.0-release.zip"
	onedpl_cmd_zip := exec.Command("curl", "-L", onedpl_zip_url, "-o", "package.zip")
	err = onedpl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	onedpl_bin_url := "https://github.com/oneapi-src/oneDPL/archive/refs/tags/oneDPL-2022.0.0-release.bin"
	onedpl_cmd_bin := exec.Command("curl", "-L", onedpl_bin_url, "-o", "binary.bin")
	err = onedpl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	onedpl_src_url := "https://github.com/oneapi-src/oneDPL/archive/refs/tags/oneDPL-2022.0.0-release.src.tar.gz"
	onedpl_cmd_src := exec.Command("curl", "-L", onedpl_src_url, "-o", "source.tar.gz")
	err = onedpl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	onedpl_cmd_direct := exec.Command("./binary")
	err = onedpl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: tbb")
exec.Command("latte", "install", "tbb")
}
