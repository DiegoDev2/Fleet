package main

// Cf4ocl - C Framework for OpenCL
// Homepage: https://nunofachada.github.io/cf4ocl/

import (
	"fmt"
	
	"os/exec"
)

func installCf4ocl() {
	// Método 1: Descargar y extraer .tar.gz
	cf4ocl_tar_url := "https://github.com/nunofachada/cf4ocl/archive/refs/tags/v2.1.0.tar.gz"
	cf4ocl_cmd_tar := exec.Command("curl", "-L", cf4ocl_tar_url, "-o", "package.tar.gz")
	err := cf4ocl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cf4ocl_zip_url := "https://github.com/nunofachada/cf4ocl/archive/refs/tags/v2.1.0.zip"
	cf4ocl_cmd_zip := exec.Command("curl", "-L", cf4ocl_zip_url, "-o", "package.zip")
	err = cf4ocl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cf4ocl_bin_url := "https://github.com/nunofachada/cf4ocl/archive/refs/tags/v2.1.0.bin"
	cf4ocl_cmd_bin := exec.Command("curl", "-L", cf4ocl_bin_url, "-o", "binary.bin")
	err = cf4ocl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cf4ocl_src_url := "https://github.com/nunofachada/cf4ocl/archive/refs/tags/v2.1.0.src.tar.gz"
	cf4ocl_cmd_src := exec.Command("curl", "-L", cf4ocl_src_url, "-o", "source.tar.gz")
	err = cf4ocl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cf4ocl_cmd_direct := exec.Command("./binary")
	err = cf4ocl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: opencl-headers")
exec.Command("latte", "install", "opencl-headers")
	fmt.Println("Instalando dependencia: opencl-icd-loader")
exec.Command("latte", "install", "opencl-icd-loader")
	fmt.Println("Instalando dependencia: pocl")
exec.Command("latte", "install", "pocl")
}
