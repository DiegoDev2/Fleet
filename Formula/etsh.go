package main

// Etsh - Two ports of /bin/sh from V6 UNIX (circa 1975)
// Homepage: https://etsh.dev/

import (
	"fmt"
	
	"os/exec"
)

func installEtsh() {
	// Método 1: Descargar y extraer .tar.gz
	etsh_tar_url := "https://etsh.dev/src/etsh_5.4.0/etsh-5.4.0.tar.xz"
	etsh_cmd_tar := exec.Command("curl", "-L", etsh_tar_url, "-o", "package.tar.gz")
	err := etsh_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	etsh_zip_url := "https://etsh.dev/src/etsh_5.4.0/etsh-5.4.0.tar.xz"
	etsh_cmd_zip := exec.Command("curl", "-L", etsh_zip_url, "-o", "package.zip")
	err = etsh_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	etsh_bin_url := "https://etsh.dev/src/etsh_5.4.0/etsh-5.4.0.tar.xz"
	etsh_cmd_bin := exec.Command("curl", "-L", etsh_bin_url, "-o", "binary.bin")
	err = etsh_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	etsh_src_url := "https://etsh.dev/src/etsh_5.4.0/etsh-5.4.0.tar.xz"
	etsh_cmd_src := exec.Command("curl", "-L", etsh_src_url, "-o", "source.tar.gz")
	err = etsh_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	etsh_cmd_direct := exec.Command("./binary")
	err = etsh_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
