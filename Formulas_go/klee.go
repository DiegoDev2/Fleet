package main

// Klee - Symbolic Execution Engine
// Homepage: https://klee-se.org

import (
	"fmt"
	
	"os/exec"
)

func installKlee() {
	// Método 1: Descargar y extraer .tar.gz
	klee_tar_url := "https://github.com/klee/klee/archive/refs/tags/v3.1.tar.gz"
	klee_cmd_tar := exec.Command("curl", "-L", klee_tar_url, "-o", "package.tar.gz")
	err := klee_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	klee_zip_url := "https://github.com/klee/klee/archive/refs/tags/v3.1.zip"
	klee_cmd_zip := exec.Command("curl", "-L", klee_zip_url, "-o", "package.zip")
	err = klee_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	klee_bin_url := "https://github.com/klee/klee/archive/refs/tags/v3.1.bin"
	klee_cmd_bin := exec.Command("curl", "-L", klee_bin_url, "-o", "binary.bin")
	err = klee_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	klee_src_url := "https://github.com/klee/klee/archive/refs/tags/v3.1.src.tar.gz"
	klee_cmd_src := exec.Command("curl", "-L", klee_src_url, "-o", "source.tar.gz")
	err = klee_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	klee_cmd_direct := exec.Command("./binary")
	err = klee_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: gperftools")
exec.Command("latte", "install", "gperftools")
	fmt.Println("Instalando dependencia: llvm@14")
exec.Command("latte", "install", "llvm@14")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: sqlite")
exec.Command("latte", "install", "sqlite")
	fmt.Println("Instalando dependencia: stp")
exec.Command("latte", "install", "stp")
	fmt.Println("Instalando dependencia: wllvm")
exec.Command("latte", "install", "wllvm")
	fmt.Println("Instalando dependencia: z3")
exec.Command("latte", "install", "z3")
	fmt.Println("Instalando dependencia: cryptominisat")
exec.Command("latte", "install", "cryptominisat")
	fmt.Println("Instalando dependencia: minisat")
exec.Command("latte", "install", "minisat")
	fmt.Println("Instalando dependencia: python-setuptools")
exec.Command("latte", "install", "python-setuptools")
}
