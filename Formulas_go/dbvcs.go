package main

// DbVcs - Version control for MySQL databases
// Homepage: https://github.com/infostreams/db

import (
	"fmt"
	
	"os/exec"
)

func installDbVcs() {
	// Método 1: Descargar y extraer .tar.gz
	dbvcs_tar_url := "https://github.com/infostreams/db/archive/refs/tags/1.1.tar.gz"
	dbvcs_cmd_tar := exec.Command("curl", "-L", dbvcs_tar_url, "-o", "package.tar.gz")
	err := dbvcs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dbvcs_zip_url := "https://github.com/infostreams/db/archive/refs/tags/1.1.zip"
	dbvcs_cmd_zip := exec.Command("curl", "-L", dbvcs_zip_url, "-o", "package.zip")
	err = dbvcs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dbvcs_bin_url := "https://github.com/infostreams/db/archive/refs/tags/1.1.bin"
	dbvcs_cmd_bin := exec.Command("curl", "-L", dbvcs_bin_url, "-o", "binary.bin")
	err = dbvcs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dbvcs_src_url := "https://github.com/infostreams/db/archive/refs/tags/1.1.src.tar.gz"
	dbvcs_cmd_src := exec.Command("curl", "-L", dbvcs_src_url, "-o", "source.tar.gz")
	err = dbvcs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dbvcs_cmd_direct := exec.Command("./binary")
	err = dbvcs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
