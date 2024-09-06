package main

// Opencolorio - Color management solution geared towards motion picture production
// Homepage: https://opencolorio.org/

import (
	"fmt"
	
	"os/exec"
)

func installOpencolorio() {
	// Método 1: Descargar y extraer .tar.gz
	opencolorio_tar_url := "https://github.com/AcademySoftwareFoundation/OpenColorIO/archive/refs/tags/v2.3.2.tar.gz"
	opencolorio_cmd_tar := exec.Command("curl", "-L", opencolorio_tar_url, "-o", "package.tar.gz")
	err := opencolorio_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	opencolorio_zip_url := "https://github.com/AcademySoftwareFoundation/OpenColorIO/archive/refs/tags/v2.3.2.zip"
	opencolorio_cmd_zip := exec.Command("curl", "-L", opencolorio_zip_url, "-o", "package.zip")
	err = opencolorio_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	opencolorio_bin_url := "https://github.com/AcademySoftwareFoundation/OpenColorIO/archive/refs/tags/v2.3.2.bin"
	opencolorio_cmd_bin := exec.Command("curl", "-L", opencolorio_bin_url, "-o", "binary.bin")
	err = opencolorio_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	opencolorio_src_url := "https://github.com/AcademySoftwareFoundation/OpenColorIO/archive/refs/tags/v2.3.2.src.tar.gz"
	opencolorio_cmd_src := exec.Command("curl", "-L", opencolorio_src_url, "-o", "source.tar.gz")
	err = opencolorio_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	opencolorio_cmd_direct := exec.Command("./binary")
	err = opencolorio_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pybind11")
exec.Command("latte", "install", "pybind11")
	fmt.Println("Instalando dependencia: imath")
exec.Command("latte", "install", "imath")
	fmt.Println("Instalando dependencia: little-cms2")
exec.Command("latte", "install", "little-cms2")
	fmt.Println("Instalando dependencia: minizip-ng")
exec.Command("latte", "install", "minizip-ng")
	fmt.Println("Instalando dependencia: openexr")
exec.Command("latte", "install", "openexr")
	fmt.Println("Instalando dependencia: pystring")
exec.Command("latte", "install", "pystring")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: yaml-cpp")
exec.Command("latte", "install", "yaml-cpp")
}
