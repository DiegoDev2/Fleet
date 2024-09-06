package main

// Clblast - Tuned OpenCL BLAS library
// Homepage: https://github.com/CNugteren/CLBlast

import (
	"fmt"
	
	"os/exec"
)

func installClblast() {
	// Método 1: Descargar y extraer .tar.gz
	clblast_tar_url := "https://github.com/CNugteren/CLBlast/archive/refs/tags/1.6.3.tar.gz"
	clblast_cmd_tar := exec.Command("curl", "-L", clblast_tar_url, "-o", "package.tar.gz")
	err := clblast_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	clblast_zip_url := "https://github.com/CNugteren/CLBlast/archive/refs/tags/1.6.3.zip"
	clblast_cmd_zip := exec.Command("curl", "-L", clblast_zip_url, "-o", "package.zip")
	err = clblast_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	clblast_bin_url := "https://github.com/CNugteren/CLBlast/archive/refs/tags/1.6.3.bin"
	clblast_cmd_bin := exec.Command("curl", "-L", clblast_bin_url, "-o", "binary.bin")
	err = clblast_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	clblast_src_url := "https://github.com/CNugteren/CLBlast/archive/refs/tags/1.6.3.src.tar.gz"
	clblast_cmd_src := exec.Command("curl", "-L", clblast_src_url, "-o", "source.tar.gz")
	err = clblast_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	clblast_cmd_direct := exec.Command("./binary")
	err = clblast_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: opencl-headers")
exec.Command("latte", "install", "opencl-headers")
	fmt.Println("Instalando dependencia: opencl-icd-loader")
exec.Command("latte", "install", "opencl-icd-loader")
	fmt.Println("Instalando dependencia: pocl")
exec.Command("latte", "install", "pocl")
}
