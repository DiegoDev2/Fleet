package main

// Bandicoot - C++ library for GPU accelerated linear algebra
// Homepage: https://coot.sourceforge.io/

import (
	"fmt"
	
	"os/exec"
)

func installBandicoot() {
	// Método 1: Descargar y extraer .tar.gz
	bandicoot_tar_url := "https://gitlab.com/bandicoot-lib/bandicoot-code/-/archive/1.15.0/bandicoot-code-1.15.0.tar.bz2"
	bandicoot_cmd_tar := exec.Command("curl", "-L", bandicoot_tar_url, "-o", "package.tar.gz")
	err := bandicoot_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bandicoot_zip_url := "https://gitlab.com/bandicoot-lib/bandicoot-code/-/archive/1.15.0/bandicoot-code-1.15.0.tar.bz2"
	bandicoot_cmd_zip := exec.Command("curl", "-L", bandicoot_zip_url, "-o", "package.zip")
	err = bandicoot_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bandicoot_bin_url := "https://gitlab.com/bandicoot-lib/bandicoot-code/-/archive/1.15.0/bandicoot-code-1.15.0.tar.bz2"
	bandicoot_cmd_bin := exec.Command("curl", "-L", bandicoot_bin_url, "-o", "binary.bin")
	err = bandicoot_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bandicoot_src_url := "https://gitlab.com/bandicoot-lib/bandicoot-code/-/archive/1.15.0/bandicoot-code-1.15.0.tar.bz2"
	bandicoot_cmd_src := exec.Command("curl", "-L", bandicoot_src_url, "-o", "source.tar.gz")
	err = bandicoot_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bandicoot_cmd_direct := exec.Command("./binary")
	err = bandicoot_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: clblas")
exec.Command("latte", "install", "clblas")
	fmt.Println("Instalando dependencia: openblas")
exec.Command("latte", "install", "openblas")
	fmt.Println("Instalando dependencia: opencl-headers")
exec.Command("latte", "install", "opencl-headers")
	fmt.Println("Instalando dependencia: opencl-icd-loader")
exec.Command("latte", "install", "opencl-icd-loader")
	fmt.Println("Instalando dependencia: pocl")
exec.Command("latte", "install", "pocl")
}
