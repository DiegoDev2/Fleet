package main

// Bagit - Library for creation, manipulation, and validation of bags
// Homepage: https://libraryofcongress.github.io/bagit-python/

import (
	"fmt"
	
	"os/exec"
)

func installBagit() {
	// Método 1: Descargar y extraer .tar.gz
	bagit_tar_url := "https://files.pythonhosted.org/packages/e5/99/927b704237a1286f1022ea02a2fdfd82d5567cfbca97a4c343e2de7e37c4/bagit-1.8.1.tar.gz"
	bagit_cmd_tar := exec.Command("curl", "-L", bagit_tar_url, "-o", "package.tar.gz")
	err := bagit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bagit_zip_url := "https://files.pythonhosted.org/packages/e5/99/927b704237a1286f1022ea02a2fdfd82d5567cfbca97a4c343e2de7e37c4/bagit-1.8.1.zip"
	bagit_cmd_zip := exec.Command("curl", "-L", bagit_zip_url, "-o", "package.zip")
	err = bagit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bagit_bin_url := "https://files.pythonhosted.org/packages/e5/99/927b704237a1286f1022ea02a2fdfd82d5567cfbca97a4c343e2de7e37c4/bagit-1.8.1.bin"
	bagit_cmd_bin := exec.Command("curl", "-L", bagit_bin_url, "-o", "binary.bin")
	err = bagit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bagit_src_url := "https://files.pythonhosted.org/packages/e5/99/927b704237a1286f1022ea02a2fdfd82d5567cfbca97a4c343e2de7e37c4/bagit-1.8.1.src.tar.gz"
	bagit_cmd_src := exec.Command("curl", "-L", bagit_src_url, "-o", "source.tar.gz")
	err = bagit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bagit_cmd_direct := exec.Command("./binary")
	err = bagit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
