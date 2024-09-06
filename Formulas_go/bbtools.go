package main

// Bbtools - Brian Bushnell's tools for manipulating reads
// Homepage: https://jgi.doe.gov/data-and-tools/bbtools/

import (
	"fmt"
	
	"os/exec"
)

func installBbtools() {
	// Método 1: Descargar y extraer .tar.gz
	bbtools_tar_url := "https://downloads.sourceforge.net/bbmap/BBMap_39.08.tar.gz"
	bbtools_cmd_tar := exec.Command("curl", "-L", bbtools_tar_url, "-o", "package.tar.gz")
	err := bbtools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bbtools_zip_url := "https://downloads.sourceforge.net/bbmap/BBMap_39.08.zip"
	bbtools_cmd_zip := exec.Command("curl", "-L", bbtools_zip_url, "-o", "package.zip")
	err = bbtools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bbtools_bin_url := "https://downloads.sourceforge.net/bbmap/BBMap_39.08.bin"
	bbtools_cmd_bin := exec.Command("curl", "-L", bbtools_bin_url, "-o", "binary.bin")
	err = bbtools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bbtools_src_url := "https://downloads.sourceforge.net/bbmap/BBMap_39.08.src.tar.gz"
	bbtools_cmd_src := exec.Command("curl", "-L", bbtools_src_url, "-o", "source.tar.gz")
	err = bbtools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bbtools_cmd_direct := exec.Command("./binary")
	err = bbtools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
