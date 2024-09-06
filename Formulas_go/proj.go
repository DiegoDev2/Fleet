package main

// Proj - Cartographic Projections Library
// Homepage: https://proj.org/

import (
	"fmt"
	
	"os/exec"
)

func installProj() {
	// Método 1: Descargar y extraer .tar.gz
	proj_tar_url := "https://github.com/OSGeo/PROJ/releases/download/9.4.1/proj-9.4.1.tar.gz"
	proj_cmd_tar := exec.Command("curl", "-L", proj_tar_url, "-o", "package.tar.gz")
	err := proj_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	proj_zip_url := "https://github.com/OSGeo/PROJ/releases/download/9.4.1/proj-9.4.1.zip"
	proj_cmd_zip := exec.Command("curl", "-L", proj_zip_url, "-o", "package.zip")
	err = proj_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	proj_bin_url := "https://github.com/OSGeo/PROJ/releases/download/9.4.1/proj-9.4.1.bin"
	proj_cmd_bin := exec.Command("curl", "-L", proj_bin_url, "-o", "binary.bin")
	err = proj_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	proj_src_url := "https://github.com/OSGeo/PROJ/releases/download/9.4.1/proj-9.4.1.src.tar.gz"
	proj_cmd_src := exec.Command("curl", "-L", proj_src_url, "-o", "source.tar.gz")
	err = proj_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	proj_cmd_direct := exec.Command("./binary")
	err = proj_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libtiff")
exec.Command("latte", "install", "libtiff")
}
