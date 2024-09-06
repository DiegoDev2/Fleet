package main

// Libpipeline - C library for manipulating pipelines of subprocesses
// Homepage: https://libpipeline.gitlab.io/libpipeline/

import (
	"fmt"
	
	"os/exec"
)

func installLibpipeline() {
	// Método 1: Descargar y extraer .tar.gz
	libpipeline_tar_url := "https://download.savannah.nongnu.org/releases/libpipeline/libpipeline-1.5.8.tar.gz"
	libpipeline_cmd_tar := exec.Command("curl", "-L", libpipeline_tar_url, "-o", "package.tar.gz")
	err := libpipeline_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libpipeline_zip_url := "https://download.savannah.nongnu.org/releases/libpipeline/libpipeline-1.5.8.zip"
	libpipeline_cmd_zip := exec.Command("curl", "-L", libpipeline_zip_url, "-o", "package.zip")
	err = libpipeline_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libpipeline_bin_url := "https://download.savannah.nongnu.org/releases/libpipeline/libpipeline-1.5.8.bin"
	libpipeline_cmd_bin := exec.Command("curl", "-L", libpipeline_bin_url, "-o", "binary.bin")
	err = libpipeline_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libpipeline_src_url := "https://download.savannah.nongnu.org/releases/libpipeline/libpipeline-1.5.8.src.tar.gz"
	libpipeline_cmd_src := exec.Command("curl", "-L", libpipeline_src_url, "-o", "source.tar.gz")
	err = libpipeline_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libpipeline_cmd_direct := exec.Command("./binary")
	err = libpipeline_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
