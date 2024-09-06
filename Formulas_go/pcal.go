package main

// Pcal - Generate Postscript calendars without X
// Homepage: https://pcal.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installPcal() {
	// Método 1: Descargar y extraer .tar.gz
	pcal_tar_url := "https://downloads.sourceforge.net/project/pcal/pcal/pcal-4.11.0/pcal-4.11.0.tgz"
	pcal_cmd_tar := exec.Command("curl", "-L", pcal_tar_url, "-o", "package.tar.gz")
	err := pcal_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pcal_zip_url := "https://downloads.sourceforge.net/project/pcal/pcal/pcal-4.11.0/pcal-4.11.0.tgz"
	pcal_cmd_zip := exec.Command("curl", "-L", pcal_zip_url, "-o", "package.zip")
	err = pcal_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pcal_bin_url := "https://downloads.sourceforge.net/project/pcal/pcal/pcal-4.11.0/pcal-4.11.0.tgz"
	pcal_cmd_bin := exec.Command("curl", "-L", pcal_bin_url, "-o", "binary.bin")
	err = pcal_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pcal_src_url := "https://downloads.sourceforge.net/project/pcal/pcal/pcal-4.11.0/pcal-4.11.0.tgz"
	pcal_cmd_src := exec.Command("curl", "-L", pcal_src_url, "-o", "source.tar.gz")
	err = pcal_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pcal_cmd_direct := exec.Command("./binary")
	err = pcal_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
