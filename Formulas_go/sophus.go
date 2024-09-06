package main

// Sophus - C++ implementation of Lie Groups using Eigen
// Homepage: https://github.com/strasdat/Sophus

import (
	"fmt"
	
	"os/exec"
)

func installSophus() {
	// Método 1: Descargar y extraer .tar.gz
	sophus_tar_url := "https://github.com/strasdat/Sophus/archive/refs/tags/1.24.6.tar.gz"
	sophus_cmd_tar := exec.Command("curl", "-L", sophus_tar_url, "-o", "package.tar.gz")
	err := sophus_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sophus_zip_url := "https://github.com/strasdat/Sophus/archive/refs/tags/1.24.6.zip"
	sophus_cmd_zip := exec.Command("curl", "-L", sophus_zip_url, "-o", "package.zip")
	err = sophus_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sophus_bin_url := "https://github.com/strasdat/Sophus/archive/refs/tags/1.24.6.bin"
	sophus_cmd_bin := exec.Command("curl", "-L", sophus_bin_url, "-o", "binary.bin")
	err = sophus_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sophus_src_url := "https://github.com/strasdat/Sophus/archive/refs/tags/1.24.6.src.tar.gz"
	sophus_cmd_src := exec.Command("curl", "-L", sophus_src_url, "-o", "source.tar.gz")
	err = sophus_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sophus_cmd_direct := exec.Command("./binary")
	err = sophus_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: ceres-solver")
exec.Command("latte", "install", "ceres-solver")
	fmt.Println("Instalando dependencia: eigen")
exec.Command("latte", "install", "eigen")
	fmt.Println("Instalando dependencia: fmt")
exec.Command("latte", "install", "fmt")
}
