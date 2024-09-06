package main

// Kin - Sane PBXProj files
// Homepage: https://github.com/Serchinastico/Kin

import (
	"fmt"
	
	"os/exec"
)

func installKin() {
	// Método 1: Descargar y extraer .tar.gz
	kin_tar_url := "https://files.pythonhosted.org/packages/4f/36/dcb0e16c5634d58d0ef2f771fe1e608264698394f4a184afc289d9a85bb8/kin-2.1.10.tar.gz"
	kin_cmd_tar := exec.Command("curl", "-L", kin_tar_url, "-o", "package.tar.gz")
	err := kin_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kin_zip_url := "https://files.pythonhosted.org/packages/4f/36/dcb0e16c5634d58d0ef2f771fe1e608264698394f4a184afc289d9a85bb8/kin-2.1.10.zip"
	kin_cmd_zip := exec.Command("curl", "-L", kin_zip_url, "-o", "package.zip")
	err = kin_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kin_bin_url := "https://files.pythonhosted.org/packages/4f/36/dcb0e16c5634d58d0ef2f771fe1e608264698394f4a184afc289d9a85bb8/kin-2.1.10.bin"
	kin_cmd_bin := exec.Command("curl", "-L", kin_bin_url, "-o", "binary.bin")
	err = kin_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kin_src_url := "https://files.pythonhosted.org/packages/4f/36/dcb0e16c5634d58d0ef2f771fe1e608264698394f4a184afc289d9a85bb8/kin-2.1.10.src.tar.gz"
	kin_cmd_src := exec.Command("curl", "-L", kin_src_url, "-o", "source.tar.gz")
	err = kin_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kin_cmd_direct := exec.Command("./binary")
	err = kin_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
