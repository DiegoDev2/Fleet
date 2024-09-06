package main

// NameThatHash - Modern hash identification system
// Homepage: https://nth.skerritt.blog/

import (
	"fmt"
	
	"os/exec"
)

func installNameThatHash() {
	// Método 1: Descargar y extraer .tar.gz
	namethathash_tar_url := "https://files.pythonhosted.org/packages/7a/d6/5bea2b09a8b4dbfd92610432dbbcdda9f983be3de770a296df957fed5d06/name_that_hash-1.11.0.tar.gz"
	namethathash_cmd_tar := exec.Command("curl", "-L", namethathash_tar_url, "-o", "package.tar.gz")
	err := namethathash_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	namethathash_zip_url := "https://files.pythonhosted.org/packages/7a/d6/5bea2b09a8b4dbfd92610432dbbcdda9f983be3de770a296df957fed5d06/name_that_hash-1.11.0.zip"
	namethathash_cmd_zip := exec.Command("curl", "-L", namethathash_zip_url, "-o", "package.zip")
	err = namethathash_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	namethathash_bin_url := "https://files.pythonhosted.org/packages/7a/d6/5bea2b09a8b4dbfd92610432dbbcdda9f983be3de770a296df957fed5d06/name_that_hash-1.11.0.bin"
	namethathash_cmd_bin := exec.Command("curl", "-L", namethathash_bin_url, "-o", "binary.bin")
	err = namethathash_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	namethathash_src_url := "https://files.pythonhosted.org/packages/7a/d6/5bea2b09a8b4dbfd92610432dbbcdda9f983be3de770a296df957fed5d06/name_that_hash-1.11.0.src.tar.gz"
	namethathash_cmd_src := exec.Command("curl", "-L", namethathash_src_url, "-o", "source.tar.gz")
	err = namethathash_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	namethathash_cmd_direct := exec.Command("./binary")
	err = namethathash_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
