package main

// Mlx - Array framework for Apple silicon
// Homepage: https://github.com/ml-explore/mlx

import (
	"fmt"
	
	"os/exec"
)

func installMlx() {
	// Método 1: Descargar y extraer .tar.gz
	mlx_tar_url := "https://github.com/ml-explore/mlx/archive/refs/tags/v0.13.0.tar.gz"
	mlx_cmd_tar := exec.Command("curl", "-L", mlx_tar_url, "-o", "package.tar.gz")
	err := mlx_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mlx_zip_url := "https://github.com/ml-explore/mlx/archive/refs/tags/v0.13.0.zip"
	mlx_cmd_zip := exec.Command("curl", "-L", mlx_zip_url, "-o", "package.zip")
	err = mlx_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mlx_bin_url := "https://github.com/ml-explore/mlx/archive/refs/tags/v0.13.0.bin"
	mlx_cmd_bin := exec.Command("curl", "-L", mlx_bin_url, "-o", "binary.bin")
	err = mlx_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mlx_src_url := "https://github.com/ml-explore/mlx/archive/refs/tags/v0.13.0.src.tar.gz"
	mlx_cmd_src := exec.Command("curl", "-L", mlx_src_url, "-o", "source.tar.gz")
	err = mlx_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mlx_cmd_direct := exec.Command("./binary")
	err = mlx_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: nlohmann-json")
exec.Command("latte", "install", "nlohmann-json")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
