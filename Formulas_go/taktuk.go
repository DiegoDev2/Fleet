package main

// Taktuk - Deploy commands to (a potentially large set of) remote nodes
// Homepage: https://taktuk.gitlabpages.inria.fr/

import (
	"fmt"
	
	"os/exec"
)

func installTaktuk() {
	// Método 1: Descargar y extraer .tar.gz
	taktuk_tar_url := "https://deb.debian.org/debian/pool/main/t/taktuk/taktuk_3.7.7.orig.tar.gz"
	taktuk_cmd_tar := exec.Command("curl", "-L", taktuk_tar_url, "-o", "package.tar.gz")
	err := taktuk_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	taktuk_zip_url := "https://deb.debian.org/debian/pool/main/t/taktuk/taktuk_3.7.7.orig.zip"
	taktuk_cmd_zip := exec.Command("curl", "-L", taktuk_zip_url, "-o", "package.zip")
	err = taktuk_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	taktuk_bin_url := "https://deb.debian.org/debian/pool/main/t/taktuk/taktuk_3.7.7.orig.bin"
	taktuk_cmd_bin := exec.Command("curl", "-L", taktuk_bin_url, "-o", "binary.bin")
	err = taktuk_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	taktuk_src_url := "https://deb.debian.org/debian/pool/main/t/taktuk/taktuk_3.7.7.orig.src.tar.gz"
	taktuk_cmd_src := exec.Command("curl", "-L", taktuk_src_url, "-o", "source.tar.gz")
	err = taktuk_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	taktuk_cmd_direct := exec.Command("./binary")
	err = taktuk_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
