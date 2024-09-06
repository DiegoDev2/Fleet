package main

// Libblastrampoline - Using PLT trampolines to provide a BLAS and LAPACK demuxing library
// Homepage: https://github.com/JuliaLinearAlgebra/libblastrampoline

import (
	"fmt"
	
	"os/exec"
)

func installLibblastrampoline() {
	// Método 1: Descargar y extraer .tar.gz
	libblastrampoline_tar_url := "https://github.com/JuliaLinearAlgebra/libblastrampoline/archive/refs/tags/v5.11.0.tar.gz"
	libblastrampoline_cmd_tar := exec.Command("curl", "-L", libblastrampoline_tar_url, "-o", "package.tar.gz")
	err := libblastrampoline_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libblastrampoline_zip_url := "https://github.com/JuliaLinearAlgebra/libblastrampoline/archive/refs/tags/v5.11.0.zip"
	libblastrampoline_cmd_zip := exec.Command("curl", "-L", libblastrampoline_zip_url, "-o", "package.zip")
	err = libblastrampoline_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libblastrampoline_bin_url := "https://github.com/JuliaLinearAlgebra/libblastrampoline/archive/refs/tags/v5.11.0.bin"
	libblastrampoline_cmd_bin := exec.Command("curl", "-L", libblastrampoline_bin_url, "-o", "binary.bin")
	err = libblastrampoline_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libblastrampoline_src_url := "https://github.com/JuliaLinearAlgebra/libblastrampoline/archive/refs/tags/v5.11.0.src.tar.gz"
	libblastrampoline_cmd_src := exec.Command("curl", "-L", libblastrampoline_src_url, "-o", "source.tar.gz")
	err = libblastrampoline_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libblastrampoline_cmd_direct := exec.Command("./binary")
	err = libblastrampoline_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openblas")
	exec.Command("latte", "install", "openblas").Run()
}
