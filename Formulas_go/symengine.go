package main

// Symengine - Fast symbolic manipulation library written in C++
// Homepage: https://www.sympy.org/en/index.html

import (
	"fmt"
	
	"os/exec"
)

func installSymengine() {
	// Método 1: Descargar y extraer .tar.gz
	symengine_tar_url := "https://github.com/symengine/symengine/archive/refs/tags/v0.12.0.tar.gz"
	symengine_cmd_tar := exec.Command("curl", "-L", symengine_tar_url, "-o", "package.tar.gz")
	err := symengine_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	symengine_zip_url := "https://github.com/symengine/symengine/archive/refs/tags/v0.12.0.zip"
	symengine_cmd_zip := exec.Command("curl", "-L", symengine_zip_url, "-o", "package.zip")
	err = symengine_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	symengine_bin_url := "https://github.com/symengine/symengine/archive/refs/tags/v0.12.0.bin"
	symengine_cmd_bin := exec.Command("curl", "-L", symengine_bin_url, "-o", "binary.bin")
	err = symengine_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	symengine_src_url := "https://github.com/symengine/symengine/archive/refs/tags/v0.12.0.src.tar.gz"
	symengine_cmd_src := exec.Command("curl", "-L", symengine_src_url, "-o", "source.tar.gz")
	err = symengine_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	symengine_cmd_direct := exec.Command("./binary")
	err = symengine_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cereal")
exec.Command("latte", "install", "cereal")
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: flint")
exec.Command("latte", "install", "flint")
	fmt.Println("Instalando dependencia: gmp")
exec.Command("latte", "install", "gmp")
	fmt.Println("Instalando dependencia: libmpc")
exec.Command("latte", "install", "libmpc")
	fmt.Println("Instalando dependencia: llvm")
exec.Command("latte", "install", "llvm")
	fmt.Println("Instalando dependencia: mpfr")
exec.Command("latte", "install", "mpfr")
	fmt.Println("Instalando dependencia: zstd")
exec.Command("latte", "install", "zstd")
	fmt.Println("Instalando dependencia: z3")
exec.Command("latte", "install", "z3")
}
