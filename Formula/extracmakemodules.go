package main

// ExtraCmakeModules - Extra modules and scripts for CMake
// Homepage: https://api.kde.org/frameworks/extra-cmake-modules/html/index.html

import (
	"fmt"
	
	"os/exec"
)

func installExtraCmakeModules() {
	// Método 1: Descargar y extraer .tar.gz
	extracmakemodules_tar_url := "https://download.kde.org/stable/frameworks/6.5/extra-cmake-modules-6.5.0.tar.xz"
	extracmakemodules_cmd_tar := exec.Command("curl", "-L", extracmakemodules_tar_url, "-o", "package.tar.gz")
	err := extracmakemodules_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	extracmakemodules_zip_url := "https://download.kde.org/stable/frameworks/6.5/extra-cmake-modules-6.5.0.tar.xz"
	extracmakemodules_cmd_zip := exec.Command("curl", "-L", extracmakemodules_zip_url, "-o", "package.zip")
	err = extracmakemodules_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	extracmakemodules_bin_url := "https://download.kde.org/stable/frameworks/6.5/extra-cmake-modules-6.5.0.tar.xz"
	extracmakemodules_cmd_bin := exec.Command("curl", "-L", extracmakemodules_bin_url, "-o", "binary.bin")
	err = extracmakemodules_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	extracmakemodules_src_url := "https://download.kde.org/stable/frameworks/6.5/extra-cmake-modules-6.5.0.tar.xz"
	extracmakemodules_cmd_src := exec.Command("curl", "-L", extracmakemodules_src_url, "-o", "source.tar.gz")
	err = extracmakemodules_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	extracmakemodules_cmd_direct := exec.Command("./binary")
	err = extracmakemodules_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: sphinx-doc")
	exec.Command("latte", "install", "sphinx-doc").Run()
}
