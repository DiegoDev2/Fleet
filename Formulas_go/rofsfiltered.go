package main

// RofsFiltered - Filtered read-only filesystem for FUSE
// Homepage: https://github.com/gburca/rofs-filtered/

import (
	"fmt"
	
	"os/exec"
)

func installRofsFiltered() {
	// Método 1: Descargar y extraer .tar.gz
	rofsfiltered_tar_url := "https://github.com/gburca/rofs-filtered/archive/refs/tags/rel-1.7.tar.gz"
	rofsfiltered_cmd_tar := exec.Command("curl", "-L", rofsfiltered_tar_url, "-o", "package.tar.gz")
	err := rofsfiltered_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rofsfiltered_zip_url := "https://github.com/gburca/rofs-filtered/archive/refs/tags/rel-1.7.zip"
	rofsfiltered_cmd_zip := exec.Command("curl", "-L", rofsfiltered_zip_url, "-o", "package.zip")
	err = rofsfiltered_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rofsfiltered_bin_url := "https://github.com/gburca/rofs-filtered/archive/refs/tags/rel-1.7.bin"
	rofsfiltered_cmd_bin := exec.Command("curl", "-L", rofsfiltered_bin_url, "-o", "binary.bin")
	err = rofsfiltered_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rofsfiltered_src_url := "https://github.com/gburca/rofs-filtered/archive/refs/tags/rel-1.7.src.tar.gz"
	rofsfiltered_cmd_src := exec.Command("curl", "-L", rofsfiltered_src_url, "-o", "source.tar.gz")
	err = rofsfiltered_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rofsfiltered_cmd_direct := exec.Command("./binary")
	err = rofsfiltered_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: libfuse@2")
exec.Command("latte", "install", "libfuse@2")
}
