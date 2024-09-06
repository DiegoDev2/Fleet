package main

// LibvisualProjectm - Visualization plug-in for projectM support from Libvisual
// Homepage: https://github.com/projectM-visualizer/frontend-libvisual-plug-in

import (
	"fmt"
	
	"os/exec"
)

func installLibvisualProjectm() {
	// Método 1: Descargar y extraer .tar.gz
	libvisualprojectm_tar_url := "https://github.com/projectM-visualizer/frontend-libvisual-plug-in/archive/refs/tags/v2.1.1.tar.gz"
	libvisualprojectm_cmd_tar := exec.Command("curl", "-L", libvisualprojectm_tar_url, "-o", "package.tar.gz")
	err := libvisualprojectm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libvisualprojectm_zip_url := "https://github.com/projectM-visualizer/frontend-libvisual-plug-in/archive/refs/tags/v2.1.1.zip"
	libvisualprojectm_cmd_zip := exec.Command("curl", "-L", libvisualprojectm_zip_url, "-o", "package.zip")
	err = libvisualprojectm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libvisualprojectm_bin_url := "https://github.com/projectM-visualizer/frontend-libvisual-plug-in/archive/refs/tags/v2.1.1.bin"
	libvisualprojectm_cmd_bin := exec.Command("curl", "-L", libvisualprojectm_bin_url, "-o", "binary.bin")
	err = libvisualprojectm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libvisualprojectm_src_url := "https://github.com/projectM-visualizer/frontend-libvisual-plug-in/archive/refs/tags/v2.1.1.src.tar.gz"
	libvisualprojectm_cmd_src := exec.Command("curl", "-L", libvisualprojectm_src_url, "-o", "source.tar.gz")
	err = libvisualprojectm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libvisualprojectm_cmd_direct := exec.Command("./binary")
	err = libvisualprojectm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libvisual-plugins")
exec.Command("latte", "install", "libvisual-plugins")
	fmt.Println("Instalando dependencia: xorg-server")
exec.Command("latte", "install", "xorg-server")
	fmt.Println("Instalando dependencia: libvisual")
exec.Command("latte", "install", "libvisual")
	fmt.Println("Instalando dependencia: projectm")
exec.Command("latte", "install", "projectm")
}
