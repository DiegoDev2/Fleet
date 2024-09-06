package main

// Cromwell - Workflow Execution Engine using Workflow Description Language
// Homepage: https://github.com/broadinstitute/cromwell

import (
	"fmt"
	
	"os/exec"
)

func installCromwell() {
	// Método 1: Descargar y extraer .tar.gz
	cromwell_tar_url := "https://github.com/broadinstitute/cromwell/releases/download/87/cromwell-87.jar"
	cromwell_cmd_tar := exec.Command("curl", "-L", cromwell_tar_url, "-o", "package.tar.gz")
	err := cromwell_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cromwell_zip_url := "https://github.com/broadinstitute/cromwell/releases/download/87/cromwell-87.jar"
	cromwell_cmd_zip := exec.Command("curl", "-L", cromwell_zip_url, "-o", "package.zip")
	err = cromwell_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cromwell_bin_url := "https://github.com/broadinstitute/cromwell/releases/download/87/cromwell-87.jar"
	cromwell_cmd_bin := exec.Command("curl", "-L", cromwell_bin_url, "-o", "binary.bin")
	err = cromwell_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cromwell_src_url := "https://github.com/broadinstitute/cromwell/releases/download/87/cromwell-87.jar"
	cromwell_cmd_src := exec.Command("curl", "-L", cromwell_src_url, "-o", "source.tar.gz")
	err = cromwell_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cromwell_cmd_direct := exec.Command("./binary")
	err = cromwell_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: sbt")
	exec.Command("latte", "install", "sbt").Run()
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
