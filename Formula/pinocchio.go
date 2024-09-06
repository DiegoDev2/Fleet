package main

// Pinocchio - Efficient and fast C++ library implementing Rigid Body Dynamics algorithms
// Homepage: https://stack-of-tasks.github.io/pinocchio

import (
	"fmt"
	
	"os/exec"
)

func installPinocchio() {
	// Método 1: Descargar y extraer .tar.gz
	pinocchio_tar_url := "https://github.com/stack-of-tasks/pinocchio/releases/download/v3.2.0/pinocchio-3.2.0.tar.gz"
	pinocchio_cmd_tar := exec.Command("curl", "-L", pinocchio_tar_url, "-o", "package.tar.gz")
	err := pinocchio_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pinocchio_zip_url := "https://github.com/stack-of-tasks/pinocchio/releases/download/v3.2.0/pinocchio-3.2.0.zip"
	pinocchio_cmd_zip := exec.Command("curl", "-L", pinocchio_zip_url, "-o", "package.zip")
	err = pinocchio_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pinocchio_bin_url := "https://github.com/stack-of-tasks/pinocchio/releases/download/v3.2.0/pinocchio-3.2.0.bin"
	pinocchio_cmd_bin := exec.Command("curl", "-L", pinocchio_bin_url, "-o", "binary.bin")
	err = pinocchio_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pinocchio_src_url := "https://github.com/stack-of-tasks/pinocchio/releases/download/v3.2.0/pinocchio-3.2.0.src.tar.gz"
	pinocchio_cmd_src := exec.Command("curl", "-L", pinocchio_src_url, "-o", "source.tar.gz")
	err = pinocchio_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pinocchio_cmd_direct := exec.Command("./binary")
	err = pinocchio_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: doxygen")
	exec.Command("latte", "install", "doxygen").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: boost-python3")
	exec.Command("latte", "install", "boost-python3").Run()
	fmt.Println("Instalando dependencia: console_bridge")
	exec.Command("latte", "install", "console_bridge").Run()
	fmt.Println("Instalando dependencia: eigen")
	exec.Command("latte", "install", "eigen").Run()
	fmt.Println("Instalando dependencia: eigenpy")
	exec.Command("latte", "install", "eigenpy").Run()
	fmt.Println("Instalando dependencia: hpp-fcl")
	exec.Command("latte", "install", "hpp-fcl").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: urdfdom")
	exec.Command("latte", "install", "urdfdom").Run()
	fmt.Println("Instalando dependencia: octomap")
	exec.Command("latte", "install", "octomap").Run()
}
