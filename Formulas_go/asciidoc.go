package main

// Asciidoc - Formatter/translator for text files to numerous formats
// Homepage: https://asciidoc-py.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installAsciidoc() {
	// Método 1: Descargar y extraer .tar.gz
	asciidoc_tar_url := "https://files.pythonhosted.org/packages/1d/e7/315a82f2d256e9270977aa3c15e8fe281fd7c40b8e2a0b97e0cb61ca8fa0/asciidoc-10.2.1.tar.gz"
	asciidoc_cmd_tar := exec.Command("curl", "-L", asciidoc_tar_url, "-o", "package.tar.gz")
	err := asciidoc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	asciidoc_zip_url := "https://files.pythonhosted.org/packages/1d/e7/315a82f2d256e9270977aa3c15e8fe281fd7c40b8e2a0b97e0cb61ca8fa0/asciidoc-10.2.1.zip"
	asciidoc_cmd_zip := exec.Command("curl", "-L", asciidoc_zip_url, "-o", "package.zip")
	err = asciidoc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	asciidoc_bin_url := "https://files.pythonhosted.org/packages/1d/e7/315a82f2d256e9270977aa3c15e8fe281fd7c40b8e2a0b97e0cb61ca8fa0/asciidoc-10.2.1.bin"
	asciidoc_cmd_bin := exec.Command("curl", "-L", asciidoc_bin_url, "-o", "binary.bin")
	err = asciidoc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	asciidoc_src_url := "https://files.pythonhosted.org/packages/1d/e7/315a82f2d256e9270977aa3c15e8fe281fd7c40b8e2a0b97e0cb61ca8fa0/asciidoc-10.2.1.src.tar.gz"
	asciidoc_cmd_src := exec.Command("curl", "-L", asciidoc_src_url, "-o", "source.tar.gz")
	err = asciidoc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	asciidoc_cmd_direct := exec.Command("./binary")
	err = asciidoc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: docbook")
exec.Command("latte", "install", "docbook")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: source-highlight")
exec.Command("latte", "install", "source-highlight")
}
