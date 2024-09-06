package main

// Curaengine - C++ 3D printing GCode generator
// Homepage: https://github.com/Ultimaker/CuraEngine

import (
	"fmt"
	
	"os/exec"
)

func installCuraengine() {
	// Método 1: Descargar y extraer .tar.gz
	curaengine_tar_url := "https://github.com/Ultimaker/CuraEngine/archive/refs/tags/4.13.1.tar.gz"
	curaengine_cmd_tar := exec.Command("curl", "-L", curaengine_tar_url, "-o", "package.tar.gz")
	err := curaengine_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	curaengine_zip_url := "https://github.com/Ultimaker/CuraEngine/archive/refs/tags/4.13.1.zip"
	curaengine_cmd_zip := exec.Command("curl", "-L", curaengine_zip_url, "-o", "package.zip")
	err = curaengine_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	curaengine_bin_url := "https://github.com/Ultimaker/CuraEngine/archive/refs/tags/4.13.1.bin"
	curaengine_cmd_bin := exec.Command("curl", "-L", curaengine_bin_url, "-o", "binary.bin")
	err = curaengine_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	curaengine_src_url := "https://github.com/Ultimaker/CuraEngine/archive/refs/tags/4.13.1.src.tar.gz"
	curaengine_cmd_src := exec.Command("curl", "-L", curaengine_src_url, "-o", "source.tar.gz")
	err = curaengine_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	curaengine_cmd_direct := exec.Command("./binary")
	err = curaengine_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
