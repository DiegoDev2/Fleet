package main

// Papilo - Parallel Presolve for Integer and Linear Optimization
// Homepage: https://www.scipopt.org

import (
	"fmt"
	
	"os/exec"
)

func installPapilo() {
	// Método 1: Descargar y extraer .tar.gz
	papilo_tar_url := "https://github.com/scipopt/papilo/archive/refs/tags/v2.3.0.tar.gz"
	papilo_cmd_tar := exec.Command("curl", "-L", papilo_tar_url, "-o", "package.tar.gz")
	err := papilo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	papilo_zip_url := "https://github.com/scipopt/papilo/archive/refs/tags/v2.3.0.zip"
	papilo_cmd_zip := exec.Command("curl", "-L", papilo_zip_url, "-o", "package.zip")
	err = papilo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	papilo_bin_url := "https://github.com/scipopt/papilo/archive/refs/tags/v2.3.0.bin"
	papilo_cmd_bin := exec.Command("curl", "-L", papilo_bin_url, "-o", "binary.bin")
	err = papilo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	papilo_src_url := "https://github.com/scipopt/papilo/archive/refs/tags/v2.3.0.src.tar.gz"
	papilo_cmd_src := exec.Command("curl", "-L", papilo_src_url, "-o", "source.tar.gz")
	err = papilo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	papilo_cmd_direct := exec.Command("./binary")
	err = papilo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: boost")
exec.Command("latte", "install", "boost")
	fmt.Println("Instalando dependencia: gcc")
exec.Command("latte", "install", "gcc")
	fmt.Println("Instalando dependencia: gmp")
exec.Command("latte", "install", "gmp")
	fmt.Println("Instalando dependencia: openblas")
exec.Command("latte", "install", "openblas")
	fmt.Println("Instalando dependencia: tbb")
exec.Command("latte", "install", "tbb")
}
