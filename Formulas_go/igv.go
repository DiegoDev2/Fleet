package main

// Igv - Interactive Genomics Viewer
// Homepage: https://igv.org/doc/desktop/

import (
	"fmt"
	
	"os/exec"
)

func installIgv() {
	// Método 1: Descargar y extraer .tar.gz
	igv_tar_url := "https://data.broadinstitute.org/igv/projects/downloads/2.18/IGV_2.18.2.zip"
	igv_cmd_tar := exec.Command("curl", "-L", igv_tar_url, "-o", "package.tar.gz")
	err := igv_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	igv_zip_url := "https://data.broadinstitute.org/igv/projects/downloads/2.18/IGV_2.18.2.zip"
	igv_cmd_zip := exec.Command("curl", "-L", igv_zip_url, "-o", "package.zip")
	err = igv_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	igv_bin_url := "https://data.broadinstitute.org/igv/projects/downloads/2.18/IGV_2.18.2.zip"
	igv_cmd_bin := exec.Command("curl", "-L", igv_bin_url, "-o", "binary.bin")
	err = igv_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	igv_src_url := "https://data.broadinstitute.org/igv/projects/downloads/2.18/IGV_2.18.2.zip"
	igv_cmd_src := exec.Command("curl", "-L", igv_src_url, "-o", "source.tar.gz")
	err = igv_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	igv_cmd_direct := exec.Command("./binary")
	err = igv_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
