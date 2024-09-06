package main

// Jam - Make-like build tool
// Homepage: https://www.perforce.com/documentation/jam-documentation

import (
	"fmt"
	
	"os/exec"
)

func installJam() {
	// Método 1: Descargar y extraer .tar.gz
	jam_tar_url := "https://swarm.workshop.perforce.com/downloads/guest/perforce_software/jam/jam-2.6.1.zip"
	jam_cmd_tar := exec.Command("curl", "-L", jam_tar_url, "-o", "package.tar.gz")
	err := jam_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jam_zip_url := "https://swarm.workshop.perforce.com/downloads/guest/perforce_software/jam/jam-2.6.1.zip"
	jam_cmd_zip := exec.Command("curl", "-L", jam_zip_url, "-o", "package.zip")
	err = jam_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jam_bin_url := "https://swarm.workshop.perforce.com/downloads/guest/perforce_software/jam/jam-2.6.1.zip"
	jam_cmd_bin := exec.Command("curl", "-L", jam_bin_url, "-o", "binary.bin")
	err = jam_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jam_src_url := "https://swarm.workshop.perforce.com/downloads/guest/perforce_software/jam/jam-2.6.1.zip"
	jam_cmd_src := exec.Command("curl", "-L", jam_src_url, "-o", "source.tar.gz")
	err = jam_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jam_cmd_direct := exec.Command("./binary")
	err = jam_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
