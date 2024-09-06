package main

// Scip - Solver for mixed integer programming and mixed integer nonlinear programming
// Homepage: https://scipopt.org

import (
	"fmt"
	
	"os/exec"
)

func installScip() {
	// Método 1: Descargar y extraer .tar.gz
	scip_tar_url := "https://scipopt.org/download/release/scip-9.1.0.tgz"
	scip_cmd_tar := exec.Command("curl", "-L", scip_tar_url, "-o", "package.tar.gz")
	err := scip_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	scip_zip_url := "https://scipopt.org/download/release/scip-9.1.0.tgz"
	scip_cmd_zip := exec.Command("curl", "-L", scip_zip_url, "-o", "package.zip")
	err = scip_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	scip_bin_url := "https://scipopt.org/download/release/scip-9.1.0.tgz"
	scip_cmd_bin := exec.Command("curl", "-L", scip_bin_url, "-o", "binary.bin")
	err = scip_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	scip_src_url := "https://scipopt.org/download/release/scip-9.1.0.tgz"
	scip_cmd_src := exec.Command("curl", "-L", scip_src_url, "-o", "source.tar.gz")
	err = scip_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	scip_cmd_direct := exec.Command("./binary")
	err = scip_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: cppad")
	exec.Command("latte", "install", "cppad").Run()
	fmt.Println("Instalando dependencia: gmp")
	exec.Command("latte", "install", "gmp").Run()
	fmt.Println("Instalando dependencia: ipopt")
	exec.Command("latte", "install", "ipopt").Run()
	fmt.Println("Instalando dependencia: openblas")
	exec.Command("latte", "install", "openblas").Run()
	fmt.Println("Instalando dependencia: papilo")
	exec.Command("latte", "install", "papilo").Run()
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
	fmt.Println("Instalando dependencia: soplex")
	exec.Command("latte", "install", "soplex").Run()
	fmt.Println("Instalando dependencia: tbb")
	exec.Command("latte", "install", "tbb").Run()
	fmt.Println("Instalando dependencia: gcc")
	exec.Command("latte", "install", "gcc").Run()
}
