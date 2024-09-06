package main

// Soplex - Optimization package for solving linear programming problems (LPs)
// Homepage: https://soplex.zib.de/

import (
	"fmt"
	
	"os/exec"
)

func installSoplex() {
	// Método 1: Descargar y extraer .tar.gz
	soplex_tar_url := "https://soplex.zib.de/download/release/soplex-7.1.0.0.tgz"
	soplex_cmd_tar := exec.Command("curl", "-L", soplex_tar_url, "-o", "package.tar.gz")
	err := soplex_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	soplex_zip_url := "https://soplex.zib.de/download/release/soplex-7.1.0.0.tgz"
	soplex_cmd_zip := exec.Command("curl", "-L", soplex_zip_url, "-o", "package.zip")
	err = soplex_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	soplex_bin_url := "https://soplex.zib.de/download/release/soplex-7.1.0.0.tgz"
	soplex_cmd_bin := exec.Command("curl", "-L", soplex_bin_url, "-o", "binary.bin")
	err = soplex_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	soplex_src_url := "https://soplex.zib.de/download/release/soplex-7.1.0.0.tgz"
	soplex_cmd_src := exec.Command("curl", "-L", soplex_src_url, "-o", "source.tar.gz")
	err = soplex_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	soplex_cmd_direct := exec.Command("./binary")
	err = soplex_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: gmp")
	exec.Command("latte", "install", "gmp").Run()
	fmt.Println("Instalando dependencia: tbb")
	exec.Command("latte", "install", "tbb").Run()
}
