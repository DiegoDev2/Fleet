package main

// Superlu - Solve large, sparse nonsymmetric systems of equations
// Homepage: https://portal.nersc.gov/project/sparse/superlu/

import (
	"fmt"
	
	"os/exec"
)

func installSuperlu() {
	// Método 1: Descargar y extraer .tar.gz
	superlu_tar_url := "https://github.com/xiaoyeli/superlu/archive/refs/tags/v7.0.0.tar.gz"
	superlu_cmd_tar := exec.Command("curl", "-L", superlu_tar_url, "-o", "package.tar.gz")
	err := superlu_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	superlu_zip_url := "https://github.com/xiaoyeli/superlu/archive/refs/tags/v7.0.0.zip"
	superlu_cmd_zip := exec.Command("curl", "-L", superlu_zip_url, "-o", "package.zip")
	err = superlu_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	superlu_bin_url := "https://github.com/xiaoyeli/superlu/archive/refs/tags/v7.0.0.bin"
	superlu_cmd_bin := exec.Command("curl", "-L", superlu_bin_url, "-o", "binary.bin")
	err = superlu_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	superlu_src_url := "https://github.com/xiaoyeli/superlu/archive/refs/tags/v7.0.0.src.tar.gz"
	superlu_cmd_src := exec.Command("curl", "-L", superlu_src_url, "-o", "source.tar.gz")
	err = superlu_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	superlu_cmd_direct := exec.Command("./binary")
	err = superlu_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: gcc")
	exec.Command("latte", "install", "gcc").Run()
	fmt.Println("Instalando dependencia: openblas")
	exec.Command("latte", "install", "openblas").Run()
}
