package main

// Lit - Portable tool for LLVM- and Clang-style test suites
// Homepage: https://llvm.org

import (
	"fmt"
	
	"os/exec"
)

func installLit() {
	// Método 1: Descargar y extraer .tar.gz
	lit_tar_url := "https://files.pythonhosted.org/packages/47/b4/d7e210971494db7b9a9ac48ff37dfa59a8b14c773f9cf47e6bda58411c0d/lit-18.1.8.tar.gz"
	lit_cmd_tar := exec.Command("curl", "-L", lit_tar_url, "-o", "package.tar.gz")
	err := lit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lit_zip_url := "https://files.pythonhosted.org/packages/47/b4/d7e210971494db7b9a9ac48ff37dfa59a8b14c773f9cf47e6bda58411c0d/lit-18.1.8.zip"
	lit_cmd_zip := exec.Command("curl", "-L", lit_zip_url, "-o", "package.zip")
	err = lit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lit_bin_url := "https://files.pythonhosted.org/packages/47/b4/d7e210971494db7b9a9ac48ff37dfa59a8b14c773f9cf47e6bda58411c0d/lit-18.1.8.bin"
	lit_cmd_bin := exec.Command("curl", "-L", lit_bin_url, "-o", "binary.bin")
	err = lit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lit_src_url := "https://files.pythonhosted.org/packages/47/b4/d7e210971494db7b9a9ac48ff37dfa59a8b14c773f9cf47e6bda58411c0d/lit-18.1.8.src.tar.gz"
	lit_cmd_src := exec.Command("curl", "-L", lit_src_url, "-o", "source.tar.gz")
	err = lit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lit_cmd_direct := exec.Command("./binary")
	err = lit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: llvm")
exec.Command("latte", "install", "llvm")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
