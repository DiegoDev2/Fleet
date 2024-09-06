package main

// Makeicns - Create icns files from the command-line
// Homepage: http://www.amnoid.de/icns/makeicns.html

import (
	"fmt"
	
	"os/exec"
)

func installMakeicns() {
	// Método 1: Descargar y extraer .tar.gz
	makeicns_tar_url := "https://distfiles.macports.org/makeicns/makeicns-1.4.10a.tar.bz2"
	makeicns_cmd_tar := exec.Command("curl", "-L", makeicns_tar_url, "-o", "package.tar.gz")
	err := makeicns_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	makeicns_zip_url := "https://distfiles.macports.org/makeicns/makeicns-1.4.10a.tar.bz2"
	makeicns_cmd_zip := exec.Command("curl", "-L", makeicns_zip_url, "-o", "package.zip")
	err = makeicns_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	makeicns_bin_url := "https://distfiles.macports.org/makeicns/makeicns-1.4.10a.tar.bz2"
	makeicns_cmd_bin := exec.Command("curl", "-L", makeicns_bin_url, "-o", "binary.bin")
	err = makeicns_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	makeicns_src_url := "https://distfiles.macports.org/makeicns/makeicns-1.4.10a.tar.bz2"
	makeicns_cmd_src := exec.Command("curl", "-L", makeicns_src_url, "-o", "source.tar.gz")
	err = makeicns_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	makeicns_cmd_direct := exec.Command("./binary")
	err = makeicns_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
