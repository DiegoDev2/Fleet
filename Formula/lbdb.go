package main

// Lbdb - Little brother's database for the mutt mail reader
// Homepage: https://www.spinnaker.de/lbdb/

import (
	"fmt"
	
	"os/exec"
)

func installLbdb() {
	// Método 1: Descargar y extraer .tar.gz
	lbdb_tar_url := "https://www.spinnaker.de/lbdb/download/lbdb-0.54.tar.gz"
	lbdb_cmd_tar := exec.Command("curl", "-L", lbdb_tar_url, "-o", "package.tar.gz")
	err := lbdb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lbdb_zip_url := "https://www.spinnaker.de/lbdb/download/lbdb-0.54.zip"
	lbdb_cmd_zip := exec.Command("curl", "-L", lbdb_zip_url, "-o", "package.zip")
	err = lbdb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lbdb_bin_url := "https://www.spinnaker.de/lbdb/download/lbdb-0.54.bin"
	lbdb_cmd_bin := exec.Command("curl", "-L", lbdb_bin_url, "-o", "binary.bin")
	err = lbdb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lbdb_src_url := "https://www.spinnaker.de/lbdb/download/lbdb-0.54.src.tar.gz"
	lbdb_cmd_src := exec.Command("curl", "-L", lbdb_src_url, "-o", "source.tar.gz")
	err = lbdb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lbdb_cmd_direct := exec.Command("./binary")
	err = lbdb_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: abook")
	exec.Command("latte", "install", "abook").Run()
	fmt.Println("Instalando dependencia: khard")
	exec.Command("latte", "install", "khard").Run()
}
