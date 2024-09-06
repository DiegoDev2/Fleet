package main

// Diamond - Accelerated BLAST compatible local sequence aligner
// Homepage: https://www.wsi.uni-tuebingen.de/lehrstuehle/algorithms-in-bioinformatics/software/diamond/

import (
	"fmt"
	
	"os/exec"
)

func installDiamond() {
	// Método 1: Descargar y extraer .tar.gz
	diamond_tar_url := "https://github.com/bbuchfink/diamond/archive/refs/tags/v2.1.9.tar.gz"
	diamond_cmd_tar := exec.Command("curl", "-L", diamond_tar_url, "-o", "package.tar.gz")
	err := diamond_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	diamond_zip_url := "https://github.com/bbuchfink/diamond/archive/refs/tags/v2.1.9.zip"
	diamond_cmd_zip := exec.Command("curl", "-L", diamond_zip_url, "-o", "package.zip")
	err = diamond_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	diamond_bin_url := "https://github.com/bbuchfink/diamond/archive/refs/tags/v2.1.9.bin"
	diamond_cmd_bin := exec.Command("curl", "-L", diamond_bin_url, "-o", "binary.bin")
	err = diamond_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	diamond_src_url := "https://github.com/bbuchfink/diamond/archive/refs/tags/v2.1.9.src.tar.gz"
	diamond_cmd_src := exec.Command("curl", "-L", diamond_src_url, "-o", "source.tar.gz")
	err = diamond_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	diamond_cmd_direct := exec.Command("./binary")
	err = diamond_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
