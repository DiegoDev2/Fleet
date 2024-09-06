package main

// Arb - C library for arbitrary-precision interval arithmetic
// Homepage: https://arblib.org

import (
	"fmt"
	
	"os/exec"
)

func installArb() {
	// Método 1: Descargar y extraer .tar.gz
	arb_tar_url := "https://github.com/fredrik-johansson/arb/archive/refs/tags/2.23.0.tar.gz"
	arb_cmd_tar := exec.Command("curl", "-L", arb_tar_url, "-o", "package.tar.gz")
	err := arb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	arb_zip_url := "https://github.com/fredrik-johansson/arb/archive/refs/tags/2.23.0.zip"
	arb_cmd_zip := exec.Command("curl", "-L", arb_zip_url, "-o", "package.zip")
	err = arb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	arb_bin_url := "https://github.com/fredrik-johansson/arb/archive/refs/tags/2.23.0.bin"
	arb_cmd_bin := exec.Command("curl", "-L", arb_bin_url, "-o", "binary.bin")
	err = arb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	arb_src_url := "https://github.com/fredrik-johansson/arb/archive/refs/tags/2.23.0.src.tar.gz"
	arb_cmd_src := exec.Command("curl", "-L", arb_src_url, "-o", "source.tar.gz")
	err = arb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	arb_cmd_direct := exec.Command("./binary")
	err = arb_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: flint")
exec.Command("latte", "install", "flint")
	fmt.Println("Instalando dependencia: gmp")
exec.Command("latte", "install", "gmp")
	fmt.Println("Instalando dependencia: mpfr")
exec.Command("latte", "install", "mpfr")
}
