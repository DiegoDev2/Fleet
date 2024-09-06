package main

// PicardTools - Tools for manipulating HTS data and formats
// Homepage: https://broadinstitute.github.io/picard/

import (
	"fmt"
	
	"os/exec"
)

func installPicardTools() {
	// Método 1: Descargar y extraer .tar.gz
	picardtools_tar_url := "https://github.com/broadinstitute/picard/releases/download/3.2.0/picard.jar"
	picardtools_cmd_tar := exec.Command("curl", "-L", picardtools_tar_url, "-o", "package.tar.gz")
	err := picardtools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	picardtools_zip_url := "https://github.com/broadinstitute/picard/releases/download/3.2.0/picard.jar"
	picardtools_cmd_zip := exec.Command("curl", "-L", picardtools_zip_url, "-o", "package.zip")
	err = picardtools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	picardtools_bin_url := "https://github.com/broadinstitute/picard/releases/download/3.2.0/picard.jar"
	picardtools_cmd_bin := exec.Command("curl", "-L", picardtools_bin_url, "-o", "binary.bin")
	err = picardtools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	picardtools_src_url := "https://github.com/broadinstitute/picard/releases/download/3.2.0/picard.jar"
	picardtools_cmd_src := exec.Command("curl", "-L", picardtools_src_url, "-o", "source.tar.gz")
	err = picardtools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	picardtools_cmd_direct := exec.Command("./binary")
	err = picardtools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
