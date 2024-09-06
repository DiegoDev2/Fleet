package main

// ArcadeLearningEnvironment - Platform for AI research
// Homepage: https://github.com/Farama-Foundation/Arcade-Learning-Environment

import (
	"fmt"
	
	"os/exec"
)

func installArcadeLearningEnvironment() {
	// Método 1: Descargar y extraer .tar.gz
	arcadelearningenvironment_tar_url := "https://github.com/Farama-Foundation/Arcade-Learning-Environment/archive/refs/tags/v0.9.1.tar.gz"
	arcadelearningenvironment_cmd_tar := exec.Command("curl", "-L", arcadelearningenvironment_tar_url, "-o", "package.tar.gz")
	err := arcadelearningenvironment_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	arcadelearningenvironment_zip_url := "https://github.com/Farama-Foundation/Arcade-Learning-Environment/archive/refs/tags/v0.9.1.zip"
	arcadelearningenvironment_cmd_zip := exec.Command("curl", "-L", arcadelearningenvironment_zip_url, "-o", "package.zip")
	err = arcadelearningenvironment_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	arcadelearningenvironment_bin_url := "https://github.com/Farama-Foundation/Arcade-Learning-Environment/archive/refs/tags/v0.9.1.bin"
	arcadelearningenvironment_cmd_bin := exec.Command("curl", "-L", arcadelearningenvironment_bin_url, "-o", "binary.bin")
	err = arcadelearningenvironment_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	arcadelearningenvironment_src_url := "https://github.com/Farama-Foundation/Arcade-Learning-Environment/archive/refs/tags/v0.9.1.src.tar.gz"
	arcadelearningenvironment_cmd_src := exec.Command("curl", "-L", arcadelearningenvironment_src_url, "-o", "source.tar.gz")
	err = arcadelearningenvironment_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	arcadelearningenvironment_cmd_direct := exec.Command("./binary")
	err = arcadelearningenvironment_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pybind11")
exec.Command("latte", "install", "pybind11")
	fmt.Println("Instalando dependencia: python-setuptools")
exec.Command("latte", "install", "python-setuptools")
	fmt.Println("Instalando dependencia: numpy")
exec.Command("latte", "install", "numpy")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: sdl2")
exec.Command("latte", "install", "sdl2")
}
