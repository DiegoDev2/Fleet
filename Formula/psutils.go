package main

// Psutils - Utilities for manipulating PostScript documents
// Homepage: https://github.com/rrthomas/psutils

import (
	"fmt"
	
	"os/exec"
)

func installPsutils() {
	// Método 1: Descargar y extraer .tar.gz
	psutils_tar_url := "https://files.pythonhosted.org/packages/ff/46/8b697d7976ceccd4971886f04b57ec3ef46d8976b2beefa97892bfa35271/pspdfutils-3.3.5.tar.gz"
	psutils_cmd_tar := exec.Command("curl", "-L", psutils_tar_url, "-o", "package.tar.gz")
	err := psutils_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	psutils_zip_url := "https://files.pythonhosted.org/packages/ff/46/8b697d7976ceccd4971886f04b57ec3ef46d8976b2beefa97892bfa35271/pspdfutils-3.3.5.zip"
	psutils_cmd_zip := exec.Command("curl", "-L", psutils_zip_url, "-o", "package.zip")
	err = psutils_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	psutils_bin_url := "https://files.pythonhosted.org/packages/ff/46/8b697d7976ceccd4971886f04b57ec3ef46d8976b2beefa97892bfa35271/pspdfutils-3.3.5.bin"
	psutils_cmd_bin := exec.Command("curl", "-L", psutils_bin_url, "-o", "binary.bin")
	err = psutils_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	psutils_src_url := "https://files.pythonhosted.org/packages/ff/46/8b697d7976ceccd4971886f04b57ec3ef46d8976b2beefa97892bfa35271/pspdfutils-3.3.5.src.tar.gz"
	psutils_cmd_src := exec.Command("curl", "-L", psutils_src_url, "-o", "source.tar.gz")
	err = psutils_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	psutils_cmd_direct := exec.Command("./binary")
	err = psutils_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libpaper")
	exec.Command("latte", "install", "libpaper").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
