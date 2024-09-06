package main

// Gaffitter - Efficiently fit files/folders to fixed size volumes (like DVDs)
// Homepage: https://gaffitter.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installGaffitter() {
	// Método 1: Descargar y extraer .tar.gz
	gaffitter_tar_url := "https://downloads.sourceforge.net/project/gaffitter/gaffitter/1.0.0/gaffitter-1.0.0.tar.gz"
	gaffitter_cmd_tar := exec.Command("curl", "-L", gaffitter_tar_url, "-o", "package.tar.gz")
	err := gaffitter_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gaffitter_zip_url := "https://downloads.sourceforge.net/project/gaffitter/gaffitter/1.0.0/gaffitter-1.0.0.zip"
	gaffitter_cmd_zip := exec.Command("curl", "-L", gaffitter_zip_url, "-o", "package.zip")
	err = gaffitter_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gaffitter_bin_url := "https://downloads.sourceforge.net/project/gaffitter/gaffitter/1.0.0/gaffitter-1.0.0.bin"
	gaffitter_cmd_bin := exec.Command("curl", "-L", gaffitter_bin_url, "-o", "binary.bin")
	err = gaffitter_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gaffitter_src_url := "https://downloads.sourceforge.net/project/gaffitter/gaffitter/1.0.0/gaffitter-1.0.0.src.tar.gz"
	gaffitter_cmd_src := exec.Command("curl", "-L", gaffitter_src_url, "-o", "source.tar.gz")
	err = gaffitter_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gaffitter_cmd_direct := exec.Command("./binary")
	err = gaffitter_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
