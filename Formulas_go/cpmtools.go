package main

// Cpmtools - Tools to access CP/M file systems
// Homepage: http://www.moria.de/~michael/cpmtools/

import (
	"fmt"
	
	"os/exec"
)

func installCpmtools() {
	// Método 1: Descargar y extraer .tar.gz
	cpmtools_tar_url := "http://www.moria.de/~michael/cpmtools/files/cpmtools-2.23.tar.gz"
	cpmtools_cmd_tar := exec.Command("curl", "-L", cpmtools_tar_url, "-o", "package.tar.gz")
	err := cpmtools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cpmtools_zip_url := "http://www.moria.de/~michael/cpmtools/files/cpmtools-2.23.zip"
	cpmtools_cmd_zip := exec.Command("curl", "-L", cpmtools_zip_url, "-o", "package.zip")
	err = cpmtools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cpmtools_bin_url := "http://www.moria.de/~michael/cpmtools/files/cpmtools-2.23.bin"
	cpmtools_cmd_bin := exec.Command("curl", "-L", cpmtools_bin_url, "-o", "binary.bin")
	err = cpmtools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cpmtools_src_url := "http://www.moria.de/~michael/cpmtools/files/cpmtools-2.23.src.tar.gz"
	cpmtools_cmd_src := exec.Command("curl", "-L", cpmtools_src_url, "-o", "source.tar.gz")
	err = cpmtools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cpmtools_cmd_direct := exec.Command("./binary")
	err = cpmtools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: libdsk")
exec.Command("latte", "install", "libdsk")
}
