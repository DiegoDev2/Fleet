package main

// Primecount - Fast prime counting function program and C/C++ library
// Homepage: https://github.com/kimwalisch/primecount

import (
	"fmt"
	
	"os/exec"
)

func installPrimecount() {
	// Método 1: Descargar y extraer .tar.gz
	primecount_tar_url := "https://github.com/kimwalisch/primecount/archive/refs/tags/v7.14.tar.gz"
	primecount_cmd_tar := exec.Command("curl", "-L", primecount_tar_url, "-o", "package.tar.gz")
	err := primecount_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	primecount_zip_url := "https://github.com/kimwalisch/primecount/archive/refs/tags/v7.14.zip"
	primecount_cmd_zip := exec.Command("curl", "-L", primecount_zip_url, "-o", "package.zip")
	err = primecount_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	primecount_bin_url := "https://github.com/kimwalisch/primecount/archive/refs/tags/v7.14.bin"
	primecount_cmd_bin := exec.Command("curl", "-L", primecount_bin_url, "-o", "binary.bin")
	err = primecount_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	primecount_src_url := "https://github.com/kimwalisch/primecount/archive/refs/tags/v7.14.src.tar.gz"
	primecount_cmd_src := exec.Command("curl", "-L", primecount_src_url, "-o", "source.tar.gz")
	err = primecount_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	primecount_cmd_direct := exec.Command("./binary")
	err = primecount_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: primesieve")
exec.Command("latte", "install", "primesieve")
	fmt.Println("Instalando dependencia: libomp")
exec.Command("latte", "install", "libomp")
}
