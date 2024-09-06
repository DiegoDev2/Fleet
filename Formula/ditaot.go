package main

// DitaOt - DITA Open Toolkit is an implementation of the OASIS DITA specification
// Homepage: https://www.dita-ot.org/

import (
	"fmt"
	
	"os/exec"
)

func installDitaOt() {
	// Método 1: Descargar y extraer .tar.gz
	ditaot_tar_url := "https://github.com/dita-ot/dita-ot/releases/download/4.2.3/dita-ot-4.2.3.zip"
	ditaot_cmd_tar := exec.Command("curl", "-L", ditaot_tar_url, "-o", "package.tar.gz")
	err := ditaot_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ditaot_zip_url := "https://github.com/dita-ot/dita-ot/releases/download/4.2.3/dita-ot-4.2.3.zip"
	ditaot_cmd_zip := exec.Command("curl", "-L", ditaot_zip_url, "-o", "package.zip")
	err = ditaot_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ditaot_bin_url := "https://github.com/dita-ot/dita-ot/releases/download/4.2.3/dita-ot-4.2.3.zip"
	ditaot_cmd_bin := exec.Command("curl", "-L", ditaot_bin_url, "-o", "binary.bin")
	err = ditaot_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ditaot_src_url := "https://github.com/dita-ot/dita-ot/releases/download/4.2.3/dita-ot-4.2.3.zip"
	ditaot_cmd_src := exec.Command("curl", "-L", ditaot_src_url, "-o", "source.tar.gz")
	err = ditaot_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ditaot_cmd_direct := exec.Command("./binary")
	err = ditaot_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
