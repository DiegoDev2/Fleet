package main

// Iocextract - Defanged indicator of compromise extractor
// Homepage: https://inquest.readthedocs.io/projects/iocextract/en/latest/

import (
	"fmt"
	
	"os/exec"
)

func installIocextract() {
	// Método 1: Descargar y extraer .tar.gz
	iocextract_tar_url := "https://files.pythonhosted.org/packages/ad/4b/19934df6cd6a0f6923aabae391a67b630fdd03c12c1226377c99a747a4f1/iocextract-1.16.1.tar.gz"
	iocextract_cmd_tar := exec.Command("curl", "-L", iocextract_tar_url, "-o", "package.tar.gz")
	err := iocextract_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	iocextract_zip_url := "https://files.pythonhosted.org/packages/ad/4b/19934df6cd6a0f6923aabae391a67b630fdd03c12c1226377c99a747a4f1/iocextract-1.16.1.zip"
	iocextract_cmd_zip := exec.Command("curl", "-L", iocextract_zip_url, "-o", "package.zip")
	err = iocextract_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	iocextract_bin_url := "https://files.pythonhosted.org/packages/ad/4b/19934df6cd6a0f6923aabae391a67b630fdd03c12c1226377c99a747a4f1/iocextract-1.16.1.bin"
	iocextract_cmd_bin := exec.Command("curl", "-L", iocextract_bin_url, "-o", "binary.bin")
	err = iocextract_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	iocextract_src_url := "https://files.pythonhosted.org/packages/ad/4b/19934df6cd6a0f6923aabae391a67b630fdd03c12c1226377c99a747a4f1/iocextract-1.16.1.src.tar.gz"
	iocextract_cmd_src := exec.Command("curl", "-L", iocextract_src_url, "-o", "source.tar.gz")
	err = iocextract_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	iocextract_cmd_direct := exec.Command("./binary")
	err = iocextract_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
	exec.Command("latte", "install", "certifi").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
