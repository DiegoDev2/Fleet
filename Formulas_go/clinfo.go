package main

// Clinfo - Print information about OpenCL platforms and devices
// Homepage: https://github.com/Oblomov/clinfo

import (
	"fmt"
	
	"os/exec"
)

func installClinfo() {
	// Método 1: Descargar y extraer .tar.gz
	clinfo_tar_url := "https://github.com/Oblomov/clinfo/archive/refs/tags/3.0.23.01.25.tar.gz"
	clinfo_cmd_tar := exec.Command("curl", "-L", clinfo_tar_url, "-o", "package.tar.gz")
	err := clinfo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	clinfo_zip_url := "https://github.com/Oblomov/clinfo/archive/refs/tags/3.0.23.01.25.zip"
	clinfo_cmd_zip := exec.Command("curl", "-L", clinfo_zip_url, "-o", "package.zip")
	err = clinfo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	clinfo_bin_url := "https://github.com/Oblomov/clinfo/archive/refs/tags/3.0.23.01.25.bin"
	clinfo_cmd_bin := exec.Command("curl", "-L", clinfo_bin_url, "-o", "binary.bin")
	err = clinfo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	clinfo_src_url := "https://github.com/Oblomov/clinfo/archive/refs/tags/3.0.23.01.25.src.tar.gz"
	clinfo_cmd_src := exec.Command("curl", "-L", clinfo_src_url, "-o", "source.tar.gz")
	err = clinfo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	clinfo_cmd_direct := exec.Command("./binary")
	err = clinfo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: opencl-headers")
exec.Command("latte", "install", "opencl-headers")
	fmt.Println("Instalando dependencia: opencl-icd-loader")
exec.Command("latte", "install", "opencl-icd-loader")
	fmt.Println("Instalando dependencia: pocl")
exec.Command("latte", "install", "pocl")
}
