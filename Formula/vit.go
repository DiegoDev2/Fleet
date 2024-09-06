package main

// Vit - Full-screen terminal interface for Taskwarrior
// Homepage: https://taskwarrior.org/news/news.20140406.html

import (
	"fmt"
	
	"os/exec"
)

func installVit() {
	// Método 1: Descargar y extraer .tar.gz
	vit_tar_url := "https://files.pythonhosted.org/packages/a2/24/71ef618e17ced54d3ad706215165ebeb6ebc86f5d71ded58c4dbcba62b83/vit-2.3.2.tar.gz"
	vit_cmd_tar := exec.Command("curl", "-L", vit_tar_url, "-o", "package.tar.gz")
	err := vit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vit_zip_url := "https://files.pythonhosted.org/packages/a2/24/71ef618e17ced54d3ad706215165ebeb6ebc86f5d71ded58c4dbcba62b83/vit-2.3.2.zip"
	vit_cmd_zip := exec.Command("curl", "-L", vit_zip_url, "-o", "package.zip")
	err = vit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vit_bin_url := "https://files.pythonhosted.org/packages/a2/24/71ef618e17ced54d3ad706215165ebeb6ebc86f5d71ded58c4dbcba62b83/vit-2.3.2.bin"
	vit_cmd_bin := exec.Command("curl", "-L", vit_bin_url, "-o", "binary.bin")
	err = vit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vit_src_url := "https://files.pythonhosted.org/packages/a2/24/71ef618e17ced54d3ad706215165ebeb6ebc86f5d71ded58c4dbcba62b83/vit-2.3.2.src.tar.gz"
	vit_cmd_src := exec.Command("curl", "-L", vit_src_url, "-o", "source.tar.gz")
	err = vit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vit_cmd_direct := exec.Command("./binary")
	err = vit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: task")
	exec.Command("latte", "install", "task").Run()
}
