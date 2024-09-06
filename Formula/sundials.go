package main

// Sundials - Nonlinear and differential/algebraic equations solver
// Homepage: https://computing.llnl.gov/projects/sundials

import (
	"fmt"
	
	"os/exec"
)

func installSundials() {
	// Método 1: Descargar y extraer .tar.gz
	sundials_tar_url := "https://github.com/LLNL/sundials/releases/download/v7.1.1/sundials-7.1.1.tar.gz"
	sundials_cmd_tar := exec.Command("curl", "-L", sundials_tar_url, "-o", "package.tar.gz")
	err := sundials_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sundials_zip_url := "https://github.com/LLNL/sundials/releases/download/v7.1.1/sundials-7.1.1.zip"
	sundials_cmd_zip := exec.Command("curl", "-L", sundials_zip_url, "-o", "package.zip")
	err = sundials_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sundials_bin_url := "https://github.com/LLNL/sundials/releases/download/v7.1.1/sundials-7.1.1.bin"
	sundials_cmd_bin := exec.Command("curl", "-L", sundials_bin_url, "-o", "binary.bin")
	err = sundials_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sundials_src_url := "https://github.com/LLNL/sundials/releases/download/v7.1.1/sundials-7.1.1.src.tar.gz"
	sundials_cmd_src := exec.Command("curl", "-L", sundials_src_url, "-o", "source.tar.gz")
	err = sundials_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sundials_cmd_direct := exec.Command("./binary")
	err = sundials_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: gcc")
	exec.Command("latte", "install", "gcc").Run()
	fmt.Println("Instalando dependencia: open-mpi")
	exec.Command("latte", "install", "open-mpi").Run()
	fmt.Println("Instalando dependencia: openblas")
	exec.Command("latte", "install", "openblas").Run()
	fmt.Println("Instalando dependencia: suite-sparse")
	exec.Command("latte", "install", "suite-sparse").Run()
}
