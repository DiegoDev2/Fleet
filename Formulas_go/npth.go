package main

// Npth - New GNU portable threads library
// Homepage: https://gnupg.org/

import (
	"fmt"
	
	"os/exec"
)

func installNpth() {
	// Método 1: Descargar y extraer .tar.gz
	npth_tar_url := "https://gnupg.org/ftp/gcrypt/npth/npth-1.7.tar.bz2"
	npth_cmd_tar := exec.Command("curl", "-L", npth_tar_url, "-o", "package.tar.gz")
	err := npth_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	npth_zip_url := "https://gnupg.org/ftp/gcrypt/npth/npth-1.7.tar.bz2"
	npth_cmd_zip := exec.Command("curl", "-L", npth_zip_url, "-o", "package.zip")
	err = npth_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	npth_bin_url := "https://gnupg.org/ftp/gcrypt/npth/npth-1.7.tar.bz2"
	npth_cmd_bin := exec.Command("curl", "-L", npth_bin_url, "-o", "binary.bin")
	err = npth_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	npth_src_url := "https://gnupg.org/ftp/gcrypt/npth/npth-1.7.tar.bz2"
	npth_cmd_src := exec.Command("curl", "-L", npth_src_url, "-o", "source.tar.gz")
	err = npth_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	npth_cmd_direct := exec.Command("./binary")
	err = npth_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
