package main

// Hostdb - Generate DNS zones and DHCP configuration from hostlist.txt
// Homepage: https://code.google.com/archive/p/hostdb/

import (
	"fmt"
	
	"os/exec"
)

func installHostdb() {
	// Método 1: Descargar y extraer .tar.gz
	hostdb_tar_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/hostdb/hostdb-1.004.tgz"
	hostdb_cmd_tar := exec.Command("curl", "-L", hostdb_tar_url, "-o", "package.tar.gz")
	err := hostdb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hostdb_zip_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/hostdb/hostdb-1.004.tgz"
	hostdb_cmd_zip := exec.Command("curl", "-L", hostdb_zip_url, "-o", "package.zip")
	err = hostdb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hostdb_bin_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/hostdb/hostdb-1.004.tgz"
	hostdb_cmd_bin := exec.Command("curl", "-L", hostdb_bin_url, "-o", "binary.bin")
	err = hostdb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hostdb_src_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/hostdb/hostdb-1.004.tgz"
	hostdb_cmd_src := exec.Command("curl", "-L", hostdb_src_url, "-o", "source.tar.gz")
	err = hostdb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hostdb_cmd_direct := exec.Command("./binary")
	err = hostdb_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
