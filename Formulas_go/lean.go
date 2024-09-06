package main

// Lean - Theorem prover
// Homepage: https://leanprover-community.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installLean() {
	// Método 1: Descargar y extraer .tar.gz
	lean_tar_url := "https://github.com/leanprover-community/lean/archive/refs/tags/v3.51.1.tar.gz"
	lean_cmd_tar := exec.Command("curl", "-L", lean_tar_url, "-o", "package.tar.gz")
	err := lean_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lean_zip_url := "https://github.com/leanprover-community/lean/archive/refs/tags/v3.51.1.zip"
	lean_cmd_zip := exec.Command("curl", "-L", lean_zip_url, "-o", "package.zip")
	err = lean_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lean_bin_url := "https://github.com/leanprover-community/lean/archive/refs/tags/v3.51.1.bin"
	lean_cmd_bin := exec.Command("curl", "-L", lean_bin_url, "-o", "binary.bin")
	err = lean_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lean_src_url := "https://github.com/leanprover-community/lean/archive/refs/tags/v3.51.1.src.tar.gz"
	lean_cmd_src := exec.Command("curl", "-L", lean_src_url, "-o", "source.tar.gz")
	err = lean_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lean_cmd_direct := exec.Command("./binary")
	err = lean_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: coreutils")
exec.Command("latte", "install", "coreutils")
	fmt.Println("Instalando dependencia: gmp")
exec.Command("latte", "install", "gmp")
	fmt.Println("Instalando dependencia: jemalloc")
exec.Command("latte", "install", "jemalloc")
}
