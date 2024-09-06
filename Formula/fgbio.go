package main

// Fgbio - Tools for working with genomic and high throughput sequencing data
// Homepage: https://fulcrumgenomics.github.io/fgbio/

import (
	"fmt"
	
	"os/exec"
)

func installFgbio() {
	// Método 1: Descargar y extraer .tar.gz
	fgbio_tar_url := "https://github.com/fulcrumgenomics/fgbio/releases/download/2.3.0/fgbio-2.3.0.jar"
	fgbio_cmd_tar := exec.Command("curl", "-L", fgbio_tar_url, "-o", "package.tar.gz")
	err := fgbio_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fgbio_zip_url := "https://github.com/fulcrumgenomics/fgbio/releases/download/2.3.0/fgbio-2.3.0.jar"
	fgbio_cmd_zip := exec.Command("curl", "-L", fgbio_zip_url, "-o", "package.zip")
	err = fgbio_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fgbio_bin_url := "https://github.com/fulcrumgenomics/fgbio/releases/download/2.3.0/fgbio-2.3.0.jar"
	fgbio_cmd_bin := exec.Command("curl", "-L", fgbio_bin_url, "-o", "binary.bin")
	err = fgbio_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fgbio_src_url := "https://github.com/fulcrumgenomics/fgbio/releases/download/2.3.0/fgbio-2.3.0.jar"
	fgbio_cmd_src := exec.Command("curl", "-L", fgbio_src_url, "-o", "source.tar.gz")
	err = fgbio_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fgbio_cmd_direct := exec.Command("./binary")
	err = fgbio_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
