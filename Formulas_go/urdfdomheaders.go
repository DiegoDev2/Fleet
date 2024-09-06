package main

// UrdfdomHeaders - Headers for Unified Robot Description Format (URDF) parsers
// Homepage: https://wiki.ros.org/urdfdom_headers/

import (
	"fmt"
	
	"os/exec"
)

func installUrdfdomHeaders() {
	// Método 1: Descargar y extraer .tar.gz
	urdfdomheaders_tar_url := "https://github.com/ros/urdfdom_headers/archive/refs/tags/1.1.1.tar.gz"
	urdfdomheaders_cmd_tar := exec.Command("curl", "-L", urdfdomheaders_tar_url, "-o", "package.tar.gz")
	err := urdfdomheaders_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	urdfdomheaders_zip_url := "https://github.com/ros/urdfdom_headers/archive/refs/tags/1.1.1.zip"
	urdfdomheaders_cmd_zip := exec.Command("curl", "-L", urdfdomheaders_zip_url, "-o", "package.zip")
	err = urdfdomheaders_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	urdfdomheaders_bin_url := "https://github.com/ros/urdfdom_headers/archive/refs/tags/1.1.1.bin"
	urdfdomheaders_cmd_bin := exec.Command("curl", "-L", urdfdomheaders_bin_url, "-o", "binary.bin")
	err = urdfdomheaders_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	urdfdomheaders_src_url := "https://github.com/ros/urdfdom_headers/archive/refs/tags/1.1.1.src.tar.gz"
	urdfdomheaders_cmd_src := exec.Command("curl", "-L", urdfdomheaders_src_url, "-o", "source.tar.gz")
	err = urdfdomheaders_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	urdfdomheaders_cmd_direct := exec.Command("./binary")
	err = urdfdomheaders_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
}
