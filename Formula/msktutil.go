package main

// Msktutil - Active Directory keytab management
// Homepage: https://github.com/msktutil/msktutil

import (
	"fmt"
	
	"os/exec"
)

func installMsktutil() {
	// Método 1: Descargar y extraer .tar.gz
	msktutil_tar_url := "https://github.com/msktutil/msktutil/releases/download/1.2.1/msktutil-1.2.1.tar.bz2"
	msktutil_cmd_tar := exec.Command("curl", "-L", msktutil_tar_url, "-o", "package.tar.gz")
	err := msktutil_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	msktutil_zip_url := "https://github.com/msktutil/msktutil/releases/download/1.2.1/msktutil-1.2.1.tar.bz2"
	msktutil_cmd_zip := exec.Command("curl", "-L", msktutil_zip_url, "-o", "package.zip")
	err = msktutil_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	msktutil_bin_url := "https://github.com/msktutil/msktutil/releases/download/1.2.1/msktutil-1.2.1.tar.bz2"
	msktutil_cmd_bin := exec.Command("curl", "-L", msktutil_bin_url, "-o", "binary.bin")
	err = msktutil_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	msktutil_src_url := "https://github.com/msktutil/msktutil/releases/download/1.2.1/msktutil-1.2.1.tar.bz2"
	msktutil_cmd_src := exec.Command("curl", "-L", msktutil_src_url, "-o", "source.tar.gz")
	err = msktutil_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	msktutil_cmd_direct := exec.Command("./binary")
	err = msktutil_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
