package main

// LeelaZero - Neural Network Go engine with no human-provided knowledge
// Homepage: https://zero.sjeng.org/

import (
	"fmt"
	
	"os/exec"
)

func installLeelaZero() {
	// Método 1: Descargar y extraer .tar.gz
	leelazero_tar_url := "https://github.com/leela-zero/leela-zero.git"
	leelazero_cmd_tar := exec.Command("curl", "-L", leelazero_tar_url, "-o", "package.tar.gz")
	err := leelazero_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	leelazero_zip_url := "https://github.com/leela-zero/leela-zero.git"
	leelazero_cmd_zip := exec.Command("curl", "-L", leelazero_zip_url, "-o", "package.zip")
	err = leelazero_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	leelazero_bin_url := "https://github.com/leela-zero/leela-zero.git"
	leelazero_cmd_bin := exec.Command("curl", "-L", leelazero_bin_url, "-o", "binary.bin")
	err = leelazero_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	leelazero_src_url := "https://github.com/leela-zero/leela-zero.git"
	leelazero_cmd_src := exec.Command("curl", "-L", leelazero_src_url, "-o", "source.tar.gz")
	err = leelazero_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	leelazero_cmd_direct := exec.Command("./binary")
	err = leelazero_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: boost")
exec.Command("latte", "install", "boost")
	fmt.Println("Instalando dependencia: opencl-headers")
exec.Command("latte", "install", "opencl-headers")
	fmt.Println("Instalando dependencia: opencl-icd-loader")
exec.Command("latte", "install", "opencl-icd-loader")
	fmt.Println("Instalando dependencia: pocl")
exec.Command("latte", "install", "pocl")
}
