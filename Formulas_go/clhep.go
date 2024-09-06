package main

// Clhep - Class Library for High Energy Physics
// Homepage: https://proj-clhep.web.cern.ch/proj-clhep/

import (
	"fmt"
	
	"os/exec"
)

func installClhep() {
	// Método 1: Descargar y extraer .tar.gz
	clhep_tar_url := "https://proj-clhep.web.cern.ch/proj-clhep/dist1/clhep-2.4.7.1.tgz"
	clhep_cmd_tar := exec.Command("curl", "-L", clhep_tar_url, "-o", "package.tar.gz")
	err := clhep_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	clhep_zip_url := "https://proj-clhep.web.cern.ch/proj-clhep/dist1/clhep-2.4.7.1.tgz"
	clhep_cmd_zip := exec.Command("curl", "-L", clhep_zip_url, "-o", "package.zip")
	err = clhep_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	clhep_bin_url := "https://proj-clhep.web.cern.ch/proj-clhep/dist1/clhep-2.4.7.1.tgz"
	clhep_cmd_bin := exec.Command("curl", "-L", clhep_bin_url, "-o", "binary.bin")
	err = clhep_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	clhep_src_url := "https://proj-clhep.web.cern.ch/proj-clhep/dist1/clhep-2.4.7.1.tgz"
	clhep_cmd_src := exec.Command("curl", "-L", clhep_src_url, "-o", "source.tar.gz")
	err = clhep_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	clhep_cmd_direct := exec.Command("./binary")
	err = clhep_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
