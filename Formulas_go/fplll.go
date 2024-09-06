package main

// Fplll - Lattice algorithms using floating-point arithmetic
// Homepage: https://github.com/fplll/fplll

import (
	"fmt"
	
	"os/exec"
)

func installFplll() {
	// Método 1: Descargar y extraer .tar.gz
	fplll_tar_url := "https://github.com/fplll/fplll/releases/download/5.4.5/fplll-5.4.5.tar.gz"
	fplll_cmd_tar := exec.Command("curl", "-L", fplll_tar_url, "-o", "package.tar.gz")
	err := fplll_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fplll_zip_url := "https://github.com/fplll/fplll/releases/download/5.4.5/fplll-5.4.5.zip"
	fplll_cmd_zip := exec.Command("curl", "-L", fplll_zip_url, "-o", "package.zip")
	err = fplll_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fplll_bin_url := "https://github.com/fplll/fplll/releases/download/5.4.5/fplll-5.4.5.bin"
	fplll_cmd_bin := exec.Command("curl", "-L", fplll_bin_url, "-o", "binary.bin")
	err = fplll_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fplll_src_url := "https://github.com/fplll/fplll/releases/download/5.4.5/fplll-5.4.5.src.tar.gz"
	fplll_cmd_src := exec.Command("curl", "-L", fplll_src_url, "-o", "source.tar.gz")
	err = fplll_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fplll_cmd_direct := exec.Command("./binary")
	err = fplll_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: gmp")
exec.Command("latte", "install", "gmp")
	fmt.Println("Instalando dependencia: mpfr")
exec.Command("latte", "install", "mpfr")
}
