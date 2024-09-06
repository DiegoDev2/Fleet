package main

// Pdfalyzer - PDF analysis toolkit
// Homepage: https://github.com/michelcrypt4d4mus/pdfalyzer

import (
	"fmt"
	
	"os/exec"
)

func installPdfalyzer() {
	// Método 1: Descargar y extraer .tar.gz
	pdfalyzer_tar_url := "https://files.pythonhosted.org/packages/0a/b5/7448e055f672565654e07658fd5dad471d92542ae76a06d844d45f696503/pdfalyzer-1.15.1.tar.gz"
	pdfalyzer_cmd_tar := exec.Command("curl", "-L", pdfalyzer_tar_url, "-o", "package.tar.gz")
	err := pdfalyzer_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pdfalyzer_zip_url := "https://files.pythonhosted.org/packages/0a/b5/7448e055f672565654e07658fd5dad471d92542ae76a06d844d45f696503/pdfalyzer-1.15.1.zip"
	pdfalyzer_cmd_zip := exec.Command("curl", "-L", pdfalyzer_zip_url, "-o", "package.zip")
	err = pdfalyzer_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pdfalyzer_bin_url := "https://files.pythonhosted.org/packages/0a/b5/7448e055f672565654e07658fd5dad471d92542ae76a06d844d45f696503/pdfalyzer-1.15.1.bin"
	pdfalyzer_cmd_bin := exec.Command("curl", "-L", pdfalyzer_bin_url, "-o", "binary.bin")
	err = pdfalyzer_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pdfalyzer_src_url := "https://files.pythonhosted.org/packages/0a/b5/7448e055f672565654e07658fd5dad471d92542ae76a06d844d45f696503/pdfalyzer-1.15.1.src.tar.gz"
	pdfalyzer_cmd_src := exec.Command("curl", "-L", pdfalyzer_src_url, "-o", "source.tar.gz")
	err = pdfalyzer_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pdfalyzer_cmd_direct := exec.Command("./binary")
	err = pdfalyzer_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
