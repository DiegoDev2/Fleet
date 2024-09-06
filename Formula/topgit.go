package main

// Topgit - Git patch queue manager
// Homepage: https://github.com/mackyle/topgit

import (
	"fmt"
	
	"os/exec"
)

func installTopgit() {
	// Método 1: Descargar y extraer .tar.gz
	topgit_tar_url := "https://github.com/mackyle/topgit/archive/refs/tags/topgit-0.19.13.tar.gz"
	topgit_cmd_tar := exec.Command("curl", "-L", topgit_tar_url, "-o", "package.tar.gz")
	err := topgit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	topgit_zip_url := "https://github.com/mackyle/topgit/archive/refs/tags/topgit-0.19.13.zip"
	topgit_cmd_zip := exec.Command("curl", "-L", topgit_zip_url, "-o", "package.zip")
	err = topgit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	topgit_bin_url := "https://github.com/mackyle/topgit/archive/refs/tags/topgit-0.19.13.bin"
	topgit_cmd_bin := exec.Command("curl", "-L", topgit_bin_url, "-o", "binary.bin")
	err = topgit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	topgit_src_url := "https://github.com/mackyle/topgit/archive/refs/tags/topgit-0.19.13.src.tar.gz"
	topgit_cmd_src := exec.Command("curl", "-L", topgit_src_url, "-o", "source.tar.gz")
	err = topgit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	topgit_cmd_direct := exec.Command("./binary")
	err = topgit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
