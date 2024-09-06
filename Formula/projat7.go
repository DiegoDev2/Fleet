package main

// ProjAT7 - Cartographic Projections Library
// Homepage: https://proj.org/

import (
	"fmt"
	
	"os/exec"
)

func installProjAT7() {
	// Método 1: Descargar y extraer .tar.gz
	projat7_tar_url := "https://github.com/OSGeo/PROJ/releases/download/7.2.1/proj-7.2.1.tar.gz"
	projat7_cmd_tar := exec.Command("curl", "-L", projat7_tar_url, "-o", "package.tar.gz")
	err := projat7_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	projat7_zip_url := "https://github.com/OSGeo/PROJ/releases/download/7.2.1/proj-7.2.1.zip"
	projat7_cmd_zip := exec.Command("curl", "-L", projat7_zip_url, "-o", "package.zip")
	err = projat7_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	projat7_bin_url := "https://github.com/OSGeo/PROJ/releases/download/7.2.1/proj-7.2.1.bin"
	projat7_cmd_bin := exec.Command("curl", "-L", projat7_bin_url, "-o", "binary.bin")
	err = projat7_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	projat7_src_url := "https://github.com/OSGeo/PROJ/releases/download/7.2.1/proj-7.2.1.src.tar.gz"
	projat7_cmd_src := exec.Command("curl", "-L", projat7_src_url, "-o", "source.tar.gz")
	err = projat7_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	projat7_cmd_direct := exec.Command("./binary")
	err = projat7_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libtiff")
	exec.Command("latte", "install", "libtiff").Run()
}
