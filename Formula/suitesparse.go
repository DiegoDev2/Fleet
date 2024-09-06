package main

// SuiteSparse - Suite of Sparse Matrix Software
// Homepage: https://people.engr.tamu.edu/davis/suitesparse.html

import (
	"fmt"
	
	"os/exec"
)

func installSuiteSparse() {
	// Método 1: Descargar y extraer .tar.gz
	suitesparse_tar_url := "https://github.com/DrTimothyAldenDavis/SuiteSparse/archive/refs/tags/v7.8.2.tar.gz"
	suitesparse_cmd_tar := exec.Command("curl", "-L", suitesparse_tar_url, "-o", "package.tar.gz")
	err := suitesparse_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	suitesparse_zip_url := "https://github.com/DrTimothyAldenDavis/SuiteSparse/archive/refs/tags/v7.8.2.zip"
	suitesparse_cmd_zip := exec.Command("curl", "-L", suitesparse_zip_url, "-o", "package.zip")
	err = suitesparse_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	suitesparse_bin_url := "https://github.com/DrTimothyAldenDavis/SuiteSparse/archive/refs/tags/v7.8.2.bin"
	suitesparse_cmd_bin := exec.Command("curl", "-L", suitesparse_bin_url, "-o", "binary.bin")
	err = suitesparse_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	suitesparse_src_url := "https://github.com/DrTimothyAldenDavis/SuiteSparse/archive/refs/tags/v7.8.2.src.tar.gz"
	suitesparse_cmd_src := exec.Command("curl", "-L", suitesparse_src_url, "-o", "source.tar.gz")
	err = suitesparse_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	suitesparse_cmd_direct := exec.Command("./binary")
	err = suitesparse_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: gcc")
	exec.Command("latte", "install", "gcc").Run()
	fmt.Println("Instalando dependencia: gmp")
	exec.Command("latte", "install", "gmp").Run()
	fmt.Println("Instalando dependencia: metis")
	exec.Command("latte", "install", "metis").Run()
	fmt.Println("Instalando dependencia: mpfr")
	exec.Command("latte", "install", "mpfr").Run()
	fmt.Println("Instalando dependencia: libomp")
	exec.Command("latte", "install", "libomp").Run()
	fmt.Println("Instalando dependencia: openblas")
	exec.Command("latte", "install", "openblas").Run()
}
