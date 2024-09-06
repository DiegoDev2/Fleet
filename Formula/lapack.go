package main

// Lapack - Linear Algebra PACKage
// Homepage: https://www.netlib.org/lapack/

import (
	"fmt"
	
	"os/exec"
)

func installLapack() {
	// Método 1: Descargar y extraer .tar.gz
	lapack_tar_url := "https://github.com/Reference-LAPACK/lapack/archive/refs/tags/v3.12.0.tar.gz"
	lapack_cmd_tar := exec.Command("curl", "-L", lapack_tar_url, "-o", "package.tar.gz")
	err := lapack_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lapack_zip_url := "https://github.com/Reference-LAPACK/lapack/archive/refs/tags/v3.12.0.zip"
	lapack_cmd_zip := exec.Command("curl", "-L", lapack_zip_url, "-o", "package.zip")
	err = lapack_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lapack_bin_url := "https://github.com/Reference-LAPACK/lapack/archive/refs/tags/v3.12.0.bin"
	lapack_cmd_bin := exec.Command("curl", "-L", lapack_bin_url, "-o", "binary.bin")
	err = lapack_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lapack_src_url := "https://github.com/Reference-LAPACK/lapack/archive/refs/tags/v3.12.0.src.tar.gz"
	lapack_cmd_src := exec.Command("curl", "-L", lapack_src_url, "-o", "source.tar.gz")
	err = lapack_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lapack_cmd_direct := exec.Command("./binary")
	err = lapack_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: gcc")
	exec.Command("latte", "install", "gcc").Run()
}
