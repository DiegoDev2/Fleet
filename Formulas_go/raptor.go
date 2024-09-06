package main

// Raptor - RDF parser toolkit
// Homepage: https://librdf.org/raptor/

import (
	"fmt"
	
	"os/exec"
)

func installRaptor() {
	// Método 1: Descargar y extraer .tar.gz
	raptor_tar_url := "https://download.librdf.org/source/raptor2-2.0.16.tar.gz"
	raptor_cmd_tar := exec.Command("curl", "-L", raptor_tar_url, "-o", "package.tar.gz")
	err := raptor_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	raptor_zip_url := "https://download.librdf.org/source/raptor2-2.0.16.zip"
	raptor_cmd_zip := exec.Command("curl", "-L", raptor_zip_url, "-o", "package.zip")
	err = raptor_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	raptor_bin_url := "https://download.librdf.org/source/raptor2-2.0.16.bin"
	raptor_cmd_bin := exec.Command("curl", "-L", raptor_bin_url, "-o", "binary.bin")
	err = raptor_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	raptor_src_url := "https://download.librdf.org/source/raptor2-2.0.16.src.tar.gz"
	raptor_cmd_src := exec.Command("curl", "-L", raptor_src_url, "-o", "source.tar.gz")
	err = raptor_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	raptor_cmd_direct := exec.Command("./binary")
	err = raptor_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
