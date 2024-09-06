package main

// LiterateGit - Render hierarchical git repositories into HTML
// Homepage: https://github.com/bennorth/literate-git

import (
	"fmt"
	
	"os/exec"
)

func installLiterateGit() {
	// Método 1: Descargar y extraer .tar.gz
	literategit_tar_url := "https://files.pythonhosted.org/packages/7b/cc/1a6c994c90fa34cfa8e90e017c80f838b149fd0262daa24cdb930c091b48/literategit-0.5.0.tar.gz"
	literategit_cmd_tar := exec.Command("curl", "-L", literategit_tar_url, "-o", "package.tar.gz")
	err := literategit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	literategit_zip_url := "https://files.pythonhosted.org/packages/7b/cc/1a6c994c90fa34cfa8e90e017c80f838b149fd0262daa24cdb930c091b48/literategit-0.5.0.zip"
	literategit_cmd_zip := exec.Command("curl", "-L", literategit_zip_url, "-o", "package.zip")
	err = literategit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	literategit_bin_url := "https://files.pythonhosted.org/packages/7b/cc/1a6c994c90fa34cfa8e90e017c80f838b149fd0262daa24cdb930c091b48/literategit-0.5.0.bin"
	literategit_cmd_bin := exec.Command("curl", "-L", literategit_bin_url, "-o", "binary.bin")
	err = literategit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	literategit_src_url := "https://files.pythonhosted.org/packages/7b/cc/1a6c994c90fa34cfa8e90e017c80f838b149fd0262daa24cdb930c091b48/literategit-0.5.0.src.tar.gz"
	literategit_cmd_src := exec.Command("curl", "-L", literategit_src_url, "-o", "source.tar.gz")
	err = literategit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	literategit_cmd_direct := exec.Command("./binary")
	err = literategit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libgit2")
	exec.Command("latte", "install", "libgit2").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
}
