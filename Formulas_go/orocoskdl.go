package main

// OrocosKdl - Orocos Kinematics and Dynamics C++ library
// Homepage: https://orocos.org/

import (
	"fmt"
	
	"os/exec"
)

func installOrocosKdl() {
	// Método 1: Descargar y extraer .tar.gz
	orocoskdl_tar_url := "https://github.com/orocos/orocos_kinematics_dynamics/archive/refs/tags/v1.5.1.tar.gz"
	orocoskdl_cmd_tar := exec.Command("curl", "-L", orocoskdl_tar_url, "-o", "package.tar.gz")
	err := orocoskdl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	orocoskdl_zip_url := "https://github.com/orocos/orocos_kinematics_dynamics/archive/refs/tags/v1.5.1.zip"
	orocoskdl_cmd_zip := exec.Command("curl", "-L", orocoskdl_zip_url, "-o", "package.zip")
	err = orocoskdl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	orocoskdl_bin_url := "https://github.com/orocos/orocos_kinematics_dynamics/archive/refs/tags/v1.5.1.bin"
	orocoskdl_cmd_bin := exec.Command("curl", "-L", orocoskdl_bin_url, "-o", "binary.bin")
	err = orocoskdl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	orocoskdl_src_url := "https://github.com/orocos/orocos_kinematics_dynamics/archive/refs/tags/v1.5.1.src.tar.gz"
	orocoskdl_cmd_src := exec.Command("curl", "-L", orocoskdl_src_url, "-o", "source.tar.gz")
	err = orocoskdl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	orocoskdl_cmd_direct := exec.Command("./binary")
	err = orocoskdl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: eigen")
exec.Command("latte", "install", "eigen")
}
