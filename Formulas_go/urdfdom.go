package main

// Urdfdom - Unified Robot Description Format (URDF) parser
// Homepage: https://wiki.ros.org/urdf/

import (
	"fmt"
	
	"os/exec"
)

func installUrdfdom() {
	// Método 1: Descargar y extraer .tar.gz
	urdfdom_tar_url := "https://github.com/ros/urdfdom/archive/refs/tags/4.0.0.tar.gz"
	urdfdom_cmd_tar := exec.Command("curl", "-L", urdfdom_tar_url, "-o", "package.tar.gz")
	err := urdfdom_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	urdfdom_zip_url := "https://github.com/ros/urdfdom/archive/refs/tags/4.0.0.zip"
	urdfdom_cmd_zip := exec.Command("curl", "-L", urdfdom_zip_url, "-o", "package.zip")
	err = urdfdom_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	urdfdom_bin_url := "https://github.com/ros/urdfdom/archive/refs/tags/4.0.0.bin"
	urdfdom_cmd_bin := exec.Command("curl", "-L", urdfdom_bin_url, "-o", "binary.bin")
	err = urdfdom_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	urdfdom_src_url := "https://github.com/ros/urdfdom/archive/refs/tags/4.0.0.src.tar.gz"
	urdfdom_cmd_src := exec.Command("curl", "-L", urdfdom_src_url, "-o", "source.tar.gz")
	err = urdfdom_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	urdfdom_cmd_direct := exec.Command("./binary")
	err = urdfdom_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: console_bridge")
exec.Command("latte", "install", "console_bridge")
	fmt.Println("Instalando dependencia: tinyxml2")
exec.Command("latte", "install", "tinyxml2")
	fmt.Println("Instalando dependencia: urdfdom_headers")
exec.Command("latte", "install", "urdfdom_headers")
}
